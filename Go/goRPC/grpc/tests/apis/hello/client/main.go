package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	hellopb "grpc-notes/protos/gen/go/hello/v1"
)

func main() {
	// grpc 连接服务器
	conn, err := grpc.Dial("127.0.0.1:5200", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("conn err: %v", err)
	}
	defer conn.Close()

	// 建立客户端
	client := hellopb.NewHelloServiceClient(conn)
	// 入参
	in := &hellopb.SayHelloRequest{
		Name:  "王德发",
		Age:   3,
		Money: 5000,
	}
	// 调用
	resp, err := client.SayHello(context.Background(), in)
	if err != nil {
		grpclog.Fatalln(err)
	}
	// 输出response
	// grpclog.Println("resp:", resp.String())
	fmt.Println("resp     :", resp)
	fmt.Println("resp.code:", resp.Code)
	fmt.Println("resp.data:", resp.GetData())
}
