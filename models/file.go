package models

import (
	"FileReport/common"
	"FileReport/db"
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type FileResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func GetProductHeader() ([]entity.ProductInfo, error) {
	o := orm.NewOrm()
	var productlist []entity.ProductInfo
	productlist, err := db.GetProductHeader(o)

	return productlist, err
}
func GetRecentUpdate() ([]entity.FileRecord, error) {
	o := orm.NewOrm()
	var productlist []entity.FileRecord
	productlist, err := db.SelectRecentUpdate(o)

	return productlist, err
}

func UpdateFile(filename, productname, userid string, ma, mi int64) (result *FileResult) {
	defer common.RecoverHandler(func(rcErr error) {
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	o := orm.NewOrm()
	to, err := o.Begin()
	file_result := new(FileResult)
	if err != nil {
		panic("****UpdateFile****start the transaction failed")
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record, err := db.SelectFileInfo(to, filename, productname)

	record.Modifier = userid
	record.ModifyTime = timenow
	record.MinorVersion = ma
	record.MinorVersion = mi
	_, err = db.UpdateFile(to, record)
	if err != nil {
		panic("****UpdateFile****更新或上传文件失败" + err.Error())
		to.Rollback()
	}

	record.Operationtype = "update"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_, err = db.Record(to, record)
	if err != nil {
		panic("****UpdateFile****添加记录失败" + err.Error())

	}
	_ = to.Commit()
	file_result.Success = true
	file_result.Msg = ""
	return file_result
}

func AbolishFile(filename, productname, userid string) (result *FileResult) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	if err != nil {
		panic("****PublishFile****start the transaction failed")
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	fileinfo := entity.FileInfo{}
	fileinfo.FileName = filename
	fileinfo.ProductName = productname
	fileinfo.Modifier = userid
	fileinfo.ModifyTime = timenow
	_, err = db.AbolishFile(to, fileinfo)
	if err != nil {
		panic("****AbolishFile****废除文件失败" + err.Error())
	}
	record, err := db.SelectFileInfo(to, filename, productname)

	record.Operationtype = "abolish"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_, err = db.Record(to, record)
	if err != nil {
		panic("****AbolishFile****添加记录失败" + err.Error())

	}
	_ = to.Commit()
	file_result.Success = true
	file_result.Msg = ""
	return file_result
}
func PublishFile(filename, productname, userid string) (result *FileResult) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	if err != nil {
		panic("****PublishFile****start the transaction failed")
	}
	fileinfo := entity.FileInfo{}
	fileinfo.FileName = filename
	fileinfo.ProductName = productname
	fileinfo.Modifier = userid
	fileinfo.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	_, err = db.PublishFile(to, fileinfo)
	if err != nil {
		panic("****PublishFile****发布文件失败" + err.Error())
	}
	record, err := db.SelectFileInfo(to, filename, productname)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record.Operationtype = "abolish"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_, err = db.Record(to, record)
	if err != nil {
		panic("****PublishFile****添加记录失败" + err.Error())
	}
	_ = to.Commit()
	file_result.Success = true
	file_result.Msg = ""
	return file_result
}
