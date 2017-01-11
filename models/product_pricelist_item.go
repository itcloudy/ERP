package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 产品价格
type ProductPricelistItem struct {
	Base
}

func init() {
	orm.RegisterModel(new(ProductPricelistItem))
}

// AddProductPricelistItem insert a new ProductPricelistItem into database and returns
// last inserted Id on success.
func AddProductPricelistItem(obj *ProductPricelistItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductPricelistItemById retrieves ProductPricelistItem by Id. Returns error if
// Id doesn't exist
func GetProductPricelistItemById(id int64) (obj *ProductPricelistItem, err error) {
	o := orm.NewOrm()
	obj = &ProductPricelistItem{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductPricelistItem retrieves all ProductPricelistItem matches certain condition. Returns empty list if
// no records exist
func GetAllProductPricelistItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductPricelistItem, error) {
	var (
		objArrs   []ProductPricelistItem
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductPricelistItem))
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

// UpdateProductPricelistItem updates ProductPricelistItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductPricelistItemById(m *ProductPricelistItem) (err error) {
	o := orm.NewOrm()
	v := ProductPricelistItem{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductPricelistItem deletes ProductPricelistItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductPricelistItem(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductPricelistItem{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductPricelistItem{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
