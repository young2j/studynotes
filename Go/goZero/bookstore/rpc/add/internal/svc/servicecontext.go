package svc

import (
	"bookstore/model"
	"bookstore/rpc/add/internal/config"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// model
	Model model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// model
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
