package controllers

import (
	"FileReport/common"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	println("先走这个")
}
func (c *MainController) Finish() {
	println("调完接口走这个")
	println(beego.URLFor("MainController.GetLog"))
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Data["json"] = "看到我，就说明你这玩意调成功了\nsdasdsad\n"
	//c.ServeJSON()

}
func (c *MainController) GetLog() {
	str := c.GetString("date")
	println(str)
	/*	test, _ := models.Test()
		logs.Info(test)
		//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
		c.Data["json"] = &test*/
	c.ServeJSON()
}
func TokenCheck(funcname, token string) bool {

	if common.ArrHasStr(common.Token_Funcs, funcname) {
		//TODO
		return true
	}
	return false
}
