package db

import (
	"FileReport/common"
	"FileReport/entity"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func SelectFileInfo(o orm.Ormer, filename, productname string) (entity.FileRecord, error) {

	var file entity.FileRecord
	err := o.Raw("select file_name,major_version,minor_version,product_name,abolish_flag,publish_flag from file_info where file_name = ? and product_name = ? ", filename, productname).QueryRow(&file)
	if err != nil {
		logs.Error(err)
	}
	return file, err
}
func SelectRecentUpdate(o orm.Ormer) ([]entity.FileRecord, error) {

	var fileinfolist []entity.FileRecord
	_, err := o.Raw("select file_name,major_version,minor_version,product_name,abolish_flag,publish_flag from file_info where modify_time between date('now', \"-1 month\") and date('now',\"+1 day\")").QueryRows(&fileinfolist)
	if err != nil {
		logs.Error(err)
	}
	return fileinfolist, err
}
func GetProductHeader(o orm.Ormer) ([]entity.ProductInfo, error) {

	var productlist []entity.ProductInfo
	nums, err := o.Raw("SELECT id,product_name,last_update_time,last_creater FROM product_info ").QueryRows(&productlist)
	if err != nil || nums < 1 {
		logs.Error(err)
	}
	return productlist, err
}
func AbolishFile(to orm.TxOrmer, fileinfo entity.FileInfo) (bool, error) {

	_, err := to.Raw("Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?",
		common.Have_Abolish, common.Not_Publish, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	if err != nil {
		logs.Error(err)
	}
	return true, err
}
func PublishFile(to orm.TxOrmer, fileinfo entity.FileInfo) (bool, error) {

	_, err := to.Raw("Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?",
		common.Not_Abolish, common.Have_Publish, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	if err != nil {
		logs.Error(err)
	}
	return true, err
}
func UpdateFile(to orm.TxOrmer, fileinfo entity.FileRecord) (bool, error) {

	result, err := to.Raw("Update file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?",
		fileinfo.MajorVersion, fileinfo.MinorVersion, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	nums, err := result.RowsAffected()
	if nums < 1 {
		err = errors.New("不存在此文件")
	}
	return true, err
}
func Record(to orm.TxOrmer, record entity.FileRecord) (bool, error) {

	_, err := to.Insert(&record)
	if err != nil {
		logs.Error(err)
	}
	return true, err
}
func GetProductInfo(o orm.Ormer, productname string) (entity.ProductInfo, error) {

	var product entity.ProductInfo
	err := o.Raw("SELECT product_name,last_update_time,last_creater FROM product_info  WHERE product_name = ?", productname).QueryRow(&product)
	if err != nil {
		logs.Error(err)
	}
	return product, err
}
func InsertOrUpdateProduct(to orm.TxOrmer, productinfo entity.ProductInfo) (bool, error) {

	result, err := to.Raw("Update product_info SET last_creater = ?,last_update_time=? WHERE product_name=?", productinfo.LastCreater, productinfo.LastUpdateTime, productinfo.ProductName).Exec()
	nums, err := result.RowsAffected()
	if nums < 1 {
		_, err = to.Insert(&productinfo)
	}

	return true, err
}
func InsertOrUpdateFile(to orm.TxOrmer, fileinfo entity.FileInfo) (bool, error) {

	result, err := to.Raw("Update file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?",
		fileinfo.MajorVersion, fileinfo.MinorVersion, fileinfo.Modifier, fileinfo.ModifyTime, fileinfo.FileName, fileinfo.ProductName).Exec()
	nums, err := result.RowsAffected()
	if nums < 1 {
		_, err = to.Insert(&fileinfo)
	}
	return true, err
}

//o.Raw("Update product_info SET last_creater = ?,last_update_time=? WHERE product_name=?", creater, timenow, productname).Exec()
func GetHeader(o orm.Ormer) ([]entity.ProductStatus, error) {

	var productInfo []entity.ProductStatus
	nums, err := o.Raw("SELECT product_name, COUNT( DISTINCT file_name ) AS file_number, COUNT( DISTINCT creater ) AS person_number, COUNT( * ) AS version_count  FROM file_record  GROUP BY product_name").QueryRows(&productInfo)
	if err != nil || nums < 1 {
		logs.Error(err)
	}
	return productInfo, err
}
func GetLatestTrend(o orm.Ormer) ([]entity.FileRecord, error) {

	var productInfo []entity.FileRecord
	nums, err := o.Raw("SELECT *  FROM file_record  ").QueryRows(&productInfo)
	if err != nil || nums < 1 {
		logs.Error(err)
	}
	return productInfo, err
}
