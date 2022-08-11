package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func SendMsg(phoneNumber, templateCode, templateParam string) {
	t := FormatDate(time.Now(), YYYY_MM_DD_HH_MM_SS)
	ss := Sign("szht_chailv", "12345678", templateCode, t, templateParam)

	rm := json.RawMessage{}
	_ = rm.UnmarshalJSON([]byte(templateParam))

	s := SendMsgInfo{
		AppId:         "szht_chailv",
		System:        "差旅系统",
		PhoneNumbers:  phoneNumber,
		SignName:      "浩天业财",
		TemplateCode:  templateCode,
		TemplateParam: rm,
		OutId:         "",
		Timestamp:     t,
		Sign:          ss}
	b, _ := json.Marshal(s)
	res, err := http.Post("http://122.9.41.215:8000/v1/s/send_msg",
		"application/json",
		bytes.NewReader(b),
	)
	if err != nil {
		logs.Error(err)
	}
	bb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(err)
	}
	logs.Info(phoneNumber, string(bb))
}

//签名生成
func Sign(appId, appSecret, templateCode, timestamp, templateParam string) string {
	//sign=access_token,app_key,method,timestamp,v ，再进行MD5加密
	var sb bytes.Buffer
	sb.WriteString(appSecret)
	sb.WriteString("app_id")
	sb.WriteString(appId)
	sb.WriteString("template_code")
	sb.WriteString(templateCode)
	sb.WriteString("timestamp")
	sb.WriteString(timestamp)
	sb.WriteString("template_param")
	sb.WriteString(templateParam)
	sb.WriteString(appSecret)
	signString := sb.String()
	signMd5String := StringToMd5(signString)
	return strings.ToUpper(signMd5String)
}

//字符串md5加密
func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type SendMsgInfo struct {
	AppId         string          `json:"app_id"`
	System        string          `json:"system"`
	PhoneNumbers  string          `json:"phone_numbers"`
	SignName      string          `json:"sign_name"`
	TemplateCode  string          `json:"template_code"`
	TemplateParam json.RawMessage `json:"template_param"`
	OutId         string          `json:"out_id"`
	Timestamp     string          `json:"timestamp"`
	Sign          string          `json:"sign"`
}
