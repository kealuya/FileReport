package routers

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/controllers"
	"FileReport/entity"
	"FileReport/models"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"time"
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
var FilterLog = func(ctx *context.Context) {
	url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
	params, _ := json.Marshal(ctx.Input.RequestBody)
	outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
	token := ctx.Request.Header.Get("token")
	divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
	topDivider := "┌" + divider
	middleDivider := "├" + divider
	bottomDivider := "└" + divider
	outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(params) + "\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
	logs.Info(outputStr)
	rec := entity.Record{
		Token:      token,
		RouteName:  string(url),
		CreateTime: common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS)}
	models.SaveRecord(rec)
}

func init() {
	// beego log、db、缓存 初始化
	// 如果在main方法中通过init方式执行，log会不起作用，故明确调用
	conf.InitConf()
	//beego.InsertFilter("/*", beego.FinishRouter, FilterLog, beego.WithReturnOnOutput(false))
	namespace :=
		beego.NewNamespace("/v1",

			beego.NSNamespace("/file",
				//beego.NSBefore(AuthFilter), //拦截器，暂时关闭
				beego.NSRouter("/updateFile", &controllers.FileController{}, "post:UpdateFile"),
				beego.NSRouter("/uploadFile", &controllers.FileController{}, "post:UploadFile"),
				beego.NSRouter("/fileAuthority", &controllers.FileController{}, "post:FileAuthority"),
				beego.NSRouter("/getCurrentHeader", &controllers.FileController{}, "post:GetCurrentHeader"),
				beego.NSRouter("/getProductHeader", &controllers.FileController{}, "post:GetProductHeader"),
				beego.NSRouter("/getRecentUpdate", &controllers.FileController{}, "post:GetRecentUpdate"),
				beego.NSRouter("/newProject", &controllers.ProController{}, "post:NewProject"),
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
			beego.NSRouter("/getProjectList", &controllers.ProController{}, "post:GetProjectList"),
			beego.NSRouter("/getDocFileList", &controllers.DocController{}, "post:GetDocFileList"),
			beego.NSRouter("/loginFr", &controllers.UserController{}, "post:LoginFr"),
			beego.NSRouter("/newDoc", &controllers.DocController{}, "post:NewDoc"),
			beego.NSRouter("/updateDoc", &controllers.DocController{}, "post:UpdateDoc"),
			beego.NSRouter("/updateAuthorityDoc", &controllers.DocController{}, "post:UpdateAuthorityDoc"),
			beego.NSRouter("/myFile", &controllers.FileController{}, "post:MyFile"),
			beego.NSRouter("/fileHistory", &controllers.FileController{}, "post:FileHistory"),
		)
	beego.AddNamespace(namespace)

}
