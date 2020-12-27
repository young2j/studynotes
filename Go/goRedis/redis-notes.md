

# 前言

本文记录了在`GO`语言中`redis`的基本操作。常用的操作库有`redigo`和`go-redis`两个，目前github上`go-redis`星星数略胜一筹，所以选择后者，记录其基本命令在go中的使用。

# 安装

```SHELL
go get github.com/go-redis/redis/v8
```

# 客户端连接

## standalone连接NewClient

```GO
func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Username: "",
		Password: "",
		// default: localhost:6379
		Addr: "localhost:6381",
		DB: 1,
		PoolSize: 5,
	})
    // ping一下检查是否连通
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	// PONG
	fmt.Println(pingResult)
}
```

## cluster连接NewClusterClient

### 集群模式下连接

```GO
func main() {
	ctx := context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Username: "",
		Password: "",
		Addrs:    []string{":6381", ":6379"},
		PoolSize: 20,
	})
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	// PONG
	fmt.Println(pingResult)
}
```

### 手动创建集群连接

> 在standalone模式下也可以手动创建集群连接

```GO
func main() {
	ctx := context.Background()
	clusterSlots := func(ctx context.Context) ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			// 第一个master:slave 节点
			{
				Start: 0,
				End:   8191,
				Nodes: []redis.ClusterNode{{
					Addr: ":7000", // master
				}, {
					Addr: ":8000", // slave
				}},
			},
			// 第二个master:slave 节点
			{
				Start: 8192,
				End:   16383,
				Nodes: []redis.ClusterNode{{
					Addr: ":7001", // master
				}, {
					Addr: ":8001", // slave
				}},
			},
		}
		return slots, nil
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})
	rdb.Ping(ctx)
}
```

## sentinel连接NewFailoverClient 

> 并发安全的连接

```GO
func main() {
	ctx := context.Background()
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		Username: "",
		Password: "",
		DB: 0,
		MasterName: "master",
		SentinelAddrs: []string{":2378"},
		SentinelPassword: "",
	})
	rdb.Ping(ctx)
}
```

## shard连接NewRing

```GO
func main() {
	ctx := context.Background()
	rdb := redis.NewRing(&redis.RingOptions{
		Username: "",
		Password: "",
		DB:       0,
		PoolSize: 10,
		Addrs: map[string]string{
			"shard1": ":7000",
			"shard2": ":7001",
			"shard3": ":7002",
		},
	})
	rdb.Ping(ctx)
}
```

## 通用连接NewUniversalClient

根据传递的不同`options`对应返回不同的`client`

```go
package main

import (
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb1 := redis.NewUniversalClient(&redis.UniversalOptions{
		// 传入的addrs切片长度大于等于2，将返回一个集群客户端ClusterClient
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	defer rdb1.Close()

	rdb2 := redis.NewUniversalClient(&redis.UniversalOptions{
		// 传递了MasterName参数，将返回一个基于sentinel的FailoverClient
		MasterName: "master",
		Addrs:      []string{":26379"},
	})
	defer rdb2.Close()

	rdb3 := redis.NewUniversalClient(&redis.UniversalOptions{
		// addrs 切片长度为1， 将返回一个普通的单节点客户端NewClient
		Addrs: []string{":6379"},
	})
	defer rdb3.Close()
}
```

## 解析URL进行连接

```GO
package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
	options, err := redis.ParseURL("redis://username:password@localhost:6379/1")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(options)
	defer rdb.Close()

	rdb.Ping(context.Background())
}
```

# 基本操作

## String

```GO
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	statusCmd := rdb.Set(ctx, "language", "golang", 5*time.Second)
	fmt.Println(statusCmd.String())
	// set language golang ex 5: OK
	fmt.Println(statusCmd.Result())
	// OK <nil>

	stringCmd := rdb.Get(ctx, "language")
	fmt.Println(stringCmd.String())
	// get language: golang
	fmt.Println(stringCmd.Result())
	// golang <nil>

	stringCmd2 := rdb.GetSet(ctx, "language", "php")
	fmt.Println(stringCmd2.String())
	// getset language php: golang
	fmt.Println(stringCmd2.Result())
	// golang <nil>


	boolCmd := rdb.SetNX(ctx, "language", "go", 5*time.Second)
	fmt.Println(boolCmd.Result())
	// OK <nil>

	intCmd := rdb.StrLen(ctx,"language")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	intCmd = rdb.Append(ctx, "language", "is the best")
	fmt.Println(intCmd.Result())
	// 14 <nil>
	str, _ := rdb.Get(ctx, "language").Result()
	fmt.Println(len(str))
	// 14

	// statusCmd2 := rdb.MSet(ctx, []interface{}{"php", "world best", "go", 666})
	// statusCmd2 := rdb.MSet(ctx, map[string]interface{}{"php": "world best", "go": 666})
	statusCmd2 := rdb.MSet(ctx, "php", "world best", "go", 666) // 三种方式
	fmt.Println(statusCmd2.Result())
	// OK <nil>
	
	sliceCmd := rdb.MGet(ctx, "php", "go")
	fmt.Println(sliceCmd.Result())
	// [world best 666] <nil>

	intCmd2 := rdb.Incr(ctx,"go")
	fmt.Println(intCmd2.Result())
	// 667 <nil>

	intCmd = rdb.Decr(ctx,"go")
	fmt.Println(intCmd.Result())
	// 666 <nil>

	intCmd3 := rdb.IncrBy(ctx, "go", 333)
	fmt.Println(intCmd3.Result())
	// 999 <nil>
	
	intCmd3 = rdb.DecrBy(ctx, "go", 333)
	fmt.Println(intCmd3.Result())
	// 666 <nil>

	floatCmd := rdb.IncrByFloat(ctx, "go", 0.666)
	fmt.Println(floatCmd.Result())
	// 666.666 <nil>
}
```

## List

```GO
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	// 为已存在的列表添加值， 不存在时返回length为0， 且push不会成功,也不会打印错误
	intCmd := rdb.RPushX(ctx, "not:exists:list", "a", "b", "c")
	length, err := intCmd.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length, err)
	// 0 <nil>

	// 向列表添加元素，无论列表是否存在
	intCmd = rdb.RPush(ctx, "example:list", "a", "b", "c")
	length, err = intCmd.Result()
	fmt.Println(length, err)
	// 3 <nil>

	// 移除列表的最后一个元素
	strCmd := rdb.RPop(ctx, "example:list")
	result, err := strCmd.Result()
	fmt.Println(result, err)
	// c <nil>

	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	strCmd = rdb.RPopLPush(ctx, "example:list", "example:list2")
	result, err = strCmd.Result()
	fmt.Println(result, err)
	// b <nil>

	// 列出长度
	intCmd = rdb.LLen(ctx, "example:list2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 列出某个索引位置的值
	strCmd = rdb.LIndex(ctx, "example:list2", 0)
	fmt.Println(strCmd.Result())
	// b <nil>

	// 列出某个范围索引的值
	strSliceCmd := rdb.LRange(ctx, "example:list2", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [b] <nil>

	// 弹出最后一个元素，如果没有则会阻塞列表直到等待超时或发现可弹出元素为止。
	// 可以有多个key，当第一个key中没有值时，才会从第二个key中pop
	strSliceCmd = rdb.BRPop(ctx, 5*time.Second, "example:list", "example:list2")
	var results []string
	results, err = strSliceCmd.Result()
	fmt.Println(results, err)
	// [example:list a] <nil>

	// 将a插入到列表中的b之前, 然后返回插入后列表元素
	intCmd = rdb.LInsert(ctx, "example:list2", "before", "b", "a")
	// 等价于
	// intCmd = rdb.LInsertBefore(ctx, "example:list2", "b", "a")
	fmt.Println(intCmd.Result())
	// 2 <nil>
	fmt.Println(rdb.LRange(ctx, "example:list2", 0, -1).Result())
	// [a b] <nil>

	// 通过索引设置列表元素的值
	statusCmd := rdb.LSet(ctx, "example:list2", 0, "c")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 移除元素多少次，0代表全部移除, 返回移除元素的个数
	intCmd = rdb.LRem(ctx, "example:list2", 0, "c")
	fmt.Println(intCmd.Result())
	// 1 <nil>
}
```

## Set

```go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	// 向集合添加元素, 返回成功添加的元素个数
	intCmd := rdb.SAdd(ctx, "example:set1", "s1","s2","s3")
	intCmd = rdb.SAdd(ctx, "example:set2", "e1","e2","e3", "s1")
	fmt.Println(intCmd.Result())
	// 4 <nil>

	// 差集 返回set1中与set2不相同的元素
	strSliceCmd := rdb.SDiff(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s2 s3] <nil>

	// 交集
	strSliceCmd = rdb.SInter(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>

	// 并集
	strSliceCmd = rdb.SUnion(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s3 s1 e3 e1 s2 e2] <nil>

	// 将交集存储至storeInter集合中, 并返回交集个数
	intCmd = rdb.SInterStore(ctx, "storeInter","example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeInter").Result())
	// [s1] <nil>
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 将差集存储至storeDiff集合中, 并返回差集个数
	intCmd = rdb.SDiffStore(ctx, "storeDiff","example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeDiff").Result())
	// [s2 s3] <nil>
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 将并集存储至storeUnion集合中, 并返回并集个数
	intCmd = rdb.SUnionStore(ctx, "storeUnion","example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeUnion").Result())
	// [s3 s1 e3 e1 s2 e2] <nil>
	fmt.Println(intCmd.Result())
	// 6 <nil>

	// 判断是否是集合中的元素
	boolCmd := rdb.SIsMember(ctx,"example:set1", "s1")
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 获取成员集合数
	intCmd = rdb.SCard(ctx, "example:set1")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 获取全部成员
	strSliceCmd = rdb.SMembers(ctx, "example:set1")
	fmt.Println(strSliceCmd.Result())
	//[s2 s3 s1] <nil>

	// 获取1个随机成员
	strCmd := rdb.SRandMember(ctx, "example:set1")
	fmt.Println(strCmd.Result())
	// s2 <nil>

	// 获取随机N个成员
	strSliceCmd = rdb.SRandMemberN(ctx, "example:set1", 2)
	fmt.Println(strSliceCmd.Result())
	// [s2 s3] <nil>

	// 移除并返回一个随机成员
	strCmd = rdb.SPop(ctx, "example:set1")
	fmt.Println(strCmd.Result())
	// s1 <nil>

	// 移除指定成员, 返回移除成功的成员个数
	intCmd = rdb.SRem(ctx, "example:set1", "s1", "s2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	fmt.Println(rdb.SMembers(ctx,"example:set2").Result())
	// [e3 e1 s1 e2] <nil>
	// 扫描返回值匹配e*的1个元素
	scanCmd := rdb.SScan(ctx, "example:set2", 0 , "e*", 1)
	var res []string
	res, cursor, err := scanCmd.Result()
	fmt.Println(res, cursor, err)
	// [e3] 4 <nil> 

	// 继续从上次游标处开始扫描
	scanCmd = rdb.SScan(ctx, "example:set2", cursor, "e*", 1)
	fmt.Println(scanCmd.Result())
	// [e1] 2 <nil>
}
```

## SortedSet

```go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	// 向有序集合添加元素, 并返回成功添加的个数
	intCmd := rdb.ZAdd(ctx, "sorted:set1", &redis.Z{Score: 10, Member: "s1"}, &redis.Z{Score: 20, Member: "s2"})
	intCmd = rdb.ZAdd(ctx, "sorted:set2", &redis.Z{Score: 10, Member: "e1"}, &redis.Z{Score: 20, Member: "e2"})
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取有序集合的元素个数
	intCmd = rdb.ZCard(ctx, "sorted:set1")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取成员分数值
	floatCmd := rdb.ZScore(ctx, "sorted:set1", "s2")
	fmt.Println(floatCmd.Result())
	// 20 <nil>

	// 获取有序集合中指定分数段中元素的个数
	intCmd = rdb.ZCount(ctx, "sorted:set1", "15", "30")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 获取指定字典区间内成员数量, 包含s1及之后的所有成员
	intCmd = rdb.ZLexCount(ctx, "sorted:set1", "[s1", "-")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取指定索引区间内的成员
	strSliceCmd := rdb.ZRange(ctx, "sorted:set1", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>

	// 获取指定索引区间内带有排序分数的成员
	zSliceCmd := rdb.ZRangeWithScores(ctx, "sorted:set1", 0, -1)
	fmt.Println(zSliceCmd.Result())
	// [{10 s1} {20 s2}] <nil>

	// 获取指定分数区间内的成员
	strSliceCmd = rdb.ZRangeByScore(ctx, "sorted:set1", &redis.ZRangeBy{Max: "10"})
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>

	// 通过字典区间返回有序集合的成员。直接对成员值作区间操作, (->不包含，[->包含
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "-", Max: "(s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "-", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "(s1", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s2] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "[s1", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>

	// 获取指定成员的索引
	intCmd = rdb.ZRank(ctx, "sorted:set1", "s2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 增加指定成员的分数
	floatCmd = rdb.ZIncr(ctx, "sorted:set1", &redis.Z{Score: 40, Member: "s2"})
	fmt.Println(floatCmd.Result())
	// 60 <nil>

	// 增加指定成员的分数
	floatCmd = rdb.ZIncrBy(ctx, "sorted:set1", 0.5, "s1")
	fmt.Println(floatCmd.Result())
	// 10.5 <nil>

	// 将交集存储在新的有序集合中
	intCmd = rdb.ZInterStore(ctx, "zStoreInter", &redis.ZStore{Keys: []string{"sorted:set1", "sorted:set2"}})
	fmt.Println(intCmd.Result())
	// 0 <nil>
	fmt.Println(rdb.ZRange(ctx,"zStoreInter", 0 ,-1).Result())
	// [] <nil>
	
	// 将并集存储在新的有序集中
	intCmd = rdb.ZUnionStore(ctx, "zStoreUnion", &redis.ZStore{Keys: []string{"sorted:set1", "sorted:set2"}})
	fmt.Println(intCmd.Result())
	// 4 <nil>
	fmt.Println(rdb.ZRange(ctx,"zStoreUnion", 0 ,-1).Result())
	// [e1 s1 e2 s2] <nil>

	// 返回索引区间内的成员，通过分数从高到低排列
	strSliceCmd = rdb.ZRevRange(ctx, "zStoreUnion", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [s2 e2 s1 e1] <nil>
	
	// 返回指定分数区间内的成员，通过分数从高到低排列
	strSliceCmd =  rdb.ZRevRangeByScore(ctx, "zStoreUnion", &redis.ZRangeBy{Max: "20"})
	fmt.Println(strSliceCmd.Result())
	// [e2 s1 e1] <nil>
	
	// 按分数从高到低排列，返回指定成员排名
	intCmd = rdb.ZRevRank(ctx, "zStoreUnion", "s1")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 移除成员
	intCmd = rdb.ZRem(ctx, "zStoreUnion", "s1")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 移除指定分数段内的成员
	intCmd = rdb.ZRemRangeByScore(ctx, "zStoreUnion", "0", "25")
	fmt.Println(intCmd.Result())
	// 2 <nil>
	
	// 移除指定排名区间内的成员
	intCmd = rdb.ZRemRangeByRank(ctx, "zStoreUnion", 1, 1)
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 移除指定字典区间内的成员
	intCmd = rdb.ZRemRangeByLex(ctx, "zStoreUnion", "[e1", "(e2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 扫描返回值匹配e*的1个元素
	scanCmd := rdb.ZScan(ctx, "sorted:set2", 0 , "e*", 1)
	var res []string
	res, cursor, err := scanCmd.Result()
	fmt.Println(res, cursor, err)
	// [e1 10 e2 20] 0 <nil>

	// 继续从上次游标处开始扫描
	scanCmd = rdb.ZScan(ctx, "sorted:set2", cursor, "e*", 1)
	fmt.Println(scanCmd.Result())
	// [e1 10 e2 20] 0 <nil>
}
```

## Hash

```go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	// 设置哈希结构的field:value, 三种方式等价
	// HMSet()、HMGet()在redis3已废弃
	intCmd := rdb.HSet(ctx, "example:hash1", "name", "张三", "role", "法外狂徒")
	intCmd = rdb.HSet(ctx, "example:hash2", map[string]interface{}{"name": "李四", "role": "???"})
	intCmd = rdb.HSet(ctx, "example:hash3", []string{"name", "王五", "role", "!!!"})
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 在字段不存在时才设置field:value
	boolCmd := rdb.HSetNX(ctx, "example:hash1", "age", 16)
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 获取哈希字段数量
	intCmd = rdb.HLen(ctx, "example:hash1")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 获取指定字段的值
	strCmd := rdb.HGet(ctx, "example:hash1", "role")
	fmt.Println(strCmd.Result())
	// 法外狂徒 <nil>

	// 获取哈希字段
	strSliceCmd := rdb.HKeys(ctx, "example:hash1")
	fmt.Println(strSliceCmd.Result())
	// [name role age] <nil>

	// 获取哈希字段值
	strSliceCmd = rdb.HVals(ctx, "example:hash1")
	fmt.Println(strSliceCmd.Result())
	// [张三 法外狂徒 16] <nil>

	// 获取所有字段和值
	strStrMapCmd := rdb.HGetAll(ctx, "example:hash1")
	fmt.Println(strStrMapCmd.Result())
	// map[age:16 name:张三 role:法外狂徒] <nil>

	// 判断字段是否存在
	boolCmd = rdb.HExists(ctx, "example:hash1", "name")
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 为字段增加值
	intCmd = rdb.HIncrBy(ctx, "example:hash1", "age", 2)
	fmt.Println(intCmd.Result())
	// 18 <nil>

	// 为字段增加浮点数值
	floatCmd := rdb.HIncrByFloat(ctx, "example:hash1", "age", 0.5)
	fmt.Println(floatCmd.Result())
	// 18.5 <nil>

	// 删除字段
	intCmd = rdb.HDel(ctx, "example:hash1", "name", "role", "age")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 迭代扫描哈希键值对
	scanCmd := rdb.HScan(ctx, "example:hash2", 0, "name", 1)
	fmt.Println(scanCmd.Result())
	// [name 李四] 0 <nil>
}
```

## Stream

详见下文消息队列。

# 消息队列

## Stream简介

日常开发中，我们可以通过`redis`的`List`类型，调用其阻塞或非阻塞的API来实现一个简单的消息队列，但这种方式往往只能应对简单的功能需求。
		在Redis 5.0版本中引入了一个新的数据类型`Stream`，它在内存中以抽象的方式模拟了一种类似日志的数据结构，是一个`仅附加`的数据结构。就像在内存中以`append`模式打开了一个日志文件一样，只能源源不断地往后追加内容。`Stream`主要用于消息队列，允许消费者以阻塞的方式等待生产者向Stream中发送新的消息，还可以创建消费者组。因为`Stream`是一种类似日志的数据结构，因此在消息传递的过程中可以记录历史消息，保证消息不丢失。

`Stream`消息队列图示如下：

![](./stream.png)

* 每个`Stream`都有一个名称，就是我们通过命令设置的`REDIS KEY`。
* 通过`XADD`可以向指定的`Stream`追加一个新的消息条目。一个Stream消息条目不是简单的字符串，而是由一个或多个键值对组成的结构化数据。
* 每个消息的ID结构是由生成id时服务器的毫秒时间戳加一个单调递增的序号组成的。形如`<millisecondsTime>-<sequenceNumber>`. 借助于ID的结构，`Stream`支持按照序号id和时间进行范围查询。
* `Stream`流与阻塞`List`在行为上有如下不同：
  * Stream中默认每一个新消息条目，都会被分发到等待的每一个消费者。而List中每个消费者都会获取到不同的元素。
  * Stream中所有的消息都被无限期地附加到Stream中（除非用户明确地要求删除这些条目），也就是消息在Stream不会被删除，同时不同的消费者组会记录读取的最后一条消息的ID（last_deliver_id）。而在阻塞List中，一个消息会从列表中pop掉，即有效删除。
* 消费者组从流中获取数据，为多个消费者提供服务。具有显式地确认已经处理的消息、检查待处理的消息、申明未处理消息的能力。消费者组具有如下功能：
  * 每条消息都提供给不同的消费者，相同的消息不可能传递给多个消费者。
  * 消费者在消费者组中具有唯一的名称标识，在第一次被用到的时候自动创建，不需要显式创建。
  * 每一个消费者组都会记录获取的最后一个消息的ID，当消费者请求新消息时，只会提供大于当前ID的消息。任意一个消费者读取了消息都会使游标 `last_deliver_id` 往后移动。
  * 消息传递到消费者时，还没有被确认为已处理。消费消息需要使用特定的命令进行显式确认`ACK`。 pending_ids 记录了当前已经被读取，但是还没有 `ACK`的消息id。

- 消费者组可以同时从多个Stream中读取消息，但是必须保证在每一个Stream中消费者组的名称相同。
- 每一条消息都有一个计数器，只有在调用`XREADGROUP`读取、调用`XCLAIM`改变归属时才会增加1.

## 常用命令

**消息队列相关命令：**

- **XADD** - 添加消息到末尾
- **XTRIM** - 对流进行修剪，限制长度
- **XDEL** - 删除消息
- **XLEN** - 获取流包含的元素数量，即消息长度
- **XRANGE** - 获取消息列表，会自动过滤已经删除的消息
- **XREVRANGE** - 反向获取消息列表，ID 从大到小
- **XREAD** - 以阻塞或非阻塞方式获取消息列表

**消费者组相关命令：**

- **XGROUP CREATE** - 创建消费者组
- **XREADGROUP GROUP** - 读取消费者组中的消息
- **XACK** - 将消息标记为"已处理"
- **XGROUP SETID** - 为消费者组设置新的最后递送消息ID
- **XGROUP DELCONSUMER** - 删除消费者
- **XGROUP DESTROY** - 删除消费者组
- **XPENDING** - 显示待处理消息的相关信息
- **XCLAIM** - 转移消息的归属权
- **XINFO** - 查看流和消费者组的相关信息；
- **XINFO GROUPS** - 打印消费者组的信息；
- **XINFO STREAM** - 打印流信息

## go中操作

```GO
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	// 向stream中添加消息，结果返回消息ID
	xAddArgs := &redis.XAddArgs{
		Stream: "example:stream",
		MaxLen: 10,
		Values: map[string]interface{}{"foodId": "10001", "foodName": "麻婆豆腐"},
	}
	strCmd := rdb.XAdd(ctx, xAddArgs)
	fmt.Println(strCmd.Result())
	// 1609083771429-0 <nil>

	// 插入第二条消息
	strCmd = rdb.XAdd(ctx, &redis.XAddArgs{Stream: "example:stream", Values: []string{"foodId", "10002", "foodName", "飘香鸡翅"}})
	fmt.Println(strCmd.Result())
	// 1609083771430-0 <nil>

	// 限制stream的长度为大约10个
	intCmd := rdb.XTrimApprox(ctx, "example:stream", 10)
	fmt.Println(intCmd.Result())
	// 0 <nil>

	// 获取stream流的长度
	intCmd = rdb.XLen(ctx, "example:stream")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取消息列表
	xMessageSliceCmd := rdb.XRange(ctx, "example:stream", "-", "+")
	fmt.Println(xMessageSliceCmd.Result())
	// [{1609083771429-0 map[foodId:10001 foodName:麻婆豆腐]} {1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]}] <nil>

	// 反向获取消息列表, 第一个为最新的消息
	xMessageSliceCmd = rdb.XRevRange(ctx, "example:stream", "+", "-")
	fmt.Println(xMessageSliceCmd.Result())
	// [{1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]} {1609083771429-0 map[foodId:10001 foodName:麻婆豆腐]}] <nil>

	// 读取给定id的下一条消息
	xReadArgs := &redis.XReadArgs{
		Streams: []string{"example:stream", "1609083771429-0"},
		Count:   1,
		Block:   5 * time.Second,
	}
	xStreamSliceCmd := rdb.XRead(ctx, xReadArgs)
	fmt.Println(xStreamSliceCmd.Result())
	// [{example:stream [{1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]}]}] <nil>

	// 删除消息,把两条消息都删掉
	intCmd = rdb.XDel(ctx, "example:stream", "1609083771430-0", "1609083771429-0")
	fmt.Println(intCmd.Result())
	// 2 <nil>

  // 在stream上创建消费者组eater, 从最新的消息($表示)开始消费
	statusCmd := rdb.XGroupCreate(ctx, "example:stream", "eater", "$")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 读取消费者组eater中的未被其他消费者读取的消息>
	// 运行后会阻塞，在redis客户端命名输入 XADD "example:stream" * foodId 1003 foodName 肥宅快乐水 会获得结果
	xReadGroupArgs := &redis.XReadGroupArgs{
		Group:    "eater",                         // 消费者组
		Consumer: "eater01",                       // 消费者，用时即创建
		Streams:  []string{"example:stream", ">"}, // stream流
		Block:    0,                               // 无限等待
		NoAck:    false,                           // 需要进行确认
	}
	xStreamSliceCmd = rdb.XReadGroup(ctx, xReadGroupArgs)
	xStreamSlice, err := xStreamSliceCmd.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xStreamSlice)
	// [{example:stream [{1609086089189-0 map[foodId:1003 foodName:肥宅快乐水]}]}]

	// 确认消息为已处理
	intCmd = rdb.XAck(ctx, "example:stream", "eater", "1609086089189-0")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 设置最后递送的id为1609086089189-0
	statusCmd = rdb.XGroupSetID(ctx, "example:stream", "eater", "1609086089189-0")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 查看待处理消息
	// type XPending struct {
	// 	Count     int64
	// 	Lower     string
	// 	Higher    string
	// 	Consumers map[string]int64
	// }
	xPendingCmd := rdb.XPending(ctx, "example:stream", "eater")
	fmt.Println(xPendingCmd.Result())
	// &{1 1609086342551-0 1609086342551-0 map[eater01:1]} <nil>

	// 转移消息的归属权, 将超过两分钟仍未得到处理的消息转移至消费者eater02
	xClaimArgs := &redis.XClaimArgs{
		Stream:   "example:stream",
		Group:    "eater",
		Consumer: "eater02",
		MinIdle:  2 * time.Minute,
		Messages: []string{"1609086342551-0"},
	}
	xMessageSliceCmd = rdb.XClaim(ctx, xClaimArgs)
	fmt.Println(xMessageSliceCmd.Result())
	// [] <nil> // 没有满足要求的消息

	// 查看流信息
	// type XInfoStream struct {
	//   Length          int64
	//   RadixTreeKeys   int64
	//   RadixTreeNodes  int64
	//   Groups          int64
	//   LastGeneratedID string
	//   FirstEntry      XMessage
	//   LastEntry       XMessage
	// }
	xInfoStreamCmd := rdb.XInfoStream(ctx, "example:stream")
	fmt.Println(xInfoStreamCmd.Result())
	// &{3 1 2 1 1609086342551-0 {1609082364313-0 map[foodId:10001 foodName:麻婆豆腐]} {1609086342551-0 map[foodId:1003 foodName:肥宅快乐水]}} <nil>

	// 查看消费者组消息
	// type XInfoGroup struct {
	// 	Name            string
	// 	Consumers       int64
	// 	Pending         int64
	// 	LastDeliveredID string
	// }
	xInfoGroupCmd := rdb.XInfoGroups(ctx, "example:stream")
	fmt.Println(xInfoGroupCmd.Result())
	// [{eater 0 0 1609086089189-0}] <nil>

	// 删除消费者
	intCmd = rdb.XGroupDelConsumer(ctx, "example:stream", "eater", "eater01")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 删除消费者组
	intCmd = rdb.XGroupDestroy(ctx, "example:stream", "eater")
	fmt.Println(intCmd.Result())
	// 1 <nil>
}
```

# 发布订阅

```GO
// pub.go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func pub() {
	ctx := context.Background()
	rdbPub := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})
	intCmd := rdbPub.Publish(ctx, "example:channel", "明天放假了!!!")
	fmt.Println(intCmd.Result())
}

func main() {
	pub()
}
```

```GO
// sub.go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func sub() {
	ctx := context.Background()

	rdbSub := redis.NewClient(&redis.Options{
		Addr: "localhost:6381",
	})

	sub := rdbSub.Subscribe(ctx, "example:channel")
	// 匹配pattern进行订阅
	// sub:= rdbSub.PSubscribe(ctx, "*:channel")

	select {
	case message := <-sub.Channel():
		fmt.Println(message.Channel)
		// example:channel
		fmt.Println(message.Pattern)
		//
		fmt.Println(message.Payload)
		// 明天放假了!!!
		fmt.Println(message.PayloadSlice)
		// []
		fmt.Println(message.String())
		// Message<example:channel: 明天放假了!!!>
	}
	defer sub.Close()

	// 获取匹配pattern的所有channels
	strSliceCmd := rdbSub.PubSubChannels(ctx, "*:channel")
	fmt.Println(strSliceCmd.Result())
	// [example:channel] <nil>

	// 获取pattern的个数
	intCmd := rdbSub.PubSubNumPat(ctx)
	fmt.Println(intCmd.Result())
	// 0 <nil>

	// 获取指定channel上的订阅者个数<代码订阅1个+终端命令行订阅1个>
	strIntMapCmd := rdbSub.PubSubNumSub(ctx, "example:channel")
	fmt.Println(strIntMapCmd.Result())
	// map[example:channel:2] <nil>
}

func main() {
	sub()
}
```

# 参考资料

[https://pkg.go.dev/github.com/go-redis/redis/v8#section-documentation](https://pkg.go.dev/github.com/go-redis/redis/v8#section-documentation)

[http://www.redis.cn/topics/streams-intro.html](http://www.redis.cn/topics/streams-intro.html)

