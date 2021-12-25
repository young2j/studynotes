package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	gatewaypb "grpc-notes/protos/gen/go/gateway/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

type ProbeService struct {
	gatewaypb.UnimplementedProbeServiceServer
}

func (h *ProbeService) Ping(ctx context.Context, in *gatewaypb.PingRequest) (*gatewaypb.PingResponse, error) {
	resp := &gatewaypb.PingResponse{
		Msg: "pong",
	}
	fmt.Println("receive:", in.Msg)
	return resp, nil
}

func (h *ProbeService) Detect(ctx context.Context, in *gatewaypb.DetectRequest) (*gatewaypb.DetectResponse, error) {
	resp := &gatewaypb.DetectResponse{
		Id: in.Id + 1,
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
	// 注册grpc服务
	probeService := &ProbeService{}
	gatewaypb.RegisterProbeServiceServer(server, probeService)
	fmt.Println("grpc server listen on", listener.Addr().String())

	// grpc over tcp
	go server.Serve(listener)

	// 在grpc服务端口上 开启http网关代理访问
	conn, err := grpc.DialContext(
		context.Background(),
		":5200",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		grpclog.Fatalln(err)
	}

	// 给http网关注册路由和handler
	gwMux := runtime.NewServeMux()
	gatewaypb.RegisterProbeServiceHandler(context.Background(), gwMux, conn)

	// 自定义路由 swagger
	// gwMux.Handle("/docs", )
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	fs := http.FileServer(http.Dir("dist/"))
	mux.Handle("/dist/", http.StripPrefix("/dist/", fs))

	// 开启http网关服务
	gwServer := &http.Server{
		Addr: ":5201",
		// Handler: gwMux,
		Handler: mux,
	}
	fmt.Println("grpc-gateway server listen on", gwServer.Addr)
	gwServer.ListenAndServe()
}
