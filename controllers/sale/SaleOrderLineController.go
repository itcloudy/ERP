package sale

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// SaleOrderLineController
type SaleOrderLineController struct {
	base.BaseController
}

// Post request
func (ctl *SaleOrderLineController) Post() {
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
func (ctl *SaleOrderLineController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/sale/order/line/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if orderLine, err := md.GetSaleOrderLineByID(idInt64); err == nil {
			if err := ctl.ParseForm(&orderLine); err == nil {

				if err := md.UpdateSaleOrderLineByID(orderLine); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *SaleOrderLineController) Get() {
	ctl.PageName = "销售订单明细管理"
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
	ctl.URL = "/sale/order/line/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuSaleOrderLineActive"] = "active"
}

// Edit edit sale orde line
func (ctl *SaleOrderLineController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if orderLine, err := md.GetSaleOrderLineByID(idInt64); err == nil {
				ctl.PageAction = orderLine.Name
				ctl.Data["orderLine"] = orderLine
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_ine_form.html"
}

// Create display create page
func (ctl *SaleOrderLineController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_order_line_form.html"
}

// Detail display sale order line info
func (ctl *SaleOrderLineController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create sale order line
func (ctl *SaleOrderLineController) PostCreate() {
	orderLine := new(md.SaleOrderLine)
	if err := ctl.ParseForm(orderLine); err == nil {

		if id, err := md.AddSaleOrderLine(orderLine); err == nil {
			ctl.Redirect("/sale/order/line/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *SaleOrderLineController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetSaleOrderLineByName(name)
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

//SaleOrderLineList 获得符合要求的数据
func (ctl *SaleOrderLineController) SaleOrderLineList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.SaleOrderLine
	paginator, arrs, err := md.GetAllSaleOrderLine(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			if line.Product != nil {
				product := make(map[string]interface{})
				product["id"] = line.Product.ID
				product["name"] = line.Product.Name
				product["defaultCode"] = line.Product.DefaultCode
				oneLine["Product"] = product
			}
			oneLine["ProductName"] = line.ProductName
			oneLine["ProductCode"] = line.ProductCode
			oneLine["FirstSaleQty"] = line.FirstSaleQty
			oneLine["FirstUomStep"] = line.FirstSaleUom.Rounding
			oneLine["SecondUomStep"] = line.SecondSaleUom.Rounding
			oneLine["FirstUomName"] = line.FirstSaleUom.Name
			oneLine["SecondUomName"] = line.SecondSaleUom.Name
			oneLine["SecondSaleQty"] = line.SecondSaleQty
			oneLine["PriceUnit"] = line.PriceUnit
			oneLine["Total"] = line.Total
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

// PostList post request list sale order line
func (ctl *SaleOrderLineController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	condAnd := make(map[string]interface{})
	condOr := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if saleOrderID, err := ctl.GetInt64("saleOrderId"); err == nil {
		condAnd["SaleOrder.Id"] = saleOrderID
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
	if len(condAnd) > 0 {
		cond["and"] = condAnd
	}
	if len(condOr) > 0 {
		cond["or"] = condOr
	}
	if result, err := ctl.SaleOrderLineList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display sale order line with list
func (ctl *SaleOrderLineController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sale-order-line"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "sale/sale_order_line_list_search.html"
}
