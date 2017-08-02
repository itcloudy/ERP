package controllers

import (
	"encoding/json"
	"fmt"
	service "golangERP/services"
	"golangERP/utils"
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
		response["msg"] = "验证通过"
		data := make(map[string]interface{})
		data["user"] = &user
		response["data"] = data
	} else {
		response["code"] = utils.FailedCode
		response["msg"] = "验证失败"
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Get 注销登录请求
func (ctl *LoginContriller) Get() {
	response := make(map[string]interface{})
	var requestBody interface{}
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	fmt.Printf("%v\n", requestBody)
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
