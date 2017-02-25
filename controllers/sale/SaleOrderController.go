package sale

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// SaleOrderController sale order
type SaleOrderController struct {
	base.BaseController
}

// Post request
func (ctl *SaleOrderController) Post() {
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
func (ctl *SaleOrderController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/sale/order/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if order, err := md.GetSaleOrderByID(idInt64); err == nil {
			if err := ctl.ParseForm(&order); err == nil {

				if err := md.UpdateSaleOrderByID(order); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *SaleOrderController) Get() {
	ctl.PageName = "销售订单管理"
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
	ctl.URL = "/sale/order/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuSaleOrderActive"] = "active"
}

// Edit edit sale order
func (ctl *SaleOrderController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if order, err := md.GetSaleOrderByID(idInt64); err == nil {
				ctl.PageAction = order.Name
				ctl.Data["Order"] = order

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_form.html"
}

// Create display sale order create page
func (ctl *SaleOrderController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_form.html"
}

// Detail display sale order info
func (ctl *SaleOrderController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create sale order
func (ctl *SaleOrderController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	saleOrder := new(md.SaleOrder)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), saleOrder); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(category)).Type().Name()
		if id, err = md.AddSaleOrder(saleOrder, &ctl.User); err == nil {
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

// Validator js valid
func (ctl *SaleOrderController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetSaleOrderByName(name)
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

//SaleOrderList 获得符合要求的数据
func (ctl *SaleOrderController) SaleOrderList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.SaleOrder
	paginator, arrs, err := md.GetAllSaleOrder(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["CreateDate"] = line.CreateDate.Format("2006-01-02 15:04:05")
			if line.SalesMan != nil {
				oneLine["SalesMan"] = line.SalesMan.NameZh
			}
			if line.Partner != nil {
				oneLine["Partner"] = line.Partner.Name
			}
			if line.StockWarehouse != nil {
				oneLine["StockWarehouse"] = line.StockWarehouse.Name
			}
			if line.Company != nil {
				oneLine["Company"] = line.Company.Name
			}
			if line.State != nil {
				oneLine["State"] = line.State.Name
			}
			oneLine["PickingPolicy"] = line.PickingPolicy

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
func (ctl *SaleOrderController) PostList() {
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
	if result, err := ctl.SaleOrderList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display sale order with list
func (ctl *SaleOrderController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sale-order"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "sale/sale_order_list_search.html"
}
