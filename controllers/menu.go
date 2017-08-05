package controllers

import (
	"encoding/json"
	"fmt"
	md "golangERP/models"
	service "golangERP/services"
	"golangERP/utils"
)

// MenuController 菜单模块
type MenuController struct {
	BaseController
}

// Post get menus by permissions
func (ctl *MenuController) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]interface{}
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	fmt.Printf("%+v\n", requestBody)
	// groups := requestBody["groups"].()
	var (
		err     error
		isAdmin bool
		menus   []md.BaseMenu
		groups  []int64
	)
	if _, ok := requestBody["isAdmin"]; ok {
		isAdmin = (requestBody["isAdmin"]).(bool)
	} else {
		isAdmin = false
	}

	if _, ok := requestBody["groups"]; ok {

	}
	data := make(map[string]interface{})
	if menus, err = service.ServiceGetMenus(isAdmin, groups); err == nil {
		data["menus"] = menus
		response["data"] = data
		response["code"] = utils.SuccessCode
		response["msg"] = "菜单获取成功"
	}
	if err != nil {
		response["code"] = utils.FailedCode
		response["msg"] = "菜单获取失败"
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()
}
