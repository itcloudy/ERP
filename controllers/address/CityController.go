package address

import (
	"encoding/json"
	cb "goERP/controllers/base"
	md "goERP/models"
	"strings"
)

type CityController struct {
	cb.BaseController
}

func (ctl *CityController) Post() {
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
func (ctl *CityController) Get() {
	ctl.PageName = "城市管理"
	ctl.URL = "/address/city/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuCityActive"] = "active"
	ctl.GetList()
	ctl.PageAction = "列表"
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction

}
func (ctl *CityController) Validator() {
	name := strings.TrimSpace(ctl.GetString("Name"))
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := md.GetAddressCityByName(name)
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

// 获得符合要求的城市数据
func (ctl *CityController) cityList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var cities []md.AddressCity
	paginator, cities, err := md.GetAllAddressCity(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, city := range cities {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = city.Name
			oneLine["Province"] = city.Province.Name
			oneLine["Country"] = city.Province.Country.Name
			oneLine["Id"] = city.Id
			oneLine["id"] = city.Id
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
func (ctl *CityController) PostList() {
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
	if result, err := ctl.cityList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *CityController) GetList() {

	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-city"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "address/city_list_search.html"
}
