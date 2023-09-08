/*
 * File: main.go
 * Created Date: 2023-04-13 11:33:00
 * Author: ysj
 * Description:
 * 读不到key 对应的配置，是因为 Viper 依赖 crypt 库，
 * 而 crypt 截至目前还不支持新版 ETCD 的 API。
 *
 * 通过重新实现remoteConfigFactory接口
	* type remoteConfigFactory interface {
	*     Get(rp RemoteProvider) (io.Reader, error)
	*     Watch(rp RemoteProvider) (io.Reader, error)
	*     WatchChannel(rp RemoteProvider) (<-chan *RemoteResponse, chan bool)
	* }
	*
	* 把加解密的部分去掉
*/
package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	// clientv3 "go.etcd.io/etcd/client/v3"
	// "github.com/bketelsen/crypt/backend/etcd"
)

type Conf struct {
	AiCates []int  `json:"aiCates"`
	Limit   string `json:"limit"`
}

func main() {
	// etcdv3
	// ctx := context.TODO()
	// cli, err := clientv3.New(clientv3.Config{
	// 	Endpoints:   []string{"localhost:2379"},
	// 	DialTimeout: 5 * time.Second,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// defer cli.Close()

	// // crypt etcd
	// cli, err := etcd.New([]string{"http://127.0.0.1:2379"})
	// if err != nil {
	// 	panic(err)
	// }

	// conf := `
	// {
	// 	"aiCates": [1, 2, 3],
	// 	"limit": 100
	// }
	// `
	// // _, err = cli.Put(ctx, "/ecm/conf.json", conf)
	// err = cli.Set("/ecm/conf.json", []byte(conf))
	// if err != nil {
	// 	panic(err)
	// }

	vp := viper.New()
	vp.Debug()
	err := vp.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", "/ecm/conf.json")
	if err != nil {
		panic(err)
	}
	vp.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	err = vp.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	for range time.Tick(time.Second * 1) {
		err = vp.WatchRemoteConfig()
		if err != nil {
			panic(err)
		}
		conf := new(Conf)
		err = vp.Unmarshal(conf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("conf: %+v\n", conf)
		// settings := vp.AllSettings()
		// fmt.Printf("settings: %v\n", settings)
		// aiCates := vp.GetIntSlice("aiCates")
		// fmt.Printf("aiCates: %v\n", aiCates)
		// limit := vp.GetInt("limit")
		// fmt.Printf("limit: %v\n", limit)
	}

}
