package svc

import (
	"bookstore/model"
	"bookstore/rpc/check/internal/config"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// Model
	Model model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// Model
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
