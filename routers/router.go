package routers

import (
	"FileReport/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/file",
				beego.NSRouter("/getCurrentHeader", &controllers.FlieController{}, "post:GetCurrentHeader"),
			),
			//beego.NSNamespace("/file",
			//	beego.NSRouter("/uploadFile", &controllers.FlieController{}, "post:UploadFile"),
			//),
			//beego.NSNamespace("/file",
			//	beego.NSRouter("/updateFile", &controllers.FlieController{}, "post:UpdateFile"),
			//),
			beego.NSNamespace("/file",
				beego.NSRouter("/publishFile", &controllers.FlieController{}, "post:PublishFile"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/abolishFile", &controllers.FlieController{}, "post:AbolishFile"),
			),
			//beego.NSNamespace("/file",
			//	beego.NSRouter("/authorityFile", &controllers.FlieController{}, "post:AuthorityFile"),
			//),
			beego.NSNamespace("/file",
				beego.NSRouter("/getProductHeader", &controllers.FlieController{}, "post:GetProductHeader"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/getRecentUpdate", &controllers.FlieController{}, "post:GetRecentUpdate"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/getLatestTrend", &controllers.FlieController{}, "post:GetLatestTrend"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/getActiveProduct", &controllers.FlieController{}, "post:GetActiveProduct"),
			),
			beego.NSNamespace("/user",
				beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			),
			beego.NSNamespace("/user",
				beego.NSRouter("/addPeople", &controllers.UserController{}, "post:AddPeople"),
			),
			beego.NSNamespace("/user",
				beego.NSRouter("/editPeople", &controllers.UserController{}, "post:EditPeople"),
			),
			// =========================
			beego.NSRouter("/getCaptcha", &controllers.UserController{}, "post:GetCaptcha"),
			beego.NSRouter("/loginFr", &controllers.UserController{}, "post:LoginFr"),
			beego.NSRouter("/uploadFile", &controllers.FlieController{}, "post:UploadFile"),
			beego.NSRouter("/fileAuthority", &controllers.FlieController{}, "post:FileAuthority"),
			beego.NSRouter("/myFile", &controllers.FlieController{}, "post:MyFile"),
			beego.NSRouter("/updateFile", &controllers.FlieController{}, "post:UpdateFile"),
		)
	beego.AddNamespace(namespace)

	/*		beego.Router("/", &controllers.MainController{})
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
	beego.Router("/user/editPeople", &controllers.UserController{}, "post:EditPeople")*/
	/*namespace :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/test",
				beego.NSRouter("/Test", &TestController{}, "get:Test"),
			),)
	beego.AddNamespace(namespace)*/
}
