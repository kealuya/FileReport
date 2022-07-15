package db

import (
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func SelectRecentUpdate(o orm.Ormer) ([]entity.FileInfo, error) {

	var fileinfolist []entity.FileInfo
	_, err := o.Raw("select * from file_info where modify_time between date('now', \"-1 month\") and date('now')").QueryRows(&fileinfolist)
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

//o.Raw("SELECT * FROM product_info WHERE product_name = ?", productname).QueryRows(&productlist)
