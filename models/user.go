package models

import (
	"FileReport/db"
	"FileReport/entity"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

func Login(userid, password string) (entity.UserInfo, error) {
	o := orm.NewOrm()
	userinfo := entity.UserInfo{}
	userinfo, err := db.SelectUser(o, userid)

	if err != nil || userinfo.Userid == "" {
		return entity.UserInfo{}, errors.New("不存在该用户")
	} else if userinfo.Password != password {
		return entity.UserInfo{}, errors.New("密码错误")
	}

	return userinfo, nil
}
func AddPeople(userid, password, username, userrole string) (entity.UserInfo, error) {
	o := orm.NewOrm()
	//不好用
	/*	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		user := entity.UserInfo{}
		user.Userid = userid
		user.Password = password
		user.UserRole = userrole
		user.Username = username

		// insert data
		// Using txOrm to execute SQL
		e := db.AddUser(txOrm, user)
		// if e != nil the transaction will be rollback
		// or it will be committed
		return e
	})*/
	to, err := o.Begin()
	user := entity.UserInfo{}
	user.Userid = userid
	user.Password = password
	user.UserRole = userrole
	user.Username = username

	err = db.AddUser(to, user)
	//另一个写库操作，待补充
	/*	user1 := user
			user1.Userid = "dsad"
		err1 := db.AddUser(to, user1)
		if err1 != nil {
			to.Rollback()
			return entity.UserInfo{}, err1
		} else */if err != nil {
		to.Rollback()
	} else {
		to.Commit()
	}
	userinfo, err := db.SelectUser(o, userid)
	return userinfo, err
}
