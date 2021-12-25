/*
 * Created Date: 2021-07-16 10:33:41
 * Author: ysj
 * Description: '配置初始化和读取'
 */
package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

var config map[string]interface{}

func init() {
	viper.AddConfigPath("/conf")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在
			fmt.Println("配置文件不存在")
		} else {
			// 发生了其他错误
			fmt.Println("发生了其他错误, error:", err.Error())
		}
	}
	// 运行环境
	ENV := os.Getenv("BROADCAST_ENV")
	grpclog.Infof("BROADCAST_ENV: %s", ENV)
	switch ENV {
	case "prod":
		config = viper.GetStringMap("prod")
		config["mode"] = "prod"
	case "test":
		config = viper.GetStringMap("test")
		config["mode"] = "test"
	default:
		config = viper.GetStringMap("dev")
		config["mode"] = "dev"
	}

}

func Get(key string) interface{} {
	if v, ok := config[key]; ok {
		return v
	}
	return nil
}

func GetString(key string) string {
	v := Get(key)
	return v.(string)
}

func GetDefaultString(key string, value string) string {
	v := Get(key)
	if v == nil {
		return value
	}
	return v.(string)
}

func GetInt(key string) int {
	v := Get(key)
	return v.(int)
}
func GetDefaultInt(key string, value int) int {
	v := Get(key)
	if v == nil {
		return value
	}
	return v.(int)
}

func GetBool(key string) bool {
	v := Get(key)
	return v.(bool)
}

func GetStringMap(key string) map[string]interface{} {
	v := Get(key)
	return v.(map[string]interface{})
}

func GetDefaultStringMap(key string, value map[string]interface{}) map[string]interface{} {
	v := Get(key)
	if v == nil {
		return value
	}
	return v.(map[string]interface{})
}
