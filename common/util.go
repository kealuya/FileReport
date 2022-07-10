package common

import (
	"github.com/beego/beego/v2/core/logs"
	jsoniter "github.com/json-iterator/go"
)

func Marshal(v interface{}) []byte {
	jsonByte, err := jsoniter.Marshal(v)
	if err != nil {

		logs.Error("★★★错误error::", err, "jsoniter.Marshal发生错误")
		return []byte{}
	}
	return jsonByte
}
