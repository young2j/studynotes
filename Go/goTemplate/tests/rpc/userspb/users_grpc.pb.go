// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	// UpsertCustomer
	UpsertCustomer(ctx context.Context, in *UpsertCustomerReq, opts ...grpc.CallOption) (*UpsertCustomerResp, error)
	// 新增Customer
	AddCustomer(ctx context.Context, in *AddCustomerReq, opts ...grpc.CallOption) (*AddCustomerResp, error)
	// 删除Customer
	DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerResp, error)
	// 修改Customer
	ChangeCustomer(ctx context.Context, in *ChangeCustomerReq, opts ...grpc.CallOption) (*ChangeCustomerResp, error)
	// 查询Customer
	GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerResp, error)
	// Customer列表
	ListCustomers(ctx context.Context, in *ListCustomersReq, opts ...grpc.CallOption) (*ListCustomersResp, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) UpsertCustomer(ctx context.Context, in *UpsertCustomerReq, opts ...grpc.CallOption) (*UpsertCustomerResp, error) {
	out := new(UpsertCustomerResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/UpsertCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) AddCustomer(ctx context.Context, in *AddCustomerReq, opts ...grpc.CallOption) (*AddCustomerResp, error) {
	out := new(AddCustomerResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/AddCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*DeleteCustomerResp, error) {
	out := new(DeleteCustomerResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ChangeCustomer(ctx context.Context, in *ChangeCustomerReq, opts ...grpc.CallOption) (*ChangeCustomerResp, error) {
	out := new(ChangeCustomerResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/ChangeCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerResp, error) {
	out := new(GetCustomerResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListCustomers(ctx context.Context, in *ListCustomersReq, opts ...grpc.CallOption) (*ListCustomersResp, error) {
	out := new(ListCustomersResp)
	err := c.cc.Invoke(ctx, "/users.UsersService/ListCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	// UpsertCustomer
	UpsertCustomer(context.Context, *UpsertCustomerReq) (*UpsertCustomerResp, error)
	// 新增Customer
	AddCustomer(context.Context, *AddCustomerReq) (*AddCustomerResp, error)
	// 删除Customer
	DeleteCustomer(context.Context, *DeleteCustomerReq) (*DeleteCustomerResp, error)
	// 修改Customer
	ChangeCustomer(context.Context, *ChangeCustomerReq) (*ChangeCustomerResp, error)
	// 查询Customer
	GetCustomer(context.Context, *GetCustomerReq) (*GetCustomerResp, error)
	// Customer列表
	ListCustomers(context.Context, *ListCustomersReq) (*ListCustomersResp, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) UpsertCustomer(context.Context, *UpsertCustomerReq) (*UpsertCustomerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCustomer not implemented")
}
func (UnimplementedUsersServiceServer) AddCustomer(context.Context, *AddCustomerReq) (*AddCustomerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomer not implemented")
}
func (UnimplementedUsersServiceServer) DeleteCustomer(context.Context, *DeleteCustomerReq) (*DeleteCustomerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedUsersServiceServer) ChangeCustomer(context.Context, *ChangeCustomerReq) (*ChangeCustomerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCustomer not implemented")
}
func (UnimplementedUsersServiceServer) GetCustomer(context.Context, *GetCustomerReq) (*GetCustomerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedUsersServiceServer) ListCustomers(context.Context, *ListCustomersReq) (*ListCustomersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomers not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_UpsertCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpsertCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/UpsertCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpsertCustomer(ctx, req.(*UpsertCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_AddCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).AddCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/AddCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).AddCustomer(ctx, req.(*AddCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).DeleteCustomer(ctx, req.(*DeleteCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ChangeCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ChangeCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/ChangeCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ChangeCustomer(ctx, req.(*ChangeCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetCustomer(ctx, req.(*GetCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ListCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/ListCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ListCustomers(ctx, req.(*ListCustomersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertCustomer",
			Handler:    _UsersService_UpsertCustomer_Handler,
		},
		{
			MethodName: "AddCustomer",
			Handler:    _UsersService_AddCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _UsersService_DeleteCustomer_Handler,
		},
		{
			MethodName: "ChangeCustomer",
			Handler:    _UsersService_ChangeCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _UsersService_GetCustomer_Handler,
		},
		{
			MethodName: "ListCustomers",
			Handler:    _UsersService_ListCustomers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}