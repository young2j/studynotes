package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ratelimit "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc_otgrpc "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	"grpc-notes/pkg/apis/gateway"
	"grpc-notes/pkg/apis/hello"
	"grpc-notes/pkg/apis/user/jwt"

	// "grpc-notes/pkg/middleware/cred"
	"grpc-notes/pkg/middleware/auth"
	"grpc-notes/pkg/middleware/opentracing"
	"grpc-notes/pkg/middleware/ratelimit"
	"grpc-notes/pkg/middleware/recovery"
	"grpc-notes/pkg/middleware/zap"
)

// 开启grpc服务
func startGrpcServer() {
	// 开启tcp监听
	listener, err := net.Listen("tcp", ":5200")
	if err != nil {
		grpclog.Fatalf("listen failed. err: %v", err)
	}

	// 开启grpc服务, 添加各种拦截器/中间价
	server := grpc.NewServer(
		// cred.TLSInterceptor(), // tls
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(auth.AuthFunc),                           // 用户验证
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),      //错误恢复
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),                     // 日志
			grpc_ratelimit.StreamServerInterceptor(ratelimit.Limiter{}),                // 限流
			grpc_otgrpc.OpenTracingStreamServerInterceptor(opentracing.JaegerTracer()), // 链路追踪
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.AuthFunc),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_ratelimit.UnaryServerInterceptor(ratelimit.Limiter{}),
			grpc_otgrpc.OpenTracingServerInterceptor(opentracing.JaegerTracer()),
		)),
	)

	// 注册grpc服务
	{
		gateway.RegisterGrpcService(server)
		hello.RegisterGrpcService(server)
		jwt.RegisterGrpcService(server)
		// others...
	}

	// grpc over tcp
	go server.Serve(listener)
	fmt.Println("grpc server listen on", listener.Addr().String())
}

// 开启http网关服务
func startGatewayServer() {
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

	// http网关路由表, 将请求的方法和路径作为元信息进行传递
	gwMux := runtime.NewServeMux(
		runtime.WithMetadata(
			func(ctx context.Context, r *http.Request) metadata.MD {
				metaData := make(map[string]string)
				if method, ok := runtime.RPCMethod(ctx); ok {
					metaData["method"] = method
				}
				if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
					metaData["pattern"] = pattern
				}
				return metadata.New(metaData)
			},
		),
	)
	// 给http网关注册路由和handler
	{
		gateway.RegisterGatewayService(gwMux, conn)
		hello.RegisterGatewayService(gwMux, conn)
		jwt.RegisterGatewayService(gwMux, conn)
		// others...
	}

	// 自定义路由 swagger
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	fs := http.FileServer(http.Dir("swagger/"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	mux.Handle("/docs", http.RedirectHandler("/swagger", http.StatusFound))

	// 开启http网关服务
	gwServer := &http.Server{
		Addr: ":5201",
		// Handler: gwMux,
		Handler: mux,
	}
	fmt.Println("grpc-gateway server listen on", gwServer.Addr)
	gwServer.ListenAndServe()
}

func main() {
	startGrpcServer()
	startGatewayServer()
}
