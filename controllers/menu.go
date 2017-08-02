package controllers

import "encoding/json"

// MenuController 菜单模块
type MenuController struct {
	BaseController
}

// Post get menus by permissions
func (ctl *MenuController) Post() {
	response := make(map[string]interface{})
	var requestBody map[string]string
	json.Unmarshal(ctl.Ctx.Input.RequestBody, &requestBody)
	response["post"] = requestBody
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
