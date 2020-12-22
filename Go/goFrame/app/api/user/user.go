package user

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"gf-app/app/service/user"
	"gf-app/library/response"
)

// C 用户API管理对象
type C struct{}

// SignUp :
// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   passport  formData string  true "用户账号名称"
// @param   password  formData string  true "用户密码"
// @param   password2 formData string  true "确认密码"
// @param   nickname  formData string false "用户昵称"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) SignUp(r *ghttp.Request) {
	var (
		data        *SignUpRequest
		signUpParam *user.SignUpParam
	)
	// 结构体转换可以使用GetStruct或者Parse方法，其中Parse同时可以执行数据校验。
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &signUpParam); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	if err := user.SignUp(signUpParam); err != nil {
		response.JSONExit(r, 1, err.Error())
	} else {
		response.JSONExit(r, 0, "ok")
	}
}


// SignIn :
// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   passport formData string true "用户账号"
// @param   password formData string true "用户密码"
// @router  /user/signin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) SignIn(r *ghttp.Request) {
	var (
		data *SignInRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
		response.JSONExit(r, 1, err.Error())
	} else {
		response.JSONExit(r, 0, "ok")
	}
}

// IsSignedIn :
// @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /user/issignedin [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (c *C) IsSignedIn(r *ghttp.Request) {
	response.JSONExit(r, 0, "", user.IsSignedIn(r.Session))
}


// SignOut :
// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /user/signout [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (c *C) SignOut(r *ghttp.Request) {
	if err := user.SignOut(r.Session); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	response.JSONExit(r, 0, "ok")
}



// CheckPassport  :
// @summary 检测用户账号接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   passport query string true "用户账号"
// @router  /user/checkpassport [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (c *C) CheckPassport(r *ghttp.Request) {
	var (
		data *CheckPassportRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	if data.Passport != "" && !user.CheckPassport(data.Passport) {
		response.JSONExit(r, 1, "账号已经存在", false)
	}
	response.JSONExit(r, 0, "", true)
}

// CheckNickName :
// @summary 检测用户昵称接口(唯一性校验)
// @tags    用户服务
// @produce json
// @param   nickname query string true "用户昵称"
// @router  /user/checknickname [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *C) CheckNickName(r *ghttp.Request) {
	var (
		data *CheckNickNameRequest
	)
	if err := r.Parse(&data); err != nil {
		response.JSONExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !user.CheckNickName(data.Nickname) {
		response.JSONExit(r, 1, "昵称已经存在", false)
	}
	response.JSONExit(r, 0, "ok", true)
}


// Profile :
// @summary 获取用户详情信息
// @tags    用户服务
// @produce json
// @router  /user/profile [GET]
// @success 200 {object} user.Entity "用户信息"
func (c *C) Profile(r *ghttp.Request) {
	response.JSONExit(r, 0, "", user.GetProfile(r.Session))
}
