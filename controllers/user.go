package controllers

import (
	"FileReport/common"
	"FileReport/conf"
	"FileReport/db"
	"FileReport/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (uCtrl *UserController) GetCaptcha() {
	logs.Info("init GetCaptcha")
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()
	phoneInfo := make(map[string]any)

	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &phoneInfo)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal, string(res))
		return
	}
	// 不定义struct，明确接受参数
	phoneNumber := phoneInfo["phoneNumber"].(string)

	// 判断是否是既存用户
	has, _, err_ObtainUser := models.ObtainUser(phoneNumber)

	if err_ObtainUser != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_ObtainUser.Error())
		logs.Error(err_ObtainUser)
		return
	}
	// 如果没有用户
	if !has {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("您不是社内人员，无法登录")
		return
	}
	// 缓存获取用户
	_, isHasKey := conf.CacheFr.GetIfPresent(phoneNumber)
	if !isHasKey {
		//	没有申请过验证码
		captcha := common.GenerateNumberCode(6)
		common.SendMsg(phoneNumber, "SMS_243890174", fmt.Sprintf("{\"code\":\"%s\"}", captcha))
		conf.CacheFr.Put(phoneNumber, captcha)
	} else {
		//	1分钟内已经申请过验证码
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("请不要重复提交申请验证码")
		return
	}

}

func (uCtrl *UserController) LoginFr() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	loginInfo := make(map[string]any)
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &loginInfo)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}

	// 不定义struct，明确接受参数
	phoneNumber := loginInfo["phoneNumber"].(string)
	captcha := loginInfo["captcha"].(string)

	cacheCaptcha, isHasKey := conf.CacheFr.GetIfPresent(phoneNumber)
	if !isHasKey {
		//	验证码已经过期
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("验证码已经过期")
		return
	} else {
		//	获取验证码，比对验证码正确性
		if cacheCaptcha == captcha {
			// 登录成功
			user, err_UserLoginHandler := models.UserLoginHandler(phoneNumber)
			if err_UserLoginHandler != nil {
				resJson.Success = false
				resJson.Msg = fmt.Sprintf("系统错误 : %s", err_UserLoginHandler.Error())
				logs.Error(err_UserLoginHandler)
			}
			resJson.Data = user
			// 登录成功，key无效化
			conf.CacheFr.Invalidate(phoneNumber)
			return
		} else {
			// 登录失败
			// 验证码不正确
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("验证码填写错误")
			return
		}
	}

}

func (uCtrl *UserController) Login() {
	resJson := NewJsonStruct(nil)
	defer func() {
		uCtrl.Data["json"] = resJson
		uCtrl.ServeJSON()
	}()

	userInfoRequestKey := UserInfo{}
	token := uCtrl.Ctx.Request.Header.Get("token")
	res := uCtrl.Ctx.Input.RequestBody
	err_Unmarshal := json.Unmarshal(res, &userInfoRequestKey)
	if err_Unmarshal != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("系统错误 : %s", err_Unmarshal.Error())
		logs.Error(err_Unmarshal)
		return
	}
	if token != "" {
		tokenback, err_Token := ParseToken(token)
		if err_Token != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("Token错误 : %s", err_Token.Error())
			logs.Error(err_Unmarshal)
			return
		}
		userinfo := db.SelectUser(tokenback.UserId)
		resJson.Data = userinfo
	} else {
		userinfo, err_Login := models.Login(
			userInfoRequestKey.Userid,
			userInfoRequestKey.Password)
		if err_Login != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("登录失败 : %s", err_Login.Error())
			logs.Error(err_Login)
			return
		}
		token, exdate, err_Token := GetnerateToken(userInfoRequestKey.Userid)
		if err_Token != nil {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("token获取失败 : %s", err_Token.Error())
			logs.Error(err_Token)
			return
		}
		result_token := TokenStruct{}
		result_token.Token = token
		result_token.ExpiresAt = exdate
		result_token.UserInfo = userinfo
		resJson.Data = result_token
	}

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
