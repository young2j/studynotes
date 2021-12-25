<p style="text-align:center;font-size:28px;font-weight:bold;">grpc</p>

# protobuf

## 安装protoc

```shell
# linux
apt install -y protobuf-compiler
# mac
brew install protobuf

# windows
# https://github.com/protocolbuffers/protobuf/releases上下载pre-compiled binaries

protoc --version
```

## 安装go插件

```shell
# 安装grpc
go get google.golang.org/grpc
# 安装插件
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 更新环境变量以使protoc能够找到插件
export PATH="$PATH:$(go env GOPATH)/bin"
```

## proto基本规则

```protobuf
syntax = "proto3";  // 指定proto2、proto3
package hello; // 指定包名
option go_package = ".;hello"; 
// option java_package = "hello";

// 可以使用import语句导入其它proto文件
// import "others.proto";
// import alias "others.proto";
import "google/protobuf/timestamp.proto";

// ------------定义RPC服务接口service--------------
// 1. protocol buffer编译器会根据所选择的不同语言生成对应的服务接口代码
// 2. 接口方法需要定义请求参数HelloRequest以及返回参数HelloResponse
service HelloService {  // service 采用驼峰命名 
    rpc SayHello(SayHelloRequest) returns(SayHelloResponse){}
}


//--------------定义请求和响应Message---------------
message SayHelloRequest { // message 采用首字母大写驼峰命名
 // 字段命名规则：修饰符 类型 名称(小写+下划线) = 字段编号或默认值;
    //-- 修饰符：
        // required : proto2
        // optional : proto2
        // repeated : proto2/proto3
    //-- 字段编号：
        // 有了该值，通信双方才能互相识别对方的字段，相同的编码值，其限定修饰符和数据类型必须相同，编码值的取值范围为 1~2^29-1(536870911)
        // 其中 1~15的编码时间和空间效率都是最高的，编码值越大，其编码的时间和空间效率就越低，所以建议把经常要传递的值把其字段编码设置为1-15之间的值
        //  19000~19999 编码值为Google protobuf 系统内部保留值，建议不要在自己的项目中使用
    //-- 字段默认值：
        // 当在传递数据时，对于required数据类型，如果用户没有设置值，则使用默认值传递到对端
   int32   age = 1;      // 必须字段
   optional int64   count = 2;    // 可选字段
   repeated double  money = 3;    // 重复字段 也是可选的,但可以包含多个值,可看作是在传递一个数组的值, 规范为使用复数
    float   score = 4;
    string  name = 5;
    bool    fat = 6;
    bytes   char = 7;
  
    // 定义枚举型：规范为采用驼峰命名，字段使用 枚举名前缀_全大写_加下划线命名, 必须有0号字段，为默认值
    enum Status {
        STATUS_OK_UNSPECIFIED = 0;
        STATUS_FAIL = 1;
    }
    Status status = 8;
    
    // 可以任意嵌套message 定义在内外都可以, 但内部声明的message只可在内部直接使用
    message NestMessage {
        bool is_nest = 1;
    }
    NestMessage nest_message = 9;
    
    // 定义map 类型, k,v 也可以是定义的message类型, 不能为repeated
    map<string, int32> map_field = 10;

    // 可以使用包名+消息名的方式来使用类型
    // others.foo.Bar from_other = 11;

    // oneof 最多可以同时设置一个字段, 设置 oneof 的任何成员会自动清除所有其他成员
    oneof result {
        string result_a = 12;
        int32 result_b = 13;
        google.protobuf.Timestamp result_c = 14;
    }
    // 保留字段名和字段号, 将不能使用
    reserved 20, 21;
    reserved "field_name";
}

message SayHelloResponse {
    string code = 1;
    message Data {
        int32 age = 1;
        string name = 2;
        double money = 3;
    }
    Data data = 2;
} 
```

## protoc命令行

```shell
protoc
-I PATH  指定proto文件的搜索路径，可指定多个。若不指定，默认值为当前路径。 是 --proto_path=PATH的简写。
```

```shell
# 由 protoc-gen-go 插件支持的参数
--go_out 指定go文件输出路径, 输出的文件名为 *.pb.go 。
				 主要是根据proto文件中定义的message、field、sevice等，生成golang类型定义，以及各种字段的getter和setter；
				 
--go-grpc_out 指定go_grpc文件的输出路径。
					主要是根据proto文件定义生成的rpc客户端和服务端接口，需要服务端和客户端分别去实现和调用。
					
--go_opt paths=import 默认方式，输出的pb文件将放置在go包目录中；
				 module=$PREFIX 输出的pb文件将放置在导入的go包路径去除PREFIX后的目录中；
				 paths=source_relative 输出的pb文件将放置在proto文件目录中；
				 
				 M*.proto=package_path 使用M指定proto文件的go包导入路径，优先级高于proto文件中的指定。
				 例如 --go_opt=Mprotos/bar.proto=example.com/project/protos/foo 

--go-grpc_opt 同--go_opt
```

## protoc示例

当前目录为`protos/`

```shell
# proto > option go_package = ".;hello";
$ protoc --go_out=. hello.proto
../protos
├── hello.pb.go
└── hello.proto

# proto > option go_package = ".;hello";
$ protoc --go-grpc_out=. hello.proto
../protos
├── hello.proto
└── hello_grpc.pb.go

# proto > option go_package = ".;hello";
$ protoc --go_out=. --go-grpc_out=. hello.proto 
../protos
├── hello.pb.go # 包含proto文件定义的所有请求和响应的消息类型（message types）
├── hello.proto  # proto文件
└── hello_grpc.pb.go # 包含客户端需要调用的、服务端需要实现的接口方法

# proto > option go_package = ".;hello"
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=import hello.proto
../protos
├── hello.pb.go
├── hello.proto
└── hello_grpc.pb.go

# proto > option go_package = "..;hello"
# 将在上层目录中生成pb文件
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=import hello.proto
├── hello.pb.go
├── hello_grpc.pb.go
└── protos
    └── hello.proto
    
# 如果在proto中指定包路径为/路径形式，将在proto同级目录下生成包目录
# proto > option go_package = "example/hello";
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=import hello.proto
../protos
├── example
│   └── hello
│       ├── hello.pb.go
│       └── hello_grpc.pb.go
└── hello.proto


# proto > option go_package = "example/hello";
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative hello.proto
../protos
├── example #  --go-grpc_opt=paths=import
│   └── hello
│       └── hello_grpc.pb.go
├── hello.pb.go # --go_opt=paths=source_relative
└── hello.proto

# proto > option go_package = "example/hello";
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative hello.proto
../protos
├── hello.pb.go
├── hello.proto
└── hello_grpc.pb.go

# proto > option go_package = "example/hello";
# 将去除example前缀，生成hello目录
$ protoc --go_out=. --go-grpc_out=. --go_opt=module=example --go-grpc_opt=module=example hello.proto
../protos
├── hello
│   ├── hello.pb.go
│   └── hello_grpc.pb.go
└── hello.proto
```

# Buf

buf是一个`proto`模块化管理工具，通过提供一个`yaml`配置文件来更加方便的编译proto文件，避免了`protoc`冗长的命令行参数。

## 安装Buf

命令无法安装，就直接去官网下载安装文件 https://docs.buf.build/installation

```shell
# mac
brew tap bufbuild/buf
brew install buf
```

## 使用Buf

### buf config init

在目录`protos/`中新建两个`api`模块：`helloapis、gatewayapis`, 并初始化生成配置文件`buf.yaml`

```shell
$ mkdir helloapis gatewayapis
# 在helloapis 和 gatewayapis目录中分别执行
$ buf config init

../protos
├── gatewayapis
│   └── buf.yaml
└── helloapis
    └── buf.yaml
```

```yaml
# buf.yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

```shell
# 创建示例proto
$ mkdir -p helloapis/hello/v1
$ mkdir -p gatewayapis/gateway/v1

../protos
├── gatewayapis
│   ├── buf.yaml
│   └── gateway
│       └── v1
│           └── gateway.proto
└── helloapis
    ├── buf.yaml
    └── hello
        └── v1
            └── hello.proto
```

### buf ls-files

```shell
# 列出*.proto文件

$ buf ls-files
gatewayapis/gateway/v1/gateway.proto
helloapis/hello/v1/hello.proto
```

### buf.work.yaml

```shell
# 创建工作空间workspace，将两个protoapis模块纳入工作空间进行管理，启到了protoc -IPATH的作用
$ vim buf.work.yaml
version: v1
directories:
  - gatewayapis
  - helloapis
 
# workspace 目录结构
../protos
├── buf.work.yaml
├── gatewayapis
│   ├── buf.yaml
│   └── gateway
│       └── v1
│           └── gateway.proto
└── helloapis
    ├── buf.yaml
    └── hello
        └── v1
            └── hello.proto
```

### buf build

- 根据`buf.yaml`配置去发现所有 `.proto` 文件。
- 将 `.proto` 文件复制到内存中。
- 编译所有 `.proto` 文件。
- 将编译结果输出到可配置的位置（默认为`/dev/null`）

#### 管理.proto文件

```shell
$ buf ls-files
gatewayapis/gateway/v1/gateway.proto
helloapis/hello/v1/hello.proto

# 在helloapis模块的buf.yaml中添加build排除项
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
 build:
   excludes:
     - hello/v1
 # 然后执行ls， helloapis的proto文件将被排除在构建之外
 $ buf ls-files
 gatewayapis/gateway/v1/gateway.proto
```

#### 管理导入依赖

假如需要在`gatewayapis`中导入`helloapis`的`proto`定义：

```shell
# 首先将helloapis从工作空间中注释掉
version: v1
directories:
  - gatewayapis
  # - helloapis

# 然后在gatewayapis/gateway/v1/gateway.proto中添加导入
import "hello/v1/hello.proto";

# 执行build，将提示不存在hello.proto
$ buf build
gatewayapis/gateway/v1/gateway.proto:5:8:hello/v1/hello.proto: does not exist

# 最后再将helloapis从工作空间中取消注释
version: v1
directories:
  - gatewayapis
  - helloapis
# 执行build，就不会出错
$ buf build
```

#### 查看构建结果

```shell
# 执行构建, 可以根据后缀识别构建格式，下面两个是一样的
$ buf build -o image.json
$ buf build -o image -format=json
```

### buf lint

格式检查

```shell
$ buf lint
gatewayapis/gateway/v1/gateway.proto:5:1:Import "hello/v1/hello.proto" is unused.
```

lint排除规则

```yaml
# 可以添加lint排除规则 
# buf.yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
  # 添加lint排除规则, 不推荐
  except:
    - PACKAGE_VERSION_SUFFIX
    - FIELD_LOWER_SNAKE_CASE
    - SERVICE_SUFFIX
  # 忽略指定proto文件lint
  ignore:
     - google/type/datetime.proto
```

### buf generate

```yaml
# 创建buf.gen.yaml
version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
	
	# $ yarn global add grpc-tools
  - name: js
    out: gen/js
    opt: import_style=commonjs,binary
  - name: js-grpc
    out: gen/js
    opt: import_style=commonjs,binary
    path: /usr/local/bin/grpc_tools_node_protoc_plugin
```

 ```shell
 $ buf generate
 gen
 ├── go
 │   ├── gateway
 │   │   └── v1
 │   │       ├── gateway.pb.go
 │   │       └── gateway_grpc.pb.go
 │   └── hello
 │       └── v1
 │           ├── hello.pb.go
 │           └── hello_grpc.pb.go
 └── js
     ├── gatewayapis
     │   └── gateway
     │       └── v1
     │           ├── gateway_grpc_pb.js
     │           └── gateway_pb.js
     └── helloapis
         └── hello
             └── v1
                 ├── hello_grpc_pb.js
                 └── hello_pb.js
 ```
### buf breaking

列出重要改动

```shell
$ buf breaking --against *.git
```
# grpc-gateway

-> 当 HTTP 请求到达 `gRPC-Gateway` 时，它将 `JSON` 数据解析为` protobuf` 消息。
-> 然后使用解析的 `protobuf `消息给到 `Go gRPC` 客户端发出正常的`gRPC`请求。
-> `gRPC` 服务器处理请求并以 `protobuf `二进制格式返回响应到`gRPC`客户端。
->  `gRPC` 客户端将解析` protobuf `消息并返回给` gRPC-Gateway`。
-> `gRPC-Gateway` 将 `protobuf` 消息编码为 `JSON` 并返回给HTTP客户端。

## 安装gateway

```shell
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
```

以`gatewayapis`为例:

```shell
../protos
├── buf.gen.yaml
├── buf.work.yaml
├── gatewayapis
│   ├── buf.lock
│   ├── buf.yaml
│   └── gateway
│       └── v1
│           └── gateway.proto
```

## 添加option标注

```protobuf
// gateway.proto
syntax = "proto3";
package gateway.v1;
option go_package="grpc-notes/protos/gatewayapis/gateway/v1;gatewayv1";
 
import "google/api/annotations.proto"; // 引入google的proto

service ProbeService {
    rpc Ping (PingRequest) returns (PingResponse){
    		// 添加到这里
        option (google.api.http) = {
            post: "/v1/probe/ping"
            body:"*"
        };
    }
}

message PingRequest {
    string msg = 1;
}
message PingResponse {
    string msg = 1;
}
```

## 添加模块依赖

```yaml
# buf.yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
# 添加依赖
deps:
  - buf.build/googleapis/googleapis
```

```shell
# 更新依赖，检查格式
$ cd gatewayapis
$ buf mod update
$ buf lint
```

## 重新生成`stub`

```yaml
# buf.gen.yaml 中添加内容

version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
  
  # $ yarn global add grpc-tools
  - name: js
    out: gen/js
    opt: import_style=commonjs,binary
  - name: js-grpc
    out: gen/js
    opt: import_style=commonjs,binary
    path: /usr/local/bin/grpc_tools_node_protoc_plugin
	
	# 添加下面这段 
  - name: grpc-gateway
    out: gen/go
    opt: paths=source_relative
```

```shell
$ buf generate

../protos
├── buf.gen.yaml
├── buf.work.yaml
├── gatewayapis
│   ├── buf.lock
│   ├── buf.yaml
│   └── gateway
│       └── v1
│           └── gateway.proto
├── gen
│   ├── go
│   │   ├── gateway
│   │   │   └── v1
│   │   │       ├── gateway.pb.go
│   │   │       ├── gateway.pb.gw.go # new add
│   │   │       └── gateway_grpc.pb.go
```

## 访问http接口

```shell
$ curl -XPOST http://localhost:5201/v1/probe/ping -d '{"msg":"ping"}'
{"msg":"pong"}
```

## 生成REST接口的三种方式

### 1. 自定义option标注

这种方式如上所述，需要在`proto`文件的接口中添加`option annotations`，然后会生成指定路由。

### 2. 默认路由映射

这种方式通过在`buf.gen.yaml`中添加`generate_unbound_methods=true` opt，自动生成http路由映射，路由格式为`/包名.服务名/服务方法名`，请求方法均为`POST`。例如在`gateway.proto`中新增`Detect`方法，其路由默认为: `POST gateway.v1.ProbeService/Detect `

```protobuf
//  gatewayapis/gateway/v1/gateway.proto
syntax = "proto3";
package gateway.v1;

service ProbeService {
    rpc Ping (PingRequest) returns (PingResponse){   
        option (google.api.http) = {
            post: "/v1/probe/ping"
            body:"*"
        };
    }
    rpc Detect (DetectRequest) returns (DetectResponse);
}
```

```yaml
# buf.gen.yaml
version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false

  # $ yarn global add grpc-tools
  - name: js
    out: gen/js
    opt: import_style=commonjs,binary
  - name: js-grpc
    out: gen/js
    opt: import_style=commonjs,binary
    path: /usr/local/bin/grpc_tools_node_protoc_plugin

  - name: grpc-gateway
    out: gen/go
    opt:
      - paths=source_relative
      - logtostderr=true
      - generate_unbound_methods=true
```

启动server，进行请求：

```shell
# 请求自定义路由
$ curl -XPOST http://localhost:5201/v1/probe/ping -d '{"msg":"ping"}'                
{"msg":"pong"}

# 自定义路由会覆盖默认路由
$ curl -XPOST http://localhost:5201/gateway.v1.ProbeService/Ping  -d '{"msg":"ping"}'
{"code":5,"message":"Not Found","details":[]}

# 请求默认路由
$ curl -XPOST http://localhost:5201/gateway.v1.ProbeService/Detect -d '{"id":1}'
{"id":"req.id=1"}
```

### 3. 路由映射配置文件

单独使用url映射配置，再使用`buf generate`会出问题，不知道怎么解决。建议单独用`protoc` 指定`url mapping`来生成`gw stub`，或者直接采用1和2的方式。

```yaml
# protos/helloapis/hello/urls.yaml
type: google.api.Service
config_version: 3

http:
  # https://github.com/googleapis/googleapis/blob/master/google/api/http.proto for details
  rules:
    - selector: hello.v1.HelloService.SayHello
      post: /v1/hello/sayhello
      body: "*"
      additional_bindings:
      - get: /v1/hello/sayhello/{name}
```

```yaml
# buf.gen.yaml
version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false

  # $ yarn global add grpc-tools
  - name: js
    out: gen/js
    opt: import_style=commonjs,binary
  - name: js-grpc
    out: gen/js
    opt: import_style=commonjs,binary
    path: /usr/local/bin/grpc_tools_node_protoc_plugin

  - name: grpc-gateway
    out: gen/go
    opt:
      - paths=source_relative
      - logtostderr=true
      - grpc_api_configuration=helloapis/hello/urls.yaml
      - grpc_api_configuration=gatewayapis/gateway/urls.yaml
      - standalone=true
```

```shell
$ buf generate
Failure: plugin grpc-gateway: HTTP rules without a matching selector: .gateway.v1.ProbeService.Ping
```

单独生成gw stub， 但目录结构和`buf generate`的生成不一致

```shell
protoc -I . \
  --grpc-gateway_out ./gen/go/hello \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt standalone=true \
  --grpc-gateway_opt grpc_api_configuration=helloapis/hello/urls.yaml \
	helloapis/hello/v1/hello.proto
	
protoc -I . \
  --grpc-gateway_out ./gen/go/gateway \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt standalone=true \
  --grpc-gateway_opt grpc_api_configuration=gatewayapis/gateway/urls.yaml \
	gatewayapis/gateway/v1/gateway.proto
```



## Swagger

### 安装插件

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

### 添加配置

在`buf.gen.yaml`中添加配置

```yaml
version: v1
plugins:
	# ...other plugins  
  - name: openapiv2
    out: gen/swagger
   #如果你能正常生成url映射文件，可以添加如下option来生成swagger
	 # opt:
      # - grpc_api_configuration=helloapis/hello/urls.yaml
      # - grpc_api_configuration=gatewayapis/gateway/urls.yaml
```

采用默认路由mapping来生成swagger:

```yaml
version: v1
plugins:
  - name: go
    out: gen/go
    opt: paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false

  - name: js
    out: gen/js
    opt: import_style=commonjs,binary
  - name: js-grpc
    out: gen/js
    opt: import_style=commonjs,binary
    path: /usr/local/bin/grpc_tools_node_protoc_plugin

  - name: grpc-gateway
    out: gen/go
    opt:
      - paths=source_relative
      - logtostderr=true
      - generate_unbound_methods=true
  
  - name: openapiv2
    out: gen/swagger
    strategy: all
    opt: allow_merge=true,merge_file_name=apis
```

生成文件

```shell
$ buf generate
../protos
├── buf.gen.yaml
├── buf.work.yaml
├── gatewayapis
│   ├── buf.lock
│   ├── buf.yaml
│   └── gateway
│       ├── urls.yaml
│       └── v1
│           └── gateway.proto
├── gen
│   ├── go
│   │   ├── gateway
│   │   │   └── v1
│   │   │       ├── gateway.pb.go
│   │   │       ├── gateway.pb.gw.go
│   │   │       └── gateway_grpc.pb.go
│   │   └── hello
│   │       └── v1
│   │           ├── hello.pb.go
│   │           ├── hello.pb.gw.go
│   │           └── hello_grpc.pb.go
│   ├── js
│   │   ├── gateway
│   │   │   └── v1
│   │   │       ├── gateway_grpc_pb.js
│   │   │       └── gateway_pb.js
│   │   └── hello
│   │       └── v1
│   │           ├── hello_grpc_pb.js
│   │           └── hello_pb.js
│   └── swagger
│       └── apis.swagger.json
└── helloapis
    ├── buf.lock
    ├── buf.yaml
    └── hello
        ├── urls.yaml
        └── v1
            └── hello.proto
```

### 开启swagger服务

#### 拷贝`swagger-ui`

从`https://github.com/swagger-api/swagger-ui/tree/master/dist`拷贝`dist`目录并重命名为`swagger`，将上述生成的`apis.swagger.json`文件拷贝至`swagger`目录中，然后 在其中的`index.html`文件中，修改url的值为`apis.swagger.json`：

```html
<script>
      window.onload = function () {
        // Begin Swagger UI call region
        const ui = SwaggerUIBundle({
          url: 'apis.swagger.json',
          dom_id: '#swagger-ui',
          deepLinking: true,
          presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
          plugins: [SwaggerUIBundle.plugins.DownloadUrl],
          layout: 'StandaloneLayout',
        });
        // End Swagger UI call region

        window.ui = ui;
      };
</script>
```

#### 启动swagger文件服务

在服务端代码中添加如下代码段:

```go
// 自定义路由 swagger
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	fs := http.FileServer(http.Dir("swagger/"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	mux.Handle("/docs", http.RedirectHandler("/swagger", http.StatusFound))

	// 开启http网关服务
	gwServer := &http.Server{
		Addr: ":5201",
		// Handler: gwMux,
		Handler: mux,
	}
	gwServer.ListenAndServe()
```

启动服务后，浏览器打开`http://localhost:5201/swagger`，即可看到swagger接口文档UI。

# validate

## 安装插件

```shell
$ go get github.com/envoyproxy/protoc-gen-validate
```

## 添加校验规则

各种类型详细的校验规则参考`grpc-validate.md`, 官方地址(https://github.com/envoyproxy/protoc-gen-validate)

```protobuf
// gatewayapis/gateway/v1/gateway.proto
syntax = "proto3";
package gateway.v1;

import "google/api/annotations.proto";
import "validate/validate.proto";

service ProbeService {
    rpc Ping (PingRequest) returns (PingResponse){   
        option (google.api.http) = {
            post: "/v1/probe/ping"
            body:"*"
        };
    }
    rpc Detect (DetectRequest) returns (DetectResponse);
}

message PingRequest {
    string msg = 1 [(validate.rules).string.max_len=5]; //
}
message PingResponse {
    string msg = 1 [(validate.rules).string={min_len: 5, max_len: 10}]; //
}

message DetectRequest {
    int32 id=1 [(validate.rules).int32.gte=10]; //
}

message DetectResponse {
    int32 id=1 [(validate.rules).int32={gte: 5, lt: 10}]; //
}
```

## 添加生成规则

```yaml
# protos/buf.gen.yaml
version: v1
plugins:
	# others...
  # 添加下述内容
  - name: validate
    out: gen/go
    opt:
      - lang=go
      - paths=source_relative
```

```shell
$ buf generate --exclude-path thirdapis/openapiv2,thirdapis/google,thirdapis/validate;
protos
├── buf.gen.yaml
├── buf.work.yaml
├── gatewayapis
│   ├── buf.lock
│   ├── buf.yaml
│   └── gateway
│       ├── urls.yaml
│       └── v1
│           └── gateway.proto
├── gen
│   ├── go
│   │   ├── gateway
│   │   │   └── v1
│   │   │       ├── gateway.pb.go # type define
│   │   │       ├── gateway.pb.gw.go # gateway
│   │   │       ├── gateway.pb.validate.go # validate
│   │   │       └── gateway_grpc.pb.go # grpc method
...
│   ├── js
└── thirdapis
```

## 代码参数校验

在代码方法实现中，入参和出参具有了`.Validate()、ValidateAll()`两个方法：

```go
func (h *ProbeService) Detect(ctx context.Context, in *gatewaypb.DetectRequest) (*gatewaypb.DetectResponse, error) {
	// validate request
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}
	// resp
	resp := &gatewaypb.DetectResponse{
		Id: fmt.Sprintf("req.id=%d", in.Id),
	}
	// validate response
	if err := resp.ValidateAll(); err != nil {
		return nil, err
	}
	return resp, nil
}
```

# grpc-middleware

## 安装middleware

```shell
go get github.com/grpc-ecosystem/go-grpc-middleware
```

## jwt

```shell
go get github.com/golang-jwt/jwt
```

## zipkin

```shell
go get -u github.com/openzipkin/zipkin-go-opentracing
```

## jaeger

```shell
go get -u 	"github.com/opentracing/opentracing-go"
go get -u 	"github.com/uber/jaeger-client-go"
go get github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc
```

```shell
docker run -d --name jaeger \       
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
  
# UI
http://localhost:16686/
```

