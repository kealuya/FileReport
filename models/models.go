package models

import (
	"FileReport/db/handler"
	"FileReport/entity"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
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
func GetProductHeader() ([]entity.ProductInfo, error) {
	o := orm.NewOrm()
	var productlist []entity.ProductInfo
	productnums, err := o.Raw("SELECT * FROM product_info ").QueryRows(&productlist)
	if err != nil || productnums < 1 {
		return productlist, err
	}
	return productlist, nil
}
func Upload(filename, productname, ismajor, creater string) (string, error) {
	o := orm.NewOrm()
	/*	id, err := o.Insert(file)
		if err == nil {
			fmt.Println(id)
		}*/
	var productlist []entity.ProductInfo
	product := new(entity.ProductInfo)
	productnums, err := o.Raw("SELECT * FROM product_info WHERE product_name = ?", productname).QueryRows(&productlist)

	timenow := time.Now().Format("2006-01-02 15:04:05")
	if productnums < 1 {
		product.ProductName = productname
		product.LastCreater = creater
		product.LastUpdateTime = timenow
		//_, err := o.Raw("INSERT file_info SET modifier = ? WHERE id=?", creater, "1").Exec()
		_, err := o.Insert(product)
		if err != nil {
			panic("sss")
		}
	} else {
		_, err := o.Raw("Update product_info SET last_creater = ?,last_update_time=? WHERE product_name=?", creater, timenow, productname).Exec()
		if err != nil {
			panic("sss")
		}
	}
	var fileinfolist []entity.FileInfo

	filenums, err := o.Raw("SELECT major_version, minor_version FROM file_info WHERE file_name = ? and product_name=?", filename, productname).QueryRows(&fileinfolist)
	fileinfo := entity.FileInfo{}
	major := int8(1)
	minor := int8(1)
	if filenums > 0 {
		fileinfo = fileinfolist[0]
		if ismajor == "yes" {
			major = fileinfo.MajorVersion + 1
		} else {
			major = fileinfo.MajorVersion
			minor = fileinfo.MinorVersion + 1
		}
	}
	fileinfo.FileName = filename
	fileinfo.Creater = creater
	fileinfo.MajorVersion = major
	fileinfo.MinorVersion = minor
	fileinfo.Modifier = creater
	fileinfo.ProductName = productname
	fileinfo.CreateTime = timenow
	fileinfo.ModifyTime = timenow

	res, err := o.Insert(&fileinfo)
	if err == nil {
		fmt.Println("mysql row affected nums: ", res)
	}
	//res, err := o.Raw("inset into file_info SET name = ?", "your").Exec()
	//if err == nil {
	//	num, _ := res.RowsAffected()
	//	fmt.Println("mysql row affected nums: ", num)
	//}

	return "success", err

}
