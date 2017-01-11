package product

import "goERP/controllers/base"

type ProductPackagingController struct {
	base.BaseController
}

func (ctl *ProductPackagingController) Get() {
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
}
func (ctl *ProductPackagingController) List() {

}
