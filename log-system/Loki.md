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

