package controllers

import (
	"encoding/json"
	service "golangERP/services"
	"golangERP/utils"
	"strconv"
)

// LoginContriller 登录模块
type LoginContriller struct {
	BaseController
}

// Post 登录请求
func (ctl *LoginContriller) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	if user, ok := service.ServiceUserLogin(username, password); ok {
		user.Password = ""
		response["code"] = utils.SuccessCode
		response["msg"] = utils.SuccessMsg
		data := make(map[string]interface{})
		data["user"] = &user
		response["data"] = data
		service.ServiceGetGroups(true, 0)
	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Get 注销登录请求
func (ctl *LoginContriller) Get() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	if ID, err := strconv.ParseInt(IDStr, 10, 64); err == nil {
		service.ServiceUserLogout(ID)
		response["code"] = utils.SuccessCode
		response["msg"] = utils.SuccessMsg

	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
