package stock

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

// StockMoveController
type StockMoveController struct {
	base.BaseController
}

// Post request
func (ctl *StockMoveController) Post() {
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
func (ctl *StockMoveController) Put() {
	id := ctl.Ctx.Input.Param(":id")
	ctl.URL = "/stock/move/"
	if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
		if stockMove, err := md.GetStockMoveByID(idInt64); err == nil {
			if err := ctl.ParseForm(&stockMove); err == nil {

				if err := md.UpdateStockMoveByID(stockMove); err == nil {
					ctl.Redirect(ctl.URL+id+"?action=detail", 302)
				}
			}
		}
	}
	ctl.Redirect(ctl.URL+id+"?action=edit", 302)

}

// Get request
func (ctl *StockMoveController) Get() {
	ctl.PageName = "出库订单明细管理"
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
	ctl.URL = "/stock/move/"
	ctl.Data["URL"] = ctl.URL

	ctl.Data["MenuStockMoveActive"] = "active"
}

// Edit edit stock orde line
func (ctl *StockMoveController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	stockMoveInfo := make(map[string]interface{})
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if stockMove, err := md.GetStockMoveByID(idInt64); err == nil {
				ctl.PageAction = stockMove.Name
				stockMoveInfo["name"] = stockMove.Name
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Data["stockMove"] = stockMoveInfo
	ctl.Layout = "base/base.html"
	ctl.TplName = "stock/stock_move_form.html"
}

// Create display create page
func (ctl *StockMoveController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.TplName = "stock/stock_move_form.html"
}

// Detail display stock order line info
func (ctl *StockMoveController) Detail() {
	//获取信息一样，直接调用Edit
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}

// PostCreate post request create stock order line
func (ctl *StockMoveController) PostCreate() {
	stockMove := new(md.StockMove)
	if err := ctl.ParseForm(stockMove); err == nil {

		if id, err := md.AddStockMove(stockMove, &ctl.User); err == nil {
			ctl.Redirect("/stock/move/"+strconv.FormatInt(id, 10)+"?action=detail", 302)
		} else {
			ctl.Get()
		}
	} else {
		ctl.Get()
	}
}

// Validator js valid
func (ctl *StockMoveController) Validator() {
	name := ctl.GetString("name")
	name = strings.TrimSpace(name)
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	obj, err := md.GetStockMoveByName(name)
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

//StockMoveList 获得符合要求的数据
func (ctl *StockMoveController) StockMoveList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.StockMove
	paginator, arrs, err := md.GetAllStockMove(query, exclude, condMap, fields, sortby, order, offset, limit)
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

// PostList post request list stock order line
func (ctl *StockMoveController) PostList() {
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
	if result, err := ctl.StockMoveList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

// GetList display stock with list
func (ctl *StockMoveController) GetList() {
	viewType := ctl.Input().Get("view")
	if viewType == "" || viewType == "table" {
		ctl.Data["ViewType"] = "table"
	}
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-stock-move"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "stock/stock_move_list_search.html"
}
