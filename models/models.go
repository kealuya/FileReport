package models

import (
	"FileReport/base_struct"
	"FileReport/db/handler"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	handler.RegisterDB()
}
func Test() (base_struct.FileInfo, error) {
	o := orm.NewOrm()
	fileInfo := base_struct.FileInfo{}
	err := o.Read(&fileInfo)

	return fileInfo, err
}
func TestView() ([]base_struct.ProductInfo, error) {
	o := orm.NewOrm()
	var productInfo []base_struct.ProductInfo
	nums, err := o.Raw("SELECT * FROM product_info ").QueryRows(&productInfo)
	if err == nil {
		fmt.Println("user nums: ", nums)
	}
	return productInfo, err

}
