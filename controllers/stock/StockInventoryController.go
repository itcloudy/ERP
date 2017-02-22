package stock

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// StockInventoryController stock inventory
type StockInventoryController struct {
	base.BaseController
}

// Post request
func (ctl *StockInventoryController) Post() {
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
func (ctl *StockInventoryController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/stock/inventory/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if inventory, err := md.GetStockInventoryByID(idInt64); err == nil {
			if err := ctl.ParseForm(&inventory); err == nil {

				if err := md.UpdateStockInventoryByID(inventory); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *StockInventoryController) Get() {
	ctl.PageName = "盘点管理"
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
	ctl.URL = "/stock/inventory/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuStockInventoryActive"] = "active"
}

// Edit edit stock inventory
func (ctl *StockInventoryController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	inventoryInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if inventory, err := md.GetStockInventoryByID(idInt64); err == nil {
				ctl.PageAction = inventory.Name
				inventoryInfo["name"] = inventory.Name

			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["inventory"] = inventoryInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "stock/stock_inventory_form.html"
}

// Create display stock inventory create page
func (ctl *StockInventoryController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "stock/stock_inventory_form.html"
}

// Detail display stock inventory info
func (ctl *StockInventoryController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create stock inventory
func (ctl *StockInventoryController) PostCreate() {
	inventory := new(md.StockInventory)
	if err := ctl.ParseForm(inventory); err == nil {

		if id, err := md.AddStockInventory(inventory, &ctl.User); err == nil {
			ctl.Redirect("/stock/inventory/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *StockInventoryController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetStockInventoryByName(name)
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

//StockInventoryList 获得符合要求的数据
func (ctl *StockInventoryController) StockInventoryList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, inventory []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.StockInventory
	paginator, arrs, err := md.GetAllStockInventory(query, exclude, condMap, fields, sortby, inventory, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["name"] = line.Name
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

// PostList post request json response
func (ctl *StockInventoryController) PostList() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})

	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	inventory := make([]string, 0, 1)
	offset, _ := ctl.GetInt64("offset")
	limit, _ := ctl.GetInt64("limit")
	inventoryStr := ctl.GetString("inventory")
	sortStr := ctl.GetString("sort")
	if inventoryStr != "" && sortStr != "" {
		sortby = append(sortby, sortStr)
		inventory = append(inventory, inventoryStr)
	} else {
		sortby = append(sortby, "Id")
		inventory = append(inventory, "desc")
	}
	if result, err := ctl.StockInventoryList(query, exclude, cond, fields, sortby, inventory, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display stock inventory with list
func (ctl *StockInventoryController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-stock-inventory"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "stock/stock_inventory_list_search.html"
}
