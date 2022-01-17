/*
 * File: main.go
 * Created Date: 2022-01-17 10:51:16
 * Author: ysj
 * Description:  casbin 基本使用
 */

package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	// 默认fileAdapter
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}
	// 在后缀将参数传入NewEnforceContext，例如2或3，它将创建 r2,p2,等。
	enforceContext := casbin.NewEnforceContext("2")
	// You can also specify a certain type individually
	enforceContext.EType = "e"
	// Don't pass in EnforceContext,the default is r,p,e,m
	e.Enforce("alice", "data2", "read") // true
	// pass in EnforceContext
	OK, _ := e.Enforce(enforceContext, struct{ Age int }{Age: 70}, "/data1", "read") //false
	fmt.Println(OK)
	OK, _ = e.Enforce(enforceContext, struct{ Age int }{Age: 30}, "/data1", "read") //true
	fmt.Println(OK)

}
