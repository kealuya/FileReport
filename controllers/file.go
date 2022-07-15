package controllers

import (
	"FileReport/common"
	"FileReport/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type FlieController struct {
	beego.Controller
}

func (fCtrl *FlieController) GetProductHeader() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()

	}()
	result, err := models.GetProductHeader()
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取项目状态Header失败 : %s", err.Error())
		return
	}
	resJson.Data = result
}
func (fCtrl *FlieController) GetRecentUpdate() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()

	}()
	
	filelist, err := models.GetRecentUpdate()
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最近更新文档失败 : %s", err.Error())
		return
	}
	resJson.Data = filelist
}
