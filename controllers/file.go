package controllers

import (
	"FileReport/common"
	"FileReport/models"
	"encoding/json"
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
func (fCtrl *FlieController) UploadFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()
	}()
	var UploadFileRequestKey = new(FileInfo)
	res := fCtrl.Ctx.Input.RequestBody
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
		UploadFileRequestKey.Userid)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", err.Error())
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) AbolishFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()
	}()
	var AbolishFileRequestKey = new(FileInfo)
	res := fCtrl.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &AbolishFileRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	result := models.AbolishFile(
		AbolishFileRequestKey.FlieName,
		AbolishFileRequestKey.ProductName,
		AbolishFileRequestKey.Userid)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", err.Error())
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) PublishFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()
	}()
	var PublishFileRequestKey = new(FileInfo)
	res := fCtrl.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &PublishFileRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	result := models.PublishFile(
		PublishFileRequestKey.FlieName,
		PublishFileRequestKey.ProductName,
		PublishFileRequestKey.Userid)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", err.Error())
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) UpdateFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()
	}()
	var UpdateFileRequestKey = new(FileInfo)
	res := fCtrl.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &UpdateFileRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	version := UpdateFileRequestKey.Version
	major := common.StringToInt64(version[1:2])
	minor := common.StringToInt64(version[3:5])
	result := models.UpdateFile(
		UpdateFileRequestKey.FlieName,
		UpdateFileRequestKey.ProductName,
		UpdateFileRequestKey.Userid,
		major,
		minor)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", err.Error())
		return
	}
	resJson.Data = result

}
