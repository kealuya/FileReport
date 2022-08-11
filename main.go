package main

import (
	_ "FileReport/conf"
	_ "FileReport/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/", "static/dist")
	beego.Run()

}
