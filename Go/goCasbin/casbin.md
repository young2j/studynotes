# 安装

```shell
go get github.com/casbin/casbin/v2
```

# 基本使用

Casbin使用配置文件来设置访问控制模式。它有两个配置文件：

* `model.conf`：存储了访问模型

* `policy.csv`：存储了特定的用户权限配置。

 基本上，我们只需要一个主要结构：**enforcer(执行器)**。 当构建这个结构时，`model.conf`和`policy.csv`将被加载。要新建一个Casbin执行器，需要提供一个[Model](https://casbin.org/docs/zh-CN/supported-models)和一个[Adapter](https://casbin.org/docs/zh-CN/adapters)。

新建一个`model.conf`和`policy.csv`文件：

```ini
# model.conf
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`
```

```go
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
// &{./policy.csv}
// true
// 允许alice读取data1
// [true false false]
```

# Model

访问控制模型被抽象为基于 **`PERM (Policy, Effect, Request, Matcher)`** 的一个文件。

`PERM(策略、效果、请求、匹配)`描述了资源与用户之间的关系。例如基本的`ACL Model`:

```ini
# model.conf
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`
```

## 请求

基本请求是一个元组对象，至少需要主题(`sub`, 访问实体)、对象(`obj`, 访问资源) 和动作(`act`, 访问方式)。

例如，一个请求可能长这样： `r={sub,obj,act}`

它定义了访问控制进行匹配的参数名称和顺序。

## 策略

它定义了在策略文件(`policy.csv`或db文件)中字段的名称和顺序。策略可以多个。

例如： `p={sub, obj, act}` 或 `p2={sub, obj, act, eft}`

> 注：如果未定义`eft`，则策略文件中的结果字段将不会被读取；且匹配的策略结果将默认被允许(eft默认等于allow)。

## 匹配器

匹配请求(`request`)和策略(`policy`)的规则。

例如， `r.sub == p.sub && r.obj == p.obj && r.act == p.act` ：意味着如果请求的参数(访问实体，访问资源和访问方式)匹配，则通过校验。

## 效果

对匹配结果再次作出逻辑组合判断。

例如， `e = some (where (p.eft == allow))`：意味着如果匹配的策略结果有一些是允许的，那么最终结果为真。

 `e = some (where (p.eft == allow)) && !some(where (p.eft == deny)` ：意味着当匹配策略均为允许（没有任何否认）时为真。


