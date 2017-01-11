package base

import (
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

func (ctl *UserController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/user/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if user, err := md.GetUserById(idInt64); err == nil {
			if err := ctl.ParseForm(&user); err == nil {
				var upateField []string
				if departmentId, err := ctl.GetInt64("department"); err == nil {
					if department, err := md.GetDepartmentById(departmentId); err == nil {
						user.Department = department
						upateField = append(upateField, "Department")
					}
				}
				groupIdsStr := ctl.GetStrings("group")
				var groupIds []int64
				for _, el := range groupIdsStr {
					if idInt64, err := strconv.ParseInt(el, 10, 64); err == nil {
						groupIds = append(groupIds, idInt64)
					}
				}
				if len(groupIds) > 0 {
					var groups []*md.Group
					for _, groupId := range groupIds {
						if group, err := md.GetGroupById(groupId); err == nil {
							groups = append(groups, group)
						}
					}
					user.Groups = groups
					upateField = append(upateField, "Groups")
				}
				if positionId, err := ctl.GetInt64("position"); err == nil {
					if position, err := md.GetPositionById(positionId); err == nil {
						user.Position = position
						upateField = append(upateField, "Position")
					}
				}
				if err := md.UpdateUserById(user); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}
func (ctl *UserController) Get() {
	ctl.PageName = "用户管理"
	ctl.URL = "/user/"
	ctl.Data["URL"] = ctl.URL

	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	case "changepasswd":
		ctl.ChangePwd()
	default:
		ctl.GetList()
	}
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction

}
func (ctl *UserController) Post() {
	action := ctl.Input().Get("action")
	ctl.URL = "/user/"
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *UserController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"

	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["MenuSelfInfoActive"] = "active"
	ctl.Data["Action"] = "detail"
}
func (ctl *UserController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-user"
	ctl.Data["MenuUserActive"] = "active"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/user_list_search.html"
}
func (ctl *UserController) Validator() {
	recordID, _ := ctl.GetInt64("recordId")
	name := strings.TrimSpace(ctl.GetString("Name"))
	result := make(map[string]bool)
	obj, err := md.GetUserByName(name)
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
func (ctl *UserController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	name := strings.TrimSpace(ctl.Input().Get("Name"))
	if name != "" {
		query["Name"] = name
	}
	if result, err := ctl.userList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *UserController) userList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var users []md.User
	paginator, users, err := md.GetAllUser(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, user := range users {

			oneLine := make(map[string]interface{})
			oneLine["Name"] = user.Name
			oneLine["NameZh"] = user.NameZh
			if user.Department != nil {
				oneLine["Department"] = user.Department.Name
			} else {
				oneLine["Department"] = "-"
			}
			if user.Position != nil {
				oneLine["Position"] = user.Position.Name
			} else {
				oneLine["Position"] = "-"
			}

			oneLine["Email"] = user.Email
			oneLine["Mobile"] = user.Mobile
			oneLine["Tel"] = user.Tel
			oneLine["IsAdmin"] = user.IsAdmin
			oneLine["Active"] = user.Active
			oneLine["Qq"] = user.Qq
			oneLine["Id"] = user.Id
			oneLine["id"] = user.Id
			oneLine["Wechat"] = user.WeChat

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

func (ctl *UserController) ChangePwd() {
	ctl.Data["MenuChangePwdActive"] = "active"
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_change_password_form.html"
}

func (ctl *UserController) GetCreate() {
	ctl.Data["Readonly"] = false

	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) PostCreate() {

	user := new(md.User)
	if err := ctl.ParseForm(user); err == nil {

		if deparentId, err := ctl.GetInt64("Department"); err == nil {
			if department, err := md.GetDepartmentById(deparentId); err == nil {
				user.Department = department
			}
		}
		if positionId, err := ctl.GetInt64("Position"); err == nil {
			if position, err := md.GetPositionById(positionId); err == nil {
				user.Position = position
			}
		}
		if id, err := md.AddUser(user); err == nil {
			ctl.Redirect(ctl.URL+strconv.FormatInt(id, 10)+"?action=detail", 302)

		}
	}

}
func (ctl *UserController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	userInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if user, err := md.GetUserById(idInt64); err == nil {
				ctl.PageAction = user.Name + "(" + user.NameZh + ")"
				userInfo["Id"] = user.Id
				userInfo["Name"] = user.Name
				userInfo["NameZh"] = user.NameZh
				userInfo["Email"] = user.Email
				userInfo["Mobile"] = user.Mobile
				userInfo["Qq"] = user.Qq
				userInfo["Wechat"] = user.WeChat
				userInfo["Tel"] = user.Tel
				department := make(map[string]string)
				if user.Department != nil {
					department["Id"] = strconv.FormatInt(user.Department.Id, 10)
					department["Name"] = user.Department.Name
					userInfo["Department"] = department
				}
				groups := make([]interface{}, 0, 4)
				for _, group := range user.Groups {
					oneLine := make(map[string]interface{})
					oneLine["Id"] = group.Id
					oneLine["Name"] = group.Name
					oneLine["Description"] = group.Description
					groups = append(groups, oneLine)
				}
				userInfo["Groups"] = groups
				position := make(map[string]string)
				if user.Position != nil {
					position["Id"] = strconv.FormatInt(user.Position.Id, 10)
					position["Name"] = user.Position.Name
					userInfo["Position"] = position
				}
			}
		}
	}
	ctl.Data["RecordId"] = id
	ctl.Data["Action"] = "edit"
	ctl.Data["User"] = userInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
func (ctl *UserController) Show() {
	ctl.Data["MenuSelfInfoActive"] = "active"
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/user_form.html"
}
