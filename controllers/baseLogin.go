package controllers

import (
	"encoding/json"
	service "golangERP/services"
	"golangERP/utils"
	"strconv"
)

// LoginController 登录模块
type LoginController struct {
	BaseController
}

// Post 登录请求
func (ctl *LoginController) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	if user, ok := service.ServiceUserLogin(username, password); ok {
		user.Password = ""
		ctl.SetSession("User", *user)
		response["code"] = utils.SuccessCode
		response["msg"] = utils.SuccessMsg
		data := make(map[string]interface{})
		data["user"] = &user
		response["data"] = data
		if groups, err := service.ServiceGetUserGroups(user.IsAdmin, user.ID); err == nil {
			leng := len(groups)
			groupIDs := make([]int64, leng, leng)
			for index, group := range groups {
				groupIDs[index] = group.ID
			}
			if len(groupIDs) == 0 {
				data["groups"] = make([]int, 0, 0)
			} else {
				data["groups"] = groupIDs
			}
		}
	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
		
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Get 注销登录请求
func (ctl *LoginController) Get() {
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
