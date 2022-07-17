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

func GetProductHeader() (result *FileResult) {
	o := orm.NewOrm()
	defer common.RecoverHandler(func(rcErr error) {
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	productlist, err := db.GetProductHeader(o)
	if err != nil {
		panic(err)
	}
	file_result.Success = true
	file_result.Msg = ""
	file_result.Data = productlist
	return file_result
}
func GetRecentUpdate() (result *FileResult) {
	o := orm.NewOrm()
	defer common.RecoverHandler(func(rcErr error) {
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	productlist, err := db.SelectRecentUpdate(o)
	if err != nil {
		panic(err)
	}
	file_result.Success = true
	file_result.Msg = ""
	file_result.Data = productlist
	return file_result
}

func UpdateFile(filename, productname, userid string, ma, mi int64) (result *FileResult) {
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
		panic("****UpdateFile****start the transaction failed")
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record, _ := db.SelectFileInfo(o, filename, productname)

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
	record, _ := db.SelectFileInfo(o, filename, productname)

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
	record, _ := db.SelectFileInfo(o, filename, productname)
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
func Upload(filename, productname, ismajor, userid string) (result *FileResult) {
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
	product, _ := db.GetProductInfo(o, productname)

	timenow := time.Now().Format("2006-01-02 15:04:05")
	product.ProductName = productname
	product.LastCreater = userid
	product.LastUpdateTime = timenow
	_, err = db.InsertOrUpdateProduct(to, product)
	fileinfo := entity.FileInfo{}
	record, _ := db.SelectFileInfo(o, filename, productname)

	major := int64(1)
	minor := int64(1)
	if fileinfo.FileName != "" {
		if ismajor == "yes" {
			major = record.MajorVersion + 1
		} else {
			major = record.MajorVersion
			minor = record.MinorVersion + 1
		}
	}
	fileinfo.FileName = filename
	fileinfo.Creater = userid
	fileinfo.MajorVersion = major
	fileinfo.MinorVersion = minor
	fileinfo.Modifier = userid
	fileinfo.ProductName = productname
	fileinfo.CreateTime = timenow
	fileinfo.ModifyTime = timenow
	_, err = db.InsertOrUpdateFile(to, fileinfo)
	if err != nil {
		panic("****UploadFile****添加或更新文件版本失败" + err.Error())

	}
	record.FileName = filename
	record.ProductName = productname
	record.Operationtype = "upload"
	record.MajorVersion = major
	record.MinorVersion = minor
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_, err = db.Record(to, record)
	if err != nil {
		panic("****UploadFile****添加记录失败" + err.Error())

	}
	_ = to.Commit()
	file_result.Success = true
	file_result.Msg = ""
	return file_result

}
func GetHeader() (result *FileResult) {
	o := orm.NewOrm()
	defer common.RecoverHandler(func(rcErr error) {
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	productlist, err := db.GetHeader(o)
	if err != nil {
		panic(err)
	}
	file_result.Success = true
	file_result.Msg = ""
	file_result.Data = productlist
	return file_result

}
func GetLatestTrend() (result *FileResult) {
	o := orm.NewOrm()
	defer common.RecoverHandler(func(rcErr error) {
		fileresult := new(FileResult)
		fileresult.Success = false
		fileresult.Msg = rcErr.Error()
		result = fileresult
	})
	file_result := new(FileResult)
	productlist, err := db.GetLatestTrend(o)
	if err != nil {
		panic(err)
	}
	file_result.Success = true
	file_result.Msg = ""
	file_result.Data = productlist
	return file_result

}
