package models

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/entity"
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
