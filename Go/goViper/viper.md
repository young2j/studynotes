# 安装

```shell
go get github.com/spf13/viper
```

# 配置读取的优先级

- explicit call to `Set`
- flag
- env
- config
- key/value store
- default

```go
// 一个viper实例只支持一个配置文件
```

