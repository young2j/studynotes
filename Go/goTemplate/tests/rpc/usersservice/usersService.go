// Code generated by goctl. DO NOT EDIT!
// Source: users.proto

package usersservice

import (
	"context"

	"go-template/rpc/userspb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCustomerReq     = userspb.AddCustomerReq
	AddCustomerResp    = userspb.AddCustomerResp
	ChangeCustomerReq  = userspb.ChangeCustomerReq
	ChangeCustomerResp = userspb.ChangeCustomerResp
	CustomerInfo       = userspb.CustomerInfo
	DeleteCustomerReq  = userspb.DeleteCustomerReq
	DeleteCustomerResp = userspb.DeleteCustomerResp
	GetCustomerReq     = userspb.GetCustomerReq
	GetCustomerResp    = userspb.GetCustomerResp
	ListCustomersReq   = userspb.ListCustomersReq
	ListCustomersResp  = userspb.ListCustomersResp
	Operator           = userspb.Operator
	UpsertCustomerReq  = userspb.UpsertCustomerReq
	UpsertCustomerResp = userspb.UpsertCustomerResp

	UsersService interface {
		//  UpsertCustomer
		UpsertCustomer(ctx context.Context, in *UpsertCustomerReq, opts ...grpc.CallOption) (*UpsertCustomerResp, error)
		//  新增Customer
		AddCustomer(ctx context.Context, in *AddCustomerReq, opts ...grpc.CallOption) (*AddCustomerResp, error)
		//  删除Customer
		DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerResp, error)
		//  修改Customer
		ChangeCustomer(ctx context.Context, in *ChangeCustomerReq, opts ...grpc.CallOption) (*ChangeCustomerResp, error)
		//  查询Customer
		GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerResp, error)
		//  Customer列表
		ListCustomers(ctx context.Context, in *ListCustomersReq, opts ...grpc.CallOption) (*ListCustomersResp, error)
	}

	defaultUsersService struct {
		cli zrpc.Client
	}
)

func NewUsersService(cli zrpc.Client) UsersService {
	return &defaultUsersService{
		cli: cli,
	}
}

//  UpsertCustomer
func (m *defaultUsersService) UpsertCustomer(ctx context.Context, in *UpsertCustomerReq, opts ...grpc.CallOption) (*UpsertCustomerResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.UpsertCustomer(ctx, in, opts...)
}

//  新增Customer
func (m *defaultUsersService) AddCustomer(ctx context.Context, in *AddCustomerReq, opts ...grpc.CallOption) (*AddCustomerResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.AddCustomer(ctx, in, opts...)
}

//  删除Customer
func (m *defaultUsersService) DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.DeleteCustomer(ctx, in, opts...)
}

//  修改Customer
func (m *defaultUsersService) ChangeCustomer(ctx context.Context, in *ChangeCustomerReq, opts ...grpc.CallOption) (*ChangeCustomerResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.ChangeCustomer(ctx, in, opts...)
}

//  查询Customer
func (m *defaultUsersService) GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.GetCustomer(ctx, in, opts...)
}

//  Customer列表
func (m *defaultUsersService) ListCustomers(ctx context.Context, in *ListCustomersReq, opts ...grpc.CallOption) (*ListCustomersResp, error) {
	client := userspb.NewUsersServiceClient(m.cli.Conn())
	return client.ListCustomers(ctx, in, opts...)
}
