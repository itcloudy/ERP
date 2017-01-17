package product

import (
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type ProductAttributeController struct {
	base.BaseController
}

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
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/product/attribute/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuProductAttributeActive"] = "active"
}
func (ctl *ProductAttributeController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	attributeInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if attribute, err := md.GetProductAttributeByID(idInt64); err == nil {
				ctl.PageAction = attribute.Name
				attributeInfo["name"] = attribute.Name
				attributeInfo["code"] = attribute.Code
				attributeInfo["sequence"] = attribute.Sequence
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["Attribute"] = attributeInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_form.html"
}
func (ctl *ProductAttributeController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_form.html"
}
func (ctl *ProductAttributeController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductAttributeController) PostCreate() {
	attribute := new(md.ProductAttribute)
	if err := ctl.ParseForm(attribute); err == nil {

		if id, err := md.AddProductAttribute(attribute); err == nil {
			ctl.Redirect("/product/attribute/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}
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
func (ctl *ProductAttributeController) productAttributeList(query map[string]interface{}, exclude map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductAttribute
	paginator, arrs, err := md.GetAllProductAttribute(query, exclude, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["code"] = line.Code
			oneLine["sequence"] = line.Sequence
			mapValues := make(map[int64]string)
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
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
func (ctl *ProductAttributeController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
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
	if result, err := ctl.productAttributeList(query, exclude, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

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
