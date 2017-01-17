package base

import (
	"encoding/json"
	md "goERP/models"
	"strconv"
	"strings"
)

// CompanyController 公司
type CompanyController struct {
	BaseController
}

// Post 请求 公司
func (ctl *CompanyController) Post() {
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

// Put request
func (ctl *CompanyController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/company/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if company, err := md.GetCompanyByID(idInt64); err == nil {
			if err := ctl.ParseForm(&company); err == nil {
				if parentID, err := ctl.GetInt64("parent"); err == nil {
					if parent, err := md.GetCompanyByID(parentID); err == nil {
						company.Parent = parent
					}
				}
				if err := md.UpdateCompanyByID(company); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *CompanyController) Get() {
	ctl.PageName = "公司管理管理"
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
	ctl.URL = "/company/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuCompanyActive"] = "active"
}

// Edit company
func (ctl *CompanyController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	companyInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {

			if company, err := md.GetCompanyByID(idInt64); err == nil {

				companyInfo["name"] = company.Name
				parent := make(map[string]interface{})
				if company.Parent != nil {
					parent["id"] = company.Parent.ID
					parent["name"] = company.Parent.Name
				}
				companyInfo["parent"] = parent
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["Category"] = companyInfo
	ctl.Layout = "base/base.html"
	ctl.PageAction = "编辑"
	ctl.TplName = "user/company_form.html"
}

// Detail display company info
func (ctl *CompanyController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.PageAction = "详情"
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

//PostCreate 请求创建产品分类
func (ctl *CompanyController) PostCreate() {
	company := new(md.Company)
	if err := ctl.ParseForm(company); err == nil {
		if parentID, err := ctl.GetInt64("parent"); err == nil {
			if parent, err := md.GetCompanyByID(parentID); err == nil {
				company.Parent = parent
			}
		}
		if id, err := md.AddCompany(company); err == nil {
			ctl.Redirect("/company/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Create page
func (ctl *CompanyController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "user/company_form.html"
}

// Validator js validator
func (ctl *CompanyController) Validator() {
	name := ctl.GetString("name")
	recordID, _ := ctl.GetInt64("recordID")
	name = strings.TrimSpace(name)
	result := make(map[string]bool)
	obj, err := md.GetCompanyByName(name)
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
func (ctl *CompanyController) companyList(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.Company
	paginator, arrs, err := md.GetAllCompany(query, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		// result["recordsFiltered"] = paginator.TotalCount
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
			if line.Parent != nil {
				oneLine["parent"] = line.Parent.Name
			} else {
				oneLine["parent"] = "-"
			}
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

// PostList post request  json response
func (ctl *CompanyController) PostList() {
	query := make(map[string]string)
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	if result, err := ctl.companyList(query, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display company table
func (ctl *CompanyController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-company"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "user/company_list_search.html"
}
