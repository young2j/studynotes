/*
 * File: gateway.go
 * Created Date: 2021-12-23 09:55:31
 * Author: ysj
 * Description:
 */

package gateway

import (
	"context"
	"fmt"

	gatewaypb "grpc-notes/protos/gen/go/gateway/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// 服务体
type ProbeService struct {
	gatewaypb.UnimplementedProbeServiceServer
}

// 注册grpc服务
func RegisterGrpcService(server *grpc.Server) {
	probeService := &ProbeService{}
	gatewaypb.RegisterProbeServiceServer(server, probeService)
}

// 注册http网关服务
func RegisterGatewayService(mux *runtime.ServeMux, conn *grpc.ClientConn) {
	gatewaypb.RegisterProbeServiceHandler(context.Background(), mux, conn)
}

/***************************** 实现接口方法 ****************************/

func (h *ProbeService) Ping(ctx context.Context, in *gatewaypb.PingRequest) (*gatewaypb.PingResponse, error) {
	// validate request
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}
	// resp
	resp := &gatewaypb.PingResponse{
		Msg: "pong",
	}
	// validate response
	// if err := resp.Validate(); err != nil {
	// 	return nil, err
	// }

	fmt.Println("receive:", in.Msg)
	return resp, nil
}

func (h *ProbeService) Detect(ctx context.Context, in *gatewaypb.DetectRequest) (*gatewaypb.DetectResponse, error) {
	// validate request
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}
	// resp
	resp := &gatewaypb.DetectResponse{
		Id: in.Id,
	}
	// validate response
	if err := resp.ValidateAll(); err != nil {
		return nil, err
	}
	return resp, nil
}
