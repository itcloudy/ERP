package base

import (
	"bytes"
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

// SourceController 公司
type SourceController struct {
	BaseController
}

func (ctl *SourceController) Post() {
	ctl.URL = "/source/"
	ctl.Data["URL"] = ctl.URL
	action := ctl.Input().Get("action")
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
func (ctl *SourceController) Get() {
	ctl.URL = "/source/"
	ctl.PageName = "系统资源管理"
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
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuSourceActive"] = "active"
}

// Put 修改产品款式
func (ctl *SourceController) Put() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	source := new(md.Source)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), source); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(source)).Type().Name()
		if id, err = md.AddSource(source, &ctl.User); err == nil {
			result["code"] = "success"
			result["location"] = ctl.URL + strconv.FormatInt(id, 10) + "?action=detail"
		} else {
			result["code"] = "failed"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
		}
	}
	if err != nil {
		result["code"] = "failed"
		result["debug"] = err.Error()
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *SourceController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	source := new(md.Source)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), source); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(source)).Type().Name()

		if id, err = md.AddSource(source, &ctl.User); err == nil {
			result["code"] = "success"
			result["location"] = ctl.URL + strconv.FormatInt(id, 10) + "?action=detail"
		} else {
			result["code"] = "failed"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
		}
	} else {
		result["code"] = "failed"
		result["message"] = "请求数据解析失败"
		result["debug"] = err.Error()
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *SourceController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if source, err := md.GetSourceByID(idInt64); err == nil {
				ctl.PageAction = source.Name
				ctl.Data["Source"] = source
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["FormField"] = "form-edit"
	ctl.Layout = "base/base.html"
	ctl.TplName = "security/source_form.html"
}
func (ctl *SourceController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["FormTreeField"] = "form-tree-edit"
	ctl.Data["Action"] = "detail"
}
func (ctl *SourceController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["FormTreeField"] = "form-tree-create"
	ctl.TplName = "security/source_form.html"
}

func (ctl *SourceController) Validator() {
	name := strings.TrimSpace(ctl.GetString("Name"))
	modelName := strings.TrimSpace(ctl.GetString("ModelName"))
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	result["valid"] = true
	if name != "" {
		obj, err := md.GetSourceByName(name)
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
	}
	if modelName != "" {
		obj, err := md.GetSourceByModelName(modelName)
		if err != nil {
			result["valid"] = true
		} else {
			if obj.ModelName == modelName {
				if recordID == obj.ID {
					result["valid"] = true
				} else {
					result["valid"] = false
				}
			} else {
				result["valid"] = true
			}
		}
	}

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的款式数据
func (ctl *SourceController) sourceList(query map[string]interface{}, exclude map[string]interface{}, cond map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.Source
	paginator, arrs, err := md.GetAllSource(query, exclude, cond, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["ModelName"] = line.ModelName
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
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
func (ctl *SourceController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	condAnd := make(map[string]interface{})
	condOr := make(map[string]interface{})
	filterMap := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	if ID, err := ctl.GetInt64("Id"); err == nil {
		query["Id"] = ID
	}
	if name := strings.TrimSpace(ctl.GetString("Name")); name != "" {
		condAnd["Name.icontains"] = name
	}
	if identity := strings.TrimSpace(ctl.GetString("ModelName")); identity != "" {
		condAnd["Identity"] = identity
	}
	filter := ctl.GetString("filter")
	if filter != "" {
		json.Unmarshal([]byte(filter), &filterMap)
	}
	if filterName, ok := filterMap["Name"]; ok {
		filterName = strings.TrimSpace(filterName.(string))
		if filterName != "" {
			condAnd["Name.icontains"] = filterName
		}
	}
	if filterIdentity, ok := filterMap["Identity"]; ok {
		filterIdentity = strings.TrimSpace(filterIdentity.(string))
		if filterIdentity != "" {
			condAnd["Identity"] = filterIdentity
		}
	}
	if len(condAnd) > 0 {
		cond["and"] = condAnd
	}
	if len(condOr) > 0 {
		cond["or"] = condOr
	}
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	orderStr := ctl.GetString("order")
	sortStr := ctl.GetString("sort")
	if orderStr != "" && sortStr != "" {
		sortby = append(sortby, sortStr)
		order = append(order, orderStr)
	} else {
		sortby = append(sortby, "Id")
		order = append(order, "desc")

	}
	if result, err := ctl.sourceList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *SourceController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-source"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "security/source_list_search.html"
}
