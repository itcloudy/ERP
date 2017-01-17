package address

import (
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strings"
)

type ProvinceController struct {
	base.BaseController
}

func (ctl *ProvinceController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	default:
		ctl.PostList()
	}
}
func (ctl *ProvinceController) Get() {
	ctl.PageName = "省份管理"
	ctl.URL = "/address/city/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProvinceActive"] = "active"
	ctl.GetList()
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
}
func (ctl *ProvinceController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	name := strings.TrimSpace(ctl.GetString("Name"))
	if name != "" {
		query["Name"] = name
	}
	if result, err := ctl.provinceList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *ProvinceController) Validator() {
	name := ctl.GetString("Name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetPositionByName(name)
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

// 获得符合要求的地区数据
func (ctl *ProvinceController) provinceList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var provinces []md.AddressProvince
	paginator, provinces, err := md.GetAllAddressProvince(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, province := range provinces {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = province.Name
			oneLine["Country"] = province.Country.Name
			oneLine["ID"] = province.ID
			oneLine["id"] = province.ID

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

func (ctl *ProvinceController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-province"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "address/province_list_search.html"
}
