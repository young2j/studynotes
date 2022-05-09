# 背景

在开发中，我们很容易通过docker启动一个普通的mongodb数据库服务。但是有时候为了保持与线上环境一致，或者为了利用mongodb副本集的某些特性，我们需要在本地部署mongodb副本集。副本集往往需要启动多个mongodb服务作为副本集成员，而通常用于开发的笔记本资源比较有限。鉴于此，官方文档给了解决办法，可以直接将一个单节点mongodb服务转换为单节点副本集(`standlone replica set`)(https://www.mongodb.com/docs/manual/tutorial/convert-standalone-to-replica-set/)

# 启动步骤

按照官方文档的说明，如果利用docker部署服务，那么依次有如下步骤:

* 第一步, 假如已经存在一个运行中的普通mongodb容器服务。此时，需要关闭服务，并通过**指定`--replSet`参数**重启该服务或者重新启动一个新的mongodb容器。

  假如mongodb的服务名及容器名均为`mongodb_rs`，运行端口映射为`27017:27017`，副本集名称为`rs0`，数据存储目录指定为`/srv/mongodb/db0`，数据卷挂载目录为`./data:/srv/mongodb/db0`。那么`docker-compose.yaml `文件可编写如下:

  ```yaml
  version: "3"
  services:  
    mongodb_rs:
      network_mode: bridge
      container_name: mongodb_rs
      image: mongo:latest
      ports:
        - "27017:27017"
      restart: always
      # environment:
      #   MONGO_INITDB_ROOT_USERNAME: username
      #   MONGO_INITDB_ROOT_PASSWORD: pwd
      command: mongod --port 27017 --replSet rs0 --dbpath /srv/mongodb/db0
      volumes:
        - ./data:/srv/mongodb/db0
  ```

* 第二步，执行如下命令启动mongodb服务

  ```shell
  docker-compose up -d mongodb_rs
  ```

* 第三步，进入容器mongosh，执行初始化副本集命令

  ```shell
  docker exec -it mongodb_rs mongo
  ```

  ```shell
  # mongosh
  rs.initiate()
  
  # ---
  # > rs.initiate()
  # {
  #	 "info2" : "no configuration specified. Using a default configuration for the set",
  #	 "me" : "f76081e20602:27017",
  #	 "ok" : 1
  # }
  # rs0:SECONDARY>
  # rs0:PRIMARY>|
  ```

* 第四步，退出容器，容器服务正常运行

# 可能遇到的问题

按照上述步骤执行后，通常情况下容器服务可以正常运行，应用程序可以正常进行连接，到这里基本就成功了。以golang代码测试：

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017/?replicaSet=rs0")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	colls, _ := client.Database("admin").ListCollectionNames(context.TODO(), bson.M{})
	fmt.Printf("colls: %v\n", colls)
}
// colls: [system.keys system.version]
```

但是有时候可能会出现本地程序代码无法连接副本集服务，控制台会报类似连接错误的问题，报错的原因在于`副本集无法识别成员host`。

在单节点副本集下，本机既是主也是从，在容器的`mongo shell`中可进行查看, `members`只有一个成员，其`name`为`"f76081e20602:27017"`,  所以如果你遇到无法连接或者其他类似错误，根本原因在于本地启动的这个副本集无法识别`f76081e20602`这个`host`。

```shell
rs.status()
---
rs0:PRIMARY> rs.status()
{
	"set" : "rs0",
	"date" : ISODate("2022-05-06T18:59:21.417Z"),
# ...	
	"members" : [
		{
			"_id" : 0,
			"name" : "f76081e20602:27017",
			"health" : 1,
			"state" : 1,
			"stateStr" : "PRIMARY",
			"uptime" : 1433,
			"optime" : {
				"ts" : Timestamp(1651863555, 1),
				"t" : NumberLong(1)
			},
		# ...
		}
	],
# ...
}
```

# 解决办法

曾经遇到上述问题，百度csdn上有多篇内容一样的文章，都说这种情况需要将应用程序也通过容器进行启动，并将应用程序与mongdb副本集服务置于同一个docker网络中，就可以正常连接了。这样做确实也可行，但似乎过于麻烦了，有点走歪路的感觉。

从上述内容已经知道是副本集成员`host`的识别问题，那么在初始化mongodb副本集时，我们可以显式的去指定成员`host`，不使用默认的副本集配置。具体而言，将启动步骤中的第三步更改为:

* 进入`mongosh`

  ```shell
  docker exec -it mongodb_rs mongo
  ```

* 自定义配置

  ```shell
  # mongosh
  conf = {
     _id : "rs0",
     members: [
        { _id: 0, host: "<本机ip地址>:27017" },
     ]
  }
  ```

* 初始化副本集

  ```shell
  # mongosh
  rs.initiate(conf)
  ```

如此，通过指定成员ip，mongo单节点副本集就可以准确的识别到副本集成员，对于多节点副本集如果出现连接问题，此方法同样适用。