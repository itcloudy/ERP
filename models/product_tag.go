package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// 产品标签
type ProductTag struct {
	Base
	Name     string            `orm:"size(20);unique"`        //产品标签名称
	Type     string            `orm:"size(20);default(\"\")"` //标签类型，前端显示采用select
	Products []*ProductProduct `orm:"rel(m2m)"`               //产品规格
}

func init() {
	orm.RegisterModel(new(ProductTag))
}

// AddProductTag insert a new ProductTag into database and returns
// last inserted Id on success.
func AddProductTag(obj *ProductTag) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetProductTagById retrieves ProductTag by Id. Returns error if
// Id doesn't exist
func GetProductTagById(id int64) (obj *ProductTag, err error) {
	o := orm.NewOrm()
	obj = &ProductTag{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductTag retrieves all ProductTag matches certain condition. Returns empty list if
// no records exist
func GetAllProductTag(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductTag, error) {
	var (
		objArrs   []ProductTag
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductTag))
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

// UpdateProductTag updates ProductTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductTagById(m *ProductTag) (err error) {
	o := orm.NewOrm()
	v := ProductTag{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetProductTagByName retrieves ProductTag by Name. Returns error if
// Name doesn't exist
func GetProductTagByName(name string) (obj *ProductTag, err error) {
	o := orm.NewOrm()
	obj = &ProductTag{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeleteProductTag deletes ProductTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductTag(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductTag{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductTag{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
