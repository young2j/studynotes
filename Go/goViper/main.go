package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在
			fmt.Println("配置文件不存在")
		} else {
			// 发生了其他错误
			fmt.Println("发生了其他错误")
		}
	}

	// 从io.Reader读取配置
	yamlExample := []byte(`
	version: 3.5
	services:
  	dev_celery:
  	  image: hub.qixincha.com/macaw:latest
  	  command: supervisord -c /code/deploy/test_celery_supervisord.conf
  	  volumes:
  	    - .:/code
  	  extra_hosts:
  	    - prerender.qixincha.com:193.112.173.62
  	  network_mode: bridge
  	  environment:
  	    - TZ=Asia/Shanghai
  	  container_name: dev_celery
  	  sysctls:
  	    net.core.somaxconn: 16384
	`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	viper.Get("services.dev_celery.image")

	// 覆写配置值
	viper.SetDefault("version", "3.0")
	viper.Set("version", "3.5")

	// work with 环境变量，大小写敏感
	viper.SetEnvPrefix("env")              // 设置环境变量前缀，将获取以ENV_为前缀的环境变量
	viper.BindEnv("VIPER", "env1", "env2") // 会使用前缀
	viper.AutomaticEnv()                   // 会使用前缀, will check for an environment variable any time a viper.Get request is made.

	// work with flags
	flag.Int("flagname", 1234, "help for flagname")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.GetInt("flagname")

	// 输出配置文件
	viper.WriteConfig()                                 // 按照定义的配置名、扩展以及路径输出配置文件，如果存在会进行覆盖
	viper.SafeWriteConfig()                             // 同上，但不会覆盖
	viper.WriteConfigAs("./config/writeConfig")         // 配置文件另存为，会进行覆盖
	viper.SafeWriteConfigAs("./config/safeWriteConfig") // 配置文件另存为，但不会进行覆盖

	// 配置监控，实时热更新
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生了变更:", e.Name, e.Op)
	})

	// 配置反序列化
	type config struct {
		Port    int
		Name    string
		PathMap string `mapstructure:"path_map"`
	}

	var C config

	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}

	// // 使用远程配置
	// // etcd
	// viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
	// viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	// viper.ReadRemoteConfig()
	// // consul
	// viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
	// viper.SetConfigType("json") // Need to explicitly set this to json
	// viper.ReadRemoteConfig()

}
