package base

import (
	"bytes"
	"encoding/json"
	md "goERP/models"
)

// RecordController 登录日志
type RecordController struct {
	BaseController
}

// Get 请求
func (ctl *RecordController) Get() {
	ctl.PageName = "登陆记录管理"
	ctl.URL = "/record/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuRecordActive"] = "active"
	ctl.GetList()
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()

}

// Post 请求
func (ctl *RecordController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "table":
		ctl.PostList()
	case "one":
		ctl.GetOneRecord()
	default:
		ctl.PostList()
	}
}

// GetOneRecord 获得一条记录
func (ctl *RecordController) GetOneRecord() {

}

//PostList Post 请求获得登录日志列表json数据
func (ctl *RecordController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
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
	if result, err := ctl.recordList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *RecordController) recordList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var records []md.Record
	paginator, records, err := md.GetAllRecord(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		tableLines := make([]interface{}, 0, 4)
		for _, record := range records {
			oneLine := make(map[string]interface{})
			oneLine["Email"] = record.User.Email
			oneLine["Mobile"] = record.User.Mobile
			oneLine["Name"] = record.User.Name
			oneLine["NameZh"] = record.User.NameZh
			oneLine["UserAgent"] = record.UserAgent
			oneLine["CreateDate"] = record.CreateDate.Format("2006-01-02 15:04:05")
			oneLine["Logout"] = record.Logout.Format("2006-01-02 15:04:05")
			oneLine["Ip"] = record.IP
			oneLine["ID"] = record.ID
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

// GetList 显示table数据
func (ctl *RecordController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-record"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/record_list_search.html"
}
