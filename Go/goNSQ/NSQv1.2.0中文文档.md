<p style='text-align:center;font-size:30px;font-weight:bold'>NSQ v1.2.0中文文档</p>

# 概览

## 快速开始

下述步骤将在你的本地机器上运行一个小型**NSQ**集群，贯穿了消息的发布、消费以及归档至磁盘。

1. 首先跟随[安装说明](https://nsq.io/deployment/installing.html)文档进行安装；

2. 在shell中开启`nsqlookupd`:

   ```shell
   $ nsqlookupd
   ```

3. 在另一个shell中开启`nsqd`:

   ```shell
   $ nsqd --lookupd-tcp-address=127.0.0.1:4160
   ```

4. 在另一个shell中开启`nsqadmin`：

   ```shell
   $ nsqadmin --lookupd-http-address=127.0.0.1:4161
   ```

5. 发布一个初始消息（也就是在集群中创建一个主题）

   ```shell
   $ curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test'
   ```

6. 最后，在另一个shell中开启`nsq_to_file`:

   ```shell
   $ nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
   ```

7. 发布更多消息到`nsqd`:

   ```shell
   $ curl -d 'hello world 2' 'http://127.0.0.1:4151/pub?topic=test'
   $ curl -d 'hello world 3' 'http://127.0.0.1:4151/pub?topic=test'
   ```

8. 为了验证上述工作是否如预期进行，可以打开浏览器地址http://127.0.0.1:4171/ 通过`nsqadmin`UI界面查看统计信息。当然，你也可以检查写到`/tmp`目录下的日志文件(`test.*.log`)内容。

这里有个比较重要的内容是，客户端并没有明确地指明`test`主题从哪里产生，`nsq_to_file`将从`nsqlookupd`提取这些信息，即使正处于连接中，也不会有消息丢失。

## 特征与保证

`NSQ`是一个实时分布式消息平台。

### 特征

* 支持无SPOF的分布式拓扑
* 水平扩展(无代理，可无缝地向集群中添加更多节点)
* 消息传递低延迟推送（[性能](https://nsq.io/overview/performance.html)）
* 负载均衡以及多种风格消息路由组合
* 擅长流式处理（高吞吐量）和面向作业（低吞吐量）的工作负载
* 主要在内存中（超出某个水平，消息将以透明的方式保存在磁盘上）
* 运行时服务发现，对于消费者查找生产者[(nsqlookupd)](https://github.com/nsqio/nsq/tree/master/nsqlookupd/README.md)
* 传输层安全 （TLS）
* 数据格式不可知
* 依赖少，易于部署，具有合理、清晰有界的默认配置
* 简单 TCP 协议支持任何语言的客户端库
* 用于统计信息、管理操作和生产者（无需发布客户端库）的 HTTP 接口
* 与实时检测的[statsd](https://github.com/etsy/statsd/)集成
* 强大的群集管理接口 （[nsqadmin](https://github.com/nsqio/nsq/tree/master/nsqadmin/README.md))

### 保证

> 这部分建议只看看加粗部分就可以了，其他部分不是啥重点，水平有限翻译过来也很拗口，或者直接去看原文。

与任何分布式系统一样，实现你的目标都需要一个明智的权衡过程。通过透明地权衡这些折衷的现实，我们希望对NSQ在生产中部署时的行为设定如下期望：

* **消息不是持久化的(默认)**
  尽管系统支持通过`--mem-queue-size`设置"释放阀(release valve)",消息将被透明地保存在磁盘上，但NSQ主要还是一个基于内存的消息平台。

  **`--mem-queue-size`可以被设置为0以确保所有到来的消息被持久化到磁盘**。在这种情况下，如果节点发生故障，那么您很容易感知到面临的故障会减少（例如OS或底层IO子系统是否发生故障）

  并**没有内置的复制集。**但是，管理这种权衡的方法有很多种，例如部署拓扑和技术以容错的方式主动持久化主题到磁盘。

* **消息至少传递一次**

  与上述密切相关，假定给定的`nsqd`节点不会失败。

  这意味着，**由于各种原因（客户端超时、断开连接、重新入队等），消息可以多次传递**。而执行幂等操作或删除重复信息是客户端的职责。

* **收到的消息是无序的**

  **不能依赖传递给消费者的消息顺序。**

  与消息传递的字面意义类似，消息是重新排列队列的结果，是内存和磁盘存储的组合。事实上每一个`nsqd`节点之间并不共享任何东西。

  通过在你的消费者中引入一个延迟窗口来接收消息并且在处理这些消息之前（尽管为了保持这些不变的消息，必须丢弃掉位于该窗口之外的消息）进行排序，以此来实宽松的排序(即对于给定的消费者，它的消息是有序的，但是在整个集群中却不能保证)是相对直截了当的。

* **消费者最终可以找到所有主题生产者**

  发现服务[(nsqlookupd)](https://github.com/nsqio/nsq/tree/master/nsqlookupd/README.md)被设计为最终一致。 `nsqlookupd`节点不会共同维护状态或应答查询。

  网络分区不会影响可用性，从这层意义上来说，分区的两侧仍然可以应答查询。部署拓扑对缓解这些类型的问题具有最重要的作用。

## 常见问题

### 部署

- **`nsqd`的推荐拓扑是什么？**

  我们强烈建议**在生成消息的任何服务旁边运行一个`nsqd`**

  `nsqd`是一个相对轻量级的进程，可限定内存占用，这使得它非常适合"与别人好好玩耍"（*“playing nice with others”*.）。

  这种模式有助于将消息流构建为消费问题，而不是生产问题。

  另一个好处是，这种模式在给定主机上为主题形成了一个独立的、分片的数据孤岛。

  注意：这不是一个绝对的要求，它只是更为简单（见下面的问题）。

- **为什么生产者不能使用`nsqlookupd`来查找发布到哪儿？**

  因为必须告诉消费者在哪里找到他们需要的主题，NSQ提倡**消费者端发现**模型，以减轻前期配置负担。

  然而，对于服务应发布到哪里，这并没有提供任何解决问题的手段。这是鸡和蛋的问题，主题在发布之前并不存在。

  通过对`nsqd`的共同定位（请参阅上面的问题），你完全回避了此问题（你的服务只是发布到本地`nsqd`）， NSQ 的运行时发现系统自然地进行工作。

- **我只想在单个节点上使用 `nsqd`作为工作队列， 这是一个合适的用例吗？**

  是的， `nsqd`在单节点上也可以运行得很好。

  `nsqlookupd`更有利于在较大的分布式环境中使用。

- **我应该运行多少 `nsqlookupd`?**

  通常只有几个，具体取决于集群大小、`nsqd`节点数和消费者数量以及你所需的容错性。

  对于几百台主机和数千个消费者，部署3个或5个就可以很好地工作。

  `nsqlookupd`节点**不需要**协作应答查询。集群中的元数据最终是*一致的*。

### 发布

- **我需要客户端库来发布消息吗？**

  不！只需使用 HTTP端点来发布（ `/pub`和`/mpub` ）。它很简单，很容易，而且几乎在任何编程环境中都无处不在。

  事实上，绝大多数 NSQ 部署都使用 HTTP 来发布。

- **为什么要强制客户端处理TCP 协议的 `PUB` 和`MPUB`命令的响应？**

  我们认为 NSQ 的默认操作模式应优先考虑安全性，我们希望协议简单且一致。

- **`PUB` 或`MPUB`什么时候可能失败？**

  1. 主题名称的格式不正确（字符/长度限制）。请参阅[主题和频道名称规格](https://nsq.io/clients/tcp_protocol_spec.html#notes)。
  2. 消息太大（此限制作为参数被暴露给`nsqd` ）。
  3. 主题正在被删除。
  4. `nsqd`正在清理退出。
  5. 发布期间与客户端连接相关的任何失败时。

  1和 2 应视为编程错误。3和4是不常见的。5是任何基于 TCP 协议自然有的。

- **如何缓解上述第3个问题？**

  删除主题是一个相对不频繁的操作。如果需要删除主题，请协调时间，删除要经过足够的时间，以便发布所引出的主题创建永远不会执行。

### 设计与理论

- **如何推荐命名主题和通道？**

  主题名称应描述流中的数据。

  通道名称应描述其消费者所执行的工作。

  例如，好的主题名称可以为 `encodes`，`decodes`，`api_requests`，`page_views`；好的通道名称为 `archive`，`analytics_increment`，`spam_analysis`。

- **单个`nsqd`可以支持的主题和通道的数量是否有任何限制？**

  没有施加内置限制。它仅受限于`nsqd`所运行主机的内存和 CPU （每个客户端的 CPU 使用率极大地减少[#236](https://github.com/nsqio/nsq/pull/236)）。

- **如何向群集宣布新主题？**

  第一次`PUB`或者`SUB`一个主题将在`nsqd`上创建主题。主题元数据接下来将传播到配置的`nsqlookupd` 。其他订阅者将定期查询`nsqlookupd` 来发现此主题。

- **NSQ 能做 Rpc 吗？**

  这是可能的，但是在设计NSQ时并未考虑到该用例。

  我们打算发布一些文档，以了解如何构建该内容，但在此期间，如果您有兴趣，请伸出援手。

### pynsq 特定的问题

- **你为什么强迫我使用Tornado？**

  `pynsq`最初打算作为以消费者为导向的库，在 Python 的异步框架（特别是由于 NSQ 面向推送的协议）中使用 NSQ 协议要简单得多。

  Tornado的 API 很简单，性能相当好。

- **Tornado的 IOLoop需要发布吗？**

  不需要，`nsqd`暴露了 HTTP 端点(`/PUB`和`/MPUB`)以非常简单的用于编程语言不可知的发布（agnostic publishing）。

  如果你担心 HTTP 的开销，那没有必要。此外，`/mpub`通过批量发布（原子发布！）来减少 HTTP 的开销。

- **我什么时候要使用`Writer`?**

  当高性能和低开销是一个优先事项。

  `Writer`使用TCP协议`PUB`和`MPUB`命令，与其他HTTP同行相比拥有更小的开销。

- **如果我只想 "fire and forget" （我可以容忍消息丢失！**

  使用 `Writer`并且对发布方法不指定回调。

  注意：这只有利于产生更简单的`pynsq`客户端代码，在幕后仍要处理来自`nsqd`的响应（就是说这样做没有性能什么优势）。

特别感谢Dustin Oprea ([@DustinOprea](https://twitter.com/DustinOprea))启动这个常见问题

## 性能

主仓库有一个脚本(`bench/bench.py`)，它会自动执行一个EC2上的分布式基准测试。它会引导N个节点，一些节点正在运行`nsqd`，一些节点在负载生成PUB和SUB应用，然后解析这些节点的输出以供总体聚合。

### 安装

下面运行的命令反应了6个`c3.2xlarge`的默认参数，最便宜的实例类型支持1gbit的链接。3个节点分别运行了一个`nsqd`实例，其余节点运行着`bench_reader(SUB)`实例和`bench_writer(PUB)`实例，以此生成依赖于基准测试模式的负载。

```shell
$ ./bench/bench.py --access-key=... --secret-key=... --ssh-key-name=...
[I 140917 10:58:10 bench:102] launching 6 instances
[I 140917 10:58:12 bench:111] waiting for instances to launch...
...
[I 140917 10:58:37 bench:130] (1) bootstrapping ec2-54-160-145-64.compute-1.amazonaws.com (i-0a018ce1)
[I 140917 10:59:37 bench:130] (2) bootstrapping ec2-54-90-195-149.compute-1.amazonaws.com (i-0f018ce4)
[I 140917 11:00:00 bench:130] (3) bootstrapping ec2-23-22-236-55.compute-1.amazonaws.com (i-0e018ce5)
[I 140917 11:00:41 bench:130] (4) bootstrapping ec2-23-23-40-113.compute-1.amazonaws.com (i-0d018ce6)
[I 140917 11:01:10 bench:130] (5) bootstrapping ec2-54-226-180-44.compute-1.amazonaws.com (i-0c018ce7)
[I 140917 11:01:43 bench:130] (6) bootstrapping ec2-54-90-83-223.compute-1.amazonaws.com (i-10018cfb)
```

### 生产者吞吐量

此基准测试仅测量了生产者的吞吐量，没有额外的负载。消息的大小是100字节，并且消息分布于3个主题中。

```shell
$ ./bench/bench.py --access-key=... --secret-key=... --ssh-key-name=... --mode=pub --msg-size=100 run
[I 140917 12:39:37 bench:140] launching nsqd on 3 host(s)
[I 140917 12:39:41 bench:163] launching 9 producer(s) on 3 host(s)
...
[I 140917 12:40:20 bench:248] [bench_writer] 10.002s - 197.463mb/s - 2070549.631ops/s - 4.830us/op
```

进入速度(`ingress`)约为`2.07mm` msgs/sec，消耗了总计197mb/s的带宽。

### 生产者和消费者吞吐量

此基准通过为生产者和消费者提供服务，更准确地反映了真实情况。同样，消息大小为 100 字节，消息分布在 3 个主题中，每个主题具有单个*通道*（每个通道 24 个客户端）。

```
$ ./bench/bench.py --access-key=... --secret-key=... --ssh-key-name=... --msg-size=100 run
[I 140917 12:41:11 bench:140] launching nsqd on 3 host(s)
[I 140917 12:41:15 bench:163] launching 9 producer(s) on 3 host(s)
[I 140917 12:41:22 bench:186] launching 9 consumer(s) on 3 host(s)
...
[I 140917 12:41:55 bench:248] [bench_reader] 10.252s - 76.946mb/s - 806838.610ops/s - 12.706us/op
[I 140917 12:41:55 bench:248] [bench_writer] 10.030s - 80.315mb/s - 842149.615ops/s - 11.910us/op
```

在大约`842k `和 `806k` msgs/s 的入口和出口时，消耗了 总计156mb/s 的带宽，我们现在在`nsqd`节点上最大化了 CPU 容量。通过引入消费者，`nsqd`需要维护每个通道的消息传递，因此负载自然会更高。

消费者的数量略低于生产者，因为消费者发送的命令数量是生产者的两倍（必须为每条消息发送一个`FIN`命令），从而影响了吞吐量。

再添加 2 个节点（一个`nsqd`和一个负载生成(load-generating)）达到超过 `1mm` msgs/s:

```shell
$ ./bench/bench.py --access-key=... --secret-key=... --ssh-key-name=... --msg-size=100 run
[I 140917 13:38:28 bench:140] launching nsqd on 4 host(s)
[I 140917 13:38:32 bench:163] launching 16 producer(s) on 4 host(s)
[I 140917 13:38:43 bench:186] launching 16 consumer(s) on 4 host(s)
...
[I 140917 13:39:12 bench:248] [bench_reader] 10.561s - 100.956mb/s - 1058624.012ops/s - 9.976us/op
[I 140917 13:39:12 bench:248] [bench_writer] 10.023s - 105.898mb/s - 1110408.953ops/s - 9.026us/op
```

### 单节点性能

免责声明：请记住**，NSQ**旨在以分布式方式使用。单节点性能虽然很重要，但不是我们所要实现的一切。此外， 基准测试是愚蠢的， 但这里多少做个展示：

- 2012 MacBook Air i7 2ghz

- go1.2

- NSQ v0.2.24

- 200 byte message

**GOMAXPROCS=1（1个发布者、1个消费者）**

```shell
$ ./bench.sh 
results...
PUB: 2014/01/12 22:09:08 duration: 2.311925588s - 82.500mb/s - 432539.873ops/s - 2.312us/op
SUB: 2014/01/12 22:09:19 duration: 6.009749983s - 31.738mb/s - 166396.273ops/s - 6.010us/op
```

**GOMAXPROCS=4（4个发布者，4个消费者)**

```shell
$ ./bench.sh 
results...
PUB: 2014/01/13 16:58:05 duration: 1.411492441s - 135.130mb/s - 708469.965ops/s - 1.411us/op
SUB: 2014/01/13 16:58:16 duration: 5.251380583s - 36.321mb/s - 190426.114ops/s - 5.251us/op
```

## 设计

注：有关随附的视觉插图，请参阅此[幻灯片组](https://speakerdeck.com/snakes/nsq-nyc-golang-meetup)。

> 不翻墙你是打不开的。

**NSQ**是[simplequeue](https://github.com/bitly/simplehttp/tree/master/simplequeue) ([simplehttp](https://github.com/bitly/simplehttp)的一部分)的后继者，因此设计为（没有特定顺序）：

- 支持高可用且消除了 SPOFs 的拓扑
- 满足了更强有力的保证消息传递的需要
- 限制了单个进程的内存占用（通过将某些消息持久化到磁盘）
- 极大地简化了生产者和消费者的配置要求
- 提供了简单直接的升级路径
- 提高了效率

### 简化配置和管理

单个`nsqd`实例设计为一次处理多个数据流。数据流称为"主题"，主题有 1 个或多个"通道"（channels）。每个通道接收主题的所有消息副本。实际上，一个通道对应着一个下游消费主题的服务。

主题和通道并未优先配置。通过发布到命名主题或订阅有关命名主题的通道，主题在首次使用时创建。而通过订阅命名通道，通道在首次使用时创建。

主题和通道所有缓冲区数据彼此独立，以防止缓慢的消费导致其他通道的积压（这同样适用于主题级别）。

一个通道通常可以连接多个客户端。假设所有连接的客户端处于准备接收消息的状态，则每条消息都将传递到一个随机的客户端。例如：

![nsqd clients](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

总之，消息来自主题 -> 通道（每个通道接收该主题的所有消息副本），但从通道 -> 消费者，消息是均匀分布的（每个消费者接收该通道的部分消息）。

**NSQ**还包括一个辅助应用程序`nsqlookupd` ，它提供了一个目录服务，消费者可以在其中查找到提供给他们订阅主题的`nsqd`实例地址。在配置方面，这使消费者与生产者分离（他们只需要知道在哪里联系`nsqlookupd`通用实例，从不是彼此之间直接联系），降低了复杂性和维护成本。

在底层，每个`nsqd`与`nsqlookupd`之间都具有一个 TCP 长连接(long-lived)，并定期推送其状态给`nsqlookupd`。此数据用于`nsqlookupd`将哪些`nsqd`地址通知给消费者。对于消费者，将暴露 HTTP 端点`/lookup`以进行轮询。

要引入新的主题消费者，只需启动一个配置了`nsqlookupd`实例地址的**NSQ**客户端。添加新的消费者或新的发布者无需更改配置，从而大大降低了开销和复杂性。

注意：在将来的版本中，启发式`nsqlookupd` 可能基于深度、连接的客户端数量或其他"智能"策略来返回`nsqd`地址。当前的实现就是全部。归根结底，目标是深度保持在接近于零的水平，确保所有生产者都能够被订阅。

需要注意的重要点是，`nsqd`和 `nsqlookupd`守护进程被设计为独立运行，同类进程之间没有沟通与协作。

我们也认为，通过一种方法来查看、思考和整体管理集群非常重要。我们为了做到这一点而构建了`nsqadmin`。它提供了一个 Web UI 来浏览主题/通道/消费者的层次结构，并检查每个层的深度和其他关键统计信息。此外，它支持一些管理命令，如删除和清空通道（当通道中的消息可以安全地抛出以将深度带回 0 时，这是一个有用的工具）。

![nsqadmin](https://media.tumblr.com/tumblr_mbmsd6YMfS1qj3yp2.png)

### 简单直接的升级路径

这是我们的最高优先事项之一。我们的生产系统处理着大量的流量，全部都基于我们现有的消息工具，因此我们需要一种缓慢且有条理的方法来升级基础架构的特定部分，带来的影响微乎其微。

首先，在消息生产者方面，我们构建了`nsqd`来匹配[simplequeue](https://github.com/bitly/simplehttp/tree/master/simplequeue)。具体地说，`nsqd` 暴露了 HTTP `/pub`端点，就像`simplequeue`一样发送二进制数据（需要注意的是，端点需要一个额外的查询参数来指定"主题"）。想要切换服务，向`nsqd`发布消息，只需进行少量的代码更改。

其次，我们构建了与现有库功能和习惯相匹配的Python 和 Go 库。通过将代码更改限制为自举（bootstrap），从而缓解了在消费者端的过渡。所有业务逻辑保持不变。

最后，我们构建了将新旧组件粘合在一起的实用程序。这些都可在存储库中的`examples`目录中获得：

- `nsq_to_file`- 将给定主题的所有消息持久化写入文件
- `nsq_to_http`- 对主题中的所有消息向(多个)端点执行 HTTP 请求

### 消除 SPOFs

**NSQ**专为分布式使用而设计。`nsqd` 客户端（通过 TCP）连接到所有实例，并提供特定的主题消息。没有中间人，没有消息中间件，也没有SPOFs：

![nsq clients](https://media.tumblr.com/tumblr_mat85kr5td1qj3yp2.png)

此拓扑消除了对单个的、聚合的源进行链接的必要。相反，直接从**所有生产者**处进行消费。从技术上讲，哪个客户端连接到哪个**NSQ**并不重要，只要有足够的客户端连接到所有的生产者来满足消息量，就能保证消息最终得到处理。

对于 `nsqdlookupd`，通过运行多个实例实现了高可用。它们之间不直接相互通信，并且数据最终被认为是一致的。消费者会轮询其配置的所有`nsqlookupd`实例并联合(union)响应结果。陈旧过时、无法访问或其他节点故障都不会使系统停止运行。

### 消息传递保证

**NSQ**保证消息将至少**传递一次，**尽管消息可能重复。消费者应该对此有所意料，并进行重复消息删除或执行幂等操作。

此保证强制作为协议的一部分，其工作方式如下（假设客户端已成功连接并订阅了主题）：

1. 客户端表示他们已准备好接收消息
2. **NSQ**发送消息并临时将数据存储在本地（在重新入队或超时时）
3. 客户端答复 FIN（完成finish）或 REQ（重新入队re-queue），分别表示成功或失败。如果**NSQ**超时(可配置)未收到客户端的答复，消息也会自动重新入队）

这可确保导致消息丢失的唯一边缘情况是`nsqd`进程未正常关闭(unclean shutdown)。在这种情况下，内存中的任何消息（或任何未刷新到磁盘的缓冲消息）都会丢失。

如果防止消息丢失至关重要，则即使此边缘情况也可以缓和。一种解决方案是支持接收相同消息副本的冗余`nsqd`对（在单独的主机上）。由于已将消费者写为幂等的，因此在这些消息上重复执行两次操作不会对下游产生影响，系统可以承受任何单点故障而不丢失消息。

重要的是**NSQ**提供了构建基础（building blocks），以支持各种生产用例和持久化深度的配置。

### 可限定内存占用

`nsqd`提供了一个配置选项`--mem-queue-size`，用以确定给定的队列在内存中保留消息的数量。如果队列的深度超过此阈值，则消息将透明地写入磁盘。这使得`nsqd`进程的内存占用量限制为：`mem-queue-size * #_of_channels_and_topics`

![message overflow](https://media.tumblr.com/tumblr_mavte17V3t1qj3yp2.png)

此外，精明的人可能已经发现了一个便捷的方式，即通过设置此值为较低的数（比如 1甚至 0），以获得更高的传递保证。磁盘队列旨在承受非正常启动（尽管消息可能传递两次）。

此外，与消息传递保证有关的是，正常关闭(clean shutdowns)（通过向`nsqd`进程发送TERM 信号）可以安全地持久化当前内存中的、正在传递的、延迟的和各种内部缓冲的消息。

请注意，主题或通道的名称以`#ephemeral`结尾的，将不会缓冲到磁盘，相反在超过`mem-queue-size`后会被删除消息。这使的不需要消息保证的消费者也能够订阅频道。这些临时的(ephemeral)通道在最后一个客户端断开连接时也会消失。对于临时主题，这意味着至少有一个通道已创建、使用和删除（通常是临时通道）。

### 效率

**NSQ**设计为使用"像 memcached 一样"的命令协议来进行通信，具有简单的以大小为前缀的响应结果。所有的消息数据都保存在核心中，包括尝试次数、时间戳等元数据。这消除了服务端与客户端之间来来回回的数据复制，这是重新入队消息时上一个工具链的固有属性。这也简化了客户端，因为它们不再需要负责维护消息状态。

此外，通过降低配置复杂性，设置和开发的时间大大缩短（尤其是在主题消费者 >1 的情况下）。

对于数据协议，我们做了一个关键的设计决策，那就是通过将数据推送到客户端而不是等待客户端来提取，以此最大化性能和吞吐量。这个概念，我们称之为`RDY`状态，本质上是客户端流控的一种形式。

当客户端连接到并订阅通道时，它被放置在`RDY` 0 的状态。这意味着不会向客户端发送任何消息。当客户端准备好接收消息时，它会发送一个命令，将状态更新到某个它准备处理(多少条消息)的#状态，例如#100。在没有任何额外命令的情况下，100 条消息将推送到客户端（服务端将会为该客户端进行RDY 计数）。

客户端库的设计是在计数达到配置的`max-in-flight`（适当考虑与多个`nsqd`实例的连接，适当拆分）的约25% 时，发送命令来更新RDY计数。

![nsq protocol](https://media.tumblr.com/tumblr_mataigNDn61qj3yp2.png)

这是一个显著的性能旋钮，因为一些下游系统能够更容易地批量处理消息，并受益于更高的`max-in-flight`。

值得注意的是，因为它既有缓冲又有推送功能，并且能够满足对流（通道）的独立副本需求，因此我们生成了一个类似于`simplequeue`和`pubsub` 组合的守护程序，我们传统上会维护上面讨论过的较旧的工具链，这在简化我们系统的拓扑方面非常强大。

### Go

我们很早就做出了一个战略决策，在Go 中构建NSQ [核心](https://golang.org/)。我们最近写了关于我们使用Go的[博客](https://word.bitly.com/post/29550171827/go-go-gadget)并提到这个项目 - 浏览该帖，了解我们对语言的思考可能会有所帮助。

关于**NSQ**，Go 通道（不要与**NSQ**通道混淆）和语言内置的并发功能非常适合`nsqd`的内部工作。我们利用缓冲通道来管理内存中的消息队列，并无缝地将溢出的消息写入磁盘。

通过标准库，可以轻松地编写网络层和客户端代码。内置内存和 cpu 分析钩子突显了优化机会，并且集成需要很少的精力。我们还发现，在隔离中测试组件、使用接口模拟类型以及迭代构建功能非常容易。

## 内部实现

NSQ 由 3 个守护进程组成：

- **[nsqd](https://nsq.io/components/nsqd.html)**是消息接收、排队和传递消息给客户端的守护进程。
- **[nsqlookupd](https://nsq.io/components/nsqlookupd.html)**是管理拓扑信息并提供最终一致的发现服务的守护进程。
- **[nsqadmin](https://nsq.io/components/nsqadmin.html)**是一个 Web UI，可以实时省察集群（并执行各种管理任务）。

NSQ 中的数据流被模型化为流和消费者的树结构。一个**主题**代表一种数据流。**通道**是订阅特定主题的消费者的逻辑分组。

![topics/channels](https://f.cloud.github.com/assets/187441/1700696/f1434dc8-6029-11e3-8a66-18ca4ea10aca.gif)

单个**nsqd**可以包含多个主题，每个主题可以具有多个通道。通道接收主题的所有消息副本，启用多播样式(multicast style)传递，而通道上的每条消息都在其订阅者之间分发，从而实现负载均衡。

这些基元构成了一个强大的框架，用于表达[各种简单而复杂的拓扑](https://nsq.io/deployment/topology_patterns.html)。

有关 NSQ 设计的信息，请参阅[设计文档](https://nsq.io/overview/design.html)。

### 主题和通道

主题和通道是 NSQ 的核心基元(primitives)，它很好地体现了系统设计如何无缝地转换为 Go 的功能。

Go 的通道（因此称为"go-chan"，用于消除歧义）是表达队列的自然方式，因此，NSQ 主题/通道的核心只是消息结构体`Message`指针的缓冲"go-chan"。缓冲的大小等于配置参数`--mem-queue-size`

在从网络中读取数据后，将消息发布到主题涉及如下行为：

1. `Message`结构体的实例化（以及消息体`[]byte`的分配）
2. 读锁(read-lock)获得`Topic`
3. 读锁(read-lock)检查发布能力
4. 发送给一个缓冲"go-chan"

若要将消息从主题发送到其通道，该主题不能依赖于典型的 go-chan 接收语义，因为在 go-chan 上接收的多个 goroutine将分发消息，而所需的最终结果是将每条消息复制到每个通道（goroutine）。

相反，每个主题维护了 3 个主要 的goroutine：

第一个称为`router` ，负责从传入 go-chan 的消息中读取最新发布的消息，并将其存储在队列（内存或磁盘）中。

第二个称为`messagePump` ，负责复制并推送消息到上述通道。

第三个负责 磁盘队列(`DiskQueue`)IO，稍后将讨论.

通道要复杂一些，但共有一个底层目标：暴露一个单输入和单输出的go-chan (以抽象出在内部消息可能存在于内存或磁盘中的事实)：

![queue goroutine](https://f.cloud.github.com/assets/187441/1698990/682fc358-5f76-11e3-9b05-3d5baba67f13.png)

此外，每个通道维护 2 个按时间排序的优先级队列，负责延迟和在途(in-flight)消息的超时（ 附带2个监视它们的goroutine）。

通过管理每个通道的数据结构（而不是依赖于 Go 运行时的全局计时器调度）并行化得以改进。

**注：**在内部，Go 运行时使用单优先级队列和 goroutine来管理计时器。这支持（但不限于）整个`time`包。它通常不需要时间排序的优先级队列，但重要的是要记住，它是一个单一的数据结构，只有一个锁，可能会影响`GOMAXPROCS > 1`时的性能。请参阅[runtime/time.go](https://github.com/golang/go/blob/release-branch.go1.9/src/runtime/time.go#L92)。（在 go-1.10+中不再是了）

### 后台/磁盘队列

NSQ 的设计目标之一是限制内存中保留的消息数。它通过`DiskQueue`(主题或通道的第三个主要goroutine)透明地将消息溢出写入磁盘。

由于内存队列仅仅是一个go-chan，如果可能的话，它首先会尝试将消息路由至内存，然后回调至磁盘：

```go
for msg := range c.incomingMsgChan {
	select {
	case c.memoryMsgChan <- msg:
	default:
		err := WriteMessageToBackend(&msgBuf, msg, c.backend)
		if err != nil {
			// ... handle errors ...
		}
	}
}
```

利用 Go 的`select`语句，只需几行代码就可以表示此功能：上述`default`语句只有在`memoryMsgChan`满了之后才会执行。

NSQ 也有临时**主题/通道**的概念。它们丢弃了(discard)消息溢出（而不是写入磁盘），并在不再有订阅的客户端时消失。这是 Go 接口的完美用例。主题和通道具有一个声明为接口而不是具体类型的结构体成员`Backend`。普通主题和通道使用的是`DiskQueue`, 而临时主题和通道暂存在 `DummyBackendQueue`中，实现了无操作后端(Backend)。

### 减少GC压力

在任何垃圾回收环境中，都会受限于吞吐量（执行有用的工作）、延迟（响应能力）和驻留集大小（resident set size, 占用空间）。

自 Go 1.2 起，the GC is mark-and-sweep (parallel), non-generational, non-compacting, stop-the-world and mostly precise（翻译不来，自行体会🤪）。它大多是精确的，因为剩余的工作没有及时完成（预定要到Go1.3）。

Go 的GC肯定会继续改进，但普遍的事实是：创造的垃圾越少，收集垃圾的时间就越少。

首先，了解GC在实际工作负载下的表现非常重要。为此，**nsqd**以[statsd](https://github.com/etsy/statsd/)格式（与其他内部指标一起）发布了 GC 统计信息。**nsqadmin**展示了这些指标的图表，让您深入了解 GC 在频率和持续时间方面的影响：

![single node view](https://f.cloud.github.com/assets/187441/1699828/8df666c6-5fc8-11e3-95e6-360b07d3609d.png)

为了真正减少垃圾，您需要知道垃圾的生成地点。Go 工具链提供了答案：

1. 使用[`test`](https://golang.org/pkg/testing/)包和`go test -benchmem`来基准测试热代码路径（to benchmark hot code paths）。它描述了每次迭代的分配数（基准测试可以与[`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp) 进行比较运行。
2. 使用`go build -gcflags -m`编译 ，输出[escape analysis](https://en.wikipedia.org/wiki/Escape_analysis)的结果。

有鉴于此，以下优化证明对**nsqd 很有用**：

1. 避免`[]byte`转换`string`。
2. 重用缓冲区(buffers)或对象(objects)（之后可能会是[`sync.Pool`](https://groups.google.com/forum/#!topic/golang-dev/kJ_R6vYVYHU)也就是[issue 4720](https://code.google.com/p/go/issues/detail?id=4720)）。
3. 预分配切片(在 `make`中指定容量)，并始终知道条目的数量和大小。
4. 对各种可配置的`dials`（例如消息大小）应用合理的限制。
5. 避免装箱`boxing`（使用 `interface{}`） 或不必要的类型包装（如"多个值"的`go-chan`结构体）。
6. 避免在热代码路径（hot code paths）中使用`defer`。

### TCP 协议

[NSQ TCP 协议](https://nsq.io/clients/tcp_protocol_spec.html)是使用 GC 优化概念以产生巨大效果的光辉示例。

该协议用长度前缀帧(length prefixed frames)进行结构化，使得它编码和解码简单直接，性能好：

```
[x][x][x][x][x][x][x][x][x][x][x][x]...
|  (int32) ||  (int32) || (binary)
|  4-byte  ||  4-byte  || N-byte
------------------------------------...
    size      frame ID     data
```

由于帧组件的确切类型和大小是提前知道的，我们可以避免[`encoding/binary`](https://golang.org/pkg/encoding/binary/)包的便利性[`Read（）`](https://golang.org/pkg/encoding/binary/#Read)和[`Write（）`](https://golang.org/pkg/encoding/binary/#Write)封装（及其无关的接口查找和转换），相反直接调用合适的二[`binary.BigEndian`](https://golang.org/pkg/encoding/binary/#ByteOrder)方法。

为了减少套接字 IO 的系统调用，客户端`net.Conn`用[`bufio.Reader和bufio.Writer`进行封装。`Reader`暴露了[`ReadSlice（）`](https://golang.org/pkg/bufio/#Reader.ReadSlice)方法，它重用了其内部缓冲区。这几乎消除了在读取套接字时的分配，大大降低了GC压力。这是可能的，因为与大多数命令关联的数据不会转义（在边缘情况下不是如此，数据会被显式复制）。

在更底层，一个`MessageID`被声明为`[16]byte`，能够使用它作为`map`键（切片不能用作map键）。但是，由于从套接字读取的数据存储为`[]byte` ，而不是通过分配`string`键产生垃圾，并且为了避免`MessageID`从切片到背后数组的复制，使用了`unsafe`包来将切片直接转换为一个`MessageID` 

```go
id := *(*nsq.MessageID)(unsafe.Pointer(&msgID))
```

**注意：**这是一个hack。如果编译器对此进行了优化，就没有必要这样做。并且[打开的issue 3512 ](https://code.google.com/p/go/issues/detail?id=3512)有可能解决此问题。值得一读[issue 5376](https://code.google.com/p/go/issues/detail?id=5376)，其中谈到"const like"`byte`类型可以在接受字符串的地方与其互换使用的可能性，而无需分配和复制。

同样，Go 标准库仅提供了数字转换`string`的方法。为了避免`string`分配，**nsqd**在操作`[]byte`上直接使用了[自定义基10转换方法](https://github.com/nsqio/nsq/blob/v1.2.0/internal/protocol/byte_base10.go#L9-L29)转换方法。

这些看起来像微优化，但TCP协议包含一些最热门的代码路径(hottest code paths)。总的来说，以每秒数万条消息的速度，它们对分配和开销的量有重大影响：

```
benchmark                    old ns/op    new ns/op    delta
BenchmarkProtocolV2Data           3575         1963  -45.09%

benchmark                    old ns/op    new ns/op    delta
BenchmarkProtocolV2Sub256        57964        14568  -74.87%
BenchmarkProtocolV2Sub512        58212        16193  -72.18%
BenchmarkProtocolV2Sub1k         58549        19490  -66.71%
BenchmarkProtocolV2Sub2k         63430        27840  -56.11%

benchmark                   old allocs   new allocs    delta
BenchmarkProtocolV2Sub256           56           39  -30.36%
BenchmarkProtocolV2Sub512           56           39  -30.36%
BenchmarkProtocolV2Sub1k            56           39  -30.36%
BenchmarkProtocolV2Sub2k            58           42  -27.59%
```

### HTTP

NSQ 的 HTTP API 构建在 Go 的[`net/http 包`](https://golang.org/pkg/net/http/)之上。因为它只是*HTTP，*所以几乎在任何现代编程环境中都可以利用它，而无需特殊的客户端库。

它的简单性掩盖了它的力量，因为 Go 的 HTTP 工具箱中最有趣的方面之一是它支持的广泛调试功能。[ `net/http/pprof`](https://golang.org/pkg/net/http/pprof/)包直接与本机 HTTP 服务器集成，公开终结点以检索 CPU、堆、转到例和操作系统线程配置文件。可以直接从该工具定位这些目标：`go`

```
$ go tool pprof http://127.0.0.1:4151/debug/pprof/profile
```

对于调试和分析正在运行的进程来说，这是一*个极其宝贵的*价值！

此外，终结点在 JSON 或打印的漂亮文本中返回大量指标，使管理员能够轻松地从命令行实时进行反省：`/stats`

```
$ watch -n 0.5 'curl -s http://127.0.0.1:4151/stats | grep -v connected'
```

这将生成连续输出，如：

```
[page_views     ] depth: 0     be-depth: 0     msgs: 105525994 e2e%: 6.6s, 6.2s, 6.2s
    [page_view_counter        ] depth: 0     be-depth: 0     inflt: 432  def: 0    re-q: 34684 timeout: 34038 msgs: 105525994 e2e%: 5.1s, 5.1s, 4.6s
    [realtime_score           ] depth: 1828  be-depth: 0     inflt: 1368 def: 0    re-q: 25188 timeout: 11336 msgs: 105525994 e2e%: 9.0s, 9.0s, 7.8s
    [variants_writer          ] depth: 0     be-depth: 0     inflt: 592  def: 0    re-q: 37068 timeout: 37068 msgs: 105525994 e2e%: 8.2s, 8.2s, 8.2s

[poll_requests  ] depth: 0     be-depth: 0     msgs: 11485060 e2e%: 167.5ms, 167.5ms, 138.1ms
    [social_data_collector    ] depth: 0     be-depth: 0     inflt: 2    def: 3    re-q: 7568  timeout: 402   msgs: 11485060 e2e%: 186.6ms, 186.6ms, 138.1ms

[social_data    ] depth: 0     be-depth: 0     msgs: 60145188 e2e%: 199.0s, 199.0s, 199.0s
    [events_writer            ] depth: 0     be-depth: 0     inflt: 226  def: 0    re-q: 32584 timeout: 30542 msgs: 60145188 e2e%: 6.7s, 6.7s, 6.7s
    [social_delta_counter     ] depth: 17328 be-depth: 7327  inflt: 179  def: 1    re-q: 155843 timeout: 11514 msgs: 60145188 e2e%: 234.1s, 234.1s, 231.8s

[time_on_site_ticks] depth: 0     be-depth: 0     msgs: 35717814 e2e%: 0.0ns, 0.0ns, 0.0ns
    [tail821042#ephemeral     ] depth: 0     be-depth: 0     inflt: 0    def: 0    re-q: 0     timeout: 0     msgs: 33909699 e2e%: 0.0ns, 0.0ns, 0.0ns
```

最后，每个新的 Go 版本通常都会带来[可衡量的性能提升](https://github.com/davecheney/autobench)。当重新编译对最新版本的 Go 提供免费提升时， 它总是很好！

## 依赖[锚点链接：依赖项](https://nsq.io/overview/internals.html#dependencies)

来自其他生态系统，Go 关于管理依赖性的理念（或缺乏依赖性）需要一点时间来习惯。

NSQ 从单一的巨型存储库，*具有相对导入*和内部包之间的很少或没有分离，发展到完全采用在结构和依赖关系管理方面建议的最佳做法。

有两个主要的思想学说：

1. **供应商：**在应用程序存储库的正确修订版中复制依赖项，并修改导入路径以引用本地副本。
2. **虚拟 Env**： 列出所需的依赖项的修订，并在生成时生成包含这些固定依赖项的原始环境。`GOPATH`

**注：**这真的只适用于*二进制*包，因为对于可导入的包来说，就使用哪个依赖项的版本做出中间决策是没有意义的。

NSQ 使用上述方法 （2）。（它首先使用[gpm，](https://github.com/pote/gpm)然后[dep](https://github.com/golang/dep)，现在使用[Go 模块](https://github.com/golang/go/wiki/Modules)）。

## 测试[锚点链接： 测试](https://nsq.io/overview/internals.html#testing)

Go 为编写测试和基准提供了可靠的内置支持，并且由于 Go 使建模并发操作变得非常简单，因此在测试环境中构建**nsqd 的完整**实例是微不足道的。

但是，初始实现有一个方面在测试中变得有问题：全局状态。最明显的罪犯是使用全局变量，该变量在运行时引用**nsqd**的实例，即 。`var nsqd *NSQd`

某些测试会通过使用短格式变量赋值（即 ）无意中将此全局变量屏蔽到其本地作用域中。这意味着全局引用没有指向当前正在运行的实例，这违反了测试。`nsqd := NewNSQd(...)`

要解决此问题，将传递一个包含配置元数据和对父**nsqd 的引用的结构**。对全局状态的所有引用都替换为此本地，允许子级（主题、通道、协议处理程序等）安全地访问此数据，并使其测试更加可靠。`Context``Context`

## 鲁棒性[锚点链接：鲁棒性](https://nsq.io/overview/internals.html#robustness)

面对不断变化的网络条件或意外事件，系统无法可靠，它是在分布式生产环境中性能不好的系统。

NSQ 的设计和实施方式允许系统容忍故障，并且以一致、可预测和毫不奇怪的方式运行。

总体理念是快速失败，将错误视为致命错误，并提供调试确实发生的任何问题的手段。

但是，为了*做出反应*，你需要能够检测*出异常*情况...

### 心跳和超时[锚点链接：心跳和超时](https://nsq.io/overview/internals.html#heartbeats-and-timeouts)

NSQ TCP 协议是面向推送的。连接、握手和订阅后，使用者将置于 的状态。当使用者准备好接收消息时，它会将该状态更新为它愿意接受的消息数。NSQ 客户端库在幕后持续管理此内容，从而产生由流控制的消息流。`RDY``0``RDY`

定期**，nsqd**将通过连接发送检测信号。客户端可以配置检测信号之间的间隔，但**nsqd**在发送下一个检测信号之前需要响应。

应用程序级别检测信号和状态的组合[避免了线头阻塞](https://en.wikipedia.org/wiki/Head-of-line_blocking)，否则会使检测信号变得毫无用处（即，如果使用者在处理消息流时落后，操作系统的接收缓冲区将填满，阻塞检测信号）。`RDY`

为了保证进度，所有网络 IO 都与相对于配置的检测信号间隔的截止时间绑定。这意味着您可以从字面上拔下**nsqd**和使用者之间的网络连接，它将检测并正确处理错误。

当检测到致命错误时，客户端连接被强制关闭。飞行中的消息将排时，重新排队，以交付给其他使用者。最后，记录错误并增加各种内部指标。

### 管理 Go 例程[锚点链接：管理 go 例程](https://nsq.io/overview/internals.html#managing-goroutines)

这是令人惊讶的容易*开始*去例程。不幸的是，要协调它们的清理并非易事。避免僵局也是具有挑战性的。通常，这可以归结为一个排序问题，即一个 go 例程接收 go-chan 退出*之前*，上游 go 例程发送它。

为什么在乎呢？很简单，一个孤立的戈例程是一个*内存泄漏*。长时间运行的守护程序中的内存泄漏是不好的，尤其是当预期当其他所有进程都失败时，您的进程将保持稳定。

使事情更加复杂，典型的**nsqd** *进程在消息*传递中涉及许多程序。在内部，消息"所有权"经常发生变化。为了能够干净地关闭，考虑所有进程内消息*非常重要*。

虽然没有任何魔法子弹，但以下技术使得它更容易管理...

#### 等待组[锚点链接：等待组](https://nsq.io/overview/internals.html#waitgroups)

同步[`包`](https://golang.org/pkg/sync/)提供[`同步。WaitGroup`](https://golang.org/pkg/sync/#WaitGroup)， 可用于对实时访问数进行记帐（并提供等待退出时的手段）。

为了减少典型的样板**，nsqd**使用此包装器：

```
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

// can be used as follows:
wg := WaitGroupWrapper{}
wg.Wrap(func() { n.idPump() })
...
wg.Wait()
```

#### 退出信号[锚点链接：退出信令](https://nsq.io/overview/internals.html#exit-signaling)

在多个子 go 例程中触发事件的最简单方法是提供一个准备就绪时关闭的 go-chan。所有挂起的接收将激活 go-chan，而不必向每个 go-chan 发送单独的信号。

```
func work() {
    exitChan := make(chan int)
    go task1(exitChan)
    go task2(exitChan)
    time.Sleep(5 * time.Second)
    close(exitChan)
}
func task1(exitChan chan int) {
    <-exitChan
    log.Printf("task1 exiting")
}

func task2(exitChan chan int) {
    <-exitChan
    log.Printf("task2 exiting")
}
```

#### 同步退出[锚点链接：同步退出](https://nsq.io/overview/internals.html#synchronizing-exit)

实现可靠、无死锁的退出路径是相当困难的，该路径占所有飞行中消息。一些提示：

1. 理想情况下， 负责发送 go - chan 的 go 例程也应该负责关闭它。
2. 如果邮件不能丢失，请确保清空相关的 go-chans（尤其是未缓冲邮件！），以确保发件人能够取得进展。
3. 或者，如果消息不再相关，则发送单个 go-chan 时应转换为 添加退出信号（如上所述），以保证进度。`select`
4. 一般顺序应为：
   1. 停止接受新连接（关闭侦听器）
   2. 信号出口到子程序（见上文）
   3. 等待 go 例程出口（见上文）`WaitGroup`
   4. 恢复缓冲数据
   5. 刷新留给磁盘的任何东西

#### 测 井[锚点链接：日志记录](https://nsq.io/overview/internals.html#logging)

最后，最重要的工具，在您的处置***是记录你的戈例程的入口和出口！\***它使*在死*锁或泄漏的情况下识别罪魁祸首变得无限容易。

**nsqd**日志行包括将 go 例程与其同级（和父级）关联的信息，例如客户端的远程地址或主题/通道名称。

日志是冗长的，但不是冗长的，到日志是压倒性的点。有一条细线，但**nsqd**倾向于在*发生*故障时在日志中包含更多信息，而不是试图以牺牲有用性为代价来减少聊天。

# 组件

## nsqd

## nsqlookupd

## nsqadmin

## utilities

# 客户端

## 客户端库

## 构建客户端库

## TCP协议规格

# 部署

## 安装

## 生产配置

## 拓扑模式

## docker

# 链接

## Talks and Slides

## Release Notes

## Github Repo

## Issues

## Google Group