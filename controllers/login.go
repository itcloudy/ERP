package controllers

import (
	"encoding/json"
	"fmt"
)

// LoginContriller 登录模块
type LoginContriller struct {
	BaseController
}

// Post 登录请求
func (ctl *LoginContriller) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]interface{}
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	username := requestBody["username"]
	password := requestBody["password"]
	fmt.Println(username)
	fmt.Println(password)
	fmt.Printf("%v\n", requestBody)
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
