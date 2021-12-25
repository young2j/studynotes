/*
 * File: types.go
 * Created Date: 2021-12-25 12:13:35
 * Author: ysj
 * Description:  类型定义
 */
package jwt

import (
	"grpc-notes/pkg/utils"
	"time"

	jwtgo "github.com/golang-jwt/jwt"
)

// 用户信息
type UserInfo struct {
	Name string
	Id   int
}

// jwt 声明
type JWTClaims struct {
	UserInfo
	jwtgo.StandardClaims
}

// 新建一个jwt声明
func NewJWTClaims(userInfo UserInfo) JWTClaims {
	now := utils.GetTimeNow()
	return JWTClaims{
		userInfo,
		jwtgo.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(12 * time.Hour).Unix(),
		},
	}
}
