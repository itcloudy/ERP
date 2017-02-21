package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// PurchaseOrderState 订单状态
type PurchaseOrderState struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string    `orm:"default()" json:"name"`            //状态名称
	Active     bool      `orm:"default(true)" json:"Active"`          //是否有效

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时

}

func init() {
	orm.RegisterModel(new(PurchaseOrderState))
}

// AddPurchaseOrderState insert a new PurchaseOrderState into database and returns
// last inserted ID on success.
func AddPurchaseOrderState(obj *PurchaseOrderState) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetPurchaseOrderStateByID retrieves PurchaseOrderState by ID. Returns error if
// ID doesn't exist
func GetPurchaseOrderStateByID(id int64) (obj *PurchaseOrderState, err error) {
	o := orm.NewOrm()
	obj = &PurchaseOrderState{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllPurchaseOrderState retrieves all PurchaseOrderState matches certain condition. Returns empty list if
// no records exist
func GetAllPurchaseOrderState(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []PurchaseOrderState, error) {
	var (
		objArrs   []PurchaseOrderState
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(PurchaseOrderState))
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

// UpdatePurchaseOrderStateByID updates PurchaseOrderState by ID and returns error if
// the record to be updated doesn't exist
func UpdatePurchaseOrderStateByID(m *PurchaseOrderState) (err error) {
	o := orm.NewOrm()
	v := PurchaseOrderState{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetPurchaseOrderStateByName retrieves PurchaseOrderState by Name. Returns error if
// Name doesn't exist
func GetPurchaseOrderStateByName(name string) (obj *PurchaseOrderState, err error) {
	o := orm.NewOrm()
	obj = &PurchaseOrderState{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeletePurchaseOrderState deletes PurchaseOrderState by ID and returns error if
// the record to be deleted doesn't exist
func DeletePurchaseOrderState(id int64) (err error) {
	o := orm.NewOrm()
	v := PurchaseOrderState{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PurchaseOrderState{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
