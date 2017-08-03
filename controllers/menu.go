package controllers

import "encoding/json"
import "fmt"

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
	userID := int64((requestBody["userID"]).(float64))
	isAdmin := (requestBody["isAdmin"]).(bool)
	fmt.Println(userID)
	fmt.Println(isAdmin)
	response["post"] = requestBody
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
