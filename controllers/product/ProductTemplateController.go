package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"

	"strconv"
	"strings"
)

type ProductTemplateController struct {
	base.BaseController
}

func (ctl *ProductTemplateController) Post() {
	ctl.URL = "/product/template/"
	ctl.Data["URL"] = ctl.URL
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "attribute":
		ctl.ProductTemplateAttributes()
	case "create":
		ctl.PostCreate()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductTemplateController) Get() {
	ctl.URL = "/product/template/"
	ctl.PageName = "产品款式管理"
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
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductTemplateActive"] = "active"
}

// Put 修改产品款式
func (ctl *ProductTemplateController) Put() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	template := new(md.ProductTemplate)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), template); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if id, err = md.AddProductTemplate(template, &ctl.User); err == nil {
			result["code"] = "success"
			result["location"] = ctl.URL + strconv.FormatInt(id, 10) + "?action=detail"
		} else {
			result["code"] = "failed"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
		}
	}
	if err != nil {
		result["code"] = "failed"
		result["debug"] = err.Error()
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *ProductTemplateController) ProductTemplateAttributes() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if ID, err := ctl.GetInt64("Id"); err == nil {
		query["Id"] = ID
	}
	result := make(map[string]interface{})
	if paginator, arrs, err := md.GetAllProductAttributeLine(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		if jsonResult, er := json.Marshal(&paginator); er == nil {
			result["paginator"] = string(jsonResult)
			result["total"] = paginator.TotalCount
		}
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			attributes := make(map[string]string)
			if line.Attribute != nil {
				attributes["id"] = strconv.FormatInt(line.Attribute.ID, 10)
				attributes["name"] = line.Attribute.Name
			}
			tmpValues := make(map[string]string)
			if line.ProductTemplate != nil {
				tmpValues["id"] = strconv.FormatInt(line.ProductTemplate.ID, 10)
				tmpValues["name"] = line.ProductTemplate.Name
			}
			attributeValuesLines := make([]interface{}, 0, 4)
			attributeValues := line.AttributeValues
			if attributeValues != nil {
				for _, line := range attributeValues {
					mapAttributeValues := make(map[string]string)
					mapAttributeValues["id"] = strconv.FormatInt(line.ID, 10)
					mapAttributeValues["name"] = line.Name
					attributeValuesLines = append(attributeValuesLines, oneLine)
				}

			}
			oneLine["Attribute"] = attributes
			oneLine["ProductTemplate"] = tmpValues
			oneLine["AttributeValues"] = attributeValuesLines

			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			tableLines = append(tableLines, oneLine)
		}
		result["data"] = tableLines
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()

}
func (ctl *ProductTemplateController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	template := new(md.ProductTemplate)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), template); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if id, err = md.AddProductTemplate(template, &ctl.User); err == nil {
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
func (ctl *ProductTemplateController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if template, err := md.GetProductTemplateByID(idInt64); err == nil {
				ctl.PageAction = template.Name
				ctl.Data["Tp"] = template
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["FormField"] = "form-edit"
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_template_form.html"
}
func (ctl *ProductTemplateController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["FormTreeField"] = "form-tree-edit"
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductTemplateController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["FormTreeField"] = "form-tree-create"
	ctl.TplName = "product/product_template_form.html"
}

func (ctl *ProductTemplateController) Validator() {
	name := strings.TrimSpace(ctl.GetString("Name"))
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetProductTemplateByName(name)
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

// 获得符合要求的款式数据
func (ctl *ProductTemplateController) productTemplateList(query map[string]interface{}, exclude map[string]interface{}, cond map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductTemplate
	paginator, arrs, err := md.GetAllProductTemplate(query, exclude, cond, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["Sequence"] = line.Sequence
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["SaleOk"] = line.SaleOk
			oneLine["Active"] = line.Active
			oneLine["DefaultCode"] = line.DefaultCode
			oneLine["ProductMethod"] = line.ProductMethod
			oneLine["ProductType"] = line.ProductType
			oneLine["VariantCount"] = line.VariantCount
			if line.Category != nil {
				category := make(map[string]interface{})
				category["id"] = line.Category.ID
				category["name"] = line.Category.Name
				oneLine["Category"] = category
			}
			if line.FirstSaleUom != nil {
				firstSaleUom := make(map[string]interface{})
				firstSaleUom["id"] = line.FirstSaleUom.ID
				firstSaleUom["name"] = line.FirstSaleUom.Name
				oneLine["FirstSaleUom"] = firstSaleUom
			}
			if line.SecondSaleUom != nil {
				secondSaleUom := make(map[string]interface{})
				secondSaleUom["id"] = line.SecondSaleUom.ID
				secondSaleUom["name"] = line.SecondSaleUom.Name
				oneLine["SecondSaleUom"] = secondSaleUom
			}
			if line.FirstPurchaseUom != nil {
				firstPurchaseUom := make(map[string]interface{})
				firstPurchaseUom["id"] = line.FirstPurchaseUom.ID
				firstPurchaseUom["name"] = line.FirstPurchaseUom.Name
				oneLine["FirstPurchaseUom"] = firstPurchaseUom
			}
			if line.SecondPurchaseUom != nil {
				secondPurchaseUom := make(map[string]interface{})
				secondPurchaseUom["id"] = line.SecondPurchaseUom.ID
				secondPurchaseUom["name"] = line.SecondPurchaseUom.Name
				oneLine["SecondPurchaseUom"] = secondPurchaseUom
			}
			attributeLines := line.AttributeLines
			if len(attributeLines) > 0 {
				attributeMap := make([]string, 0, 0)
				for _, item := range attributeLines {
					attributeMap = append(attributeMap, item.Attribute.Name)
				}
				oneLine["Attributes"] = attributeMap
			}
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
func (ctl *ProductTemplateController) PostList() {
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
	if name := strings.TrimSpace(ctl.GetString("Name")); name != "" {
		condAnd["Name.icontains"] = name
	}
	if defaultCode := strings.TrimSpace(ctl.GetString("DefaultCode")); defaultCode != "" {
		condOr["DefaultCode.icontains"] = defaultCode
	}
	filter := ctl.GetString("filter")
	if filter != "" {
		json.Unmarshal([]byte(filter), &filterMap)
	}
	// 对filterMap进行判断
	if filterActive, ok := filterMap["Active"]; ok {
		condAnd["Active"] = filterActive
	}
	if filterSaleOk, ok := filterMap["SaleOk"]; ok {
		condAnd["SaleOk"] = filterSaleOk
	}
	if filterName, ok := filterMap["Name"]; ok {
		filterName = strings.TrimSpace(filterName.(string))
		if filterName != "" {
			condAnd["Name.icontains"] = filterName
		}
	}
	if filterDefaultCode, ok := filterMap["DefaultCode"]; ok {

		filterDefaultCode = strings.TrimSpace(filterDefaultCode.(string))
		if filterDefaultCode != "" {
			condAnd["DefaultCode.icontains"] = filterDefaultCode
		}
	}
	if filterCategory, ok := filterMap["Category"]; ok {
		filterCategoryID := int64(filterCategory.(float64))
		if filterCategoryID > 0 {
			lineIdsArr := make([]int64, 0, 0)
			lineIdsArr = append(lineIdsArr, filterCategoryID)
			if _, arrs, err := md.GetAllChildCategorys(filterCategoryID); err == nil {
				for _, item := range arrs {
					lineIdsArr = append(lineIdsArr, item.ID)
				}
			}
			condAnd["Category.in"] = lineIdsArr
		}
	}
	if filterProductType, ok := filterMap["ProductType"]; ok {
		filterProductType = strings.TrimSpace(filterProductType.(string))
		if filterProductType != "" {
			condAnd["ProductType"] = filterProductType
		}
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
	if result, err := ctl.productTemplateList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductTemplateController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-template"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_template_list_search.html"
}
