package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// PurchaseConfig 销售设置
type PurchaseConfig struct {
	Base
	Name    string   `orm:"unique"`
	Company *Company `orm:"rel(fk)"` //公司
}

func init() {
	orm.RegisterModel(new(PurchaseConfig))
}

// AddPurchaseConfig insert a new PurchaseConfig into database and returns
// last inserted ID on success.
func AddPurchaseConfig(obj *PurchaseConfig) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetPurchaseConfigByID retrieves PurchaseConfig by ID. Returns error if
// ID doesn't exist
func GetPurchaseConfigByID(id int64) (obj *PurchaseConfig, err error) {
	o := orm.NewOrm()
	obj = &PurchaseConfig{Base: Base{ID: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetPurchaseConfigByName retrieves PurchaseOrder by Name. Returns error if
// Name doesn't exist
func GetPurchaseConfigByName(name string) (obj *PurchaseConfig, err error) {
	o := orm.NewOrm()
	obj = &PurchaseConfig{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllPurchaseConfig retrieves all PurchaseConfig matches certain condition. Returns empty list if
// no records exist
func GetAllPurchaseConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []PurchaseConfig, error) {
	var (
		objArrs   []PurchaseConfig
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(PurchaseConfig))
	qs = qs.RelatedSel()

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
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
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
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
		paginator = utils.GenPaginator(limit, offset, cnt)
	}
	if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
		paginator.CurrentPageSize = num
	}
	return paginator, objArrs, err
}

// UpdatePurchaseConfigByID updates PurchaseConfig by ID and returns error if
// the record to be updated doesn't exist
func UpdatePurchaseConfigByID(m *PurchaseConfig) (err error) {
	o := orm.NewOrm()
	v := PurchaseConfig{Base: Base{ID: m.ID}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePurchaseConfig deletes PurchaseConfig by ID and returns error if
// the record to be deleted doesn't exist
func DeletePurchaseConfig(id int64) (err error) {
	o := orm.NewOrm()
	v := PurchaseConfig{Base: Base{ID: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PurchaseConfig{Base: Base{ID: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
