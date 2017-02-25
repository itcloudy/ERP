package base

import (
	"bytes"
	"encoding/json"
	md "goERP/models"

	"fmt"
	"strconv"
	"strings"
)

type PartnerController struct {
	BaseController
}

func (ctl *PartnerController) Post() {
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
func (ctl *PartnerController) Put() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	partner := new(md.Partner)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), partner); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if id, err = md.UpdatePartner(partner, &ctl.User); err == nil {
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
func (ctl *PartnerController) Get() {
	style := ctl.Input().Get("type")
	switch style {
	case "customer":
		ctl.PageName = "客户管理"
		ctl.Data["MenuCustomerActive"] = "active"
		ctl.Data["IsCustomer"] = true

	case "supplier":
		ctl.PageName = "供应商管理"
		ctl.Data["MenuSupplierActive"] = "active"
		ctl.Data["IsCustomer"] = false

	}
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
	ctl.URL = "/partner/"
	ctl.Data["URL"] = ctl.URL
}
func (ctl *PartnerController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if partner, err := md.GetPartnerByID(idInt64); err == nil {
				ctl.PageAction = partner.Name
				ctl.Data["Partner"] = partner
			}
		}
	}
	ctl.Data["FormField"] = "form-edit"
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"

	ctl.TplName = "partner/partner_form.html"
}

func (ctl *PartnerController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

//post请求创建产品分类
func (ctl *PartnerController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	fmt.Println(postData)
	partner := new(md.Partner)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), partner); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(partner)).Type().Name()
		if id, err = md.AddPartner(partner, &ctl.User); err == nil {
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
func (ctl *PartnerController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Data["FormField"] = "form-create"
	ctl.Layout = "base/base.html"
	ctl.TplName = "partner/partner_form.html"
}
func (ctl *PartnerController) Validator() {
	name := ctl.GetString("name")
	recordID, _ := ctl.GetInt64("recordID")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := md.GetPartnerByName(name)
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
func (ctl *PartnerController) partnerList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.Partner
	paginator, arrs, err := md.GetAllPartner(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["IsCompany"] = line.IsCompany
			oneLine["IsSupplier"] = line.IsSupplier
			oneLine["IsCustomer"] = line.IsCustomer
			oneLine["Mobile"] = line.Mobile
			oneLine["Tel"] = line.Tel
			oneLine["Email"] = line.Email
			oneLine["Qq"] = line.Qq
			oneLine["WeChat"] = line.WeChat
			if line.Parent != nil {
				parent := make(map[string]interface{})
				parent["id"] = line.Parent.ID
				parent["name"] = line.Parent.Name
				oneLine["Parent"] = parent
			}
			b := bytes.Buffer{}
			if line.Country != nil {
				country := make(map[string]interface{})
				country["id"] = line.Country.ID
				country["name"] = line.Country.Name
				oneLine["Country"] = country
				b.WriteString(line.Country.Name)
			}
			if line.Province != nil {
				province := make(map[string]interface{})
				province["id"] = line.Province.ID
				province["name"] = line.Province.Name
				oneLine["Province"] = province
				b.WriteString(line.Province.Name)
			}
			if line.City != nil {
				city := make(map[string]interface{})
				city["id"] = line.City.ID
				city["name"] = line.City.Name
				oneLine["City"] = city
				b.WriteString(line.City.Name)
			}
			if line.District != nil {
				district := make(map[string]interface{})
				district["id"] = line.District.ID
				district["name"] = line.District.Name
				oneLine["District"] = district
				b.WriteString(line.District.Name)
			}
			oneLine["Street"] = line.Street
			b.WriteString(line.Street)
			oneLine["Address"] = b.String()

			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
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
func (ctl *PartnerController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	condAnd := make(map[string]interface{})
	condOr := make(map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	orderStr := ctl.GetString("order")
	sortStr := ctl.GetString("sort")
	if name := strings.TrimSpace(ctl.GetString("Name")); name != "" {
		condAnd["Name.icontains"] = name
	}
	if isCustomer, err := ctl.GetBool("IsCustomer"); err == nil {
		condAnd["IsCustomer"] = isCustomer
	}
	if isSupplier, err := ctl.GetBool("IsSupplier"); err == nil {
		condAnd["IsSupplier"] = isSupplier
	}
	if orderStr != "" && sortStr != "" {
		sortby = append(sortby, sortStr)
		order = append(order, orderStr)
	} else {
		sortby = append(sortby, "Id")
		order = append(order, "desc")
	}
	if len(condAnd) > 0 {
		cond["and"] = condAnd
	}
	if len(condOr) > 0 {
		cond["or"] = condOr
	}
	if result, err := ctl.partnerList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *PartnerController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-partner"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "partner/partner_list_search.html"
}
