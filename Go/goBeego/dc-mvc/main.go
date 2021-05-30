package main

import (
	_ "dc-mvc/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

