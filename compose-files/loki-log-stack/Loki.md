# Loki

# Components

## Distributor

负责处理客户端日志，是日志处理的第一步。distributor一旦接收到日志数据，就会将其分割为日志块(batches)，同时并行地发送到多个`ingester`。

distributor与ingester直接通过grpc进行交流通信。

## Hashing

distributor 使用一致性hash(`in conjunction with a configurable replication factor `)来决定哪一个ingester服务实例应该接受日志数据。

hash值的计算基于日志的标签`labels`和租户`ID`。

所有的ingesters通过他们自己的token将自身注册到hash环中。这个hash环存储在Consul中，用来实现一致性hash。distributor会寻找到与日志的hash值最相近的token，然后将日志数据发送给持有该token的ingester。

## Quorum consistency

所有的distributors对同一哈希环共享访问权限，因此可以将日志的写请求发送给任何distributor。

为了确保查询结果的一致性，Loki 在读写时使用 [Dynamo style](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf) Quorum一致性。就是说，写入日志的请求在响应之前，distributor将等待至少一半的ingester写入数据作出积极响应。

## Ingester

ingester负责将日志数据写入持久存储后台。

ingester会按照时间戳升序来验证日志的顺序，如果接收到的日志不符合期望的顺序，将拒绝写入日志并向用户返回错误。

日志根据不同的、唯一的标签`labels`或标签组合以数据块的形式`chunks`暂存在内存中, 然后刷新至存储后台。

如果一个ingester处理崩溃了或者异常退出，所有未刷新至存储的日志数据将丢失。Loki通常对每一个日志启用3个副本来防止丢失的风险。

## Timestamp Ordering

时间戳顺序的处理规则：

对于接收的每一条日志，都要求其(纳秒)时间戳晚于上一条日志记录。当同一数据流中两条日志具有相同的时间戳时：

1. 如果日志的时间戳与上一条日志的时间戳相同，且日志文本内容也相同，则认为是重复日志，该条日志将被忽略；
2. 如果日志的时间戳与上一条日志的时间戳相同，但日志文本内容不相同，则认为是合理写入；

## handoff

默认地， 如果一个ingester被关闭且试图离开hash环，在数据刷新前Loki将等待是否有新的ingester加入并初始化一个交接`handoff`。交接的意思是，如果有新的ingester加入，被关闭的ingester所拥有的tokens和内存数据块都会转移给新的ingester。

这个机制是为了避免ingester关闭时刷新所有的数据块，这个处理是很缓慢的。

## Querier

querier具体处理实际的 [LogQL](https://grafana.com/docs/loki/latest/logql/)查询。它首先会在所有ingester的内存数据中查找结果，然后才会去后台存储中查询。

## Query frontend

是一个可选的组件，主要作用是调度查询请求，在可能的时候进行并行和缓存。

## Chunk Store

Chunk Store是Loki的持久化数据存储，用来支持交互式查询和持续写入，而无需进行后台维护。与其他核心组件不一样，Chunk Store并不是一个单独的服务、job或者进程，而是嵌入在ingester和querier之间的一个库。它由两部分组成：

* 数据库索引`index`: 索引可以通过 [DynamoDB from Amazon Web Services](https://aws.amazon.com/dynamodb), [Bigtable from Google Cloud Platform](https://cloud.google.com/bigtable), or [Apache Cassandra](https://cassandra.apache.org/) 进行备份。
* 数据块自身的存储`Key-Value:`  可以存储至DynamoDB, Bigtable, Cassandra again, 或者对象存储 [Amazon * S](https://aws.amazon.com/s3)中。



# 安装

更多安装方式

https://grafana.com/docs/loki/latest/installation/docker/

## docker

https://github.com/grafana/loki/tree/master/production

```yaml
version: "3"

networks:
  loki:

services:
  loki:
    image: grafana/loki:2.0.0
    ports:
      - "3100:3100"
    command: -config.file=/etc/loki/local-config.yaml
    networks:
      - loki

  promtail:
    image: grafana/promtail:2.0.0
    volumes:
      - /var/log:/var/log
    command: -config.file=/etc/promtail/config.yml
    networks:
      - loki

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    networks:
      - loki
```

# Label

标签是键值对，可以被定义为任何值。标签键值的组合定义了数据流，是描述一个数据流的元数据。如果改变了其中任何一个值，就代表会新建一个数据流。

标签是 Loki 日志数据的索引。它们用于查找压缩日志内容，这些内容单独存储为块。标签和值的每个唯一组合都定义了流，流的日志被批处理、压缩和存储为块。

标签的定义方式有两种：

* 静态定义：固定的标签值，每一个日志行都会加上相同的标签
* 动态定义：根据日志行的内容而定，每一个日志行会加上相同的标签Key

## 静态定义

```yaml
scrape_configs:
# http://localhost:3100/metrics可以直接看到Loki系统日志
- job_name: system # jobname用于区别从不同的项目收集日志， system定义的是Loki自身产生的日志
  static_configs:
  - targets: # 可选配置。老版本的promtail需要进行该配置。Prometheus的服务发现也需要这个入口
      - localhost
    labels: # 日志标签，标签{env="dev" job=”syslog”}会添加到每一行日志上，好的命名方式通常为：环境名-job名或者app名
      job: varlogs  
      env: dev
      __path__: /var/log/*log # 日志存储的路径，Promtail将会从这个路径去收集日志
```

## 动态定义

我们可以提取日志行的内容作为标签的值。假如有这样格式的日志:

```verilog
11.11.11.11 - frank [25/Jan/2000:14:00:01 -0500] "GET /1986.js HTTP/1.1" 200 932 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; de; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7 GTB6"
```

通过正则表达式的捕获组提取请求的方法和响应状态码作为标签：

```yaml
- job_name: system
   pipeline_stages:
      - regex:
        expression: "^(?P<ip>\\S+) (?P<identd>\\S+) (?P<user>\\S+) \\[(?P<timestamp>[\\w:/]+\\s[+\\-]\\d{4})\\] \"(?P<action>\\S+)\\s?(?P<path>\\S+)?\\s?(?P<protocol>\\S+)?\" (?P<status_code>\\d{3}|-) (?P<size>\\d+|-)\\s?\"?(?P<referer>[^\"]*)\"?\\s?\"?(?P<useragent>[^\"]*)?\"?$"
    - labels:
        action:
        status_code:
   static_configs:
   - targets:
      - localhost
     labels:
      job: apache
      env: dev
      __path__: /var/log/apache.log
```

然后，在Loki中会得到如下格式的日志行:

```verilog
{job=”apache”,env=”dev”,action=”GET”,status_code=”200”} 11.11.11.11 - frank [25/Jan/2000:14:00:01 -0500] "GET /1986.js HTTP/1.1" 200 932 "-" "Mozilla/5.0 (Windows; U; Windows NT 5.1; de; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7 GTB6"
```

## 需要注意的问题

**在Loki种定义更多的标签，并不能使得查询更快。**

全文索引中索引的字段越多，查询的效率会更快，但存在的缺陷是，日志数据的全文索引大小通常与日志数据本身相同或更大，随着存储更多的日志，索引会快速变大（所以你会听到说elasticsearch很耗内存）。

Loki使用标签label作为索引，定义很多标签，查询会更快吗？通常不会。因为只要有一个label的值不同，就会创建一个新的数据流，然后存储在新的数据块中。假如有M种请求方法、N种状态码，则会创建M\*N个数据流。如果再使用ip作为label，则10万个用户就会创建100000\*M\*N个数据流，多标签的结果就是：巨量索引+巨量小存储块=>使得Loki的性能极低。

所以，**绝对不要使用取值很多的字段作为标签**，比如ip，每一个ip都会创建一个数据流。

为了保证效率，**Loki是通过查询并行化的方式来解决问题的**。Loki将查询分解为小块，并并行调度它们。例如如下查询语句：

```verilog
{job=”apache”} |= “11.11.11.11”
```

在幕后，Loki 首先根据标签索引匹配到`job="apache"`的每个数据流区块，然后并行化的在这些数据块中查找IP地址为“11.11.11.11”的日志行。

所以Loki的索引是很轻量的，存储增长也非常缓慢。

# web界面

## Loki

* Loki系统日志

   http://localhost:3100/metrics 

## promtail

* service-discovery：展示所有发现的targets，以及target被删掉的原因。
  http://localhost:9080/service-discovery

* targets：只展示活跃的targets及其标签、文件等信息

  http://localhost:9080/targets

# 最佳实践

 

* **静态标签是好的选择**

  比如使用host, application, 和environment 都不错。

* **谨慎使用动态标签**

  能不使用则不使用，除非有必要。

* **标签值始终是有限的**

  动态标签值的取值数最好在10个以内，但也有例外，比如主机数host，有1000台机器，对应1000个host也是正常的。

* 注意Loki的不同客户端(负责日志收集的如promtail、Fluentd、Fluent Bit等)可能会用到的动态标签

  应该始终检查是否有不当的动态标签，比如请求的`requestId`, 这个应该作为内容过滤选项，而不是标签。

* 配置缓存

  Loki 可以在多个级别缓存数据，从而显著提高性能。

* **保证每个日志流的时间顺序**

  如果某条日志的时间戳小于上一条，则该条日志会被拒绝写入。

* **使用`chunk_target_size`**

  Loki 默认`chunk_target_size=1536000`, 即将每个数据流存储为目标压缩大小1.5M的数据块，Loki将不断使用原始的日志数据（5-10倍压缩率，7.5-10M大小）来填充这个数据块以达到目标压缩大小，处理效率高。

* 使用`-print-config-stderr` 或者 `-log-config-reverse-order`选项

  这两个选项主要用于在Loki启动时便于观察整个配置对象。

