package routers

import (
	"FileReport/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/log/getlog", &controllers.MainController{}, "get:GetLog")
	beego.Router("/log/getview", &controllers.MainController{}, "get:GetView")
	beego.Router("/log/uploadFile", &controllers.MainController{}, "post:UploadFile")
	/*namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/test",
				beego.NSRouter("/Test", &TestController{}, "get:Test"),
			),)
	beego.AddNamespace(namespace)*/
}
