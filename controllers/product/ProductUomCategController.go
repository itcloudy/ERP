package product

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type ProductUomCategController struct {
	base.BaseController
}

func (ctl *ProductUomCategController) Post() {
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
func (ctl *ProductUomCategController) Get() {
	ctl.PageName = "单位类别管理"
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

	ctl.URL = "/product/uomcateg/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuProductUomCategActive"] = "active"
}
func (ctl *ProductUomCategController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/product/uomcateg/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if uomCateg, err := md.GetProductUomCategByID(idInt64); err == nil {
			if err := ctl.ParseForm(&uomCateg); err == nil {

				if err := md.UpdateProductUomCategByID(uomCateg); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)
}
func (ctl *ProductUomCategController) Validator() {
	name := ctl.GetString("name")
	recordID, _ := ctl.GetInt64("recordID")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := md.GetProductUomCategByName(name)
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
func (ctl *ProductUomCategController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	uomCateg := new(md.ProductUomCateg)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), uomCateg); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(uuomCategom)).Type().Name()
		if id, err = md.AddProductUomCateg(uomCateg, &ctl.User); err == nil {
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
func (ctl *ProductUomCategController) productUomCategList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {
	var arrs []md.ProductUomCateg
	paginator, arrs, err := md.GetAllProductUomCateg(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			uoms := line.Uoms
			mapValues := make(map[int64]string)
			for _, line := range uoms {
				mapValues[line.ID] = line.Name
			}
			oneLine["uoms"] = mapValues
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
func (ctl *ProductUomCategController) PostList() {
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
	if result, err := ctl.productUomCategList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()
}
func (ctl *ProductUomCategController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if categ, err := md.GetProductUomCategByID(idInt64); err == nil {
				ctl.PageAction = categ.Name
				ctl.Data["UomCateg"] = categ
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"

	ctl.TplName = "product/product_uom_categ_form.html"
}
func (ctl *ProductUomCategController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *ProductUomCategController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-product-uom-categ"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "product/product_uom_categ_list_search.html"
}
func (ctl *ProductUomCategController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"
	ctl.PageAction = "创建"
	ctl.TplName = "product/product_uom_categ_form.html"
}
