package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type ProductUomController struct {
	base.BaseController
}

func (ctl *ProductUomController) Post() {
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
func (ctl *ProductUomController) Get() {
	ctl.PageName = "单位管理"
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
	ctl.URL = "/product/uom/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductUomActive"] = "active"
}
func (ctl *ProductUomController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/uom/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if uom, err := md.GetProductUomByID(idInt64); err == nil {
			if err := ctl.ParseForm(&uom); err == nil {

				if err := md.UpdateProductUomByID(uom); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}

func (ctl *ProductUomController) Validator() {
	name := ctl.GetString("name")
	recordID, _ := ctl.GetInt64("recordID")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := md.GetProductUomByName(name)
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
func (ctl *ProductUomController) PostCreate() {

	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	uom := new(md.ProductUom)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), uom); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(uom)).Type().Name()
		if id, err = md.AddProductUom(uom, &ctl.User); err == nil {
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

	if err := ctl.ParseForm(uom); err == nil {
		if uomCategID, err := ctl.GetInt64("category"); err == nil {
			if category, err := md.GetProductUomCategByID(uomCategID); err == nil {
				uom.Category = category
				if id, err := md.AddProductUom(uom, &ctl.User); err == nil {
					ctl.Redirect("/product/uom/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Get()

}
func (ctl *ProductUomController) productUomList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.ProductUom
	paginator, arrs, err := md.GetAllProductUom(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["active"] = line.Active
			oneLine["rounding"] = line.Rounding
			oneLine["symbol"] = line.Symbol
			switch line.Type {
			case 1:
				oneLine["type"] = "小于参考计量单位"
				oneLine["factor"] = line.Factor
			case 2:
				oneLine["type"] = "参考计量单位"
			case 3:
				oneLine["type"] = "大于参考计量单位"
				oneLine["factorInv"] = line.FactorInv
			default:
				oneLine["type"] = "参考计量单位"
			}

			oneLine["category"] = line.Category.Name
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
func (ctl *ProductUomController) PostList() {
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
	if result, err := ctl.productUomList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()
}
func (ctl *ProductUomController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if uom, err := md.GetProductUomByID(idInt64); err == nil {
				ctl.PageAction = uom.Name
				switch uom.Type {
				case 1:
					uom.TypeName = "小于参考计量单位"
				case 2:
					uom.TypeName = "参考计量单位"
				case 3:
					uom.TypeName = "大于参考计量单位"
				default:
					uom.TypeName = "参考计量单位"
				}
				ctl.Data["Uom"] = uom
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.TplName = "product/product_uom_form.html"
}
func (ctl *ProductUomController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductUomController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-uom"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_uom_list_search.html"
}
func (ctl *ProductUomController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["FormField"] = "form-create"
	ctl.Data["Readonly"] = false
	ctl.Layout = "base/base.html"
	ctl.PageAction = "创建"
	ctl.TplName = "product/product_uom_form.html"
}
