package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/entity"
	"github.com/beego/beego/v2/core/logs"
	"log"
)

func ObtainProjectList() (proList []entity.Project, funcErr error) {

	common.RecoverHandler(func(err error) {
		proList = make([]entity.Project, 0)
		funcErr = err
		return
	})

	proList = make([]entity.Project, 0)
	err_Find := conf.Engine.Find(&proList)
	common.ErrorHandler(err_Find)

	return proList, nil

}
func InsertProject(pro entity.Project) (funcErr error) {

	defer common.RecoverHandler(func(err error) {
		funcErr = err
		return
	})

	num, err_InsertOne := conf.Engine.InsertOne(pro)
	common.ErrorHandler(err_InsertOne)
	if num < 1 {
		logs.Error("Pro表存入数据失败")
		log.Panicln()
	}

	return nil

}
