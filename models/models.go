package models

import (
	"FileReport/db/handler"
	"FileReport/entity"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	handler.RegisterDB()
}

/*func Test() (entity.FileInfo, error) {
	o := orm.NewOrm()
	fileInfo := entity.FileInfo{}
	err := o.Read(&fileInfo)

	return fileInfo, err
}*/
func TestView() ([]entity.ProductStatus, error) {
	o := orm.NewOrm()
	var productInfo []entity.ProductStatus
	nums, err := o.Raw("SELECT * FROM product_status ").QueryRows(&productInfo)
	if err == nil {
		fmt.Println("user nums: ", nums)
	}
	return productInfo, err

}
