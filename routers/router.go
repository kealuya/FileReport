package routers

import (
	"FileReport/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	/*namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/test",
				beego.NSRouter("/Test", &TestController{}, "get:Test"),
			),)
	beego.AddNamespace(namespace)*/
}
