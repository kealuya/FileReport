package db

import (
	"FileReport/common"
	"FileReport/db/handler"
	"FileReport/entity"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func SelectFileInfo(filename, productname string) entity.FileRecord {
	o := orm.NewOrm()
	var file entity.FileRecord
	err := o.Raw(handler.Select_File_Info, filename, productname).QueryRow(&file)
	common.ErrorHandler(err, "从数据库获取文档信息失败")
	return file
}
func SelectRecentUpdate() []entity.FileRecord {
	o := orm.NewOrm()
	var fileinfolist []entity.FileRecord
	_, err := o.Raw(handler.Select_Recent_Update).QueryRows(&fileinfolist)
	common.ErrorHandler(err, "从数据库获取最近更新信息失败")
	return fileinfolist
}
func GetProductHeader() []entity.ProductInfo {
	o := orm.NewOrm()
	var productlist []entity.ProductInfo
	nums, err := o.Raw(handler.Select_Product_Header).QueryRows(&productlist)
	common.ErrorHandler(err, "从数据库获取项目header信息失败")
	if nums < 1 {
		logs.Error("未查到项目header")
	}
	return productlist
}
func AbolishFile(to orm.TxOrmer, fileinfo entity.FileInfo) bool {

	_, err := to.Raw(handler.Update_Abolish_File,
		common.Have_Abolish, common.Not_Publish, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	common.ErrorHandler(err, "废除文档失败")
	return true
}
func PublishFile(to orm.TxOrmer, fileinfo entity.FileInfo) bool {

	_, err := to.Raw(handler.Update_Publish_File,
		common.Not_Abolish, common.Have_Publish, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	common.ErrorHandler(err, "发布文档失败")
	return true
}
func AuthorityFile(to orm.TxOrmer, fileinfo entity.FileInfo) bool {

	_, err := to.Raw(handler.Update_Authority_File,
		fileinfo.EditFlag, fileinfo.PublishFlag, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	common.ErrorHandler(err, "发布文档失败")
	return true
}
func UpdateFile(to orm.TxOrmer, fileinfo entity.FileRecord) bool {

	result, err := to.Raw(handler.Update_File_Info,
		fileinfo.MajorVersion, fileinfo.MinorVersion, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	common.ErrorHandler(err, "更新文档信息失败")
	nums, _ := result.RowsAffected()
	if nums < 1 {
		common.ErrorHandler(errors.New("不存在此文件"), "更新文档失败")
	}

	return true
}
func Record(to orm.TxOrmer, record entity.FileRecord) bool {

	_, err := to.Insert(&record)
	common.ErrorHandler(err, "添加记录失败")
	return true
}
func GetProductInfo(productname string) entity.ProductInfo {
	o := orm.NewOrm()
	var product entity.ProductInfo
	err := o.Raw(handler.Select_ProductInfo_By_Name, productname).QueryRow(&product)
	common.ErrorHandler(err, "从数据库获取项目信息失败")
	return product
}
func InsertOrUpdateProduct(to orm.TxOrmer, productinfo entity.ProductInfo) bool {

	result, err := to.Raw(handler.Update_Product_info, productinfo.LastCreater, productinfo.LastUpdateTime, productinfo.ProductName).Exec()
	nums, _ := result.RowsAffected()
	common.ErrorHandler(err, "更新项目信息失败")
	if nums < 1 {
		_, err_Insert := to.Insert(&productinfo)
		common.ErrorHandler(err_Insert, "添加项目信息失败")
	}

	return true
}
func InsertOrUpdateFile(to orm.TxOrmer, fileinfo entity.FileInfo) bool {

	result, err := to.Raw(handler.Update_File,
		fileinfo.MajorVersion, fileinfo.MinorVersion, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	common.ErrorHandler(err, "更新文档信息失败")
	nums, _ := result.RowsAffected()
	if nums < 1 {
		_, err_Insert := to.Insert(&fileinfo)
		common.ErrorHandler(err_Insert, "添加文档信息失败")
	}
	return true
}

//func GetCurrentHeader() []entity.ProductStatus {
//	o := orm.NewOrm()
//	var productInfo []entity.ProductStatus
//	nums, err := o.Raw(handler.Select_Current_Header).QueryRows(&productInfo)
//	common.ErrorHandler(err, "从数据库获取当前状态header信息失败")
//	if nums < 1 {
//		common.ErrorHandler(errors.New("不存在此文件更新记录"), "从数据库获取当前状态header信息失败")
//	}
//	return productInfo
//}
func GetLatestTrend() []entity.FileRecord {
	o := orm.NewOrm()
	var productInfo []entity.FileRecord
	nums, err := o.Raw(handler.Select_Latest_Trend).QueryRows(&productInfo)
	common.ErrorHandler(err, "从数据库获取最新动态信息失败")
	if nums < 1 {
		common.ErrorHandler(errors.New("不存在更新记录"), "从数据库获取最新动态信息失败")
	}
	return productInfo
}
