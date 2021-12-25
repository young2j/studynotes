package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gatewaypb "grpc-notes/protos/gen/go/gateway/v1"
)

func main() {
	// grpc 连接服务器
	conn, err := grpc.Dial("127.0.0.1:5200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatalf("conn err: %v", err)
	}
	defer conn.Close()

	// 建立客户端
	client := gatewaypb.NewProbeServiceClient(conn)
	// 入参
	in := &gatewaypb.PingRequest{
		Msg: "ping",
	}
	// 调用
	resp, err := client.Ping(context.Background(), in)
	if err != nil {
		grpclog.Fatalln(err)
	}
	// 输出response
	fmt.Println("resp.Msg:", resp.Msg)
}
