package controllers

import (
	"FileReport/common"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
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
func (c *MainController) GetView() {
	str := c.GetString("date")
	println(str)
	test, _ := models.TestView()
	logs.Info(test)
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	c.Data["json"] = &test
	c.ServeJSON()
}

func (c *MainController) UploadFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		c.Data["json"] = string(common.Marshal(resJson))
		c.ServeJSON()
	}()
	var UploadFileRequestKey = new(FileInfo)
	res := c.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &UploadFileRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	/*
		file := new(entity.FileInfo)
		file.FlieName = "差旅演示"
		file.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		file.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
		file.Creater = "zxh"
		file.Modifier = "zxh"
		file.MajorVersion = 1
		file.MinorVersion = 1
		file.ProductName = "cl"*/
	result, err := models.Upload(
		UploadFileRequestKey.FlieName,
		UploadFileRequestKey.ProductName,
		UploadFileRequestKey.MajorUpdate,
		UploadFileRequestKey.Creater)
	if err != nil {
		c.Data["json"] = result
	}
	c.ServeJSON()

}
