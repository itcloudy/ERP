package product

import "goERP/controllers/base"

type ProductAttributeLineController struct {
	base.BaseController
}

func (ctl *ProductAttributeLineController) Get() {
	action := ctl.GetString(":action")
	viewType := ctl.Input().Get("view_type")
	switch action {
	case "list":
		switch viewType {
		case "list":
			ctl.List()
		default:
			ctl.List()
		}
	default:
		ctl.List()
	}
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/product/attribute/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["productRootActive"] = "active"
	ctl.Data["productAttributeLineActive"] = "active"
	ctl.Layout = "base/base.html"
}
func (ctl *ProductAttributeLineController) List() {
	ctl.PageAction = "列表"
}
