package db

import (
	"FileReport/common"
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
)

func SelectUser(userid string) entity.UserInfo {
	o := orm.NewOrm()
	userinfo := entity.UserInfo{}
	err := o.Raw("SELECT * FROM user_info WHERE userid = ? ", userid).QueryRow(&userinfo)
	common.ErrorHandler(err, "从数据库获取用户信息失败")
	return userinfo
}
func AddUser(to orm.TxOrmer, user entity.UserInfo) error {
	_, err := to.Insert(&user)
	return err
}
