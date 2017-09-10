package controllers

import (
	service "golangERP/services"
	"golangERP/utils"
)

// AddressProvinceContriller 城市模块
type AddressProvinceContriller struct {
	BaseController
}

// Get get provinces
func (ctl *AddressProvinceContriller) Get() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	var err error
	// 获得城市列表信息
	if IDStr == "" {
		query := make(map[string]interface{})
		exclude := make(map[string]interface{})
		cond := make(map[string]map[string]interface{})
		condAnd := make(map[string]interface{})
		fields := make([]string, 0, 0)
		sortby := make([]string, 0, 0)
		order := make([]string, 0, 0)
		offsetStr := ctl.Input().Get("offset")
		nameStr := ctl.Input().Get("name")

		if nameStr != "" {
			condAnd["Name__icontains"] = nameStr
		}
		if len(condAnd) > 0 {
			cond["and"] = condAnd
		}
		var offset int64
		var limit int64 = 20
		if offsetStr != "" {
			offset, _ = utils.GetInt64(offsetStr)
		}
		limitStr := ctl.Input().Get("limit")
		if limitStr != "" {
			if limit, err = utils.GetInt64(limitStr); err != nil {
				limit = 20
			}
		}
		var provinces []map[string]interface{}
		var paginator utils.Paginator
		if paginator, provinces, err = service.ServiceGetAddressProvince(&ctl.User, query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			data := make(map[string]interface{})
			data["provinces"] = &provinces
			data["paginator"] = &paginator
			response["data"] = data
		} else {
			response["code"] = utils.FailedCode
			response["msg"] = utils.FailedMsg
			response["err"] = err
		}
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
