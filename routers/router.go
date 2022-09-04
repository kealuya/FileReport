package routers

import (
	"FileReport/controllers"
	"FileReport/models"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

var AuthFilter = func(ctx *context.Context) {
	println("进拦截器")
	token := ctx.Request.Header.Get("token")
	if token != "" {
		result := models.UserExist(token)
		if result {

		} else {
			ctx.ResponseWriter.Write([]byte("不存在改成员，或改成员已失效"))
		}
	} else {
		ctx.ResponseWriter.Write([]byte("没token"))
	}
	// 权限验证过滤器内容
}

func init() {

	namespace :=
		beego.NewNamespace("/v1",

			beego.NSNamespace("/file",
				//beego.NSBefore(AuthFilter), //拦截器，暂时关闭
				beego.NSRouter("/updateFile", &controllers.FlieController{}, "post:UpdateFile"),
				beego.NSRouter("/uploadFile", &controllers.FlieController{}, "post:UploadFile"),
				beego.NSRouter("/fileAuthority", &controllers.FlieController{}, "post:FileAuthority"),
				beego.NSRouter("/getCurrentHeader", &controllers.FlieController{}, "post:GetCurrentHeader"),
				beego.NSRouter("/getProductHeader", &controllers.FlieController{}, "post:GetProductHeader"),
			),
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
			beego.NSRouter("/myFile", &controllers.FlieController{}, "post:MyFile"),
		)
	beego.AddNamespace(namespace)

}
