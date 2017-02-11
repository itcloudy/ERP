package sale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
)

// SaleCounterProductController counterProduct
type SaleCounterProductController struct {
	base.BaseController
}

// Put request
func (ctl *SaleCounterProductController) Put() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	ctl.URL = "/sale/counter/product/"
	counterProduct := new(md.SaleCounterProduct)
	var (
		err error
	)
	if err = json.Unmarshal([]byte(postData), counterProduct); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if err = md.UpdateSaleCounterProduct(counterProduct, &ctl.User); err == nil {
			result["code"] = "success"
			result["location"] = ctl.URL + strconv.FormatInt(counterProduct.ID, 10) + "?action=detail"
		} else {
			result["code"] = "failed"
			result["message"] = "数据更新失败"
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

// Get request
func (ctl *SaleCounterProductController) Get() {
	ctl.PageName = "柜台产品管理"
	ctl.URL = "/sale/counter/product/"
	ctl.Data["URL"] = ctl.URL

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
	ctl.Data["MenuSaleCounterProductActive"] = "active"

}

// Post request
func (ctl *SaleCounterProductController) Post() {
	action := ctl.Input().Get("action")
	ctl.URL = "/sale/counter/product/"
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

// Create get counterProduct create page
func (ctl *SaleCounterProductController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_counter_product_form.html"
}

// Detail display counterProduct info
func (ctl *SaleCounterProductController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// GetList display counterProduct with list
func (ctl *SaleCounterProductController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sale-counter-product"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "sale/sale_counter_product_list_search.html"
}

// Validator js valid
func (ctl *SaleCounterProductController) Validator() {

	result := make(map[string]bool)
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// PostList post request json response
func (ctl *SaleCounterProductController) PostList() {
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

	filter := ctl.GetString("filter")
	if filter != "" {
		json.Unmarshal([]byte(filter), &filterMap)
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
	if result, err := ctl.counterProductList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}
func (ctl *SaleCounterProductController) counterProductList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var counterProducts []md.SaleCounterProduct
	paginator, counterProducts, err := md.GetAllSaleCounterProduct(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, counterProduct := range counterProducts {

			oneLine := make(map[string]interface{})
			oneLine["ID"] = counterProduct.ID
			oneLine["id"] = counterProduct.ID

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

//PostCreate create counterProduct with post params
func (ctl *SaleCounterProductController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	fmt.Printf("%+v\n", postData)
	counterProduct := new(md.SaleCounterProduct)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), counterProduct); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if id, err = md.AddSaleCounterProduct(counterProduct, &ctl.User); err == nil {
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

// Edit edit counterProduct info
func (ctl *SaleCounterProductController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if counterProduct, err := md.GetSaleCounterProductByID(idInt64); err == nil {
				ctl.Data["SaleCounterProduct"] = counterProduct
				ctl.PageAction = counterProduct.SaleCounter.Name
			}
		}
	}
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["RecordID"] = id
	ctl.Data["Action"] = "edit"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_counter_product_form.html"
}
