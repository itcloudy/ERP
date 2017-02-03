package product

import (
	"bytes"
	"goERP/controllers/base"
)

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
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()
}
func (ctl *ProductPackagingController) List() {

}
