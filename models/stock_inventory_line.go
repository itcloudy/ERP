package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// StockInventoryLine 盘点明细
type StockInventoryLine struct {
	ID                  int64               `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser          *User               `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser          *User               `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate          time.Time           `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate          time.Time           `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Inventory           *StockInventory     `orm:"rel(fk)"`                              //盘点单
	Location            *StockLocation      `orm:"rel(fk);null"`                         //库位
	Product             *ProductProduct     `orm:"rel(fk);null"`                         // 产品规格
	Template            *ProductTemplate    `orm:"rel(fk);null"`                         //产品款式
	Package             *StockQuantPackage  `orm:"rel(fk);null"`                         //包装
	ProdLot             *StockProductionLot `orm:"rel(fk);null"`                         //批次、序列号
	Company             *Company            `orm:"rel(fk);null"`                         //公司
	State               string              `orm:"default(draft)" json:"State"`          //状态draft/confirm/process/done/cancel
	ProductName         string              `orm:"default()" json:"ProductName"`         //产品名称
	MeasureFirstUomQty  float64             `orm:"default(0)"`                           //计算第一单位数量
	MeasureSecondUomQty float64             `orm:"default(0)"`                           //计算第二单位数量
	CheckedFirstUomQty  float64             `orm:"default(0)"`                           //实际第一单位数量
	CheckedSecondUomQty float64             `orm:"default(0)"`                           //实际第二单位数量
	FirstUom            *ProductUom         `orm:"rel(fk)"`                              //第一单位
	SecondUom           *ProductUom         `orm:"rel(fk);null"`                         //第二单位

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	InventoryID  int64    `orm:"-" json:"Inventory"`
	LocationID   int64    `orm:"-" json:"Location"`
	ProductID    int64    `orm:"-" json:"Product"`
	TemplateID   int64    `orm:"-" json:"Template"`
	PackageID    int64    `orm:"-" json:"Package"`
	ProdLotID    int64    `orm:"-" json:"ProdLot"`
	CompanyID    int64    `orm:"-" json:"Company"`
	FirstUomID   int64    `orm:"-" json:"FirstUom"`
	SecondUomID  int64    `orm:"-" json:"SecondUom"`
}

func init() {
	orm.RegisterModel(new(StockInventoryLine))
}

// AddStockInventoryLine insert a new StockInventoryLine into database and returns
// last inserted ID on success.
func AddStockInventoryLine(obj *StockInventoryLine, addUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.CreateUser = addUser
	obj.UpdateUser = addUser
	errBegin := o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return 0, errBegin
	}
	id, err = o.Insert(obj)
	if err == nil {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetStockInventoryLineByID retrieves StockInventoryLine by ID. Returns error if
// ID doesn't exist
func GetStockInventoryLineByID(id int64) (obj *StockInventoryLine, err error) {
	o := orm.NewOrm()
	obj = &StockInventoryLine{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Company != nil {
			o.Read(obj.Company)
		}
		return obj, nil
	}
	return nil, err
}

// GetAllStockInventoryLine retrieves all StockInventoryLine matches certain condition. Returns empty list if
// no records exist
func GetAllStockInventoryLine(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []StockInventoryLine, error) {
	var (
		objArrs   []StockInventoryLine
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(StockInventoryLine))
	qs = qs.RelatedSel()

	//cond k=v cond必须放到Filter和Exclude前面
	cond := orm.NewCondition()
	if _, ok := condMap["and"]; ok {
		andMap := condMap["and"]
		for k, v := range andMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.And(k, v)
		}
	}
	if _, ok := condMap["or"]; ok {
		orMap := condMap["or"]
		for k, v := range orMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.Or(k, v)
		}
	}

	qs = qs.SetCond(cond)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	//exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Exclude(k, v)
	}

	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[i] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[0] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return paginator, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return paginator, nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	if cnt, err := qs.Count(); err == nil {
		if cnt > 0 {
			paginator = utils.GenPaginator(limit, offset, cnt)
			if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
				paginator.CurrentPageSize = num
			}
		}
	}

	return paginator, objArrs, err
}

// UpdateStockInventoryLineByID updates StockInventoryLine by ID and returns error if
// the record to be updated doesn't exist
func UpdateStockInventoryLineByID(m *StockInventoryLine) (err error) {
	o := orm.NewOrm()
	v := StockInventoryLine{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStockInventoryLine deletes StockInventoryLine by ID and returns error if
// the record to be deleted doesn't exist
func DeleteStockInventoryLine(id int64) (err error) {
	o := orm.NewOrm()
	v := StockInventoryLine{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StockInventoryLine{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
