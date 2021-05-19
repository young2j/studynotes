```shell
# 生成go module目录
mkdir cobra && go mod init cobra

# 在GOPATH下可生成cobra可执行文件
go get -u github.com/spf13/cobra

# 初始化一个cobra app
 cobra init --pkg-name cobra

# 查看目录
tree .
├── LICENSE
├── cmd
│   └── root.go
├── cobra.md
├── go.mod
├── go.sum
└── main.go

# 执行
go mod tidy

# 添加子命令, 子命令名称必须为驼峰, -p 指定父cmd， 默认为rootCmd
cobra add serve [-p rootCmd]
cobra add config [-p rootCmd]
cobra add create -p configCmd

# tree .
.
├── LICENSE
├── cmd
│   ├── config.go
│   ├── create.go
│   ├── root.go
│   └── serve.go
├── cobra.md
├── go.mod
├── go.sum
└── main.go

# 然后就可以执行如下命令
go run main.go serve
go run main.go help serve
go run main.go config
go run main.go config create
```

