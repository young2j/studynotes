package router

import (
    "gf-app/app/api/chat"
    "gf-app/app/api/curd"
    "gf-app/app/api/user"
    "gf-app/app/service/middleware"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
    s := g.Server()

    // 某些浏览器直接请求favicon.ico文件，特别是产生404时
    s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

    // 分组路由注册方式
    s.Group("/", func(group *ghttp.RouterGroup) {
        ctlChat := new(chat.C)
				ctlUser := new(user.C)
				// 所有接口均绑定了middleware.CORS允许跨域请求
				group.Middleware(middleware.CORS)
				// ALL使得该路由可以被任意的HTTP Method访问
        group.ALL("/chat", ctlChat)
        g9roup.ALL("/user", ctlUser)
        group.ALL("/curd/:table", new(curd.C))
        group.Group("/", func(group *ghttp.RouterGroup) {
						// 只有/user/profile路由需要鉴权控制。
            group.Middleware(middleware.Auth)
            group.ALL("/user/profile", ctlUser, "Profile")
        })
    })
}