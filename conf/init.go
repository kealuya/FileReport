package conf

import (
	"errors"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/goburrow/cache"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"xorm.io/xorm"
)

func InitConf() {
	initCacheFr()
	initDbEngine()
	InitLogConfig()
}

var CacheFr cache.LoadingCache

func initCacheFr() {

	load := func(k cache.Key) (cache.Value, error) {
		return "", errors.New("no")
	}
	// Create a new cache
	CacheFr = cache.NewLoadingCache(load,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterWrite(60*time.Second),
	)
}

var Engine *xorm.Engine

func initDbEngine() {
	dbFilePath, _ := config.String("dbfilepath")
	engine, err_NewEngine := xorm.NewEngine("sqlite3", dbFilePath)
	if err_NewEngine != nil {
		logs.Error("数据库初始化发生错误::", err_NewEngine)
	}
	Engine = engine
	Engine.ShowSQL(true)
}

func InitLogConfig() {
	_ = logs.SetLogger(logs.AdapterConsole)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true,"separate":["error", "warning", "info", "debug"]}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出log
	//logs.Async()
}
