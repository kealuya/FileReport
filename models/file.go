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

func GetProductHeader() (result []entity.ProductInfo, resultErr error) {

	defer common.RecoverHandler(func(rcErr error) {
		result = []entity.ProductInfo{}
		resultErr = rcErr
	})
	productlist := db.GetProductHeader()
	return productlist, nil
}
func GetRecentUpdate() (result []entity.FileRecord, resultErr error) {
	defer common.RecoverHandler(func(rcErr error) {

		result = []entity.FileRecord{}
	})
	productlist := db.SelectRecentUpdate()

	return productlist, nil
}

func UpdateFile(filename, productname, userid string, ma, mi int64) (result string, resultErr error) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = "fail"
		resultErr = rcErr
	})

	if err != nil {
		panic("****UpdateFile****start the transaction failed")
	}
	timenow := time.Now().Format("2006-01-02 15:04:05")
	record := db.SelectFileInfo(filename, productname)

	record.Modifier = userid
	record.ModifyTime = timenow
	record.MinorVersion = ma
	record.MinorVersion = mi
	_ = db.UpdateFile(to, record)

	record.Operationtype = "update"
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_ = db.Record(to, record)

	_ = to.Commit()

	return "sucess", nil
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
func Upload(filename, productname, ismajor, userid string) (result string, resultErr error) {
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
	product := db.GetProductInfo(productname)

	timenow := time.Now().Format("2006-01-02 15:04:05")
	product.ProductName = productname
	product.LastCreater = userid
	product.LastUpdateTime = timenow
	_ = db.InsertOrUpdateProduct(to, product)
	fileinfo := entity.FileInfo{}
	record := db.SelectFileInfo(filename, productname)

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
	_ = db.InsertOrUpdateFile(to, fileinfo)

	record.FileName = filename
	record.ProductName = productname
	record.Operationtype = "upload"
	record.MajorVersion = major
	record.MinorVersion = minor
	record.CreateTime = timenow
	record.Creater = userid
	record.Modifier = userid
	record.ModifyTime = timenow
	_ = db.Record(to, record)

	_ = to.Commit()
	return "success", nil

}
func GetCurrentHeader() (result []entity.ProductStatus, resultErr error) {
	defer common.RecoverHandler(func(rcErr error) {

		result = []entity.ProductStatus{}
		resultErr = rcErr
	})
	productlist := db.GetCurrentHeader()

	return productlist, nil

}
func GetLatestTrend() (result []entity.FileRecord, resultErr error) {
	defer common.RecoverHandler(func(rcErr error) {
		result = []entity.FileRecord{}
		resultErr = rcErr
	})

	productlist := db.GetLatestTrend()

	return productlist, nil

}
