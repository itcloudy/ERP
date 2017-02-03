package base

import (
	"bytes"
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

// TeamController team
type TeamController struct {
	BaseController
}

// Post request
func (ctl *TeamController) Post() {

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
func (ctl *TeamController) Get() {
	ctl.PageName = "团队管理"
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
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()
	ctl.URL = "/team/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuTeamActive"] = "active"
}

// Edit edit team
func (ctl *TeamController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	teamInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if team, err := md.GetTeamByID(idInt64); err == nil {
				ctl.PageAction = team.Name
				teamInfo["Name"] = team.Name

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["Team"] = teamInfo
	ctl.Layout = "base/base.html"

	ctl.TplName = "base/team_form.html"
}

// Detail dispaly team info
func (ctl *TeamController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// Create get create page
func (ctl *TeamController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false

	ctl.Layout = "base/base.html"
	ctl.TplName = "base/team_form.html"
}

// Validator js valid
func (ctl *TeamController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetTeamByName(name)
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
func (ctl *TeamController) teamList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var teams []md.Team
	paginator, teams, err := md.GetAllTeam(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, team := range teams {
			oneLine := make(map[string]interface{})

			oneLine["ID"] = team.ID
			oneLine["id"] = team.ID
			oneLine["name"] = team.Name

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
func (ctl *TeamController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	name := strings.TrimSpace(ctl.Input().Get("name"))
	if name != "" {
		query["name"] = name
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
	if result, err := ctl.teamList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display team with table
func (ctl *TeamController) GetList() {
	ctl.Data["tableId"] = "table-team"
	ctl.TplName = "base/base_list_view.html"
}
