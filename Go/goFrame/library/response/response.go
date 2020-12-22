package response

import "github.com/gogf/gf/net/ghttp"

// JSONResponse 返回统一的json数据结构
type JSONResponse struct {
	Code 		int 					`json:"code"`
	Message string 				`json:"message"`
	Data 		interface{} 	`json:"data"`
}

// JSON 返回Json
func JSON(r *ghttp.Request, code int, message string, data ...interface{})  {
	responseData := interface{}(nil)
	if len(data)>0{
		responseData = data[0]
	}
	r.Response.WriteJson(JSONResponse{
		Code: code,
		Message: message,
		Data: responseData,
	})
}

// JSONExit 返回Json并退出
func JSONExit(r *ghttp.Request, code int, message string, data ...interface{}){
	JSON(r,code,message,data)
	r.Exit()
}
