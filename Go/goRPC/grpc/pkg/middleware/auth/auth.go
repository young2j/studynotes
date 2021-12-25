/*
 * File: auth.go
 * Created Date: 2021-12-24 06:45:09
 * Author: ysj
 * Description: jwt验证中间件
 */

package auth

import (
	"context"

	apis_jwt "grpc-notes/pkg/apis/user/jwt"
	"grpc-notes/pkg/constants"

	jwtgo "github.com/golang-jwt/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

type ctxKey string

func AuthFunc(ctx context.Context) (context.Context, error) {
	// 给获取token的请求放行
	pattern := metautils.ExtractIncoming(ctx).Get("pattern")
	if pattern == "/auth/v1/jwt/get_token" {
		return ctx, nil
	}
	// if md, ok := metadata.FromIncomingContext(ctx); ok {
	// 	fmt.Printf("%#v\n", md)
	// }

	//获取token token=header['authorization']
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, err
	}

	// 验证token是否有效，有效则给ctx注入用户信息
	token, err := jwtgo.ParseWithClaims(tokenString, &apis_jwt.JWTClaims{}, func(t *jwtgo.Token) (interface{}, error) {
		return constants.SIGN_KEY, nil
	})
	if err != nil {
		return ctx, err
	}
	claims, ok := token.Claims.(*apis_jwt.JWTClaims)
	if !(ok && token.Valid) {
		return ctx, err
	}
	key := ctxKey("user")
	ctx = context.WithValue(ctx, key, claims.UserInfo)
	return ctx, nil
}
