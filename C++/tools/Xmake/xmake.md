<https://xmake.io/#/zh-cn/guide>

# 安装

```shell
# macos
brew install xmake

# ubuntu
sudo add-apt-repository ppa:xmake-io/xmake
sudo apt update
sudo apt install xmake
```

# 创建工程

```shell
# default binary
xmake create -l c++ -P ./hello -t [target] [target_name]

# 创建静态库
xmake create -l c++ -P ./hello_static -t static [target_name]

# 创建动态库
xmake create -l c++ -P ./hello_shared -t shared [target_name]
```

# 构建工程

```shell
xmake

# 构建为wasm
xmake f -p wasm
xmake
```

# 运行程序

```shell
xmake run|r hello
```

# 调试程序

```shell
xmake config|f -m debug 
xmake
xmake run -d hello

# 指定调试器
xmake f --debugger=gdb
xmake run -d hello
```

# 清理构建

```shell
xmake clean -a [target]
```

# 查看支持的构建策略

```shell
xmake l core.project.policy.policies
```