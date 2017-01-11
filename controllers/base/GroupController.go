package base

import (
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

type GroupController struct {
	BaseController
}

func (ctl *GroupController) Post() {

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

func (ctl *GroupController) Get() {
	ctl.PageName = "权限管理"
	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	default:
		ctl.GetList()
	}
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/group/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuGroupActive"] = "active"
}
func (ctl *GroupController) Edit() {
	id := ctl.Ctx.Input.Param(":id")

	groupInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if group, err := md.GetGroupById(idInt64); err == nil {
				ctl.PageAction = group.Name
				groupInfo["Name"] = group.Name

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Group"] = groupInfo
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_category_form.html"
}
func (ctl *GroupController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *GroupController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_category_form.html"
}
func (ctl *GroupController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := md.GetGroupByName(name)
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
func (ctl *GroupController) groupList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var groups []md.Group
	paginator, groups, err := md.GetAllGroup(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, group := range groups {
			oneLine := make(map[string]interface{})

			oneLine["Id"] = group.Id
			oneLine["id"] = group.Id
			oneLine["name"] = group.Name

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
func (ctl *GroupController) PostList() {
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
	if result, err := ctl.groupList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *GroupController) GetList() {
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-group"
	ctl.TplName = "base/base_list_view.html"
}
