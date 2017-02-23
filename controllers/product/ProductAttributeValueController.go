package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type ProductAttributeValueController struct {
	base.BaseController
}

func (ctl *ProductAttributeValueController) Post() {
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
func (ctl *ProductAttributeValueController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/attributevalue/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if attrValue, err := md.GetProductAttributeValueByID(idInt64); err == nil {
			if err := ctl.ParseForm(&attrValue); err == nil {
				if attributeID, err := ctl.GetInt64("productAttributeID"); err == nil {
					if attribute, err := md.GetProductAttributeByID(attributeID); err == nil {
						attrValue.Attribute = attribute
					}
				}
				if err := md.UpdateProductAttributeValueByID(attrValue); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}
func (ctl *ProductAttributeValueController) Get() {
	ctl.PageName = "产品属性值管理"
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
	ctl.URL = "/product/attributevalue/"
	ctl.Data["URL"] = ctl.URL
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()
	ctl.Data["MenuProductAttributeValueActive"] = "active"
}
func (ctl *ProductAttributeValueController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_value_form.html"
}
func (ctl *ProductAttributeValueController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if attributeValue, err := md.GetProductAttributeValueByID(idInt64); err == nil {
				ctl.PageAction = attributeValue.Name
				ctl.Data["ProductAttValue"] = attributeValue
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_attribute_value_form.html"
}
func (ctl *ProductAttributeValueController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductAttributeValueController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	attrValue := new(md.ProductAttributeValue)

	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), attrValue); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(attrValue)).Type().Name()
		if id, err = md.AddProductAttributeValue(attrValue, &ctl.User); err == nil {
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
func (ctl *ProductAttributeValueController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetProductAttributeValueByName(name)
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
func (ctl *ProductAttributeValueController) productAttributeValueList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductAttributeValue
	paginator, arrs, err := md.GetAllProductAttributeValue(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["Attribute"] = line.Attribute.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["ProductsCount"] = line.ProductsCount
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
func (ctl *ProductAttributeValueController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if attributeID, err := ctl.GetInt64("attributeId"); err == nil {
		query["Attribute.Id.in"] = attributeID
	}
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
	orderStr := ctl.GetString("order")
	sortStr := ctl.GetString("sort")
	if orderStr != "" && sortStr != "" {
		sortby = append(sortby, sortStr)
		order = append(order, orderStr)
	} else {
		sortby = append(sortby, "Id")
		order = append(order, "desc")
	}
	if result, err := ctl.productAttributeValueList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductAttributeValueController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-attributevalue"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_attribute_value_list_search.html"
}
