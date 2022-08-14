package models

import (
	"FileReport/common"
	"FileReport/conf"
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

/*func Upload(filename, productname, ismajor, userid string) (result string, resultErr error) {
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

}*/
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
func Upload(docinfo entity.Doc, fileinfo entity.File) (result string, resultErr error) {

	common.RecoverHandler(func(err error) {
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
	doc.CreateDate = time.Now()
	doc.ProId = docinfo.ProId

	file := new(entity.File)
	file.FileName = fileinfo.FileName
	file.Version = fileinfo.Version
	file.VersionShow = fileinfo.VersionShow
	file.UpdateDate = time.Now()
	file.UpdateUserId = fileinfo.UpdateUserId
	file.UpdateContent = fileinfo.UpdateContent
	_, err_Insert := conf.Engine.Insert(doc, file)
	common.ErrorHandler(err_Insert)

	return "success", nil
}
func FileAuthority(docinfo entity.Doc) (result string, resultErr error) {

	common.RecoverHandler(func(err error) {
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
	doc.CreateDate = time.Now()
	doc.ProId = docinfo.ProId

	_, err_Insert := conf.Engine.Where("doc_id=?", 2).Update(doc)
	common.ErrorHandler(err_Insert)

	return "success", nil
}

type DocFile struct {
	entity.Doc    `xorm:"extends"`
	Version       int       `xorm:"not null integer" json:"version,omitempty"`
	VersionShow   string    `xorm:"TEXT" json:"versionShow,omitempty"`
	UpdateDate    time.Time `xorm:"DATE" json:"updateDate"`
	UpdateUserId  string    `xorm:"text" json:"updateUserId,omitempty"`
	UpdateContent string    `xorm:"text" json:"updateContent,omitempty"`
	FileName      string    `xorm:"TEXT" json:"fileName,omitempty"`
}

func MyFile() (result []DocFile, resultErr error) {

	common.RecoverHandler(func(err error) {
		result = []DocFile{}
		resultErr = err
		return
	})

	docfile := new([]DocFile)

	err_Select := conf.Engine.Table("doc").Join("INNER", "file", "file.doc_id=doc.doc_id").Find(docfile)
	common.ErrorHandler(err_Select)

	return *docfile, nil
}
