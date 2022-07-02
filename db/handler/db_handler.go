package handler

import (
	"FileReport/base_struct"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "db/filereport.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//注册Sqlite数据库
func RegisterDB() {
	logs.Info("========DB连接初始化========")
	orm.RegisterModel(new(base_struct.FileInfo))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

}
