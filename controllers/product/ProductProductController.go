package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"sort"
	"strconv"
	"strings"
)

type ProductProductController struct {
	base.BaseController
}

func (ctl *ProductProductController) Post() {
	action := ctl.Input().Get("action")
	switch action {
	case "validator":
		ctl.Validator()
	case "table": //bootstrap table的post请求
		ctl.PostList()
	case "create":
		ctl.PostCreate()
	case "batchUpdate":
		ctl.PostBatchUpdate()
	default:
		ctl.PostList()
	}
}
func (ctl *ProductProductController) Get() {
	ctl.PageName = "产品规格管理"
	ctl.URL = "/product/product/"
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
	ctl.URL = "/product/product/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductProductActive"] = "active"

}

// PostBatchUpdate 批量操作
func (ctl *ProductProductController) PostBatchUpdate() {
	result := make(map[string]interface{})
	field := ctl.GetString("field")
	updateFields := make(map[string]interface{})
	result["code"] = "failed"
	var ids []int64
	switch field {
	case "Active":
		if active, err := ctl.GetBool("value"); err == nil {
			updateFields["Active"] = active
		} else {
			result["debug"] = err.Error()
		}
	}
	if len(updateFields) > 0 {
		idsStr := ctl.GetStrings("ids[]")
		for _, idStr := range idsStr {
			if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			query := make(map[string]interface{})
			query["Id.in"] = ids
			if err := md.BatchUpdateProductProduct(query, updateFields); err == nil {
				result["code"] = "success"
			} else {
				result["debug"] = err.Error()
			}
		}
	} else {
		result["debug"] = "参数不正确，没有可更新的信息"
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *ProductProductController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	product := new(md.ProductProduct)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), product); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(product)).Type().Name()
		if id, err = md.AddProductProduct(product, &ctl.User); err == nil {
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
func (ctl *ProductProductController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/product/"
	//需要判断文件上传时页面不用跳转的情况
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if product, err := md.GetProductProductByID(idInt64); err == nil {
			if err := ctl.ParseForm(&product); err == nil {

				if err := md.UpdateProductProductByID(product); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}
func (ctl *ProductProductController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-create"
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if product, err := md.GetProductProductByID(idInt64); err == nil {
				ctl.PageAction = product.Name
				ctl.Data["Product"] = product
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-edit"
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductProductController) Validator() {
	name := strings.TrimSpace(ctl.GetString("name"))
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	// 默认验证成功
	result["valid"] = true
	AttributeValueIds := ctl.GetStrings("AttributeValueIds[]")
	if name != "" {
		obj, err := md.GetProductProductByName(name)
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
	} else if len(AttributeValueIds) > 0 {
		sort.Strings(AttributeValueIds)
		if productTemplateID, err := ctl.GetInt64("ProductTemplateID"); err == nil {
			query := make(map[string]interface{})
			exclude := make(map[string]interface{})
			cond := make(map[string]map[string]interface{})
			fields := make([]string, 0, 0)
			sortby := make([]string, 0, 1)
			order := make([]string, 0, 1)
			query["ProductTemplate.Id"] = productTemplateID
			query["AttributeValuesString"] = strings.Join(AttributeValueIds, "-")
			if _, arrs, err := md.GetAllProductProduct(query, exclude, cond, fields, sortby, order, 0, 2); err == nil {
				if len(arrs) > 0 {
					if recordID > 0 && len(arrs) == 1 {
						if arrs[0].ID == recordID {
							result["valid"] = true
						} else {
							result["valid"] = false
						}
					} else {
						result["valid"] = false

					}
				} else {
					result["valid"] = true
				}
			} else {
				result["valid"] = false
			}
		}
	}

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductProductController) productProductList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductProduct
	paginator, arrs, err := md.GetAllProductProduct(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["SaleOk"] = line.SaleOk
			oneLine["Active"] = line.Active
			oneLine["DefaultCode"] = line.DefaultCode
			oneLine["ProductType"] = line.ProductType
			if line.Category != nil {
				category := make(map[string]interface{})
				category["id"] = line.Category.ID
				category["name"] = line.Category.Name
				oneLine["Category"] = category
			}
			if line.ProductTemplate != nil {
				productTemplate := make(map[string]interface{})
				productTemplate["id"] = line.ProductTemplate.ID
				productTemplate["name"] = line.ProductTemplate.Name
				oneLine["ProductTemplate"] = productTemplate
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
			mapValues := make(map[int64]string)
			values := line.AttributeValues
			for _, line := range values {
				mapValues[line.ID] = line.Name
			}
			oneLine["AttributeValues"] = mapValues
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
func (ctl *ProductProductController) PostList() {
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
	if result, err := ctl.productProductList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *ProductProductController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-product"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_product_list_search.html"
}
