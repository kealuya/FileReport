package controllers

import (
	"FileReport/entity"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gohouse/t"
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

type Project struct {
	ProId   int    `json:"pro_id"`
	ProName string `json:"pro_name"`
	ProLogo string `json:"pro_logo"`
}

func (uCtrl *ProController) NewProject() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	res := uCtrl.Ctx.Input.RequestBody
	project := Project{}
	err_Unmarshal := json.Unmarshal(res, &project)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}

	pro := entity.Project{
		ProName: project.ProName,
		ProLogo: project.ProLogo,
		ProId:   t.New(project.ProId).Int(),
	}

	err_InsertPro := models.InsertProject(pro)
	if err_InsertPro != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_InsertPro.Error())
		logs.Error(err_InsertPro, string(res))
		return
	}

	return
}
