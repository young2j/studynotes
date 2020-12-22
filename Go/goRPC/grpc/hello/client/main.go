package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "grpc-notes/proto/hello"
)


func main() {
	conn, err := grpc.Dial("127.0.0.1:5200", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("conn err: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	in := &pb.HelloRequest{
		Name: "王德发",
		Age: 3,
		Money: 5000000,
	}
	resp, err := client.HelloMethod(context.Background(), in)
	if err != nil {
		grpclog.Fatalln(err)
	}
	// grpclog.Println("resp:", resp.String())
	fmt.Println("resp:", resp)
	fmt.Println("resp string:", resp.String())
	fmt.Println("resp code:", resp.Code)
	fmt.Println("resp data:", resp.GetData())
}