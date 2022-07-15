package db

import (
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func SelectUser(o orm.Ormer, userid string) (entity.UserInfo, error) {

	var userinfo entity.UserInfo
	err := o.Raw("SELECT * FROM user_info where userid = ? ", userid).QueryRow(&userinfo)
	if err != nil {
		logs.Error(err)
	}
	return userinfo, err
}
func AddUser(to orm.TxOrmer, user entity.UserInfo) error {
	_, err := to.Insert(&user)
	return err
}
