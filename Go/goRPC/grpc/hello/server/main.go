package main

import (
	"context"
	"fmt"
	"net"

	"grpc-notes/proto/hello"
	pb "grpc-notes/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type HelloService struct{
	hello.UnimplementedHelloServiceServer
}

func (h *HelloService) HelloMethod(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := &pb.HelloResponse{
		Code: "200",
		Data: &pb.HelloResponse_Data{
			Name: in.Name,
			Age: in.GetAge(),
			Money: in.Money,
		},
	}
	return resp, nil
}

func (h *HelloService) mustEmbedUnimplementedHelloServiceServer() {}

func main() {
	listener, err := net.Listen("tcp", ":5200")
	if err != nil {
		grpclog.Fatalf("listen failed. err: %v", err)
	}

	server := grpc.NewServer()
	helloService := &HelloService{}

	pb.RegisterHelloServiceServer(server, helloService)
	grpclog.Println("server listen on:", listener.Addr().String())
	fmt.Println("server listen on:", listener.Addr().String())
	server.Serve(listener)
}
