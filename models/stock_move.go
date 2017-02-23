package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// StockMove  	移动明细
type StockMove struct {
	ID                 int64             `orm:"column(id);pk;auto" json:"id"`                //主键
	CreateUser         *User             `orm:"rel(fk);null" json:"-"`                       //创建者
	UpdateUser         *User             `orm:"rel(fk);null" json:"-"`                       //最后更新者
	CreateDate         time.Time         `orm:"auto_now_add;type(datetime)" json:"-"`        //创建时间
	UpdateDate         time.Time         `orm:"auto_now;type(datetime)" json:"-"`            //最后更新时间
	Sequence           int64             `orm:"default(0)" json:"Sequence"`                  //序列号
	Name               string            `json:"Name"`                                       //明细产品名称
	Priority           int64             `orm:"default(1)" json:"Priority"`                  //优先级
	Date               time.Time         `orm:" type(datetime)"`                             //预定日期
	DateExpected       time.Time         `orm:" type(datetime)"`                             //预定日期
	Product            *ProductProduct   `orm:"rel(fk)"`                                     //产品规格
	FirstUomQty        float64           `orm:"default(0)"`                                  //第一单位数量
	SecondUomQty       float64           `orm:"default(0)"`                                  //第二单位数量
	FirstUom           *ProductUom       `orm:"rel(fk)"`                                     //第一单位
	SecondUom          *ProductUom       `orm:"rel(fk);null"`                                //第二单位
	ProductTemplate    *ProductTemplate  `orm:"rel(fk);null"`                                //产品款式
	ProductPackaging   *ProductPackaging `orm:"rel(fk);null"`                                //包装类型、包装数量等属性
	LocationSrc        *StockLocation    `orm:"rel(fk);null"`                                //源库位
	LocationDest       *StockLocation    `orm:"rel(fk);null"`                                //目标库位
	Partner            *Partner          `orm:"rel(fk);null"`                                //合作伙伴
	Picking            *StockPicking     `orm:"rel(fk)"`                                     //调拨单
	State              string            `orm:"default(draft)" json:"State"`                 //状态
	Note               string            `json:"Note"`                                       //备注
	PartiallyAvailable bool              `orm:"default(false)"`                              //部分出货
	PriceUnit          float64           `orm:"default(0)" json:"PriceUnit"`                 //单价
	Company            *Company          `orm:"rel(fk)"`                                     //所属公司
	BackOrder          *StockPicking     `orm:"rel(fk);null"`                                //退货单
	Origin             string            `json:"Origin"`                                     //源数据
	ProcureMethod      string            `orm:"default(make_to_order)" json:"ProcureMethod"` //单据来源:make_to_stock/make_to_order
	Scrapped           bool              `orm:"default(false)" json:"Scrapped"`              //报废，跟LocationDest的类型一致
	Quants             []*StockQuant     `orm:"rel(m2m);rel_table(stock_quant_move_rel)"`    //迁移数量
	ReservedQuant      []*StockQuant     `orm:"reverse(many)"`                               //保留数量
	Inventory          *StockInventory   `orm:"rel(fk);null"`                                //盘点单
	WareHouse          *StockWarehouse   `orm:"rel(fk);null"`                                //仓库
	FormAction         string            `orm:"-" json:"FormAction"`                         //非数据库字段，用于表示记录的增加，修改
	ActionFields       []string          `orm:"-" json:"ActionFields"`                       //需要操作的字段,用于update时
	PickingID          int64             `orm:"-" json:"Picking"`                            //
	FirstUomID         int64             `orm:"-" json:"FirstUom"`                           //
	SecondUomID        int64             `orm:"-" json:"SecondUom"`                          //

}

func init() {
	orm.RegisterModel(new(StockMove))
}
func FirstRemainingQty(obj *StockMove) int {
	return 0
}
func SecondRemainingQty(obj *StockMove) int {
	return 0
}

// AddStockMove insert a new StockMove into database and returns
// last inserted ID on success.
func AddStockMove(obj *StockMove, addUser *User) (id int64, err error) {
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
	if obj.PickingID > 0 {
		obj.Picking, _ = GetStockPickingByID(obj.PickingID)
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

// GetStockMoveByID retrieves StockMove by ID. Returns error if
// ID doesn't exist
func GetStockMoveByID(id int64) (obj *StockMove, err error) {
	o := orm.NewOrm()
	obj = &StockMove{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllStockMove retrieves all StockMove matches certain condition. Returns empty list if
// no records exist
func GetAllStockMove(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []StockMove, error) {
	var (
		objArrs   []StockMove
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(StockMove))
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

// UpdateStockMoveByID updates StockMove by ID and returns error if
// the record to be updated doesn't exist
func UpdateStockMoveByID(m *StockMove) (err error) {
	o := orm.NewOrm()
	v := StockMove{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetStockMoveByName retrieves StockMove by Name. Returns error if
// Name doesn't exist
func GetStockMoveByName(name string) (obj *StockMove, err error) {
	o := orm.NewOrm()
	obj = &StockMove{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeleteStockMove deletes StockMove by ID and returns error if
// the record to be deleted doesn't exist
func DeleteStockMove(id int64) (err error) {
	o := orm.NewOrm()
	v := StockMove{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StockMove{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
