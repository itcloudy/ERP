package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// StockPickingType 分拣类型决定分拣视图
type StockPickingType struct {
	ID         int64             `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User             `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User             `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time         `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time         `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string            `orm:"unique" json:"Name"`                   //分拣类型名称
	Active     bool              `orm:"default(true)" json:"Active"`          //是否有效
	Code       string            `json:"Code"`                                //移库类型 incoming/outgoing/internal
	WareHouse  *StockWarehouse   `orm:"rel(fk)"`                              //仓库
	NextStep   *StockPickingType `orm:"null;rel(one)"`                        //下一步
	PrevStep   *StockPickingType `orm:"null;rel(one)"`                        //上一步
	IsStart    bool              `orm:"default(false)"`                       //流程开始
	IsEnd      bool              `orm:"default(false)"`                       //流程结束

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	WareHouseID  int64    `orm:"-" json:"WareHouse"`
	NextStepID   int64    `orm:"-" json:"NextStep"`
	PrevStepID   int64    `orm:"-" json:"PrevStep"`
}

func init() {
	orm.RegisterModel(new(StockPickingType))
}

// AddStockPickingType insert a new StockPickingType into database and returns
// last inserted ID on success.
func AddStockPickingType(obj *StockPickingType, addUser *User) (id int64, err error) {
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

	if obj.WareHouseID > 0 {
		obj.WareHouse, _ = GetStockWarehouseByID(obj.WareHouseID)
	}
	if obj.NextStepID > 0 {
		obj.NextStep, _ = GetStockPickingTypeByID(obj.NextStepID)
	}
	if obj.PrevStepID > 0 {
		obj.PrevStep, _ = GetStockPickingTypeByID(obj.PrevStepID)
	}
	id, err = o.Insert(obj)
	if err != nil {
		return 0, err
	} else {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetStockPickingTypeByID retrieves StockPickingType by ID. Returns error if
// ID doesn't exist
func GetStockPickingTypeByID(id int64) (obj *StockPickingType, err error) {
	o := orm.NewOrm()
	obj = &StockPickingType{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.WareHouse != nil {
			o.Read(obj.WareHouse)
		}
		if obj.NextStep != nil {
			o.Read(obj.NextStep)
		}
		if obj.PrevStep != nil {
			o.Read(obj.PrevStep)
		}
		return obj, nil
	}
	return nil, err
}

// GetStockPickingTypeByName retrieves StockPickingType by Name. Returns error if
// Name doesn't exist
func GetStockPickingTypeByName(name string) (*StockPickingType, error) {
	o := orm.NewOrm()
	var obj StockPickingType
	cond := orm.NewCondition()
	cond = cond.And("Name", name)
	qs := o.QueryTable(&obj)
	qs = qs.SetCond(cond)
	err := qs.One(&obj)
	return &obj, err
}

// GetAllStockPickingType retrieves all StockPickingType matches certain condition. Returns empty list if
// no records exist
func GetAllStockPickingType(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []StockPickingType, error) {
	var (
		objArrs   []StockPickingType
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(StockPickingType))
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

// UpdateStockPickingType updates StockPickingType by ID and returns error if
// the record to be updated doesn't exist
func UpdateStockPickingType(obj *StockPickingType, updateUser *User) error {
	o := orm.NewOrm()
	v := User{ID: obj.ID}
	errBegin := o.Begin()
	var err error
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return errBegin
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Update(obj, append(obj.ActionFields, "UpdateUser", "UpdateDate")...); err != nil {
			utils.LogOut("error", "update user fields failed:"+err.Error())
		}
	}
	if err != nil {
		return err
	}
	errCommit := o.Commit()
	if errCommit != nil {
		return errCommit
	}
	return nil
}

// DeleteStockPickingType deletes StockPickingType by ID and returns error if
// the record to be deleted doesn't exist
func DeleteStockPickingType(id int64) (err error) {
	o := orm.NewOrm()
	v := StockPickingType{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StockPickingType{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
