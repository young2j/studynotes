# Promtail

# 简介

promtail是日志代理，提供了日志转换、服务发现、日志指标等功能，主要负责将处理后的日志内容发送给Loki实例。通常部署到需要提取日志的应用程序所在的机器上。

目前promtail支持两种日志源：本地日志文件(`log files`)和系统日志(`systemd journal`)。

Promtal借用了[Prometheus服务发现机制](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config)，能自动发现日志文件。但此功能目前只支持`kubernetes`服务，Promtail能从`kubernetes API server`中提取所需要的`labels`。

Promtail不仅可以作为服务端来收集发送日志，同时也可以通过`loki_push_api`配置来作为一个客户端，接收来自其他Promtail或者Loki client发来的日志。

# 配置

Promtail的配置文件通常位于`/etc/promtail/config.yml`中，里面包含Promtail服务信息`server`、读取的位置文件`positions.yaml`、日志推送给Loki 实例的`url`以及具体的日志抓取配置`scrape_configs`。

> 详细的配置可查看官方文档[promtail配置](https://grafana.com/docs/loki/latest/clients/promtail/configuration/)，这里只列出通常用得上的，或者大多时候需要自行定义的内容。

## 运行时打印Promtail配置

在启动命令行后添加如下参数(flags)之一，可以让promtail在运行时打印出当前自定义的及默认的配置信息：

* `-print-config-stderr`
* `-log-config-reverse-order`
* `-print-config-stderr=true`

## 指定配置文件

在命令行后通过`-config.file`指定要加载的配置文件，如：

```shell
-config.file=/etc/promtail/config.yml
```

## 在配置文件中使用环境变量

格式：`{VAR}`或`${VAR:default_value}`

 其中，`VAR`是环境变量名，大小写敏感；`default_value`为默认值。

## server

promtail作为http服务的配置:

```yaml
server:
	http_listen_address: 127.0.0.1
  http_listen_port: 9080
  
  grpc_listen_address: 127.0.0.1
  grpc_listen_port: 0
```

## client

配置promtail如何连接到Loki实例：

```yaml
clients:
  - url: http://loki:3100/loki/api/v1/push
    basic_auth:
	  	username: <string>
  		password: <string>
    	password_file: <filename>

  	bearer_token: <secret>
  	bearer_token_file: <filename>
  	proxy_url: <string>

  	tls_config:
    	ca_file: <string>
    	cert_file: <filename>
    	key_file: <filename>
    	server_name: <string>
    	insecure_skip_verify: <boolean> | default = false
  
  	external_labels:
    	<labelname1>: <labelvalue1>
    	<labelname2>: <labelvalue2>
```

## positions

位置文件记录了promtail读取日志内容的位置信息，或者叫偏移`offset`信息，当promtail重启时，会从该文件记录的位置开始继续读取日志内容。

```yaml
# positions.yaml
positions:
  /data/coconut_api/log/archive.log: "0"
  /data/coconut_api/log/beat.log: "0"
  /data/coconut_api/log/beat.log.1: "0"
  /data/coconut_api/log/deamon_task.log: "0"
  /data/coconut_api/log/face_rec.log: "0"
  /data/coconut_api/log/in_order.log: "0"
  /data/coconut_api/log/out_order.log: "0"
```

## scrape_configs

抓取配置，这部分配置如何从目标位置抓取、转换日志。是自定义配置的主要内容。

```yaml
scrape_configs:
- job_name: <string> # 抓取的作业名，通常用于对应一个项目或应用
  pipeline_stages: # 日志动态转换流程配置，功能相对复杂强大，简单的应用日志或许都用不上此配置
  journal: # 抓取来自promtail自身的系统日志
  syslog: # 抓取syslog系统日志
  loki_push_api: # 定义如何接收其他promtail或者docker日志驱动提供的日志
  relabel_configs: # 重新定义标签配置
  static_configs: # 抓取的静态目标配置，能用它解决的，就尽量不用pipeline_stages
  file_sd_configs: # 基于文件的服务发现配置
  kubernetes_sd_configs:  # k8s服务发现配置
- job_name: <string>
	...
```

### pipeline_stages

pipeline_stages由不同的阶段组成。日志内容通过不同的阶段处理后，提取的数据将转换为临时的map结构，然后可以被后续的阶段所使用，也可以被`labels`和`output`所使用。

```yaml
 pipeline_stages: 
    - docker:
    - cri: 

    - regex:
      # 使用正则表达式命名组捕获，从日志中提取内容
      expression: "^(?s)(?P<time>\\S+?) (?P<stream>stdout|stderr) (?P<flags>\\S+?) (?P<content>.*)$"
      # 需要解析的日志内容，默认为空，解析整个日志内容
      source: content
    
    - json:
      # 将日志内容解析为json格式, 值为json日志内容中的key
      stream: stream
      timestamp: time
      output: log
      source: <string>
    
    - template:
      # 使用go的text/template来解析日志
      source: level
      template: '{{ if eq .Value "WARN" }}{{ Replace .Value "WARN" "OK" -1 }}{{ else }}{{ .Value }}{{ end }}'

    - match:
      # logQL stream 选择器
      selector: <string>
    
    - timestamp:
      source: <string>
      # [ANSIC UnixDate RubyDate RFC822 RFC822Z RFC850 RFC1123 RFC1123Z RFC3339 RFC3339Nano Unix UnixMs UnixUs UnixNs]
      format: <string>
      # Timezone Database string.
      location: "Asia/Shanghai"

    - output:
      # output定义经过前述stage处理后，最终由Loki存储的日志内容
      source: <string>

    - labels:
      # 从解析的日志内容中提取标签
      <string>: <string>

    - metrics:
      # metics定义一些统计指标，这些指标不会发往Loki，可以通过promtail的http端点/metrics查看
      <string>: [ <counter> | <gauge> | <histogram> |...]

    - tenant:
      # 设置租户ID，二者只能取其一
      source: <string>
      value: <string>
```

**pipeline_stages封装**

为了避免重复编写一些常用的日志转换配置，针对docker和cri官方提供了配置的封装。

* **docker**

  ```shell
  {"log":"level=info ts=2019-04-30T02:12:41.844179Z caller=filetargetmanager.go:180 msg=\"Adding target\"\n","stream":"stderr","time":"2019-04-30T02:12:41.8443515Z"}
  ```

  ```yaml
  - json:
      output: log
      stream: stream
      timestamp: time
  - labels:
      stream:
  - timestamp:
      source: timestamp
      format: RFC3339Nano
  - output:
      source: output
  ```

* **CRI**

  ```shell
  2019-01-01T01:00:00.000000001Z stderr P some log message
  ```

  ```yaml
  - regex:
      expression: "^(?s)(?P<time>\\S+?) (?P<stream>stdout|stderr) (?P<flags>\\S+?) (?P<content>.*)$",
  - labels:
      stream:
  - timestamp:
      source: time
      format: RFC3339Nano
  - output:
      source: content
  ```

###  journal

```yaml
  # 定义如何抓取来自promtail自身的系统日志
  journal: <journal_config>
    json: <boolean> | default = false
    labels:
      <labelname>: <labelvalue>
    # 默认/var/log/journal 和 /run/log/journal
    path: <string>
```

### syslog

```yaml
  # 定义如何抓取syslog
  syslog: <syslog_config>
    listen_address: <string>
    labels:
```

### loki_push_api

```yaml
 # 定义如何接收其他promtail或者docker日志驱动提供的日志
  loki_push_api: <loki_push_api_config>
    # 配置通server
    server: <server_config>
    labels:
    # false将使用当前的时间戳
    use_incoming_timestamp: <bool> | default = false
```

### relabel_configs

> [传送门](https://grafana.com/docs/loki/latest/clients/promtail/configuration/#relabel_configs)

```yaml
 # 重新定义标签配置
  relabel_configs:
    - [<relabel_config>]
```

### static_configs

```yaml
 # 抓取的静态目标配置
  static_configs:
    - targets: # 可选配置。老版本的promtail需要进行该配置。Prometheus的服务发现的入口，要么不配置，配置则只能是localhost
      - localhost
    labels: # 日志标签，标签{env="dev" job="job_name"}会添加到每一行日志上，好的命名方式通常为：应用名(或者job名)+环境名
      job: job_name
      env: dev
      __path__: /var/log/*log # 日志存储的路径，Promtail将会从这个路径去收集日志
```

### file_sd_configs

可以将日志抓取的配置定义在文件中，用于动态的服务发现，例如可以以列表的形式在`json`文件中定义`static_configs`:

```json
  [
    {
      "targets": [ "localhost" ],
      "labels": {
        "__path__": "<string>", ...
        "<labelname>": "<labelvalue>", ...
      }
    },
    ...
  ]
```

然后在promtail的配置文件中，通过files指定服务配置文件：

```yaml
 # 基于文件的服务发现配置
  file_sd_configs:
    - files:
      - my/path/*.json
      - my/path/*.yml
      - my/path/*.yaml
```

###  kubernetes_sd_configs

> [传送门](https://grafana.com/docs/loki/latest/clients/promtail/configuration/#kubernetes_sd_config)

```yaml
  # 定义如何发现k8s服务
  kubernetes_sd_configs:
    - [<kubernetes_sd_config>]
```

## target_config

```yaml
target_config:
# 控制当前发现的服务或正在读取文件的行为
  sync_period: "10s" # Period to resync directories
```

## 完整配置

```yaml
server:
	http_listen_address: 127.0.0.1
  http_listen_port: 9080
  
  grpc_listen_address: 127.0.0.1
  grpc_listen_port: 0
  
clients:
  - url: http://loki:3100/loki/api/v1/push
    basic_auth:
	  	username: <string>
  		password: <string>
    	password_file: <filename>

  	bearer_token: <secret>
  	bearer_token_file: <filename>
  	proxy_url: <string>

  	tls_config:
    	ca_file: <string>
    	cert_file: <filename>
    	key_file: <filename>
    	server_name: <string>
    	insecure_skip_verify: <boolean> | default = false
  
  	external_labels:
    	<labelname1>: <labelvalue1>
    	<labelname2>: <labelvalue2>

positions:
  filename: /tmp/positions.yaml
  sync_period: <duration> | default = 10s
  
scrape_configs:
- job_name: <string>
  # 定义日志如何转换
  pipeline_stages: 
    - docker: {}
    - cri: {}

    - regex:
      # 使用正则表达式命名组捕获，从日志中提取内容
      expression: "^(?s)(?P<time>\\S+?) (?P<stream>stdout|stderr) (?P<flags>\\S+?) (?P<content>.*)$"
      # 需要解析的日志内容，默认为空，解析整个日志内容
      source: content
    
    - json:
      # 将日志内容解析为json格式, 值为日志内容中的key
      stream: stream
      timestamp: time
      output: log
      source: <string>
    
    - template:
      # 使用go的text/template来解析日志
      source: level
      template: '{{ if eq .Value "WARN" }}{{ Replace .Value "WARN" "OK" -1 }}{{ else }}{{ .Value }}{{ end }}'

    - match:
      # logQL stream 选择器
      selector: <string>
    
    - timestamp:
      source: <string>
      # [ANSIC UnixDate RubyDate RFC822 RFC822Z RFC850 RFC1123 RFC1123Z RFC3339 RFC3339Nano Unix UnixMs UnixUs UnixNs]
      format: <string>
      # Timezone Database string.
      location: "Asia/Shanghai"

    - output:
      # output定义经过前述stage处理后，最终由Loki存储的日志内容
      source: <string>

    - labels:
      # 从解析的日志内容中提取标签
      <string>: <string>

    - metrics:
      # metics定义一些统计指标，这些指标不会发往Loki，可以通过promtail的http端点/metrics查看
      <string>: [ <counter> | <gauge> | <histogram> |...]

    - tenant:
      # 设置租户ID，二者只能取其一
      source: <string>
      value: <string>
  
  # 定义如何抓取来自promtail自身的系统日志
  journal: <journal_config>
    json: <boolean> | default = false
    labels:
      <labelname>: <labelvalue>
    # 默认/var/log/journal 和 /run/log/journal
    path: <string>

  # 定义如何抓取syslog
  syslog: <syslog_config>
    listen_address: <string>
    labels:

  # 定义如何接收其他promtail或者docker日志驱动提供的日志
  loki_push_api: <loki_push_api_config>
    # 配置通server
    server: <server_config>
    labels:
    # false将使用当前的时间戳
    use_incoming_timestamp: <bool> | default = false

  # 重新定义标签配置
  relabel_configs:
    - [<relabel_config>]

  # 抓取的静态目标配置
  static_configs:
    - targets: # 可选配置。老版本的promtail需要进行该配置。Prometheus的服务发现的入口，要么不配置，配置则只能是localhost
      - localhost
    labels: # 日志标签，标签{env="dev" job="job_name"}会添加到每一行日志上，好的命名方式通常为：应用名(或者job名)+环境名
      job: job_name
      env: dev
      __path__: /var/log/*log # 日志存储的路径，Promtail将会从这个路径去收集日志

  # 基于文件的服务发现配置, json文件中必须是static_configs列表，如：
  # [
  #   {
  #     "targets": [ "localhost" ],
  #     "labels": {
  #       "__path__": "<string>", ...
  #       "<labelname>": "<labelvalue>", ...
  #     }
  #   },
  #   ...
  # ]
  file_sd_configs:
    - files:
      - my/path/*.json
      - my/path/*.yml
      - my/path/*.yaml

  # 定义如何发现k8s服务
  kubernetes_sd_configs:
    - [<kubernetes_sd_config>]

target_config:
# 控制当前发现的服务或正在读取文件的行为
  sync_period: "10s" # Period to resync directories
```

# 示例

## Docker Config

docker驱动将在每个日志行中添加如下默认标签:

* `filename`: log文件地址
* `host`: 生成log文件的主机名
* `container_name`:  生成log文件的容器名

想要Promtail收集到docker容器中的日志，需要先安装docker插件。

### 安装Loki Docker插件

```shell
# 安装：
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

# 查看
docker plugin ls

# 升级
docker plugin disable loki --force
docker plugin upgrade loki grafana/loki-docker-driver:latest --grant-all-permissions
docker plugin enable loki
service docker restart

# 卸载
docker plugin disable loki --force
docker plugin rm loki
```

### 容器运行时配置

例如运行grafana时，将grafana的日志按400的batch大小发往Loki，发送请求重试5次：

```shell
docker run --log-driver=loki \
    --log-opt loki-url="http://ip_or_hostname_where_Loki_run:3100/loki/api/v1/push" \
    --log-opt loki-retries=5 \
    --log-opt loki-batch-size=400 \
    grafana/grafana
```

### daemon.json中配置

```json
{
    "debug" : true,
    "log-driver": "loki",
    "log-opts": {
        "loki-url": "http://ip_or_hostname_where_Loki_run:3100/loki/api/v1/push",
        "loki-batch-size": "400"
    }
}
```

> 注意：值需为字符串，需要加 " "。

### docker-compose.yml中配置

```yaml
version: "3"
services:
  nginx:
    image: grafana/grafana
    logging:
      driver: loki
      options:
        loki-url: http://host:3100/loki/api/v1/push
        loki-pipeline-stages: |
          - regex:
              expression: '(level|lvl|severity)=(?P<level>\w+)'
          - labels:
              level:
        loki-relabel-config: |
          - action: labelmap
            regex: swarm_stack
            replacement: namespace
          - action: labelmap
            regex: swarm_(service)
    ports:
      - "3000:3000"
```

> 具体可配置的options：[传送门](https://grafana.com/docs/loki/latest/clients/docker-driver/configuration/#supported-log-opt-options)

## Static Config

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /var/log/positions.yaml 

client:
  url: http://ip_or_hostname_where_Loki_run:3100/loki/api/v1/push

scrape_configs:
 - job_name: system
   pipeline_stages:
   static_configs:
   - targets:
      - localhost
     labels:
      job: varlogs 
      host: yourhost 
      # 路径匹配格式 https://github.com/bmatcuk/doublestar
      __path__: /var/log/*.log 
```

## Static Config without targets

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /var/log/positions.yaml

client:
  url: http://ip_or_hostname_where_Loki_run:3100/loki/api/v1/push

scrape_configs:
 - job_name: system
   pipeline_stages:
   static_configs:
   - labels:
      job: varlogs  
      host: yourhost 
      __path__: /var/log/*.log
```

## Journal Config

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://ip_or_hostname_where_loki_runns:3100/loki/api/v1/push

scrape_configs:
  - job_name: journal
    journal:
      max_age: 12h
      labels:
        job: systemd-journal
    relabel_configs:
      - source_labels: ['__journal__systemd_unit']
        target_label: 'unit'
```

## Syslog Config

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki_addr:3100/loki/api/v1/push

scrape_configs:
  - job_name: syslog
    syslog:
      listen_address: 0.0.0.0:1514
      labels:
        job: "syslog"
    relabel_configs:
      - source_labels: ['__syslog_message_hostname']
        target_label: 'host'
```

## Push Config

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0 # 0表示随机分配端口

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://ip_or_hostname_where_Loki_run:3100/loki/api/v1/push

scrape_configs:
- job_name: push1
  loki_push_api:
    server:
      http_listen_port: 3500
      grpc_listen_port: 3600
    labels:
      pushserver: push1
```



