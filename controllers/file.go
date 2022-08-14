package controllers

import (
	"FileReport/common"
	"FileReport/entity"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type FlieController struct {
	beego.Controller
}

func (fCtrl *FlieController) GetProductHeader() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()
	result, err_Product := models.GetProductHeader()
	if err_Product != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取项目header失败 : %s", err_Product.Error())
		logs.Error(err_Product)
		return
	}
	resJson.Data = result
}
func (fCtrl *FlieController) GetCurrentHeader() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()
	result, err_Header := models.GetCurrentHeader()
	if err_Header != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取当前header失败 : %s", err_Header.Error())
		logs.Error(err_Header)
		return
	}
	resJson.Data = result
}
func (fCtrl *FlieController) GetRecentUpdate() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()

	filelist, err_Recent := models.GetRecentUpdate()
	if err_Recent != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最近更新失败 : %s", err_Recent.Error())
		logs.Error(err_Recent)
		return
	}
	resJson.Data = filelist
}

/*func (fCtrl *FlieController) UploadFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()
	}()
	uploadFileRequestKey := FileInfo{}
	res := fCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &uploadFileRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}

	result, err_Upload := models.Upload(
		uploadFileRequestKey.FlieName,
		uploadFileRequestKey.ProductName,
		uploadFileRequestKey.MajorUpdate,
		uploadFileRequestKey.Userid)
	if err_Upload != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("上传文件失败 : %s", err_Upload.Error())
		logs.Error(err_Upload)
		return
	}
	resJson.Data = result

}*/
func (fCtrl *FlieController) AbolishFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()
	}()
	abolishFileRequestKey := FileInfo{}
	res := fCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &abolishFileRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	result, err_Abolish := models.AbolishFile(
		abolishFileRequestKey.FlieName,
		abolishFileRequestKey.ProductName,
		abolishFileRequestKey.Userid)
	if err_Abolish != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("废除文件失败 : %s", err_Abolish.Error())
		logs.Error(err_Abolish)
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) PublishFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()
	}()
	publishFileRequestKey := FileInfo{}
	res := fCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &publishFileRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	result, err_Publish := models.PublishFile(
		publishFileRequestKey.FlieName,
		publishFileRequestKey.ProductName,
		publishFileRequestKey.Userid)
	if err_Publish != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("发布文件失败 : %s", err_Publish.Error())
		logs.Error(err_Publish)
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) AuthorityFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()
	}()
	authorityFileFileRequestKey := AuthorityFile{}
	res := fCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &authorityFileFileRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	result, err_Authority := models.AuthorityFile(
		authorityFileFileRequestKey.FlieName,
		authorityFileFileRequestKey.ProductName,
		authorityFileFileRequestKey.Userid,
		authorityFileFileRequestKey.IsOwnerEdit,
		authorityFileFileRequestKey.IsRelease)
	if err_Authority != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("更改文件权限失败 : %s", err_Authority.Error())
		logs.Error(err_Authority)
		return
	}
	resJson.Data = result

}

func (fCtrl *FlieController) UpdateFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()
	}()
	updateFileRequestKey := FileInfo{}
	res := fCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &updateFileRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	version := updateFileRequestKey.Version
	major := common.StringToInt64(version[1:2])
	minor := common.StringToInt64(version[3:5])
	result, err_Update := models.UpdateFile(
		updateFileRequestKey.FlieName,
		updateFileRequestKey.ProductName,
		updateFileRequestKey.Userid,
		major,
		minor)
	if err_Update != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("更新文件失败 : %s", err_Update.Error())
		logs.Error(err_Update)
		return
	}
	resJson.Data = result

}
func (fCtrl *FlieController) GetLatestTrend() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()
	result, err_Lat := models.GetLatestTrend()
	if err_Lat != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最新动态失败 : %s", err_Lat.Error())
		logs.Error(err_Lat)
		return
	}
	resJson.Data = result
}
func (fCtrl *FlieController) GetActiveProduct() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()

	resJson.Data = "该接口需要明确是要咋处理再实现"
}
func (uCtrl *FlieController) UploadFile() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()
	doc := entity.Doc{}
	doc.IsRelease = "false"
	doc.IsOwnerEdit = "false"
	doc.DocName = "文档"
	doc.DocType = "ppt"
	doc.IsDiscard = "false"
	doc.OwnerId = "155"
	doc.CreateDate = time.Now()
	doc.ProId = 2

	file := entity.File{}
	file.FileName = "文件"
	file.Version = 2
	file.VersionShow = "v1.01"
	file.UpdateDate = time.Now()
	file.UpdateUserId = "155"
	file.UpdateContent = "初始化"

	//res := uCtrl.Ctx.Input.RequestBody
	/*	err_Unmarshal := json.Unmarshal(res, &phoneInfo)
		if err_Unmarshal != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
			logs.Error(err_Unmarshal, string(res))
			return
		}*/
	// 判断是否是既存用户
	_, err_Upload := models.Upload(doc, file)

	if err_Upload != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Upload.Error())
		logs.Error(err_Upload)
		return
	}

}
func (uCtrl *FlieController) FileAuthority() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()
	doc := entity.Doc{}
	doc.IsRelease = "false"
	doc.IsOwnerEdit = "false"
	doc.DocName = "文档"
	doc.DocType = "ppt"
	doc.IsDiscard = "true"
	doc.OwnerId = "155"
	doc.CreateDate = time.Now()
	doc.ProId = 2

	//res := uCtrl.Ctx.Input.RequestBody
	/*	err_Unmarshal := json.Unmarshal(res, &phoneInfo)
		if err_Unmarshal != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
			logs.Error(err_Unmarshal, string(res))
			return
		}*/
	// 判断是否是既存用户
	_, err_Authority := models.FileAuthority(doc)

	if err_Authority != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Authority.Error())
		logs.Error(err_Authority)
		return
	}

}
func (uCtrl *FlieController) MyFile() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	//res := uCtrl.Ctx.Input.RequestBody
	/*	err_Unmarshal := json.Unmarshal(res, &phoneInfo)
		if err_Unmarshal != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
			logs.Error(err_Unmarshal, string(res))
			return
		}*/
	// 判断是否是既存用户
	result, err_MyFile := models.MyFile()

	if err_MyFile != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_MyFile.Error())
		logs.Error(err_MyFile)
		return
	}
	resJson.Data = result
}
