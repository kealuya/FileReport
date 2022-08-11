package common

import (
	"errors"
	"github.com/beego/beego/v2/core/logs"
	jsoniter "github.com/json-iterator/go"
	"log"
	"math/rand"
	"reflect"
	"runtime/debug"
	"strconv"
	"time"
)

func Marshal(v interface{}) []byte {
	jsonByte, err := jsoniter.Marshal(v)
	if err != nil {

		logs.Error("★★★错误error::", err, "jsoniter.Marshal发生错误")
		return []byte{}
	}
	return jsonByte
}

//共通错误recover处理方法
func RecoverHandler(f func(err error)) {
	if err := recover(); err != nil {
		logs.Error("发生系统panic::", err)
		logs.Error(string(debug.Stack()))
		reflect.TypeOf(&err).String()
		if f != nil {
			_, ok := err.(error)
			if ok {
				f(err.(error))
			} else {
				f(errors.New(err.(string)))
			}
		}
	}
}

//共通错误error处理方法
func ErrorHandler(err error, info ...interface{}) {
	if err != nil {
		logs.Error("发生系统panic::", err, info)
		log.Panicln(err)
	}
}

//String转Int8
func StringToInt64(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 8)
	return n
}

const (
	Have_Abolish     = "1"
	Have_Abolish_Msg = "已废除"
	Not_Abolish      = "0"
	Not_Abolish_Msg  = "未废除"
	Have_Publish     = "1"
	Have_Publish_Msg = "已发布"
	Not_Publish      = "0"
	Not_Publish_Msg  = "未发布"
	OwnEdit          = "0"
	OwnEdit_Msg      = "仅自己"
	Not_OwnEdit      = "1"
	Not_OwnEdit_Msg  = "所有人"
)

func GenerateSubId(n int32) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateNumberCode(n int32) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}
