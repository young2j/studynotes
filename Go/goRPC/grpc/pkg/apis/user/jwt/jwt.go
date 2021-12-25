/*
 * File: jwt.go
 * Created Date: 2021-12-24 11:53:34
 * Author: ysj
 * Description:
 */

package jwt

import (
	"context"

	jwtgo "github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"grpc-notes/pkg/constants"
	jwtpb "grpc-notes/protos/gen/go/jwt/v1"
)

// 服务体
type JwtService struct {
	jwtpb.UnimplementedJWTServiceServer
}

// 注册grpc服务
func RegisterGrpcService(server *grpc.Server) {
	jwtService := &JwtService{}
	jwtpb.RegisterJWTServiceServer(server, jwtService)
}

// 注册http网关服务
func RegisterGatewayService(mux *runtime.ServeMux, conn *grpc.ClientConn) {
	jwtpb.RegisterJWTServiceHandler(context.Background(), mux, conn)
}

/***************************** 实现接口方法 ****************************/
// 获取token
func (h *JwtService) GetToken(ctx context.Context, in *jwtpb.GetTokenRequest) (*jwtpb.GetTokenResponse, error) {
	// validate request
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}
	// 根据in.Ticket请求用户信息
	userInfo := UserInfo{
		Name: "王冉明",
		Id:   1,
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, NewJWTClaims(userInfo))
	signedString, err := token.SignedString(constants.SIGN_KEY)
	if err != nil {
		return nil, err
	}
	// resp
	resp := &jwtpb.GetTokenResponse{
		Token: signedString,
	}
	// validate response
	if err := resp.Validate(); err != nil {
		return nil, err
	}

	return resp, nil
}

// 刷新token
func (h *JwtService) RefreshToken(ctx context.Context, in *jwtpb.RefreshTokenRequest) (*jwtpb.RefreshTokenResponse, error) {
	// validate request
	if err := in.ValidateAll(); err != nil {
		return nil, err
	}
	// 验证旧的token
	token, err := jwtgo.ParseWithClaims(in.Token, &JWTClaims{}, func(t *jwtgo.Token) (interface{}, error) {
		return constants.SIGN_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	// 验证通过，刷新token
	token = jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, NewJWTClaims(claims.UserInfo))
	signedString, err := token.SignedString(constants.SIGN_KEY)
	if err != nil {
		return nil, err
	}
	// respone
	resp := &jwtpb.RefreshTokenResponse{
		Token: signedString,
	}
	// validate response
	if err := resp.ValidateAll(); err != nil {
		return nil, err
	}
	return resp, nil
}
