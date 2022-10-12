package controllers

import (
	"FileReport/models"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ProController struct {
	beego.Controller
}

func (uCtrl *ProController) GetProjectList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	proList, err_ObtainProjectList := models.ObtainProjectList()
	if err_ObtainProjectList != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_ObtainProjectList.Error())
		logs.Error(err_ObtainProjectList)
		return
	}
	resJson.Data = proList
	return
}
