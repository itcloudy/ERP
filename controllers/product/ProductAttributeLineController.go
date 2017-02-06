package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// ProductAttributeLineController 款式属性明细
type ProductAttributeLineController struct {
	base.BaseController
}

// Post post请求 post
func (ctl *ProductAttributeLineController) Post() {
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

// Put 修改属性明细
func (ctl *ProductAttributeLineController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/category/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if category, err := md.GetProductAttributeLineByID(idInt64); err == nil {
			if err := ctl.ParseForm(&category); err == nil {

				if err := md.UpdateProductAttributeLineByID(category); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get 显示产品属性明细 get
func (ctl *ProductAttributeLineController) Get() {
	ctl.PageName = "产品类别管理"
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
	ctl.URL = "/product/attributeline/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductAttributeLineActive"] = "active"
}

// Edit 款式属性明细编辑 get
func (ctl *ProductAttributeLineController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if category, err := md.GetProductAttributeLineByID(idInt64); err == nil {

				ctl.Data["Category"] = category
			}
		}
	}
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_attribute_line_form.html"
}

// Detail 显示款式属性明细 get
func (ctl *ProductAttributeLineController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

//PostCreate 创建款式属性明细 post
func (ctl *ProductAttributeLineController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	category := new(md.ProductAttributeLine)
	var (
		err error
	)
	if err = json.Unmarshal([]byte(postData), category); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(category)).Type().Name()

	} else {
		result["code"] = "failed"
		result["message"] = "请求数据解析失败"
		result["debug"] = err.Error()
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// Create 款式属性明细创建页面 get
func (ctl *ProductAttributeLineController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_line_form.html"
}

// Validator 款式属性明细验证 post
func (ctl *ProductAttributeLineController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的款式属性明细数据
func (ctl *ProductAttributeLineController) productAttributeLineList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductAttributeLine
	paginator, arrs, err := md.GetAllProductAttributeLine(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Attribute"] = line.Attribute.Name
			oneLine["ProductTemplate"] = line.ProductTemplate.Name
			oneLine["ProductTemplate.DefaultCode"] = line.ProductTemplate.DefaultCode
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			attributeValueArrs := make([]interface{}, 0, 4)
			attributeValues := line.AttributeValues
			for _, attributeValue := range attributeValues {
				attributeValueMap := make(map[string]interface{})
				attributeValueMap["id"] = attributeValue.ID
				attributeValueMap["name"] = attributeValue.Name
				attributeValueArrs = append(attributeValueArrs, attributeValueMap)
			}
			oneLine["attributeValueArrs"] = attributeValueArrs
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

// PostList 获得多条款式属性明细 post
func (ctl *ProductAttributeLineController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if tmpID, err := ctl.GetInt64("tmpId"); err == nil {
		query["ProductTemplate.Id"] = tmpID
	}
	//排除已经选择的属性
	excludeIdsStr := ctl.GetStrings("exclude[]")
	if len(excludeIdsStr) > 0 {
		attributeIds := make([]int64, 0, 0)
		for _, attributeValueID := range excludeIdsStr {
			if idInt64, e := strconv.ParseInt(attributeValueID, 10, 64); e == nil {
				if productAttributeValue, err := md.GetProductAttributeValueByID(idInt64); err == nil {
					attributeIds = append(attributeIds, productAttributeValue.Attribute.ID)
				}
			}
		}
		if len(attributeIds) > 0 {
			exclude["Attribute.Id"] = attributeIds
		}
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
	if result, err := ctl.productAttributeLineList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()
}

// GetList 多条款式属性明细显示 get
func (ctl *ProductAttributeLineController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-attribute-line"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_attribute_line_search.html"
}
