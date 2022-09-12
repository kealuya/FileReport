package routers

import (
	"FileReport/conf"
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
	// beego log、db、缓存 初始化
	// 如果在main方法中通过init方式执行，log会不起作用，故明确调用
	conf.InitConf()

	namespace :=
		beego.NewNamespace("/v1",

			beego.NSNamespace("/file",
				//beego.NSBefore(AuthFilter), //拦截器，暂时关闭
				beego.NSRouter("/updateFile", &controllers.FileController{}, "post:UpdateFile"),
				beego.NSRouter("/uploadFile", &controllers.FileController{}, "post:UploadFile"),
				beego.NSRouter("/fileAuthority", &controllers.FileController{}, "post:FileAuthority"),
				beego.NSRouter("/getCurrentHeader", &controllers.FileController{}, "post:GetCurrentHeader"),
				beego.NSRouter("/getProductHeader", &controllers.FileController{}, "post:GetProductHeader"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/publishFile", &controllers.FileController{}, "post:PublishFile"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/abolishFile", &controllers.FileController{}, "post:AbolishFile"),
			),
			//beego.NSNamespace("/file",
			//	beego.NSRouter("/authorityFile", &controllers.FileController{}, "post:AuthorityFile"),
			//),
			beego.NSNamespace("/file",
				beego.NSRouter("/getRecentUpdate", &controllers.FileController{}, "post:GetRecentUpdate"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/getLatestTrend", &controllers.FileController{}, "post:GetLatestTrend"),
			),
			beego.NSNamespace("/file",
				beego.NSRouter("/getActiveProduct", &controllers.FileController{}, "post:GetActiveProduct"),
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
			beego.NSRouter("/newDoc", &controllers.DocController{}, "post:NewDoc"),
			beego.NSRouter("/myFile", &controllers.FileController{}, "post:MyFile"),
			beego.NSRouter("/fileHistory", &controllers.FileController{}, "post:FileHistory"),
		)
	beego.AddNamespace(namespace)

}
