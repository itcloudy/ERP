package product

import (
	"encoding/json"
	"goERP/controllers/base"

	md "goERP/models"
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
	ctl.Data["PageName"] = ctl.PageName + "\\" + ctl.PageAction
	ctl.URL = "/product/product/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductProductActive"] = "active"

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
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	productInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if product, err := md.GetProductProductByID(idInt64); err == nil {
				ctl.PageAction = product.Name
				productInfo["name"] = product.Name
				productInfo["defaultCode"] = product.DefaultCode
				productInfo["standardPrice"] = product.DefaultCode

				// 款式类别
				categ := product.Categ
				categValues := make(map[string]string)
				if categ != nil {
					categValues["id"] = strconv.FormatInt(categ.ID, 10)
					categValues["name"] = categ.Name
				}
				productInfo["category"] = categValues
				// 销售第一单位
				firstSaleUom := product.FirstSaleUom
				firstSaleUomValues := make(map[string]string)
				if firstSaleUom != nil {
					firstSaleUomValues["id"] = strconv.FormatInt(firstSaleUom.ID, 10)
					firstSaleUomValues["name"] = firstSaleUom.Name
				}
				productInfo["firstSaleUom"] = firstSaleUomValues
				// 销售第二单位
				secondSaleUom := product.SecondSaleUom
				secondSaleUomValues := make(map[string]string)
				if secondSaleUom != nil {
					secondSaleUomValues["id"] = strconv.FormatInt(secondSaleUom.ID, 10)
					secondSaleUomValues["name"] = secondSaleUom.Name
				}
				productInfo["secondSaleUom"] = secondSaleUomValues
				// 采购第一单位
				firstPurchaseUom := product.FirstPurchaseUom
				firstPurchaseUomValues := make(map[string]string)
				if firstPurchaseUom != nil {
					firstPurchaseUomValues["id"] = strconv.FormatInt(firstPurchaseUom.ID, 10)
					firstPurchaseUomValues["name"] = firstPurchaseUom.Name
				}
				productInfo["firstPurchaseUom"] = firstSaleUomValues
				// 采购第二单位
				secondPurchaseUom := product.SecondPurchaseUom
				secondPurchaseUomValues := make(map[string]string)
				if secondSaleUom != nil {
					secondPurchaseUomValues["id"] = strconv.FormatInt(secondPurchaseUom.ID, 10)
					secondPurchaseUomValues["name"] = secondPurchaseUom.Name
				}
				productInfo["secondPurchaseUom"] = secondPurchaseUomValues
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["Product"] = productInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_product_form.html"
}
func (ctl *ProductProductController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductProductController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
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
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *ProductProductController) productProductList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductProduct
	paginator, arrs, err := md.GetAllProductProduct(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["ID"] = line.ID
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
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if result, err := ctl.productProductList(query, fields, sortby, order, offset, limit); err == nil {
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
