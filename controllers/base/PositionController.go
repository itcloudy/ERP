package base

import (
	"encoding/json"
	md "goERP/models"
	"strings"
)

type PositionController struct {
	BaseController
}

func (ctl *PositionController) Post() {

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

func (ctl *PositionController) Get() {
	ctl.GetList()
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/position/"
	ctl.Data["URL"] = ctl.URL
	ctl.Layout = "base/base.html"
	ctl.Data["MenuPositionActive"] = "active"
}

func (ctl *PositionController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := md.GetPositionByName(name)
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
func (ctl *PositionController) positionList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var positions []md.Position
	paginator, positions, err := md.GetAllPosition(query, fields, sortby, order, offset, limit)

	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, position := range positions {
			oneLine := make(map[string]interface{})

			oneLine["Id"] = position.Id
			oneLine["id"] = position.Id
			oneLine["name"] = position.Name

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
func (ctl *PositionController) PostList() {
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
	if result, err := ctl.positionList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *PositionController) GetList() {
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-position"
	ctl.TplName = "base/base_list_view.html"
}
