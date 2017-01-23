package base

import (
	"encoding/json"
	md "goERP/models"
)

// SequenceController 登录日志
type SequenceController struct {
	BaseController
}

// Get 请求
func (ctl *SequenceController) Get() {
	ctl.PageName = "登陆记录管理"
	ctl.URL = "/sequence/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuRecordActive"] = "active"
	ctl.GetList()
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction

}

// Post 请求
func (ctl *SequenceController) Post() {
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
func (ctl *SequenceController) GetOneRecord() {

}

//PostList Post 请求获得登录日志列表json数据
func (ctl *SequenceController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if result, err := ctl.sequenceList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *SequenceController) sequenceList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var sequences []md.Record
	paginator, sequences, err := md.GetAllRecord(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		tableLines := make([]interface{}, 0, 4)
		for _, sequence := range sequences {
			oneLine := make(map[string]interface{})
			oneLine["Email"] = sequence.User.Email
			oneLine["Mobile"] = sequence.User.Mobile
			oneLine["Name"] = sequence.User.Name
			oneLine["NameZh"] = sequence.User.NameZh
			oneLine["UserAgent"] = sequence.UserAgent
			oneLine["CreateDate"] = sequence.CreateDate.Format("2006-01-02 15:04:05")
			oneLine["Logout"] = sequence.Logout.Format("2006-01-02 15:04:05")
			oneLine["Ip"] = sequence.IP
			oneLine["ID"] = sequence.ID
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
func (ctl *SequenceController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sequence"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/sequence_list_search.html"
}
