package base

import (
	"encoding/json"
	md "goERP/models"
	"strings"
)

// PartnerController partner
type PartnerController struct {
	BaseController
}

// Post request
func (ctl *PartnerController) Post() {

	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "selectSearch":
		ctl.PostList()
	default:
		ctl.PostList()
	}
}

// Get request
func (ctl *PartnerController) Get() {
	ctl.GetList()
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/partner/"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuPartnerActive"] = "active"
}

// Validator js vaild
func (ctl *PartnerController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetPartnerByName(name)
	if err != nil {
		result["valid"] = true
	} else {
		if obj.Name == name {
			if recordID == obj.ID {
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

// 获得符合要求的城市数据
func (ctl *PartnerController) partnerList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var partners []md.Partner
	paginator, partners, err := md.GetAllPartner(query, fields, sortby, order, offset, limit)

	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, partner := range partners {
			oneLine := make(map[string]interface{})

			oneLine["ID"] = partner.ID
			oneLine["id"] = partner.ID
			oneLine["name"] = partner.Name

			tableLines = append(tableLines, oneLine)
		}

		result["data"] = tableLines
		if jsonResult, er := json.Marshal(&paginator); er == nil {
			result["paginator"] = string(jsonResult)
			result["total"] = paginator.TotalCount
		}
	}
	return result, err
}

// PostList post request json response
func (ctl *PartnerController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	name := strings.TrimSpace(ctl.Input().Get("name"))
	if name != "" {
		query["name"] = name
	}
	if result, err := ctl.partnerList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList get table display page
func (ctl *PartnerController) GetList() {
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-partner"
	ctl.TplName = "base/base_list_view.html"
}
