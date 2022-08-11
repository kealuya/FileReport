package handler

import (
	"FileReport/entity"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/mattn/go-sqlite3"
)

var (
	_DB_NAME, _     = config.String("dbfilepath")
	_SQLITE3_DRIVER = "sqlite3"
)

//注册Sqlite数据库
func RegisterDB() {
	logs.Info("========DB连接初始化========")
	orm.RegisterModel(new(entity.FileRecord), new(entity.ProductInfo), new(entity.UserInfo), new(entity.FileInfo))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

}
