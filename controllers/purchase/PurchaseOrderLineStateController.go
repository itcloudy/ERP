package purchase

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type PurchaseOrderLineStateController struct {
	base.BaseController
}

// Post request
func (ctl *PurchaseOrderLineStateController) Post() {
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
func (ctl *PurchaseOrderLineStateController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/purchase/order/line/state/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if state, err := md.GetPurchaseOrderLineStateByID(idInt64); err == nil {
			if err := ctl.ParseForm(&state); err == nil {

				if err := md.UpdatePurchaseOrderLineStateByID(state); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *PurchaseOrderLineStateController) Get() {
	ctl.PageName = "采购订单明细状态管理"
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
	ctl.URL = "/purchase/order/line/state/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuPurchaseOrderLineStateActive"] = "active"
}

// Edit edit purchase order line state
func (ctl *PurchaseOrderLineStateController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	stateInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if state, err := md.GetPurchaseOrderLineStateByID(idInt64); err == nil {
				ctl.PageAction = state.Name
				stateInfo["name"] = state.Name
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["state"] = stateInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "purchase/purchase_order_line_state_form.html"
}

// Create display purchase order line state page
func (ctl *PurchaseOrderLineStateController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "purchase/purchase_order_line_state_form.html"
}

// Detail display purchase order line state info
func (ctl *PurchaseOrderLineStateController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create purchase order line state
func (ctl *PurchaseOrderLineStateController) PostCreate() {
	state := new(md.PurchaseOrderLineState)
	if err := ctl.ParseForm(state); err == nil {

		if id, err := md.AddPurchaseOrderLineState(state); err == nil {
			ctl.Redirect("/purchase/order/line/state/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *PurchaseOrderLineStateController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetPurchaseOrderLineStateByName(name)
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

//PurchaseOrderLineStateList 获得符合要求的数据
func (ctl *PurchaseOrderLineStateController) PurchaseOrderLineStateList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.PurchaseOrderLineState
	paginator, arrs, err := md.GetAllPurchaseOrderLineState(query, exclude, condMap, fields, sortby, order, offset, limit)
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
func (ctl *PurchaseOrderLineStateController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 1, 1)
	order := make([]string, 1, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	orderStr := ctl.GetString("order")
	sortStr := ctl.GetString("sort")
	if orderStr != "" && sortStr != "" {
		sortby[0] = sortStr
		order[0] = orderStr
	} else {
		sortby[0] = "Id"
		order[0] = "desc"
	}
	if result, err := ctl.PurchaseOrderLineStateList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display purchase order line state with list
func (ctl *PurchaseOrderLineStateController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-purchase-order-line-state"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "purchase/purchase_order_line_state_list_search.html"
}
