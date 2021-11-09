# `mysql`集群简介及对比

# 背景

最近负责的项目严重依赖`mysql`,遇到了不少性能问题. 前期通过代码优化、索引优化、读写分离等手段, 项目尚能苟延残喘. 但随着数据量与日俱增, 单纯的优化已经不能从本质上解决性能问题, 需要从数据库架构上做出改变. 于是对mysql集群进行了初步调研, 简单介绍和对比下六种集群方案.

# 集群简介

## innodb副本集(InnoDB ReplicaSet)

> 详细信息可查看 https://dev.mysql.com/doc/mysql-shell/8.0/en/mysql-innodb-replicaset.html

innodb 副本集的拓扑图与单主模式的innodb集群(见下文)非常相似,  由单个主节点(`Primary Node`, `Read and Write`)和多个从节点(`Secondary Node`,`Read Only`)组成. 应用程序可通过`MySQL Router`中间件来透明地访问 `InnoDB ReplicaSet`.

 `InnoDB ReplicaSet`主要具有如下特点:

1. 单主多从模式.
2. 所有成员(从节点)都从一个源(主节点)进行异步复制，不需要对事务达成共识.
3. 可以简单地将现有的副本纳入副本集管理, 提供了读取横向扩展的能力.

InnoDB ReplicaSet 主要具有如下限制:

1. 不具备高可用性. 由于`InnoDB ReplicaSet` 不支持多主模式, 在发生故障时需要手动进行故障转移到新的主服务器。
2. 不提供数据的一致性. 在复制延迟情况下, 副本服务器的更新可能远落后于主服务器; 在意外停止情况下, 也可能造成部分数据丢失

## innodb集群(InnoDB Cluster)

> 详细信息可查看 https://dev.mysql.com/doc/mysql-shell/8.0/en/mysql-innodb-cluster.html

下图为`InnoDB Cluster `单主模式的拓扑图, 由单个主节点(`Primary Node`, `Read and Write`)和多个从节点(`Secondary Node`,`Read Only`)组成. 与`InnoDB ReplicaSet`一样, 客户端应用程序同样通过`MySQL Router` 透明地连接到服务器实例, 但不同的是InnoDB Cluster具有内置的故障转移机制, 如果主节点出现故障，从节点会自动选举为主节点, `MySQL Router` 检测到这一点后, 会将客户端应用程序自动转发到新的主节点。

![](https://dev.mysql.com/doc/mysql-shell/8.0/en/images/innodb_cluster_overview.png)

`InnoDB Cluster`主要具有如下特点:

1. 集群可以配置为单主多从、多主多从模式.
2. 具备高可用性. 集群中的每个服务器实例都运行着组复制(`MySQL Group Replication`)，服务器属于同一组时，它们会通过选举机制自动协调, 自动进行故障转移.
3. 具有一致性. 组复制要求组中的大多数成员必须就全局事务状态及事务顺序达成一致, 在任何给定时间点保持组视图一致并可供所有服务器使用. 例如当有成员意外离开组时, 故障检测机制会检测到这一点并通知组视图更新。

`InnoDB Cluster`主要具有如下限制:

1. `InnoDB Cluster` 不管理手动配置的异步复制通道。组复制不能确保异步复制仅在主服务器上处于活动状态，并且异步复制的状态不会跨实例复制, 这可能导致复制不再起作用。
2. `InnoDB Cluster`依赖稳定且低延迟的网络,旨在部署在局域网中。在广域网上部署单个 InnoDB 集群对写入性能有显著影响。稳定且低延迟的网络对于成员服务器使用底层 `Group Replication` 技术相互通信以达成事务共识非常重要。

## innodb集群副本集(InnoDB ClusterSet)

> 详细信息可查看https://dev.mysql.com/doc/mysql-shell/8.0/en/innodb-clusterset.html

`MySQL InnoDB ClusterSet` 由主集群和副本集群构成, 主要应用于不同数据中心集群数据的复制, 为`InnoDB Cluster` 部署提供容灾能力。如果主集群由于数据中心丢失或网络连接丢失而变得不可用，可以改为激活副本集群以恢复服务的可用性。例如下图示例, 主集群及其副本集群均由三个成员服务器实例组成，一个主服务器实例和两个从服务器实例。其中, `MySQL Router`仍然可以自动将客户端应用程序路由到 `InnoDB ClusterSet` 中的正确集群。

![](https://dev.mysql.com/doc/mysql-shell/8.0/en/images/innodb_clusterset_main.png)

`InnoDB ClusterSet`主要具有如下特点:

1. 主集群和副本集群之间的紧急故障转移需要由管理员主动进行受控切换.
2. 只有主集群的主服务器接受来自客户端应用程序的写入流量. 
3. 副本集群不接受写入, 但可以被应用程序读取，异步复制的延迟可能导致数据还不完整。
4. 可以拥有的副本集群数量没有定义限制。
5. 异步复制通道将事务从主集群复制到副本集群。
6. 底层的`Group Replication`技术确保复制始终在主集群的主服务器（作为发送方）和副本集群的主服务器（作为接收方）之间进行。
7. `InnoDB ClusterSet` 优先考虑可用性而不是数据一致性，以最大限度地提高容灾能力。
8. 每个单独的 `InnoDB Cluster` 内的一致性由底层 `Group Replication` 技术保证。但复制延迟或网络分区可能使得主集群遇到问题时，部分或全部副本集群与主集群不完全一致。如果触发紧急故障转移，任何未复制的事务都有丢失的风险.

`InnoDB ClusterSet`主要具有如下限制:

1. `InnoDB Cluster` 技术以及`Group Replication` 技术的限制.
2. `InnoDB ClusterSet `不会自动故障转移到副本集群, 无法保证在发生紧急故障转移时会保留数据。
3. `InnoDB ClusterSet` 只支持异步复制，不能使用半同步复制。
4. `InnoDB ClusterSet` 仅支持`Cluster`实例的单主模式,不支持多主模式。 即只能包含一个读写主集群, 所有副本集群都是只读的, 不允许具有多个主集群的双活设置，因为在集群发生故障时无法保证数据一致性。
5. 已有的 `InnoDB Cluster` 不能用作 `InnoDB ClusterSet` 部署中的副本集群。副本集群必须从单个服务器实例启动，作为新的 `InnoDB` 集群。
6. `InnoDB ClusterSet` 不支持使用运行 MySQL Server 5.7 的实例。包含 MySQL 5.7 实例的 `InnoDB Cluster` 不能成为`InnoDB ClusterSet`部署的一部分。

## ndb集群(NDB cluster)

> 详细信息可查看 https://dev.mysql.com/doc/refman/8.0/en/mysql-cluster-overview.html

`NDB Cluster`将标准 MySQL 服务器与称为 NDB（`Network DataBase`）的内存集群存储引擎集成在一起, 是一种在无共享系统中实现内存数据库集群的技术, 提供了高可用性和数据持久性功能。

> 在无共享系统中，每个组件都应该有自己的内存和磁盘，不推荐或不支持使用网络共享、网络文件系统和 SAN 等共享存储机制。无共享架构使系统能够使用非常便宜的硬件，并且对硬件或软件的特定要求最低。

`NDB Cluster`就是指一个或多个 MySQL 服务器与 `NDB` 存储引擎的组合。集群一般包括 MySQL 服务器节点（`sql node 或 API node`, 用于访问 NDB 数据）、数据节点（`data node`, 用于存储数据）、一个或多个管理服务器节点(`mgm node`, 用于集群管理)。

![](https://dev.mysql.com/doc/refman/8.0/en/images/cluster-components-1.png)

`NDB Cluster`具有如下特点:

1. SQL 节点可以直接访问集群中的数据. 当应用程序更新数据后, 所有MySQL 服务器都可以立即看到此更改。

2. SQL 节点使用的 `mysqld` 服务器守护程序，在许多关键方面与 MySQL 8.0 发行版提供的 mysqld 二进制文件不同，两个版本的 `mysqld` 不可互换。

3. 集群可以处理单个数据节点的故障.

4. 单个节点可以停止和重新启动，然后可以重新加入集群。

5. `NDB Cluster` 表通常完全存储在内存中而不是磁盘上。但一些 `NDB Cluster` 数据可以存储在磁盘上.

6. 表和表数据存储在数据节点中, 数据节点中存储的数据可以进行镜像。

7. 一个或多个节点构成节点组，节点存储数据分区(`Partition`), 一个分区保存至少一个片段副本(`Fragment Replica`). 

   ```mysql
   节点组的数量 = 节点数 / 片段副本数
   [# of node groups] = [# of data nodes] / NoOfReplicas
   分区数 = 节点数 / LDM线程数
   [# of partitions] = [# of data nodes] * [# of LDM threads]
   ```

   四个节点、两个片段副本、单线程(ndb)的存储结构如下:

   ![](https://dev.mysql.com/doc/refman/8.0/en/images/fragment-replicas-groups-1-1.png)

8. NDB8.0 支持外键、支持`JSON data type`.

`NDB Cluster`主要具有如下限制:

1. `sql`语法不兼容
   * 不支持临时表
   * 列索引的最大长度不超过3072字节
   * 不支持全文索引
   * 不支持前缀索引, 只能索引整个列
   * 如果外键不是(关联)表的主键，则作为外键引用的列都需要是显式的唯一键。
   * 当外键引用是父表的主键时，不支持级联更新`ON UPDATE CASCADE`
   * 支持地理数据, 但不支持空间索引
   * ...
2. 数据节点的最大数量为 145 
3. 集群最大节点总数为 255。这个数字包括所有 SQL 节点（MySQL 服务器）、API 节点、数据节点和管理服务器
4. NDB存储引擎仅支持 `READ COMMITTED` 事务隔离级别
5. `NDB Cluster` 不能很好地处理大型事务；与包含大量操作的单个大事务相比，执行多个小事务且每个事务包含少量操作要更好。大型事务需要非常大量的内存。
6. `NDB Cluster` 中所有 NDB 数据库对象的最大数量（包括数据库、表和索引）限制为 20320
7. 每个表的属性（即列和索引）数最大为 512。每个键的最大属性数为 32。
8. 行数据大小最大为 30000 字节
9. 不支持半同步复制
10. SQL节点没有分布式表锁。` LOCK TABLES` 语句或`GET_LOCK() `调用仅适用于发出锁的 SQL 节点；集群中没有其他 SQL 节点“看到”这个锁

## mariadb集群(MariaDB Galera Cluster)

> 详细信息可查看 https://mariadb.com/kb/en/what-is-mariadb-galera-cluster/#galera-versions

`MariaDB Galera Cluster` 是一个几近同步的多主集群。 它仅在`Linux`上可用，并且仅支持 `InnoDB` 存储引擎.

![](https://mariadb.com/kb/en/about-mariadb-galera-cluster/+image/galera_small)

`MariaDB Galera Cluster`主要具有如下特点:

1. 几近同步地复制(`Virtually synchronous replication`)
2. 双活多主拓扑(`Active-active multi-primary topology`)
3. 任意集群节点均可读写(`Read and write to any cluster node`)
4. 自动成员控制，故障节点自动从集群中删除(`Automatic membership control, failed nodes drop from the cluster`)
5. 自动节点加入(`Automatic node joining`)
6. 真正行级并行复制(`True parallel replication, on row level`)
7. 客户端直接连接，MariaDB原生体验 (`Direct client connections, native MariaDB look & feel`)

上述特点使得集群无复制滞后、无事务丢失、读取可扩展、客户端延迟更小.

`MariaDB Galera Cluster`主要具有如下限制:

1. 目前仅适用于 `InnoDB` 存储引擎。
2. 不支持的显式锁定, 包括`LOCK TABLES, FLUSH TABLES {explicit table list} WITH READ LOCK, (GET_LOCK(), RELEASE_LOCK(),…)`. 但支持全局锁定操作，如`FLUSH TABLES WITH READ LOCK`。
3. 所有表都应该有一个主键（支持多列主键）。没有主键的表不支持 DELETE 操作。此外，表中没有主键的行在不同节点上可能以不同的顺序出现。
4. 一般查询日志和慢查询日志不能直接指向一个表。如果启用这些日志，则必须通过设置 `log_output=FILE` 将日志转发到文件。
5. 不支持 XA 事务。
6. 事务规模. 极大的事务（例如 `LOAD DATA`）可能会对节点性能产生不利影响。为避免这种情况，默认情况下 `wsrep_max_ws_rows` 和 `wsrep_max_ws_size` 系统变量将事务行数限制为 128K，将事务大小限制为 2Gb。

## PXC集群(Percona XtraDB Cluster)

> 详细信息可查看https://www.percona.com/doc/percona-xtradb-cluster/LATEST/index.html

从下图可以看到,`Percona XtraDB Cluster`与`MariaDB Galera Cluster`具有几乎相同的特点, 具体如下:

1. 没有中央管理。每个节点之间同步复制相同的数据集, 某一节点不可用时, 集群将继续运行而不会丢失任何数据
2. 每个节点均可进行读写, 当执行查询时，会直接在本地执行。 
3. 通过加入节点横向扩展读取能力。

![](https://www.percona.com/doc/percona-xtradb-cluster/LATEST/_images/cluster-diagram1.png)

`Percona XtraDB Cluster`主要具有如下限制:

1. 复制仅适用于 `InnoDB` 存储引擎。
2. 不支持如`LOCK TABLES` 和 `UNLOCK TABLES`等锁表操作, 也不支持`GET_LOCK()、RELEASE_LOCK()`等锁函数.
3. 所有表都应该有一个主键。没有主键的表不支持 DELETE 操作。
4. 查询日志不能直接存入表, 如果启用查询日志, 必须要通过设置`log_output = FILE`将其存入文件.
5. 不支持 XA 事务。
6.  `wsrep_max_ws_rows` 和 `wsrep_max_ws_size` 变量定义了最大事务规模。`LOAD DATA INFILE` 操作将会每10 000 行提交一次。因此，由于 `LOAD DATA` 导致的大事务将被拆分为一系列小事务。
7. 整个集群的写入吞吐量受到最弱节点的限制。如果一个节点变慢，整个集群就会变慢。
8. 不具有有效的写入扩展解决方案.  当添加新节点时，必须从现有节点之一复制完整数据集。

# 简要对比

1. InnoDB副本集( `InnoDB ReplicaSet`)仅提供了数据冗余, 提升了数据读取能力, 但不具备高可用性, 也不能保证数据的一致性; 
2. 相较于`InnoDB ReplicaSet`, 官方推荐尽可能部署InnoDB集群(`InnoDB Cluster`), 可部署为一主多从或者多主多从结构, 组复制(`Group Replication`)技术保障了各节点的高可用性和一致性.
3. InnoDB集群副本集(`MySQL InnoDB ClusterSet`)进一步提供了集群的容灾能力, 这种"集群的集群"主要用于大型企业不同数据中心的数据备份.
4. NDB集群(`NDB Cluster`)不同于InnoDB集群、Mariadb集群以及PXC集群, 他们每个节点都进行了数据的完整复制, 有多少个节点就有多少份数据. NDB集群在这方面更具灵活性, 可以自定义是否冗余,冗余多少. 分区存储的概念也与mongdb的分片集比较相似. 相对于mysql其他“复制方案”, NDB集群最有集群的样子.
5. MariaDB集群(`MariaDB Galera Cluster`)和PXC集群(`Percona XtraDB Cluster`)结构几乎没有什么差别, 二者均无主节点或者都为主节点, 每个节点数据完全一致, 都具有读写能力, 可以简单的通过增加节点扩展读取能力, 但写入能力无法扩展.

综上所述, 现有的mysql集群方案都重在“复制”上, 当我们仅仅只是想要提升mysql的性能而寻求集群方案时, 不得不把数据存个好几份, 对于大量时效性较短的数据而言, 用后即删, 根本没必要存储多份, 反而显得有些浪费资源. 这种情况下, 可能不进行冗余备份的NDB集群才是最佳选择?

