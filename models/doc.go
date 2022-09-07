package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/entity"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gohouse/t"
	"log"
)

func InsertDoc(doc entity.Doc) (docId int32, funcErr error) {

	common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	num, err_InsertOne := conf.Engine.InsertOne(doc)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Doc表存入数据失败")
		log.Panicln()
	}
	var resultQuery []map[string][]byte
	resultQuery, err_Query := conf.Engine.Query(`select max(doc_id) as id from Doc`)
	if err_Query != nil {
		logs.Error("Doc表获取最大DocId数据失败")
		log.Panicln()
	}
	fmt.Println(resultQuery)

	return t.New(resultQuery[0]["id"]).Int32(), nil

}

func InsertFile(file entity.File) (funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})
	num, err_InsertOne := conf.Engine.InsertOne(file)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Doc表存入数据失败")
		log.Panicln()
	}
	return nil

}
