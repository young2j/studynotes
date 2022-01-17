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
	enforcer, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(enforcer.GetAdapter())

	// 权限校验
	sub := "alice"                                // 想要访问资源的用户。
	obj := "data1"                                // 将被访问的资源。
	act := "read"                                 // 用户对资源执行的操作。
	addOK, _ := enforcer.AddPolicy(sub, obj, act) // 添加策略，允许alice read data1
	fmt.Println(addOK)

	ok, _ := enforcer.Enforce(sub, obj, act)
	if ok {
		fmt.Println("允许alice读取data1")
	} else {
		fmt.Println("拒绝请求，抛出异常")
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如 results[0] 是 {"alice", "data1", "read"} 的结果
	results, _ := enforcer.BatchEnforce([][]interface{}{
		{"alice", "data1", "read"},
		{"bob", "data2", "write"},
		{"jack", "data3", "read"},
	})

	fmt.Println(results)
}
