package controllers

import (
	"FileReport/common"
	"FileReport/entity"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gohouse/t"
	"time"
)

type DocController struct {
	beego.Controller
}
type DocFile struct {
	DocType       string `json:"docType"`
	VersionShow   string `json:"versionShow"`
	DocName       string `json:"docName"`
	FileName      string `json:"fileName"`
	OwnerId       string `json:"ownerId"`
	DocId         string `json:"docId"`
	ProId         string `json:"proId"`
	CreateDate    string `json:"createDate"`
	IsDiscard     string `json:"isDiscard"`
	IsOwnerEdit   string `json:"isOwnerEdit"`
	IsRelease     string `json:"isRelease"`
	Owner         string `json:"owner"`
	UpdateContent string `json:"updateContent"`
	UpdateDate    string `json:"updateDate"`
	UpdateUser    string `json:"updateUser"`
}

func (uCtrl *DocController) NewDoc() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	res := uCtrl.Ctx.Input.RequestBody
	docFile := DocFile{}
	err_Unmarshal := json.Unmarshal(res, &docFile)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}

	//

	// Doc 处理
	/*
		fileName , docName , docType , ownerId ,proId 传过来
	*/

	doc := entity.Doc{
		DocName:     docFile.DocName,
		OwnerId:     docFile.OwnerId,
		CreateDate:  common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS),
		DocType:     docFile.DocType,
		IsRelease:   "true",
		IsOwnerEdit: "true",
		ProId:       t.New(docFile.ProId).Int(),
		IsDiscard:   "false",
	}

	docId, err_InsertDoc := models.InsertDoc(doc)
	if err_InsertDoc != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_InsertDoc.Error())
		logs.Error(err_InsertDoc, string(res))
		return
	}

	fmt.Println(docId)
	fmt.Println(docFile)

	// todo file 处理
	file := entity.File{
		DocId:         t.New(docId).Int(),
		Version:       1,
		VersionShow:   docFile.VersionShow,
		UpdateContent: "初始化",
		UpdateUserId:  docFile.OwnerId,
		UpdateDate:    common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS),
		FileName:      docFile.FileName,
	}
	err_InsertFile := models.InsertFile(file)
	if err_InsertFile != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_InsertFile.Error())
		logs.Error(err_InsertFile, string(res))
		return
	}
	return
}

func (uCtrl *DocController) UpdateDoc() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	res := uCtrl.Ctx.Input.RequestBody
	docFile := DocFile{}
	err_Unmarshal := json.Unmarshal(res, &docFile)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}
	// DocId 必备

	version, err_GetFileVersion := models.GetFileVersion(docFile.DocId)
	if err_GetFileVersion != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_GetFileVersion.Error())
		logs.Error(err_GetFileVersion, string(res))
		return
	}
	updateVersion := t.New(version).Int() + 1

	file := entity.File{
		DocId:         t.New(docFile.DocId).Int(),
		Version:       updateVersion,
		VersionShow:   docFile.VersionShow,
		UpdateContent: docFile.UpdateContent,
		UpdateUserId:  docFile.OwnerId,
		UpdateDate:    common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS),
		FileName:      docFile.FileName,
	}
	err_InsertFile := models.InsertFile(file)
	if err_InsertFile != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_InsertFile.Error())
		logs.Error(err_InsertFile, string(res))
		return
	}
	return
}

func (uCtrl *DocController) GetDocFileList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	res := uCtrl.Ctx.Input.RequestBody
	pi := models.PagingInfo{}
	err_Unmarshal := json.Unmarshal(res, &pi)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}

	docFiles, err_GetDocFileListByCondition := models.GetDocFileListByCondition(pi)
	if err_GetDocFileListByCondition != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_GetDocFileListByCondition.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}

	count, err_GetDocCountByProId := models.GetDocCountByProId(pi)
	if err_GetDocCountByProId != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_GetDocCountByProId.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}

	dataMap := make(map[string]any)
	dataMap["docFiles"] = docFiles
	dataMap["count"] = count

	resJson.Data = dataMap
	return
}
