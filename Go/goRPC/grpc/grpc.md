<p style="text-align:center;font-size:28px;font-weight:bold;">grpc</p>

# 安装compiler

```shell
# linux
apt install -y protobuf-compiler
# mac
brew install protobuf

# windows
# https://github.com/protocolbuffers/protobuf/releases上下载pre-compiled binaries

protoc --version
```

# go

```shell
# 安装grpc
go get google.golang.org/grpc
# 安装插件
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 更新环境变量以使protoc能够找到插件
export PATH="$PATH:$(go env GOPATH)/bin"
```

```shell
# 编译proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto

proto/hello
├── hello.pb.go # 包含proto文件定义的所有请求和响应的消息类型（message types）
├── hello_grpc.pb.go # 包含客户端需要调用的、服务端需要实现的接口方法
└── hello.proto
```

