package routers

import (
	"FileReport/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/log/getlog", &controllers.MainController{}, "get:GetLog")
	beego.Router("/file/getCurrentHeader", &controllers.FlieController{}, "post:GetCurrentHeader")
	beego.Router("/file/uploadFile", &controllers.FlieController{}, "post:UploadFile")
	beego.Router("/file/updateFile", &controllers.FlieController{}, "post:UpdateFile")
	beego.Router("/file/publishFile", &controllers.FlieController{}, "post:PublishFile")
	beego.Router("/file/abolishFile", &controllers.FlieController{}, "post:AbolishFile")
	beego.Router("/file/getProductHeader", &controllers.FlieController{}, "post:GetProductHeader")
	beego.Router("/file/getRecentUpdate", &controllers.FlieController{}, "post:GetRecentUpdate")
	beego.Router("/file/getLatestTrend", &controllers.FlieController{}, "post:GetLatestTrend")
	beego.Router("/file/getActiveProduct", &controllers.FlieController{}, "post:GetActiveProduct")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/user/addPeople", &controllers.UserController{}, "post:AddPeople")
	beego.Router("/user/editPeople", &controllers.UserController{}, "post:EditPeople")
	/*namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/test",
				beego.NSRouter("/Test", &TestController{}, "get:Test"),
			),)
	beego.AddNamespace(namespace)*/
}
