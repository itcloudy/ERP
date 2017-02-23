package stock

import (
	"bytes"
	"encoding/json"
	"goERP/controllers/base"
	md "goERP/models"
	"strconv"
	"strings"
)

type StockPickingTypeController struct {
	base.BaseController
}

func (ctl *StockPickingTypeController) Post() {
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
func (ctl *StockPickingTypeController) Get() {

	ctl.PageName = "库位类型管理"
	ctl.URL = "/stock/picking/type/"
	ctl.Data["URL"] = ctl.URL
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
	case "flowChart":
		ctl.GetFlowChart()
	default:
		ctl.GetKanban()
	}
	// 标题合成
	b := bytes.Buffer{}
	b.WriteString(ctl.PageName)
	b.WriteString("\\")
	b.WriteString(ctl.PageAction)
	ctl.Data["PageName"] = b.String()
	ctl.URL = "/stock/picking/type/"
	ctl.Data["URL"] = ctl.URL
	ctl.Data["MenuStockPickingTypeActive"] = "active"

}
func (ctl *StockPickingTypeController) GetFlowChart() {

	ctl.PageAction = "流程图"
	ctl.Data["flowChartId"] = "flow-chart-stock-picking-type"
	ctl.TplName = "base/base_flow_view.html"

}
func (ctl *StockPickingTypeController) GetKanban() {

	ctl.PageAction = "看板"
	ctl.Data["KanbanId"] = "kanban-stock-picking-type"
	ctl.TplName = "base/base_kanban_view.html"

}
func (ctl *StockPickingTypeController) PostCreate() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	picking := new(md.StockPickingType)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), picking); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(picking)).Type().Name()
		if id, err = md.AddStockPickingType(picking, &ctl.User); err == nil {
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
func (ctl *StockPickingTypeController) Put() {
	result := make(map[string]interface{})
	postData := ctl.GetString("postData")
	spt := new(md.StockPickingType)
	var (
		err error
		id  int64
	)
	if err = json.Unmarshal([]byte(postData), spt); err == nil {
		// 获得struct表名
		// structName := reflect.Indirect(reflect.ValueOf(template)).Type().Name()
		if id, err = md.AddStockPickingType(spt, &ctl.User); err == nil {
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
func (ctl *StockPickingTypeController) Create() {
	ctl.Data["Action"] = "create"
	ctl.Data["Readonly"] = false
	ctl.PageAction = "创建"
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-create"
	ctl.TplName = "stock/stock_picking_type_form.html"
}
func (ctl *StockPickingTypeController) Edit() {
	id := ctl.Ctx.Input.Param(":id")
	if id != "" {
		if idInt64, e := strconv.ParseInt(id, 10, 64); e == nil {
			if picking, err := md.GetStockPickingTypeByID(idInt64); err == nil {
				ctl.PageAction = picking.Name
				ctl.Data["StockPickingType"] = picking
			}
		}
	}
	ctl.Data["Action"] = "edit"
	ctl.Data["RecordID"] = id
	ctl.Layout = "base/base.html"
	ctl.Data["FormField"] = "form-edit"
	ctl.TplName = "stock/stock_picking_type_form.html"
}
func (ctl *StockPickingTypeController) Detail() {
	ctl.Edit()
	ctl.Data["Readonly"] = true
	ctl.Data["Action"] = "detail"
}
func (ctl *StockPickingTypeController) Validator() {
	name := strings.TrimSpace(ctl.GetString("name"))
	recordID, _ := ctl.GetInt64("recordID")
	result := make(map[string]bool)
	// 默认验证成功
	result["valid"] = true
	if name != "" {
		obj, err := md.GetStockPickingTypeByName(name)
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
	}

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

// 获得符合要求的城市数据
func (ctl *StockPickingTypeController) pickingProductList(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (map[string]interface{}, error) {

	var arrs []md.StockPickingType
	paginator, arrs, err := md.GetAllStockPickingType(query, exclude, condMap, fields, sortby, order, offset, limit)
	result := make(map[string]interface{})
	if err == nil {

		//使用多线程来处理数据，待修改
		tableLines := make([]interface{}, 0, 4)
		for _, line := range arrs {
			oneLine := make(map[string]interface{})
			oneLine["Name"] = line.Name
			oneLine["Code"] = line.Code
			oneLine["ID"] = line.ID
			oneLine["id"] = line.ID
			oneLine["IsStart"] = line.IsStart
			oneLine["IsEnd"] = line.IsEnd

			if line.WareHouse != nil {
				wareHouse := make(map[string]interface{})
				wareHouse["id"] = line.WareHouse.ID
				wareHouse["name"] = line.WareHouse.Name
				oneLine["WareHouse"] = wareHouse
			}
			if line.NextStep != nil {
				nextStep := make(map[string]interface{})
				nextStep["id"] = line.NextStep.ID
				nextStep["name"] = line.NextStep.Name
				oneLine["NextStep"] = nextStep
			}
			if line.PrevStep != nil {
				prevStep := make(map[string]interface{})
				prevStep["id"] = line.PrevStep.ID
				prevStep["name"] = line.PrevStep.Name
				oneLine["PrevStep"] = prevStep
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
func (ctl *StockPickingTypeController) PostList() {
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
		order = append(order, "asc")
	}
	if result, err := ctl.pickingProductList(query, exclude, cond, fields, sortby, order, offset, limit); err == nil {
		ctl.Data["json"] = result
	}
	ctl.ServeJSON()

}

func (ctl *StockPickingTypeController) GetList() {
	ctl.PageAction = "列表"
	ctl.Data["tableId"] = "table-stock-picking-type"
	ctl.Layout = "base/base_list_view.html"
	ctl.TplName = "stock/stock_picking_type_list_search.html"
}
