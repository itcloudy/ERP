package controllers

import (
	service "golangERP/services"
	"golangERP/utils"
)

// UserController 城市模块
type UserController struct {
	BaseController
}

// Get get users
func (ctl *UserController) Get() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	var err error
	// 获得城市列表信息
	if IDStr == "" {
		query := make(map[string]interface{})
		exclude := make(map[string]interface{})
		cond := make(map[string]map[string]interface{})
		fields := make([]string, 0, 0)
		sortby := make([]string, 0, 0)
		order := make([]string, 0, 0)
		offsetStr := ctl.Input().Get("offset")
		var offset int64
		var limit int64 = 20
		if offsetStr != "" {
			offset, _ = utils.ToInt64(offsetStr)
		}
		limitStr := ctl.Input().Get("limit")
		if limitStr != "" {
			if limit, err = utils.ToInt64(limitStr); err != nil {
				limit = 20
			}
		}
		var users []map[string]interface{}
		var paginator utils.Paginator
		if paginator, users, err = service.ServiceGetUser(&ctl.User, query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			data := make(map[string]interface{})
			data["users"] = &users
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
