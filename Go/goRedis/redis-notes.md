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

解析URL进行连接

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

# 数据操作

# 自定义命令

# 发布订阅



