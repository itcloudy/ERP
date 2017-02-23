package sale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type SaleCounterController struct {
	base.BaseController
}

func (ctl *SaleCounterController) Post() {
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
func (ctl *SaleCounterController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/sale/counter/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if counter, err := md.GetSaleCounterByID(idInt64); err == nil {
			if err := ctl.ParseForm(&counter); err == nil {

				if err := md.UpdateSaleCounterByID(counter); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}
func (ctl *SaleCounterController) Get() {
	ctl.PageName = "柜台管理"
	action := ctl.Input().Get("action")
	switch action {
	case "create":
		ctl.Create()
	case "edit":
		ctl.Edit()
	case "detail":
		ctl.Detail()
	case "table":
		ctl.GetList()
	case "kanban":
		ctl.GetKanban()
	default:
		ctl.GetKanban()
	}
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()
	ctl.URL = "/sale/counter/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuSaleCounterActive"] = "active"
}
func (ctl *SaleCounterController) GetKanban() {

	ctl.PageAction = "看板"
	ctl.Data["KanbanId"] = "kanban-sale-counter"
	ctl.TplName = "base/base_kanban_view.html"
}
func (ctl *SaleCounterController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if counter, err := md.GetSaleCounterByID(idInt64); err == nil {
				ctl.PageAction = counter.Name
				ctl.Data["SaleCounter"] = counter
			}
		}
	}
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"

	ctl.TplName = "sale/sale_counter_form.html"
}

func (ctl *SaleCounterController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

//post请求创建产品分类
func (ctl *SaleCounterController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	counter := new(md.SaleCounter)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), counter); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(counter)).Type().Name()
		if id, err = md.AddSaleCounter(counter, &ctl.User); err == nil {
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
func (ctl *SaleCounterController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"
	ctl.TplName = "sale/sale_counter_form.html"
}
func (ctl *SaleCounterController) Validator() {
	name := ctl.GetString("Name")

	recordID, _ := ctl.GetInt64("recordID")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	query["Name"] = name
	if company, err := ctl.GetInt64("company"); err == nil {
		query["Company.Id"] = company
	}
	if _, arrs, err := md.GetAllSaleCounter(query, exclude, cond, fields, sortby, order, 0, 2); err == nil {
		if len(arrs) == 1 {
			if arrs[0].ID == recordID {
				result["valid"] = true
			} else {
				result["valid"] = false
			}
		} else {
			result["valid"] = true
		}
	} else {
		fmt.Println(err)
		result["valid"] = true
	}

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *SaleCounterController) SaleCounterList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.SaleCounter
	paginator, arrs, err := md.GetAllSaleCounter(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["ProductsCount"] = line.ProductsCount
			oneLine["TemplatesCount"] = line.TemplatesCount
			oneLine["Description"] = line.Description
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			if line.Company != nil {
				company := make(map[string]interface{})
				company["id"] = line.Company.ID
				company["name"] = line.Company.Name
				oneLine["Company"] = company
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
func (ctl *SaleCounterController) PostList() {
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
	if result, err := ctl.SaleCounterList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *SaleCounterController) GetList() {
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-sale-counter"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "sale/sale_counter_list_search.html"
}
