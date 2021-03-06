package models

import (
	"FileReport/common"
	"FileReport/db"
	"FileReport/entity"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func Login(userid, password string) (result entity.UserInfo, resultErr error) {
	defer common.RecoverHandler(func(rcErr error) {
		result = entity.UserInfo{}
		resultErr = rcErr
	})

	userinfo := db.SelectUser(userid)

	if userinfo.Userid == "" {
		return entity.UserInfo{}, errors.New("不存在该用户")
	}

	if userinfo.Password != password {
		return entity.UserInfo{}, errors.New("密码错误")
	}

	return userinfo, nil
}
func AddPeople(userid, password, userrole, username string) (result entity.UserInfo, resultErr error) {
	o := orm.NewOrm()
	//另一种处理事务的方法
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
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = entity.UserInfo{}
		resultErr = rcErr
	})
	if err != nil {
		panic("****AddPeople****start the transaction failed")
	}
	user := entity.UserInfo{}
	user.Userid = userid
	user.Password = password
	user.UserRole = userrole
	user.Username = username

	_ = db.AddUser(to, user)
	_ = to.Commit()
	userinfo := db.SelectUser(userid)
	return userinfo, nil
}
func EditPeople(userid, password, username, userrole, modifier string) (result string, resultErr error) {
	o := orm.NewOrm()
	to, err := o.Begin()
	defer common.RecoverHandler(func(rcErr error) {
		to.Rollback()
		result = "fail"
		resultErr = rcErr
	})
	if err != nil {
		panic("****EditPeople****start the transaction failed")
	}
	user := entity.UserInfo{}
	user.Userid = userid
	user.Password = password
	user.UserRole = userrole
	user.Username = username
	user.Modifier = modifier
	user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	_ = db.UpdateUser(to, user)

	_ = to.Commit()
	return "success", nil
}
