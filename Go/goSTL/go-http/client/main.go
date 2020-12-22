package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Get request without params
	resp, err := http.Get("http://www.topgoer.com/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("status:", resp.Status)
	fmt.Println("header:", resp.Header)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println("读取完毕：", string(buf[:n]))
		break
	}

	// Get request with query params
	uri := "http://www.topgoer.com/"
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Println("parse uri failed. err: ", err)
	}
	params := url.Values{}
	params.Set("name", "anoymous")
	params.Set("id", "1234")

	fmt.Println(u)
	u.RawQuery = params.Encode() // urlencode
	fmt.Println(u.String())

	resp, err = http.Get(u.String())
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get response failed, err:", err)
		return
	}
	fmt.Println(string(b))

	// Post request
	uri = "http://127.0.0.1:8000"
	// 表单数据 
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=anoymous&id=1"
	// resp, err = http.Post(uri, contentType, strings.NewReader(data))
	// => 等价于http.PostForm(uri,url.Values{"name":"anoymous","id":1})
	
	// JSON
	contentType := "application/json"
	data := `{"name":"anoymous","id":1}`
	resp, err = http.Post(uri, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get response failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
