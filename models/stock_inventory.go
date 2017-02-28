package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// StockInventory 盘点
type StockInventory struct {
	ID         int64                 `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User                 `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User                 `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time             `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time             `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string                `orm:"unique"`                               //盘点名称
	Date       time.Time             `orm:"auto_now;type(datetime)" json:"Date"`  //盘点日期
	Lines      []*StockInventoryLine `orm:"reverse(many)"`                        //盘点明细
	Moves      []*StockMove          `orm:"reverse(many)"`                        //移动明细
	State      string                `orm:"default(draft)" json:"State"`          //状态draft/confirm/process/done/cancel
	Company    *Company              `orm:"rel(fk)"`                              //公司
	Location   *StockLocation        `orm:"rel(fk);null"`                         //盘点库位，包括下级所有的库位
	Product    *ProductProduct       `orm:"rel(fk);null"`                         //产品规格，可以根据规格盘点
	Template   *ProductTemplate      `orm:"rel(fk);null"`                         //产品款式，可以根据款式盘点
	Package    *StockQuantPackage    `orm:"rel(fk);null"`                         //可以根据包来盘
	Lot        *StockProductionLot   `orm:"rel(fk);null"`                         //指定批量/序列号
	Filter     string                `orm:"default(all)"`                         //盘点对象all/product/template/pack/lot/partial

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	CompanyID    int64    `orm:"-" json:"Company"`
	LocationID   int64    `orm:"-" json:"Location"`
	ProductID    int64    `orm:"-" json:"Product"`
	TemplateID   int64    `orm:"-" json:"Template"`
	PackageID    int64    `orm:"-" json:"Package"`
	LotID        int64    `orm:"-" json:"Lot"`
}

func init() {
	orm.RegisterModel(new(StockInventory))
}

// AddStockInventory insert a new StockInventory into database and returns
// last inserted ID on success.
func AddStockInventory(obj *StockInventory, addUser *User) (id int64, err error) {
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

// GetStockInventoryByID retrieves StockInventory by ID. Returns error if
// ID doesn't exist
func GetStockInventoryByID(id int64) (obj *StockInventory, err error) {
	o := orm.NewOrm()
	obj = &StockInventory{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Company != nil {
			o.Read(obj.Company)
		}
		return obj, nil
	}
	return nil, err
}

// GetAllStockInventory retrieves all StockInventory matches certain condition. Returns empty list if
// no records exist
func GetAllStockInventory(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []StockInventory, error) {
	var (
		objArrs   []StockInventory
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(StockInventory))
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

// UpdateStockInventoryByID updates StockInventory by ID and returns error if
// the record to be updated doesn't exist
func UpdateStockInventoryByID(m *StockInventory) (err error) {
	o := orm.NewOrm()
	v := StockInventory{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetStockInventoryByName retrieves StockInventory by Name. Returns error if
// Name doesn't exist
func GetStockInventoryByName(name string) (obj *StockInventory, err error) {
	o := orm.NewOrm()
	obj = &StockInventory{Name: name}
	if err = o.Read(obj); err == nil {
		if obj.Company != nil {
			o.Read(obj.Company)
		}
		return obj, nil
	}
	return nil, err
}

// DeleteStockInventory deletes StockInventory by ID and returns error if
// the record to be deleted doesn't exist
func DeleteStockInventory(id int64) (err error) {
	o := orm.NewOrm()
	v := StockInventory{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StockInventory{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
