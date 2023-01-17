package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/db"
	"FileReport/db/handler"
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type FileResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func AbolishFile(filename, productname, userid string) (result string, resultErr error) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = "fail"
		resultErr = rcErr

	})
	if err != nil {
		panic("****PublishFile****start the transaction failed")
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	fileinfo := entity.FileInfo{}
	fileinfo.FileName = filename
	fileinfo.ProductName = productname
	fileinfo.Modifier = userid
	fileinfo.ModifyTime = timenow
	_ = db.AbolishFile(to, fileinfo)

	record := db.SelectFileInfo(filename, productname)

	record.Operationtype = "abolish"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_ = db.Record(to, record)

	_ = to.Commit()

	return "success", nil
}
func PublishFile(filename, productname, userid string) (result string, resultErr error) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = "fail"
		resultErr = rcErr
	})
	if err != nil {
		panic("****PublishFile****start the transaction failed")
	}
	fileinfo := entity.FileInfo{}
	fileinfo.FileName = filename
	fileinfo.ProductName = productname
	fileinfo.Modifier = userid
	fileinfo.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	_ = db.PublishFile(to, fileinfo)

	record := db.SelectFileInfo(filename, productname)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record.Operationtype = "abolish"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_ = db.Record(to, record)

	_ = to.Commit()
	return "success", nil
}
func AuthorityFile(filename, productname, userid string, IsOwnerEdit, IsRelease bool) (result string, resultErr error) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = "fail"
		resultErr = rcErr
	})
	if err != nil {
		panic("****AuthorityFile****start the transaction failed")
	}
	fileinfo := entity.FileInfo{}
	fileinfo.FileName = filename
	fileinfo.ProductName = productname
	fileinfo.Modifier = userid
	fileinfo.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	fileinfo.EditFlag = "0"
	fileinfo.PublishFlag = "0"
	if !IsOwnerEdit {
		fileinfo.EditFlag = "1"
	}
	if IsRelease {
		fileinfo.PublishFlag = "1"
	}
	_ = db.AuthorityFile(to, fileinfo)

	record := db.SelectFileInfo(filename, productname)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record.Operationtype = "authority"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_ = db.Record(to, record)

	_ = to.Commit()
	return "success", nil
}
func GetLatestTrend() (result []entity.FileRecord, resultErr error) {
	defer common.RecoverHandler(func(rcErr error) {
		result = []entity.FileRecord{}
		resultErr = rcErr
	})

	productlist := db.GetLatestTrend()

	return productlist, nil

}
func Upload(docinfo entity.Doc, fileinfo entity.File) (result string, resultErr error) {
	session := conf.Engine.NewSession()
	defer common.RecoverHandler(func(err error) {
		session.Rollback()
		result = "false"
		resultErr = err
		return
	})

	defer session.Close()
	// add Begin() before any action
	if err_Session := session.Begin(); err_Session != nil {
		common.ErrorHandler(err_Session)
	}
	doc := new(entity.Doc)
	doc.IsRelease = docinfo.IsRelease
	doc.IsOwnerEdit = docinfo.IsOwnerEdit
	doc.DocName = docinfo.DocName
	doc.DocType = docinfo.DocType
	doc.IsDiscard = docinfo.IsDiscard
	doc.OwnerId = docinfo.OwnerId
	doc.ProId = docinfo.ProId
	_, err_InsertDoc := session.Insert(doc)
	common.ErrorHandler(err_InsertDoc)
	file := new(entity.File)
	file.DocId = doc.DocId
	file.FileName = fileinfo.FileName
	file.Version = fileinfo.Version
	file.VersionShow = fileinfo.VersionShow
	file.UpdateDate = common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS)
	file.UpdateUserId = fileinfo.UpdateUserId
	file.UpdateContent = fileinfo.UpdateContent
	_, err_InsertFile := session.Insert(file)
	common.ErrorHandler(err_InsertFile)
	session.Commit()
	return "success", nil
}
func FileAuthority(docinfo entity.Doc) (result string, resultErr error) {

	defer common.RecoverHandler(func(err error) {
		result = "false"
		resultErr = err
		return
	})

	doc := new(entity.Doc)
	doc.IsRelease = docinfo.IsRelease
	doc.IsOwnerEdit = docinfo.IsOwnerEdit
	doc.DocName = docinfo.DocName
	doc.DocType = docinfo.DocType
	doc.IsDiscard = docinfo.IsDiscard
	doc.OwnerId = docinfo.OwnerId
	doc.ProId = docinfo.ProId

	_, err_Insert := conf.Engine.Where("doc_id=?", 2).Update(doc)
	common.ErrorHandler(err_Insert)

	return "success", nil
}

type DocFile struct {
	DocType           string              `json:"docType"`
	VersionShow       string              `json:"versionShow"`
	DocName           string              `json:"docName"`
	FileName          string              `json:"fileName"`
	OwnerId           string              `json:"ownerId"`
	DocId             string              `json:"docId"`
	ProId             string              `json:"proId"`
	CreateDate        string              `json:"createDate"`
	IsDiscard         string              `json:"isDiscard"`
	IsOwnerEdit       string              `json:"isOwnerEdit"`
	IsRelease         string              `json:"isRelease"`
	Owner             string              `json:"owner"`
	UpdateContent     string              `json:"updateContent"`
	UpdateDate        string              `json:"updateDate"`
	UpdateUser        string              `json:"updateUser"`
	UpdateContentList []map[string]string `json:"updateContentList"`
}

func MyFile() (result []DocFile, resultErr error) {

	defer common.RecoverHandler(func(err error) {
		result = []DocFile{}
		resultErr = err
		return
	})

	docfile := new([]DocFile)
	err_Select := conf.Engine.SQL(handler.Select_doc_File).Find(docfile)

	common.ErrorHandler(err_Select)

	return *docfile, nil
}
func FileHistory(doc_id string) (result []DocFile, resultErr error) {

	defer common.RecoverHandler(func(err error) {
		result = []DocFile{}
		resultErr = err
		return
	})

	docfile := new([]DocFile)
	err_Select := conf.Engine.Table("doc").Join("INNER", "file", "file.doc_id=doc.doc_id").Where("file.doc_id = ?", doc_id).Find(docfile)

	common.ErrorHandler(err_Select)

	return *docfile, nil
}
func UpdateFile(docinfo entity.Doc, fileinfo entity.File) (result string, resultErr error) {
	session := conf.Engine.NewSession()
	defer common.RecoverHandler(func(err error) {
		session.Rollback()
		result = "false"
		resultErr = err
		return
	})

	defer session.Close()
	// add Begin() before any action
	if err_Session := session.Begin(); err_Session != nil {
		common.ErrorHandler(err_Session)
	}
	file := new(entity.File)
	file.DocId = fileinfo.DocId
	file.FileName = fileinfo.FileName
	file.Version = fileinfo.Version
	file.VersionShow = fileinfo.VersionShow
	file.UpdateDate = common.FormatDate(time.Now(), common.YYYY_MM_DD_HH_MM_SS)
	file.UpdateUserId = fileinfo.UpdateUserId
	file.UpdateContent = fileinfo.UpdateContent
	_, err_InsertFile := session.Insert(file)
	common.ErrorHandler(err_InsertFile)
	session.Commit()
	return "success", nil
}

type ProductStatus struct {
	ProId    int    //项目id
	ProName  string //项目名称
	ProLogo  string //项目图标
	DocNums  int    //文件数量
	PeoNums  int    //维护人数
	Versions int    //发布次数
}

func GetCurrentHeader() (result []ProductStatus, resultErr error) {

	defer common.RecoverHandler(func(err error) {
		result = []ProductStatus{}
		resultErr = err
		return
	})

	productStatus := new([]ProductStatus)

	err_Select := conf.Engine.SQL(handler.Select_Current_Header).Find(productStatus)
	common.ErrorHandler(err_Select)

	return *productStatus, nil
}

type ProjectHeader struct {
	DocId        int
	UpdateUserId string
	UpdateDate   time.Time
	UserName     string
	ProName      string
	ProId        int
	PhoneNumber  string
}

func GetProductHeader() (result []ProjectHeader, resultErr error) {

	defer common.RecoverHandler(func(rcErr error) {
		result = []ProjectHeader{}
		resultErr = rcErr
	})
	productheader := new([]ProjectHeader)

	err_Select := conf.Engine.SQL(handler.Select_Product_Header).Find(productheader)
	common.ErrorHandler(err_Select)

	return *productheader, nil
}
func GetRecentUpdate() (result []DocFile, resultErr error) {

	defer common.RecoverHandler(func(err error) {
		result = []DocFile{}
		resultErr = err
		return
	})

	docfile := new([]DocFile)

	err_Select := conf.Engine.SQL(handler.Select_Last_Month_Update).Find(docfile)
	common.ErrorHandler(err_Select)

	return *docfile, nil
}
