package controllers

import (
	service "golangERP/services"
	"golangERP/utils"
)

// SaleOrderController 销售订单模块
type SaleOrderController struct {
	BaseController
}

// Delete delete order attribute value
func (ctl *SaleOrderController) Delete() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")

	if id, err := utils.ToInt64(IDStr); err == nil {
		if _, err := service.ServiceDeleteSaleOrder(&ctl.User, id); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			response["orderID"] = id
		} else {
			response["code"] = utils.FailedCode
			response["msg"] = utils.FailedMsg
			response["err"] = err.Error()
		}
	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
		response["err"] = "ID转换失败"
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Put update order
func (ctl *SaleOrderController) Put() {
	response := make(map[string]interface{})
	IDStr := ctl.Ctx.Input.Param(":id")
	if IDStr != "" {
		if id, err := utils.ToInt64(IDStr); err == nil {
			if err := service.ServiceUpdateSaleOrder(&ctl.User, ctl.Ctx.Input.RequestBody, id); err == nil {
				response["code"] = utils.SuccessCode
				response["msg"] = utils.SuccessMsg
				response["orderID"] = id
			} else {
				response["code"] = utils.FailedCode
				response["msg"] = utils.FailedMsg
				response["err"] = err.Error()
			}
		} else {
			response["code"] = utils.FailedCode
			response["msg"] = utils.FailedMsg
			response["err"] = "ID转换失败"
		}

	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
		response["err"] = "ID为空"
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Post create
func (ctl *SaleOrderController) Post() {
	response := make(map[string]interface{})
	if orderID, err := service.ServiceCreateSaleOrder(&ctl.User, ctl.Ctx.Input.RequestBody); err == nil {
		response["code"] = utils.SuccessCode
		response["msg"] = utils.SuccessMsg
		response["orderID"] = orderID
	} else {
		response["code"] = utils.FailedCode
		response["msg"] = utils.FailedMsg
		response["err"] = err.Error()
	}

	ctl.Data["json"] = response
	ctl.ServeJSON()
}

// Get get cities
func (ctl *SaleOrderController) Get() {
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
		nameStr := ctl.Input().Get("name")

		if nameStr != "" {
			condAnd["Name__icontains"] = nameStr
		}
		if len(condAnd) > 0 {
			cond["and"] = condAnd
		}
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
		var cities []map[string]interface{}
		var paginator utils.Paginator
		var access utils.AccessResult
		if access, paginator, cities, err = service.ServiceGetSaleOrder(&ctl.User, query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
			response["code"] = utils.SuccessCode
			response["msg"] = utils.SuccessMsg
			data := make(map[string]interface{})
			data["cities"] = &cities
			data["paginator"] = &paginator
			data["access"] = access
			response["data"] = data

		} else {
			response["code"] = utils.FailedCode
			response["msg"] = utils.FailedMsg
			response["err"] = err
		}
	} else {
		// 获得某个城市的信息
		if orderID, err := utils.ToInt64(IDStr); err == nil {
			if access, order, err := service.ServiceGetSaleOrderByID(&ctl.User, orderID); err == nil {
				response["code"] = utils.SuccessCode
				response["msg"] = utils.SuccessMsg
				data := make(map[string]interface{})
				data["order"] = &order
				data["access"] = access
				response["data"] = data
			} else {
				response["code"] = utils.FailedCode
				response["msg"] = utils.FailedMsg
				response["err"] = err
			}
		}
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()
}
