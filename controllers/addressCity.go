package controllers

import (
	"fmt"
)

// AddressCityContriller 城市模块
type AddressCityContriller struct {
	BaseController
}

// Get get cities
func (ctl *AddressCityContriller) Get() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	fmt.Println(IDStr)
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
