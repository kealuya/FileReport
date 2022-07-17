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
	result := models.GetProductHeader()
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取项目状态Header失败 : %s", result.Msg)
		return
	}
	resJson.Data = result
}
func (fCtrl *FlieController) GetHeader() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()

	}()
	result := models.GetHeader()
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取项当前状态Header失败 : %s", result.Msg)
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

	filelist := models.GetRecentUpdate()
	if !filelist.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最近更新文档失败 : %s", filelist.Msg)
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

	result := models.Upload(
		UploadFileRequestKey.FlieName,
		UploadFileRequestKey.ProductName,
		UploadFileRequestKey.MajorUpdate,
		UploadFileRequestKey.Userid)
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("上传文档记录失败 : %s", result.Msg)
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
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("废除文档失败 : %s", result.Msg)
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
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("发布文档失败 : %s", result.Msg)
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
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", result.Msg)
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) GetLatestTrend() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = string(common.Marshal(resJson))
		fCtrl.ServeJSON()

	}()
	result := models.GetLatestTrend()
	if !result.Success {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最新动态失败 : %s", result.Msg)
		return
	}
	resJson.Data = result
}
