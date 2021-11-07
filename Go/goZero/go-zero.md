# 安装

```shell
# 安装go-zero
go get -u github.com/tal-tech/go-zero
# 安装命令行工具
go get -u github.com/tal-tech/go-zero/tools/goctl
```

# 快速开始

```shell
goctl api new qkstart

qkstart
├── etc
│   └── qkstart-api.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── qkstarthandler.go
│   │   └── routes.go
│   ├── logic
│   │   └── qkstartlogic.go
│   ├── svc
│   │   └── servicecontext.go
│   └── types
│       └── types.go
├── qkstart.api
└── qkstart.go
```

# bookstore

## api

```shell
# bookstore/api/

# 单独生成一个api文件用于编辑
goctl api -o bookstore.api 
# 根据api文件生成go api代码
goctl api go -api bookstore.api -dir .

bookstore
├── api
│   ├── bookstore.api
│   ├── bookstore.go
│   ├── etc
│   │   └── bookstore-api.yaml
│   └── internal
│       ├── config
│       │   └── config.go
│       ├── handler
│       │   ├── addhandler.go
│       │   ├── checkhandler.go
│       │   └── routes.go
│       ├── logic
│       │   ├── addlogic.go
│       │   └── checklogic.go
│       ├── svc
│       │   └── servicecontext.go
│       └── types
│           └── types.go
└── go.mod
```

## rpc

```shell
# bookstore/rpc/

# 安装protoc
 apt install -y protobuf-compiler # Linux
 brew install protobuf # Mac
# 安装protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go

# 生成rpc模版文件.proto, 然后可编辑
goctl rpc template -o add.proto
goctl rpc template -o check.proto
# 根据编辑后的proto文件生成rpc代码
goctl rpc proto -src add.proto -dir . 
goclt rpc proto -src check.proto -dir .
```

## 在api中调用rpc

## model

```sql
-- 首先编辑sql文件: `book.sql`
CREATE TABLE `book`
(
  `book` varchar(255) NOT NULL COMMENT 'book name',
  `price` int NOT NULL COMMENT 'book price',
  PRIMARY KEY(`book`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

```shell
# 【使用ddl】 然后根据sql文件生成代码, -c 指定使用redis缓存
goctl model mysql ddl -c -src book.sql -dir .

# 【使用datasource】
goctl model mysql datasource -c -url root:rootroot@tcp(127.0.0.1:3306)/gozero -t book -dir .
```

