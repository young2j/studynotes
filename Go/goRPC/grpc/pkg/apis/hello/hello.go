/*
 * File: hello.go
 * Created Date: 2021-12-23 10:52:15
 * Author: ysj
 * Description:  
 */

package hello

import (
	"context"
	hellopb "grpc-notes/protos/gen/go/hello/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// 服务体
type HelloService struct {
	hellopb.UnimplementedHelloServiceServer
}

// 注册grpc服务
func RegisterGrpcService(server *grpc.Server) {
	helloService := &HelloService{}
	hellopb.RegisterHelloServiceServer(server, helloService)
}

// 注册http网关服务
func RegisterGatewayService(mux *runtime.ServeMux, conn *grpc.ClientConn) {
	hellopb.RegisterHelloServiceHandler(context.Background(), mux, conn)
}

/***************************** 实现接口方法 ****************************/
func (h *HelloService) SayHello(ctx context.Context, in *hellopb.SayHelloRequest) (*hellopb.SayHelloResponse, error) {
	resp := &hellopb.SayHelloResponse{
		Code: "200",
		Data: &hellopb.SayHelloResponse_Data{
			Name:  in.Name,
			Age:   in.GetAge(),
			Money: in.Money,
		},
	}
	return resp, nil
}
