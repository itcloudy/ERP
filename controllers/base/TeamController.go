package base

import (
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

type TeamController struct {
	BaseController
}

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
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/team/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuTeamActive"] = "active"
}
func (ctl *TeamController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	teamInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if team, err := md.GetTeamById(idInt64); err == nil {
				ctl.PageAction = team.Name
				teamInfo["Name"] = team.Name

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordId"] = id
	ctl.Data["Team"] = teamInfo
	ctl.Layout = "base/base.html"

	ctl.TplName = "base/team_form.html"
}
func (ctl *TeamController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *TeamController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false

	ctl.Layout = "base/base.html"
	ctl.TplName = "base/team_form.html"
}
func (ctl *TeamController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordId")
	result := make(map[string]bool)
	obj, err := md.GetTeamByName(name)
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
func (ctl *TeamController) teamList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var teams []md.Team
	paginator, teams, err := md.GetAllTeam(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, team := range teams {
			oneLine := make(map[string]interface{})

			oneLine["Id"] = team.Id
			oneLine["id"] = team.Id
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
func (ctl *TeamController) PostList() {
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
	if result, err := ctl.teamList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *TeamController) GetList() {
	ctl.Data["tableId"] = "table-team"
	ctl.TplName = "base/base_list_view.html"
}
