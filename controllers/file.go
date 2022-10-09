package controllers

import (
	"FileReport/common"
	"FileReport/entity"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type FileController struct {
	beego.Controller
}

func (fCtrl *FileController) AbolishFile() {
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
func (fCtrl *FileController) PublishFile() {
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
func (fCtrl *FileController) AuthorityFile() {
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

func (fCtrl *FileController) GetLatestTrend() {
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
func (fCtrl *FileController) GetActiveProduct() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()

	resJson.Data = "该接口需要明确是要咋处理再实现"
}
func (uCtrl *FileController) UploadFile() {

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
	doc.ProId = 2

	file := entity.File{}
	file.FileName = "文件"
	file.Version = 2
	file.VersionShow = "v1.01"
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
func (uCtrl *FileController) FileAuthority() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()
	doc := entity.Doc{}
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &doc)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}
	token := uCtrl.Ctx.Request.Header.Get("token")
	if TokenCheck("fileAuthority", token, strconv.Itoa(doc.DocId)) {
	} else {
		resJson.Success = false
		resJson.Msg = "非本人不允许操作"
		return
	}

	/*	doc.IsRelease = "false"
		doc.IsOwnerEdit = "false"
		doc.DocName = "文档"
		doc.DocType = "ppt"
		doc.IsDiscard = "true"
		doc.OwnerId = "155"
		doc.ProId = 2*/

	// 判断是否是既存用户
	_, err_Authority := models.FileAuthority(doc)

	if err_Authority != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Authority.Error())
		logs.Error(err_Authority)
		return
	}

}
func (uCtrl *FileController) MyFile() {

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
func (uCtrl *FileController) FileHistory() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	fileinfo := make(map[string]any)
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &fileinfo)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}

	// 不定义struct，明确接受参数
	docId := fileinfo["docId"].(string)
	// 判断是否是既存用户
	result, err_FileHistory := models.FileHistory(docId)

	if err_FileHistory != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_FileHistory.Error())
		logs.Error(err_FileHistory)
		return
	}
	resJson.Data = result
}
func (uCtrl *FileController) UpdateFile() {

	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()
	doc := entity.Doc{}

	file := entity.File{}
	file.DocId = 8
	file.FileName = "文件"
	file.Version = 2
	file.VersionShow = "v1.01"
	file.UpdateDate = common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS)
	file.UpdateUserId = "155"
	file.UpdateContent = "更新"

	//res := uCtrl.Ctx.Input.RequestBody
	/*	err_Unmarshal := json.Unmarshal(res, &phoneInfo)
		if err_Unmarshal != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
			logs.Error(err_Unmarshal, string(res))
			return
		}*/
	// 判断是否是既存用户
	_, err_UpdateFile := models.UpdateFile(doc, file)

	if err_UpdateFile != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_UpdateFile.Error())
		logs.Error(err_UpdateFile)
		return
	}

}
func (fCtrl *FileController) GetCurrentHeader() {

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
func (fCtrl *FileController) GetProductHeader() {
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
func (fCtrl *FileController) GetRecentUpdate() {
	resJson := NewJsonStruct(nil)
	defer func() {
		fCtrl.Data["json"] = resJson
		fCtrl.ServeJSON()

	}()
	result, err_Product := models.GetRecentUpdate()
	if err_Product != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取最近更新header失败 : %s", err_Product.Error())
		logs.Error(err_Product)
		return
	}
	resJson.Data = result
}
