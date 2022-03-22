package svc

import (
	"scana/services/access/endpoint/rpc/accessservice"
	"scana/services/users/endpoint/rpc/internal/config"
	"scana/services/users/model"

	"github.com/zeromicro/go-zero/zrpc"
)
Î
type ServiceContext struct {
	Config             config.Config
	UsersCustomerModel model.UsersCustomerModel
	UsersStaffModel    model.UsersStaffModel

	AccessRpc accessservice.AccessService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		UsersCustomerModel: model.NewUsersCustomerModel(c.Mongodb.ScanaURL, "usersCustomer", c.CacheRedis),
		UsersStaffModel:    model.NewUsersStaffModel(c.Mongodb.ScanaURL, "usersStaff", c.CacheRedis),
		AccessRpc:          accessservice.NewAccessService(zrpc.MustNewClient(c.AccessRpc)),
	}
}
