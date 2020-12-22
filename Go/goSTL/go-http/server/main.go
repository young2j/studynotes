package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 一般请求处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功.....")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	w.Write([]byte("i am go http server"))
}

// 带查询参数的处理函数
func queryHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	fmt.Println(params.Get("name"))
	fmt.Println(params.Get("id"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

// post请求处理函数
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("id"))

	// 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("get request body failed, err:", err)
		return
	}
	fmt.Println(string(b))

	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
