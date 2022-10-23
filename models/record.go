package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/entity"
	"github.com/beego/beego/v2/core/logs"
	"log"
)

func SaveRecord(rec_info entity.Record) (funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	num, err_InsertOne := conf.Engine.InsertOne(rec_info)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Record表存入数据失败")
		log.Panicln()
	}

	return nil

}
