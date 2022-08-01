package db

import (
	"FileReport/common"
	"FileReport/db/handler"
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
)

func SelectUser(userid string) entity.UserInfo {
	o := orm.NewOrm()
	userinfo := entity.UserInfo{}
	err := o.Raw(handler.Select_User, userid).QueryRow(&userinfo)
	common.ErrorHandler(err, "从数据库获取用户信息失败")
	return userinfo
}
func AddUser(to orm.TxOrmer, user entity.UserInfo) bool {
	_, err := to.Insert(&user)
	common.ErrorHandler(err, "添加人员信息失败")
	return true
}
func UpdateUser(to orm.TxOrmer, userinfo entity.UserInfo) bool {

	_, err := to.Raw(handler.Update_User,
		userinfo.UserRole, userinfo.Password, userinfo.Modifier, userinfo.ModifyTime, userinfo.Userid).Exec()
	common.ErrorHandler(err, "更改人员信息失败")
	return true
}
