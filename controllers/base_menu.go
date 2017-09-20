package controllers

import (
	"encoding/json"
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
		postGroups := utils.ToSlice(requestBody["groups"])
		for _, group := range postGroups {
			if groupID, err := utils.ToInt64(group); err == nil {
				groups = append(groups, groupID)

			}
		}
	}
	data := make(map[string]interface{})
	if !isAdmin && len(groups) == 0 {
		response["code"] = utils.FailedCode
		response["msg"] = "菜单获取失败"
	} else {
		if menus, err = service.ServiceGetMenus(isAdmin, groups); err == nil {
			if len(menus) == 0 {
				data["menus"] = make([]int, 0, 0)
			} else {
				data["menus"] = menus
			}
			response["data"] = data
			response["code"] = utils.SuccessCode
			response["msg"] = "菜单获取成功"
		}
		if err != nil {
			response["code"] = utils.FailedCode
			response["msg"] = "菜单获取失败"
			response["err"] = err
		}
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()
}
