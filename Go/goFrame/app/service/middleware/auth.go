package middleware

import (
	"net/http"
	"github.com/gogf/gf/net/ghttp"
	"gf-app/app/service/user"
)

// Auth 只有登录后才能通过请求
func Auth(r *ghttp.Request)  {
	if user.IsSignedIn(r.Session) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}