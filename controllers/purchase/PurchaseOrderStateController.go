package purchase

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type PurchaseOrderStateController struct {
	base.BaseController
}

// Post request
func (ctl *PurchaseOrderStateController) Post() {
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
func (ctl *PurchaseOrderStateController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/purchase/order/state/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if state, err := md.GetPurchaseOrderStateByID(idInt64); err == nil {
			if err := ctl.ParseForm(&state); err == nil {

				if err := md.UpdatePurchaseOrderStateByID(state); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *PurchaseOrderStateController) Get() {
	ctl.PageName = "采购订单状态管理"
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
	ctl.URL = "/purchase/order/state/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuPurchaseOrderStateActive"] = "active"
}

// Edit edite purchase order state
func (ctl *PurchaseOrderStateController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	stateInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if state, err := md.GetPurchaseOrderStateByID(idInt64); err == nil {
				ctl.PageAction = state.Name
				stateInfo["name"] = state.Name
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["state"] = stateInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "purchase/purchase_order_state_form.html"
}

// Create dislplay create page
func (ctl *PurchaseOrderStateController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "purchase/purchase_order_state_form.html"
}

// Detail display purchase order state info
func (ctl *PurchaseOrderStateController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post params create purchase order state
func (ctl *PurchaseOrderStateController) PostCreate() {
	state := new(md.PurchaseOrderState)
	if err := ctl.ParseForm(state); err == nil {

		if id, err := md.AddPurchaseOrderState(state); err == nil {
			ctl.Redirect("/purchase/order/state/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *PurchaseOrderStateController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetPurchaseOrderStateByName(name)
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

//PurchaseOrderStateList 获得符合要求的数据
func (ctl *PurchaseOrderStateController) PurchaseOrderStateList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.PurchaseOrderState
	paginator, arrs, err := md.GetAllPurchaseOrderState(query, exclude, condMap, fields, sortby, order, offset, limit)
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
func (ctl *PurchaseOrderStateController) PostList() {
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
	if result, err := ctl.PurchaseOrderStateList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display purchase order state with list
func (ctl *PurchaseOrderStateController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-purchase-order-state"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "purchase/purchase_order_state_list_search.html"
}
