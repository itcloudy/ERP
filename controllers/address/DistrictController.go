package address

import (
	"bytes"
	"encoding/json"
	cb "goERP/controllers/base"
	md "goERP/models"
	"strings"
)

type DistrictController struct {
	cb.BaseController
}

func (ctl *DistrictController) Post() {
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
func (ctl *DistrictController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	name := strings.TrimSpace(ctl.GetString("Name"))
	if name != "" {
		query["Name"] = name
	}
	orderStr := ctl.GetString("order")
	sortStr := ctl.GetString("sort")
	if orderStr != "" && sortStr != "" {
		sortby = append(sortby, sortStr)
		order = append(order, orderStr)
	} else {
		sortby = append(sortby, "Id")
		order = append(order, "desc")
	}
	if result, err := ctl.districtList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// 获得符合要求的地区数据
func (ctl *DistrictController) districtList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var districtes []md.AddressDistrict
	paginator, districtes, err := md.GetAllAddressDistrict(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {
		provinceMap := make(map[int64]string)
		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, district := range districtes {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = district.Name
			oneLine["Province"] = district.City.Province.Name

			provinceID := district.City.Province.ID
			if _, ok := provinceMap[provinceID]; ok != true {
				if province, e := md.GetAddressProvinceByID(district.City.Province.ID); e == nil {
					provinceMap[provinceID] = province.Country.Name
				}
			}
			if _, ok := provinceMap[provinceID]; ok {
				oneLine["Country"] = provinceMap[provinceID]
			}
			oneLine["City"] = district.City.Name
			oneLine["ID"] = district.ID
			oneLine["id"] = district.ID
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
func (ctl *DistrictController) Validator() {

	name := strings.TrimSpace(ctl.GetString("Name"))
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetAddressDistrictByName(name)
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

func (ctl *DistrictController) Get() {

	ctl.PageName = "区县管理"
	ctl.URL = "/address/district/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuDistrictActive"] = "active"
	ctl.GetList()
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()

}
func (ctl *DistrictController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-district"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "address/district_list_search.html"
}
