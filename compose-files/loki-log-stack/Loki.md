# Loki

# 组件

## Distributor

负责处理客户端日志，是日志处理的第一步。distributor一旦接收到日志数据，就会将其分割为日志块(batches)，同时并行地发送到多个`ingester`。

distributor与ingester直接通过grpc进行交流通信。

## Hashing

distributor 使用一致性hash(`in conjunction with a configurable replication factor `)和可配置的副本因子来决定哪一个ingester服务实例应该接受日志数据。

hash值的计算基于日志的标签`labels`和租户`ID`。

所有的ingesters通过他们自己的token(随机未签名的32位数字)将自身注册到hash环中。这个hash环存储在Consul中，用来实现一致性hash。distributor会寻找到与日志的hash值最相近的token，然后将日志数据发送给持有该token的ingester。

## Quorum consistency

所有的distributors对同一哈希环共享访问权限，因此可以将日志的写请求发送给任何distributor。

为了确保查询结果的一致性，Loki 在读写时使用 [Dynamo style](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf) Quorum一致性。就是说，写入日志的请求在响应之前，distributor将等待至少一半的ingester写入数据作出积极响应。

## Ingester

ingester负责将日志数据写入持久存储后台。

ingester会按照时间戳升序来验证日志的顺序，如果接收到的日志不符合期望的顺序，将拒绝写入日志并向用户返回错误。

日志根据不同的、唯一的标签`labels`或标签组合以数据块的形式`chunks`暂存在内存中, 然后根据配置的时间间隔刷新至存储后台。

如果一个ingester处理崩溃了或者异常退出，所有未刷新至存储的日志数据将丢失。Loki通常对每一个日志启用3个副本来防止丢失的风险。

## Timestamp Ordering

时间戳顺序的处理规则：

对于接收的每一条日志，都要求其(纳秒)时间戳晚于上一条日志记录。当同一数据流中两条日志具有相同的时间戳时：

1. 如果日志的时间戳与上一条日志的时间戳相同，且日志文本内容也相同，则认为是重复日志，该条日志将被忽略；
2. 如果日志的时间戳与上一条日志的时间戳相同，但日志文本内容不相同，则认为是合理写入；

## handoff

默认地， 如果一个ingester被关闭且试图离开hash环时，在数据刷新前该ingester将等待是否有新的ingester加入，并尝试启动一个交接`handoff`。交接的意思是，如果有新的ingester加入，被关闭的ingester所拥有的tokens和内存数据块都会转移给新的ingester。

这个机制是为了避免ingester关闭时刷新所有的数据块，这个处理是很缓慢的。

## Querier

querier具体处理实际的 [LogQL](https://grafana.com/docs/loki/latest/logql/)查询。它首先会在所有ingester的内存数据中查找结果，然后才会去后台存储中查询。

## Query frontend

是一个可选的组件，主要作用是调度查询请求，在可能的时候进行并行和缓存。

## Chunk Store

Chunk Store是Loki的持久化数据存储，用来支持交互式查询和持续写入，而无需进行后台维护。与其他核心组件不一样，Chunk Store并不是一个单独的服务、job或者进程，而是嵌入在ingester和querier之间的一个库。它由两部分组成：

* 数据库索引`index`: 索引可以通过 [DynamoDB from Amazon Web Services](https://aws.amazon.com/dynamodb), [Bigtable from Google Cloud Platform](https://cloud.google.com/bigtable), or [Apache Cassandra](https://cassandra.apache.org/) 进行备份。
* 数据块自身的存储`Key-Value:`  可以存储至DynamoDB, Bigtable, Cassandra again, 或者对象存储 [Amazon * S](https://aws.amazon.com/s3)中。

## Ruler

Loki的警报(alert)组件。负责持续评估一组可配置的查询，然后在达成某些条件时发出警报，例如错误日志的百分比很高。

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

# 配置

Loki的配置文件通常位于`/etc/loki/loki.yaml`中，如果是Loki 镜像，官方将配置文件放在镜像中的`/etc/loki/local_config.yaml`文件中（可查看官方的镜像文件[Loki Dockerfile文件](https://github.com/grafana/loki/tree/master/cmd/loki)）。配置文件包含了Loki服务信息以及组件信息，具体配置取决于Loki是以何种模式启动的。

> 详细的配置可查看官方文档，[传送门](https://grafana.com/docs/loki/latest/configuration/)

## 运行时打印Loki配置

在启动命令行后添加如下参数(flags)之一，可以让Loki在运行时打印出当前自定义的及默认的配置信息对象：

* `-print-config-stderr`
* `-log-config-reverse-order`
* `-print-config-stderr=true`

## 指定配置文件

在命令行后通过`-config.file`指定要加载的配置文件，如：

```shell
-config.file=/etc/loki/local-config.yaml
```

## 在配置文件中使用环境变量

> 该功能在Loki2.1+版本才可用

启用该功能需要在命令行后添加`-config.expand-env=true`。

在配置文件中使用环境变量：

格式：`{VAR}`或`${VAR:default_value}`

 其中，`VAR`是环境变量名，大小写敏感；`default_value`为默认值。

## 整体配置内容

```yaml
# Loki运行包含的组件或模块，支持选项有：
# all, distributor, ingester, querier, query-frontend, table-manager.
target: <string> | default = "all"

# 启用auth授权
# 如果为true，则必须在请求头中添加X-Scope-OrgID=<tenantID>
# 如果为false，则X-Scope-OrgID始终为"fake"
auth_enabled: <boolean> | default = true

# Loki服务配置
server: <server_config>

# distributor 组件配置
distributor: <distributor_config>

# querier组件配置
querier: <querier_config>

# query_frontend 组件配置
frontend: <query_frontend_config>

# 配置查询拆分、查询缓存
query_range: <queryrange_config>

# ruler组件配置，配置告警、annotation等
ruler: <ruler_config>

# 配置distributor如何连接到ingester
ingester_client: <ingester_client_config>

# 组件ingester配置，包含ingester如何注册至key-value存储
ingester: <ingester_config>

# 日志数据存储地址配置
storage_config: <storage_config>

# 在指定的存储中具体如何存储数据的配置
chunk_store_config: <chunk_store_config>

# chunk和index的 schema配置、以及schema存储地址配置
schema_config: <schema_config>

# 全局或租户限制配置
limits_config: <limits_config>

# 在querier组件中负责执行查询的worker配置
frontend_worker: <frontend_worker_config>

# table_manager配置，主要配置日志数据的保留时间
table_manager: <table_manager_config>

# 运行时配置，负责重新载入运行时的配置文件
runtime_config: <runtime_config>

# tracing配置
tracing: <tracing_config>
```

## server

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#server_config)

```yaml
# Loki服务配置
server:
  http_listen_address: <string>
  http_listen_port: <int> | default = 80
  grpc_listen_address: <string>
  grpc_listen_port: <int> | default = 9095
  # 日志级别过滤，只有在配置的level水平之上才会接收
  log_level: <string> | default = "info"
  # http路由前缀，"/api/prom/"应该已经弃用了，目前应为"/loki/api/v1"
  http_prefix: <string> | default = "/api/prom"
```

## distributor

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#distributor_config)

```yaml
# distributor 组件配置
distributor:
  ring:
    kvstore:
      # hash环的后台存储配置，使用的是键值对存储库
      # 支持 consul, etcd, inmemory, memberlist
      store: <string>

      # 存储键的前缀，必须以斜杠'/'结尾
      prefix: <string> | default = "collectors/"

      # 针对设置的不同store，配置不同的store backend，只能配置其中之一
      consul: <consul_config>
      etcd: <etcd_config>
      memberlist: <memberlist_config>
```

## querier

> querier基本不用自定义，采用默认就好，详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#querier_config)

```yaml
# querier组件配置
querier: 
	# http api: /loki/api/v1/query 查询超时
  query_timeout: <duration> | default = 1m
 
 # http api: /loki/api/v1/tail 查询持续时间
  tail_max_duration: <duration> | default = 1h
 
 # http请求延迟
  extra_query_delay: <duration> | default = 0s
 
  # 设置一个时间，当超过这个值时，查询将不会发往ingester
  # 0代表所有查询都会发往ingester
  query_ingesters_within: <duration> | default = 0s
  
  # 查询引擎配置
  engine:
    timeout: <duration> | default = 3m
    max_look_back_period: <duration> | default = 30s
```

## frontend

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#query_frontend_config)

```yaml
# query_frontend 组件配置
frontend: 
  # 每个查询端每个租户最大未处理或未完成请求数
  max_outstanding_per_tenant: <int> | default = 100

  # http response 压缩
  compress_responses: <boolean> | default = false

  # tail查询的代理url
  tail_proxy_url: <string> | default = ""
```

## query_range

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#queryrange_config)

```yaml
# 配置查询拆分、查询缓存
query_range:
  # 按照interval进行查询分割以进行并行查询。默认0，禁用此功能。
  # 应使用24小时的倍数以避免querier下载和处理同样的chunks。
  # 如果开启了缓存，这也决定了缓存keys的选择。
  split_queries_by_interval: <duration> | default = 0s

  # 是否开启结果缓存
  cache_results: <boolean> | default = false

  results_cache:
    # 结果缓存配置
    cache: <cache_config>

  # 请求的最大重试次数
  max_retries: <int> | default = 5

  # 仅针对chunks存储引擎的分片查询并行化
  parallelise_shardable_queries: <boolean> | default = false
```

## ruler

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#ruler_config)

```yaml
# ruler组件配置，配置告警
ruler: 
  storage:
    # rule后台存储类型，可取值azure, gcs, s3, swift, local
    type: local
    local:
      # rules存储文件目录
      directory: <filename> | default = ""
  
  # 存储临时rule文件的路径
  rule_path: <filename> | default = "/rules"
  
  # 发送告警通知的url，多个url以逗号分隔。
  # 每个url被视为单独的组
  alertmanager_url: <string> | default = ""
  
  # 启用DNS服务解析alert hosts
  enable_alertmanager_discovery: <boolean> | default = false
  
  # DNS解析alert hosts的刷新间隔
  alertmanager_refresh_interval: <duration> | default = 1m
  
  ring:
    kvstore:
      # hash环的后端存储，支持: consul, etcd, inmemory, memberlist, multi.
      store: <string> | default = "consul"
  
      # 存储keys的前缀，必须以'/'结尾'
      prefix: <string> | default = "rulers/"
  
      # 存储库配置,对应store的值进行配置
      consul: <consul_config>
      etcd: <etcd_config>
      multi:
        primary: <string> | default = ""
        secondary: <string> | default = ""
        mirror_enabled: <boolean> | default = false
        mirror_timeout: <duration> | default = 2s

    # 每一个ingester的tokens数量
    num_tokens: <int> | default = 128
  
  # 刷新rules组的间隔
  flush_period: <duration> | default = 1m
  
  # 是否启用rulers API
  enable_api: <boolean> | default = false
```

## frontend_worker

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#frontend_worker_config)

```yaml
# 在querier组件中负责执行查询的worker配置
frontend_worker:
  # 查询前端的地址，格式为 host:port
  frontend_address: <string> | default = ""

  # 查询的并行数
  parallelism: <int> | default = 10
```

## ingester_client

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#ingester_client_config)

```yaml
# 配置distributor如何连接到ingester
ingester_client:
  # 连接池配置
  pool_config:
    # 是否对ingester进行健康检查
    health_check_ingesters: <boolean> | default = false

    # ingester服务关闭后，经过多久清理掉客户端连接
    client_cleanup_period: <duration> | default = 15s

  # 远程客户端请求超时
  remote_timeout: <duration> | default = 5s

  # 连接ingester的grpc配置
  grpc_client_config: <grpc_client_config>
```

## ingester

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#ingester_config)

```yaml
# 组件ingester配置，包含ingester如何注册至key-value存储
ingester:
  # ingester运作的生命周期配置
  lifecycler:
  	# 当一个ingester退出时，新的ingester在60s后自动交接
  	join_after: <duration> | default = 0s
    # ingester退出前的休眠时间，主要用于确保metrics抓取
    final_sleep: <duration> | default = 30s
    ring:
      kvstore:
        # hash ring 的后端存储. 支持: consul, etcd, inmemory, memberlist
        store: <string> | default = "consul"

        # 存储keys的前缀, 必须以'/'结尾
        prefix: <string> | default = "collectors/"
        consul: <consul_config>
        etcd: <etcd_config>
        memberlist: <memberlist_config>

      # 副本因子，即负责读写的ingester数
      replication_factor: <int> | default = 3

    # 当前生命周期中，ingester加入hash ring将生成的tokens数
    num_tokens: <int> | default = 128


  # 每一个日志流的并发刷新数（chunk刷新至store）
  concurrent_flushes: <int> | default = 16

  # chunks被刷新后，在内存中继续保存多久
  chunk_retain_period: <duration> | default = 15m

  # 当chunk的大小没有达到设定的最大值，如果一直没有任何更新，
  # 在达到一定时间后将强制刷新，默认30m没有更新就会刷新至store
  chunk_idle_period: <duration> | default = 30m

  # 当接收的数据达到256kb时，将在chunk中进行压缩
  chunk_block_size: <int> | default = 262144

  # chunk压缩后的目标大小，这个值不是精确的，可能稍大，也可能明显小于目标值。
  # 默认为0，则chunk固定为10个block
  # 指定一个数值，则会根据目标大小创建blocks数
  chunk_target_size: <int> | default = 0

  # 压缩算法：
  # - `gzip` 压缩率高，但解压速度较慢(144kB/chunk)
  # - `lz4` 压缩速度最快(188kB/chunk)
  # - `snappy` 快且流行的压缩算法(272kB/chunk)
  chunk_encoding: <string> | default = gzip

  # chunks在内存中保存的时间，如果某个时序流运行时间超过了这个值，chunk将被刷新至store，然后创建一个新的chunk
  max_chunk_age: <duration> | default = 1h
```

## storage_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#storage_config)

```yaml
# 日志数据后端存储配置
storage_config:
  aws:
  bigtable:
  gcs:
  cassandra:
  swift:

  # index store in boltdb——旧版本使用
  boltdb:
    # Location of BoltDB index files.
    directory: <string>

  # index store in boltdb_shipper——新版使用
  boltdb_shipper:
    active_index_directory: /loki/boltdb-shipper-active
    cache_location: /loki/boltdb-shipper-cache
    cache_ttl: 24h  
    shared_store: filesystem

  # chunk store
  filesystem:
    # Directory to store chunks in.
    directory: <string>

  # 索引缓存校验, 其值必须小于等于ingester.chunk_idle_period
  index_cache_validity: <duration> | default = 5m

  # 每个batch提取的最大chunks数
  max_chunk_batch_size: <int> | default = 50
```

## chunk_store_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#chunk_store_config)

```yaml
# 在指定的存储中具体如何存储数据的配置
chunk_store_config:
  # chunks 缓存配置
  chunk_cache_config: <cache_config>

  # 去重缓存配置
  write_dedupe_cache_config: <cache_config>

  # chunk从更新到保存至store的最小时间
  min_chunk_age: <duration>

  # 只缓存这个时间值之前的索引，默认不启用
  cache_lookups_older_than: <duration>

  # 查询回溯的最大时间限制，默认不启用
  # 值必须小于等于table_manager.retention_period
  max_look_back_period: <duration>
```

## schema_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#schema_config)

```yaml
# chunk和index的 schema配置、以及schema存储地址配置
schema_config: <schema_config>
  configs:
    # index buckets 开始创建的日期。格式为YYYY-MM-DD
    - from: <daytime>
    # 使用何种index store. 
    # 支持aws, aws-dynamo, gcp, bigtable, bigtable-hashed,cassandra, boltdb, boltdb-shipper
    store: <string>

    # 使用何种chunk store.
    # 支持 aws, azure, gcp, bigtable, gcs, cassandra, swift , filesystem. 
    # 忽略设置将使用和store一样的值
    object_store: <string>

    # schema 版本， 当前推荐v11
    schema: <string>

    # 配置index如何更新和存储
    index:
      # 索引表前缀
      prefix: <string>
      # 索引表期间, 最好的是24小时
      # "BoltDB shipper works best with 24h periodic index files."
      # https://grafana.com/docs/loki/latest/operations/storage/boltdb-shipper/
      period: <duration> | default = 168h
      # 给所有表添加tags， map结构
      tags:
        <string>: <string>

    # 配置chunk更新和存储的细节
    chunks:
      # chunks表前缀
      prefix: <string>
      # chunks表期间
      period: <duration> | default = 168h
      # 给所有表添加tags， map结构
      tags:
        <string>: <string>

    # 创建多少分片，仅使用于schema v10及以上
    row_shards: <int> | default = 16
```

## limits_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#limits_config)

```yaml
# 全局或租户限制配置
limits_config:
 
  # 是否拒绝老样本
  reject_old_samples: <bool> | default = false

  # 拒绝之前，可接受的最大样本年龄
  reject_old_samples_max_age: <duration> | default = 336h

  # 查询返回的最大日志数
  max_entries_limit_per_query: <int> | default = 5000 

  # 每个查询可以提取的最大chunks数
  max_chunks_per_query: <int> | default = 2000000

  # frontend 可调度的最大查询并行数
  max_query_parallelism: <int> | default = 14
  
  # 每个用户日志数据流引入的速率限制，默认4mb/s
  ingestion_rate_mb: <float> | default = 4
```

## table_manager

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#table_manager_config)

```yaml
# table_manager配置，主要配置日志数据的保留时间
table_manager:
  # 删除表格保留的开与关
  retention_deletes_enabled: <boolean> | default = false

  # 被删除前，需要保留多久之前的表格。
  # 默认0s，禁止删除
  # 这个值必须是index/chunks 的table period的倍数
  retention_period: <duration> | default = 0s
```

## compactor_config

compactor是使用boltdb-shipper时特定的一个服务，主要通过去除重复索引文件、对每个索引表的boltdb文件进行合并来减少索引的大小， 可以提升查询效率。如果单个ingester每天创建超过96个文件，则强烈推荐运行一个compactor。[原文](https://grafana.com/docs/loki/latest/operations/storage/boltdb-shipper/#compactor)

```yaml
compactor:
  working_directory: /loki/boltdb-shipper-compactor
  shared_store: filesystem
```

## tracing_config

```yaml
# tracing配置
tracing: 
  # 是否启用jaeger调用跟踪
  enabled: <boolean>| default = true
```

## runtime_config

> 目前只有limits和multi kv-store可以进行运行时配置。
>
> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#runtime-configuration-file)

```yaml
# 运行时配置，负责重新载入运行时的配置文件
runtime_config:
  # 运行时配置文件
  file: <string>| default = empty
  # 检测运行时配置的间隔
  period: <duration>| default = 10s
```

## grpc_client_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#grpc_client_config)

```yaml
grpc_client_config:
  max_recv_msg_size: <int> | default = 104857600
  max_send_msg_size: <int> | default = 16777216
  use_gzip_compression: <bool> | default = false
  rate_limit: <float> | default = 0
  rate_limit_burst: <int> | default = 0
  backoff_on_ratelimits: <bool> | default = false
  backoff_config:
    min_period: <duration> | default = 100ms
    max_period: <duration> | default = 10s
    max_retries: <int> | default = 10
```

## consul_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#consul_config)

```yaml
consul_config:
  host: <string> | default = "localhost:8500"
  acl_token: <string>
  http_client_timeout: <duration> | default = 20s
  consistent_reads: <boolean> | default = true
```

## etcd_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#etcd_config)

```yaml
etcd_config:
  endpoints: <list of string> | default = []
  dial_timeout: <duration> | default = 10s
  max_retries: <int> | default = 10
```

## memberlist_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#memberlist_config)

`memberlist`配置，用以发现和连接`distributors, ingesters`以及`queriers`. 这个配置在三个组件之间是一个唯一的共享环。

## cache_config

> 详细配置[传送门](https://grafana.com/docs/loki/latest/configuration/#cache_config)

```yaml
cache_config:
  enable_fifocache: <boolean>
  default_validity: <duration>
  
  # 使用memcached时配置
  background:    
    writeback_goroutines: <int> | default = 10
    writeback_buffer: <int> = 10000

  memcached:
    expiration: <duration>
    batch_size: <int>
    parallelism: <int> | default = 100

  memcached_client:    
    host: <string>
    service: <string> | default = "memcached"
    timeout: <duration> | default = 100ms
    max_idle_conns: <int> | default = 16
    update_interval: <duration> | default = 1m
    consistent_hash: <bool>
  
  redis:    
    endpoint: <string>
    master_name: <string>
    timeout: <duration> | default = 100ms
    expiration: <duration> | default = 0s
    db: <int>
    pool_size: <int> | default = 0
    password: <string>
    enable_tls: <boolean> | default = false
    idle_timeout: <duration> | default = 0s
    max_connection_age: <duration> | default = 0s
  
  fifocache:    
    max_size_bytes: <string> | default = ""
    max_size_items: <int> | default = 0
    validity: <duration> | default = 0s
```



# LogQL

## Log Queries

查询日志内容。包含两个部分：

* ` log stream selector:`日志流选择器 。根据标签值选择日志流。

  例如`{app="rabbit", env="prod"}`

* `log pipeline:`日志管道。可以跟在日志流选择器之后。

  例如 `|="error"`

### 标签匹配运算符

| 运算符 | 含义       | 示例                |
| ------ | ---------- | ------------------- |
| `=`    | 完全相等   | `{app="rabbit"}`    |
| `!=`   | 不相等     | `{app!="rabbit"}`   |
| `=~`   | 正则匹配   | `{env=~"prod*"}`    |
| `!~`   | 正则不匹配 | `{env!~"prod_\d+"}` |

### 管道表达式

#### 行筛选表达式

| 表达式 | 含义          | 示例                                  |
| ------ | ------------- | ------------------------------------- |
| `|=`   | 行 包含       | `{job="mysql"}|="error"`              |
| `!=`   | 行 不包含     | `{job="mysql"}!="error"`              |
| `|~`   | 行 正则匹配   | `{job="mysql"} |~ ` \`total_num=\d+\` |
| `!~`   | 行 正则不匹配 | `{job="mysql"} !~`\`total_num=\d+\`   |

#### 解析器表达式

解析器表达式可以从日志内容中解析和提取标签。

| 解析器           | 日志示例                                                     | 结果示例                                                     |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `|json`          | {<br/>	"protocol": "HTTP/2.0",<br/>	"servers": ["129.0.1.1","10.2.1.3"],<br/>	"request": {<br/>		"method": "GET"<br/>	}<br/>} | {<br />     protocol="HTTP/2.0",<br />     request_method="GET"<br />}<br />数组会被忽略 |
| `|logfmt`        | "method=GET path=/ host=grafana.net"                         | {method="GET",path="/",host="grafana.net"}<br />提取所有键值对 |
| `|regexp "<re>"` | "POST /api/prom/api/v1/query_range (200) 1.5s"<br />`regexp "(?P<method>\\w+) (?P<path>[\\w/]+) \\((?P<status>\\d+?)\\) (?P<duration>.*)"` | {method="POST", path="/", status="200", duration="1.5s"}<br />使用正则命名组捕获 |

#### 标签筛选表达式

```shell
# 格式: 标签名 操作符 标签值
```

```shell
# 例如:
| duration >= 20ms or size == 20kb and method!~"2.."
| duration >= 20ms or size == 20kb | method!~"2.."
| duration >= 20ms or size == 20kb , method!~"2.."
| duration >= 20ms or size == 20kb  method!~"2.."

# and、|、空格 均表示且
# or 表示或者
```

标签值有四种类型:

* `String`:  " " 或``标示，如status="200"
* `Duration`:  单位为`ns, us (or  µs), ms, s, m, h`, 如 `cost >= 20s`
* `Number`: 数字， 如 `count==50`
* `Bytes`: 字节数。单位为`“b”, “kib”, “kb”, “mib”, “mb”, “gib”, “gb”, “tib”, “tb”, “pib”, “pb”, “eib”, “eb”`如`size>=1024kb`

#### 行格式化表达式

`|line_format`使用`text/template`模板语法格式，可以重写日志行内容.

例如使用一个请求的query和duration重写日志内容：

```shell
{container="frontend"} | logfmt | line_format "{{.query}} {{.duration}}"
```

#### 标签格式化表达式

`|label_format`可以重命名、修改或者新增标签。

例如:

```shell
# 修改src标签为dst标签
|label_format dst=src
# 使用template变量值替换后，重命名为query
| label_format query="{{ Replace .query \"\\n\" \"\" -1 }}"
```

## Metric Queries

根据日志内容进行指标计算。

### 范围矢量聚合

* `rate(log-range)`: 计算每秒日志实例数

* `count_over_time(log-range)`: 给定范围内的日志流计数

* `bytes_rate(log-range)`: 计算每个日志流每秒字节数

* `bytes_over_time(log-range)`: 计算每个日志流使用的字节数

* `absent_over_time(log-range)`: 如果传递给它的范围矢量具有任何元素，则返回空向量;如果传递给它的范围矢量没有元素，则返回值为 1 个元素为1的向量。

 **示例：**

```shell
count_over_time({job="mysql"}[5m])
```

```shell
sum by (host) (rate({job="mysql"} |= "error" != "timeout" | json | duration > 10s [1m]))
```

### unwrapped 范围聚合

unwrapped范围聚合使用提取的标签而不是日志行作为样例值。

- `duration_seconds(label_identifier)`：以秒为单位转换标签值。
- `bytes(label_identifier)`将标签值转换为字节。
- `rate(unwrapped-range)`：计算指定间隔内所有值的秒速率。
- `sum_over_time(unwrapped-range)`：指定间隔内所有值的总和。
- `avg_over_time(unwrapped-range)`：指定间隔内所有点的平均值。
- `max_over_time(unwrapped-range)`：指定间隔内所有点的最大值。
- `min_over_time(unwrapped-range)`：指定间隔内所有点的最小值
- `stdvar_over_time(unwrapped-range)`：指定间隔内值的总体标准方差。
- `stddev_over_time(unwrapped-range)`：指定间隔内值的总体标准偏差。
- `quantile_over_time(scalar,unwrapped-range)`：指定间隔内φ分位数（0≤ φ ≤1）的值。
- `absent_over_time(unwrapped-range)`：如果传递给它的范围矢量具有任何元素，则返回空向量;如果传递给它的范围矢量没有元素，则返回值为 1个值为 1 的向量。

除了 `sum_over_time`,`absent_over_time` and `rate`外，其他支持分组聚合：

```logql
<aggr-op>([parameter,] <unwrapped-range>) [without|by (<label list>)]
```

**aggr-op：**

- `sum`： 计算标签的总和
- `min`： 选择最小值而不是标签
- `max`： 选择标签上的最大值
- `avg`： 计算标签上的平均值
- `stddev`： 计算与标签的总体标准偏差
- `stdvar`： 计算标签上的总体标准方差
- `count`：计算矢量中的元素数
- `bottomk`： 按样本值选择最小的 k 元素
- `topk`： 按样本值选择最大 k 元素

示例：

```shell
quantile_over_time(0.99,
  {cluster="ops-tools1",container="ingress-nginx"}
    | json
    | __error__ = ""
    | unwrap request_time [1m])) by (path)
```

```shell
sum by (org_id) (
  sum_over_time(
  {cluster="ops-tools1",container="loki-dev"}
      |= "metrics.go"
      | logfmt
      | unwrap bytes_processed [1m])
  )
```

```shell
topk(10,sum(rate({region="us-east1"}[5m])) by (name))
```

```shell
sum(count_over_time({job="mysql"}[5m])) by (level)
```

```shell
avg(rate(({job="nginx"} |= "GET" | json | path="/home")[10s])) by (region)
```

```shell
sum(rate({app="foo"})) * 2
```

```shell
sum(rate({app="foo", level="warn"}[1m])) / sum(rate({app="foo", level="error"}[1m]))
```



# HTTP API

Loki各个组件暴露了不同的http端点。在微服务模式下所有组件都公开了如下端点：

* `GET /ready`

* `GET /metrics`

* `GET /config`

  支持参数`mode`， 取值为`diff`或`defaults`

## querier

- `GET /loki/api/v1/query`

  单时间点查询，支持如下四个查询参数：

  * `query`: logQL查询语句, 如'query={job="varlogs"}'
  * `limit`：返回的最大数量
  * `time`: 查询的Unix纳秒时间，默认当前
  * `direction`：日志顺序，支持`forward`、`backward`(默认)

- `GET /loki/api/v1/query_range`

   时间范围查询，支持如下查询参数：

  * `query`: logQL语句
  * `limit`: 返回的最大条数
  * `start`: Unix纳秒时间，默认一个小时前
  * `end`: Unix纳秒时间，默认当前
  * `step`: duration格式([0-9]\[smhdwy])或浮点秒数。仅应用于返回matrix的查询或 `metric`查询
  * `interval`: duration格式([0-9]\[smhdwy])或浮点秒数，返回>=指定间隔的数据。仅用于返回stream的查询. （实验性的，未来可能移除）
  * `direction`:日志顺序，支持`forward`、`backward`(默认)

- `GET /loki/api/v1/labels`

  查询labels列表，支持如下参数：

  * `start`: Unix纳秒时间，默认6个小时前
  * `end`: Unix纳秒时间，默认当前

- `GET /loki/api/v1/label/<name>/values`

  查询名为\<name>的label值列表，支持如下参数：

  * `start`: Unix纳秒时间，默认6个小时前
  * `end`: Unix纳秒时间，默认当前

- `GET /loki/api/v1/tail`

  查询日志流，是一个websocket接口。支持如下参数：

  * `query`: LogQL
  * `delay_for`：延迟获取日志流的秒数，默认为0，不能大于5
  * `limit`: 最大返回条数
  * `start`: Unix纳秒时间，默认1个小时前

- `GET /loki/api/v1/series`

   查询满足标签集的时间序列列表，支持如下参数：

   * `match`: 日志流选择器，如`'match={container_name=~"prometheus.*", component="server"}'、'match={app="loki"}'`，至少需要提供一个match
   * `start`: Unix纳秒时间
   * `end`: Unix纳秒时间

- `POST /loki/api/v1/series`

   通过`Content-Type: application/x-www-form-urlencoded`的方式发送POST请求， 结果同GET请求。

## distributor

- `POST /loki/api/v1/push`

  日志推送至Loki的端点，`POST body`是`snappy-compressed protobuf`消息，也可以是`json`。如果是`json`，请求头需设置`Content-Type：application/json`，body需为如下格式:

  ```json
  {
    "streams": [
      {
        "stream": {
          "label": "value"
        },
        "values": [
            [ "<unix epoch in nanoseconds>", "<log line>" ],
            [ "<unix epoch in nanoseconds>", "<log line>" ]
        ]
      }
    ]
  }
  ```

## ingester

* `POST /flush`

  刷新所有的内存chunks至后端存储，主要用于本地测试。

* `POST /ingester/flush_shutdown`

  关闭ingester，触发ingester交接，即总是会刷新内存chunks至后端存储。

## ruler

* `GET /ruler/ring`

  展示hash环的状态，是一个web页面。

* `GET /loki/api/v1/rules`

  列出所有的rules。默认是不启用的，可通过命令行参数`-experimental.ruler.enable-api`或yaml配置选项开启。

* `GET /loki/api/v1/rules/{namespace}`

  查询namespace下的所有rules。启用同上。

* `GET /loki/api/v1/rules/{namespace}/{groupName}`

  查询指定namespace、groupname下的rules组。启用同上。

* `POST /loki/api/v1/rules/{namespace}`

  启用同上。创建或更新rule组。header中需设置`Content-Type: application/yaml`, body为rules YAML定义，成功返回202. 

  如：

  ```yaml
  # Request headers:
   Content-Type: application/yaml
  
  # Request body:
  name: <string>
  interval: <duration;optional>
  rules:
    - alert: <string>
      expr: <string>
      for: <duration>
      annotations:
        <annotation_name>: <string>
      labels:
        <label_name>: <string>
  ```

* `DELETE /loki/api/v1/rules/{namespace}/{groupName}`

  删除指定namespace、groupname下的rules组。成功返回202。启用同上。

* `DELETE /loki/api/v1/rules/{namespace}`

  删除namespace下的所有rules。成功返回202。启用同上。

* `GET /prometheus/api/v1/rules`

  兼容Prometheus的端点，列出当前加载的 alerting 和recording rules。启用同上。

* `GET /prometheus/api/v1/alerts`

  兼容Prometheus的端点，列出当前活跃的 alerting。启用同上。

