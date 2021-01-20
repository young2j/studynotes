# Label

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

标签是 Loki 日志数据的索引。它们用于查找压缩日志内容，这些内容单独存储为块。标签和值的每个唯一组合都定义了流，流的日志被批处理、压缩和存储为块。

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

