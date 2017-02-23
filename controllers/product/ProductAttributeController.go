package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// ProductAttributeController 产品属性
type ProductAttributeController struct {
	base.BaseController
}

// Post post请求
func (ctl *ProductAttributeController) Post() {
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

// Put 产品属性put请求，修改属性信息
func (ctl *ProductAttributeController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/attribute/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if attribute, err := md.GetProductAttributeByID(idInt64); err == nil {
			if err := ctl.ParseForm(&attribute); err == nil {

				if err := md.UpdateProductAttributeByID(attribute); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get 产品属性get请求
func (ctl *ProductAttributeController) Get() {
	ctl.PageName = "产品属性管理"
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
	ctl.URL = "/product/attribute/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuProductAttributeActive"] = "active"
}

// Edit 产品属性编辑get请求
func (ctl *ProductAttributeController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if attribute, err := md.GetProductAttributeByID(idInt64); err == nil {
				ctl.PageAction = attribute.Name
				ctl.Data["Attribute"] = attribute

			}
		}
	}
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_form.html"
}

// Create 产品属性创建get请求页面
func (ctl *ProductAttributeController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["FormField"] = "form-create"
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_form.html"
}

// Detail 产品属性信息显示get请求，信息不可修改
func (ctl *ProductAttributeController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate 产品属性post请求创建新属性
func (ctl *ProductAttributeController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	attribute := new(md.ProductAttribute)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), attribute); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(category)).Type().Name()
		if id, err = md.AddProductAttribute(attribute, &ctl.User); err == nil {
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

// Validator 产品属性信息post请求，用于验证
func (ctl *ProductAttributeController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetProductAttributeByName(name)
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

// 获得符合要求的数据
func (ctl *ProductAttributeController) productAttributeList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductAttribute
	paginator, arrs, err := md.GetAllProductAttribute(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["Code"] = line.Code
			oneLine["Sequence"] = line.Sequence
			oneLine["ProductsCount"] = line.ProductsCount
			oneLine["TemplatesCount"] = line.TemplatesCount
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			mapValues := make(map[int64]string)
			values := line.ValueIDs
			for _, line := range values {
				mapValues[line.ID] = line.Name
			}
			oneLine["values"] = mapValues
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

// PostList 产品属性信息post请求，用于获得多条属性信息
func (ctl *ProductAttributeController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	cond := make(map[string]map[string]interface{})

	excludeIdsStr := ctl.GetStrings("exclude[]")
	var excludeIds []int64
	for _, v := range excludeIdsStr {
		if val, err := strconv.ParseInt(v, 10, 64); err == nil {
			excludeIds = append(excludeIds, val)
		}
	}
	if len(excludeIds) > 0 {
		exclude["Id.in"] = excludeIds
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
	if result, err := ctl.productAttributeList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList 产品属性信息get请求，列出产品属性
func (ctl *ProductAttributeController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-attribute"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_attribute_list_search.html"
}
