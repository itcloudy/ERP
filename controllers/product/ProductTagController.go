package product

import (
	"goERP/controllers/base"
	md "goERP/models"
	"strings"
)

type ProductTagController struct {
	base.BaseController
}

func (ctl *ProductTagController) Get() {
	ctl.PageName = "产品标签"
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
func (ctl *ProductTagController) List() {
	ctl.PageAction = "列表"
}
func (ctl *ProductTagController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := md.GetProductTagByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordID == obj.Id {
				result["valid"] = true
			} else {
				result["valid"] = false
			}

		} else {
			result["valid"] = true
		}

	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
