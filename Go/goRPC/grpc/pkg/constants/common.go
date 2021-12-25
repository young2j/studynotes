/*
 * File: common.go
 * Created Date: 2021-12-25 12:31:04
 * Author: ysj
 * Description:  这里定义公用常量、变量
 */
package constants

var (
	// jwt签名的key
	// openssl ecparam -genkey -name secp521r1 -out private.pem #生成私钥
	// openssl ec -in private.pem -pubout -out public.pem #生成公钥
	SIGN_KEY = []byte(`MIHcAgEBBEIBQyiN2WKJ6CYb8047sCW4uU9IvZhZb1UDxyyNVcyziGxW9LcXGCBo
Lrm0PPv0qMM+05riaUknuq0f1idvR+A4OvCgBwYFK4EEACOhgYkDgYYABAE3RWt5
P3T9lqF53wVef6d6sKeYXwrqfCAeAxphirBGSce2kgHvAxMQ4jm24GR6AQGUiz71
W3Dgx5ShgPatfpWC0wFtGLtzu2GgalSUJJLAxqi+Lx5MugQqDOibeg3uxoPdOM6G
wVnoE49pvoeAUU1z58ng/URn+gqrBgtBA61ITu3/AA==`)
)
