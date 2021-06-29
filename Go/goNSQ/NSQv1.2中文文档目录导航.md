#! https://zhuanlan.zhihu.com/p/352740512
# NSQv1.2.0中文文档

`nsq`是go语言写的消息队列库，目前最新稳定版为`1.2.0`。在学习nsq的过程中，顺便将nsq的官方文档翻译为了中文，下面内容是文档目录，共五个部分(除去最后一个**链接**部分，实质只有四个)，点击可跳转至相应的翻译内容。由于本人英语水平很有限，翻译有误的地方，<del>本人概不负责</del> 请以官方原文为准，英文水平好的也请尽量阅读原文: [https://nsq.io/overview/quick_start.html](https://nsq.io/overview/quick_start.html)

# 内容导航

**[概览](https://juejin.cn/post/6932864138339155982/)**

* [快速开始](https://juejin.cn/post/6932864138339155982/)
* [特征与保证](https://juejin.cn/post/6932864138339155982/)
* [常见问题](https://juejin.cn/post/6932864138339155982/)
* [性能](https://juejin.cn/post/6932864138339155982/)
* [设计](https://juejin.cn/post/6932864138339155982/)
* [内部](https://juejin.cn/post/6932864138339155982/)

**[组件](https://juejin.cn/post/6932863986308546567/)**

* [nsqd](https://juejin.cn/post/6932863986308546567/)
* [nsqlookupd](https://juejin.cn/post/6932863986308546567/)
* [nsqadmin](https://juejin.cn/post/6932863986308546567/)
* [utilities](https://juejin.cn/post/6932863986308546567/)

**[客户端](https://juejin.cn/post/6932865148784902158/)**
* [客户端库](https://juejin.cn/post/6932865148784902158/)
* [构建客户端库](https://juejin.cn/post/6932865148784902158/)
* [TCP协议规范](https://juejin.cn/post/6932865148784902158/)

**[部署](https://juejin.cn/post/6932866637519552525/)**
* [安装](https://juejin.cn/post/6932866637519552525/)
* [生产配置](https://juejin.cn/post/6932866637519552525/)
* [拓扑模式](https://juejin.cn/post/6932866637519552525/)
* [docker](https://juejin.cn/post/6932866637519552525/)

**[链接]()**

[TALKS AND SLIDES]()

- 2012-11-08 - **NYC Golang Meetup** [slides](https://speakerdeck.com/snakes/nsq-nyc-golang-meetup) ([@imsnakes](https://twitter.com/imsnakes) & [@jehiah](https://twitter.com/jehiah) of Bitly)
- 2013-02-13 - **PHP UK 2013, Planning To Fail** [slides](https://speakerdeck.com/davegardnerisme/planning-to-fail) ([@davegardnerisme](https://twitter.com/davegardnerisme) of Hailo)
- 2013-04-03 - **PhillyETE 2013, Stream Processing: Philosophy, Concepts, and Technologies** [video](https://www.infoq.com/presentations/data-streaming-nsq) [slides](https://speakerdeck.com/danielhfrank/stream-processing-philosophy-concepts-and-technologies) ([@danielhfrank](https://twitter.com/danielhfrank) of Bitly)
- 2013-05-15 - **NYC Data Engineering Meetup** [slides](https://speakerdeck.com/snakes/nsq-nyc-data-engineering-meetup) [video](https://www.youtube.com/watch?v=IkU8JsxdCAM) ([@imsnakes](https://twitter.com/imsnakes) & [@jehiah](https://twitter.com/jehiah) of Bitly)
- 2013-06-17 - **C\* Summit 2013, Big Architectures for Big Data** [slides](https://www.slideshare.net/planetcassandra/2-eric-lubow) [video](https://www.youtube.com/watch?v=dT0A0bh_CLw) ([@elubow](https://twitter.com/elubow) of simplereach)
- 2013-06-19 - **SF Golang Meetup** [video](https://plus.google.com/u/0/events/ckpnkggt52aoc7vagkctqsjg6v8) ([@imsnakes](https://twitter.com/imsnakes) of Bitly)
- 2014-01-11 - **Data Days Texas 2014** [slides](https://eric.lubow.org/presentations/data-day-texas-2014/) ([@elubow](https://twitter.com/elubow) of SimpleReach)
- 2014-04-24 - **GopherCon 2014** [video](https://confreaks.com/videos/3429-gophercon2014-spray-some-nsq-on-it) [slides](https://speakerdeck.com/snakes/spray-some-nsq-on-it) ([@imsnakes](https://twitter.com/imsnakes) of Torando Labs)
- 2015-06-02 - **Berlin Buzzwords 2015** [video](https://www.youtube.com/watch?v=OwD-W7uU2zU) [slides](https://georgi.io/scale-with-nsq) ([@GeorgiCodes](https://twitter.com/GeorgiCodes) of Bitly)
- 2015-07-30 - **Munich Node.js User Group** [video](https://www.youtube.com/watch?v=xhNapGc6SsU) ([@juliangruber](https://twitter.com/juliangruber) of NowSecure)
- 2018-06-12 - Building a distributed message processing system in Go using NSQ [slides](https://bit.ly/nsqslides) ([@GBrayUT][https://twitter.com/GBrayUT] of Walmart Labs)

[Release Notes](https://github.com/nsqio/nsq/releases)

[Github Repo](https://github.com/nsqio/nsq)

[Issues](https://github.com/nsqio/nsq/issues)

[Google Group](https://groups.google.com/group/nsq-users)