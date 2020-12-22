package middleware

import "github.com/gogf/gf/net/ghttp"


// CORS 跨域
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}