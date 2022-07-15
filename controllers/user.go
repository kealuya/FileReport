package controllers

import (
	"FileReport/common"
	"FileReport/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (uCtrl *UserController) Login() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = string(common.Marshal(resJson))
		uCtrl.ServeJSON()

	}()

	var UserInfoRequestKey = new(UserInfo)
	res := uCtrl.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &UserInfoRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	userinfo, err := models.Login(
		UserInfoRequestKey.Userid,
		UserInfoRequestKey.Password)
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败失败 : %s", err.Error())
		return
	}
	resJson.Data = userinfo
}
func (uCtrl *UserController) AddPeople() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = string(common.Marshal(resJson))
		uCtrl.ServeJSON()

	}()

	var UserInfoRequestKey = new(UserInfo)
	res := uCtrl.Ctx.Input.RequestBody
	err := json.Unmarshal(res, &UserInfoRequestKey)
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err.Error())
		return
	}
	userinfo, err := models.AddPeople(
		UserInfoRequestKey.Userid,
		UserInfoRequestKey.Password,
		UserInfoRequestKey.UserRole,
		UserInfoRequestKey.Username)
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	if err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("添加人员失败 : %s", err.Error())
		return
	}
	resJson.Data = userinfo
}
