package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// rpc conf 通过etcd自动去发现可用的add/check服务
	Add   zrpc.RpcClientConf
	Check zrpc.RpcClientConf
}
