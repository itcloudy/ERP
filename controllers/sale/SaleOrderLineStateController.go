package sale

import (
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type SaleOrderLineStateController struct {
	base.BaseController
}

// Post request
func (ctl *SaleOrderLineStateController) Post() {
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

// Put request
func (ctl *SaleOrderLineStateController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/sale/order/line/state/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if state, err := md.GetSaleOrderLineStateByID(idInt64); err == nil {
			if err := ctl.ParseForm(&state); err == nil {

				if err := md.UpdateSaleOrderLineStateByID(state); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *SaleOrderLineStateController) Get() {
	ctl.PageName = "销售订单明细状态管理"
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
	ctl.URL = "/sale/order/line/state/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuSaleOrderLineStateActive"] = "active"
}

// Edit edit sale order line state
func (ctl *SaleOrderLineStateController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	stateInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if state, err := md.GetSaleOrderLineStateByID(idInt64); err == nil {
				ctl.PageAction = state.Name
				stateInfo["name"] = state.Name
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["state"] = stateInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_line_state_form.html"
}

// Create display sale order line state page
func (ctl *SaleOrderLineStateController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_line_state_form.html"
}

// Detail display sale order line state info
func (ctl *SaleOrderLineStateController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create sale order line state
func (ctl *SaleOrderLineStateController) PostCreate() {
	state := new(md.SaleOrderLineState)
	if err := ctl.ParseForm(state); err == nil {

		if id, err := md.AddSaleOrderLineState(state); err == nil {
			ctl.Redirect("/sale/order/line/state/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *SaleOrderLineStateController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetSaleOrderLineStateByName(name)
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

//SaleOrderLineStateList 获得符合要求的数据
func (ctl *SaleOrderLineStateController) SaleOrderLineStateList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.SaleOrderLineState
	paginator, arrs, err := md.GetAllSaleOrderLineState(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
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

// PostList post request json response
func (ctl *SaleOrderLineStateController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if result, err := ctl.SaleOrderLineStateList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display sale order line state with list
func (ctl *SaleOrderLineStateController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sale-order-line-state"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "sale/sale_order_line_state_list_search.html"
}
