package controllers

import (
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (uCtrl *UserController) Login() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	userInfoRequestKey := UserInfo{}

	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &userInfoRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	userinfo, err_Login := models.Login(
		userInfoRequestKey.Userid,
		userInfoRequestKey.Password)
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	if err_Login != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("登录失败 : %s", err_Login.Error())
		logs.Error(err_Login)
		return
	}
	resJson.Data = userinfo
}
func (uCtrl *UserController) AddPeople() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()

	}()

	UserInfoRequestKey := UserInfo{}
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &UserInfoRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	userinfo, err_AddPeople := models.AddPeople(
		UserInfoRequestKey.Userid,
		UserInfoRequestKey.Password,
		UserInfoRequestKey.UserRole,
		UserInfoRequestKey.Username)
	//c.Ctx.WriteString("看到我，就说明你这玩意调成功了\nsdasdsad\n")
	if err_AddPeople != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("添加人员失败 : %s", err_AddPeople.Error())
		return
	}
	resJson.Data = userinfo
}
func (uCtrl *UserController) EditPeople() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()

	}()

	UserInfoRequestKey := UserInfo{}
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &UserInfoRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	result, err_AddPeople := models.EditPeople(
		UserInfoRequestKey.Userid,
		UserInfoRequestKey.Password,
		UserInfoRequestKey.UserRole,
		UserInfoRequestKey.Username,
		UserInfoRequestKey.Modifier)
	if err_AddPeople != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("修改人员信息失败 : %s", err_AddPeople.Error())
		return
	}
	resJson.Data = result
}
