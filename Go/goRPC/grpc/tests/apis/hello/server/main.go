package main

import (
	"context"
	"fmt"
	"net"

	hellopb "grpc-notes/protos/gen/go/hello/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type HelloService struct {
	hellopb.UnimplementedHelloServiceServer
}

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

func main() {
	// 开启tcp监听
	listener, err := net.Listen("tcp", ":5200")
	if err != nil {
		grpclog.Fatalf("listen failed. err: %v", err)
	}

	// 开启grpc服务
	server := grpc.NewServer()
	
	// 注册hello服务
	helloService := &HelloService{}
	hellopb.RegisterHelloServiceServer(server, helloService)

	grpclog.Infoln("server listen on:", listener.Addr().String())
	fmt.Println("server listen on:", listener.Addr().String())

	// grpc over tcp
	server.Serve(listener)
}
