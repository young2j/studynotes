package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// rpc 通过ServiceContext在不同业务逻辑之间传递依赖
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// rpc 通过ServiceContext在不同业务逻辑之间传递依赖
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
