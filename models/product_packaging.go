package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 产品打包方式
type ProductPackaging struct {
	Base
	Name            string
	sequence        int32            //序列号
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"` //产品款式
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"` //产品规格
	FirstQty        float64          //第一单位最大数量
	// SecondQty       float64          //第二单位最大数量

}

func init() {
	orm.RegisterModel(new(ProductPackaging))
}

// AddProductPackaging insert a new ProductPackaging into database and returns
// last inserted Id on success.
func AddProductPackaging(obj *ProductPackaging) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductPackagingById retrieves ProductPackaging by Id. Returns error if
// Id doesn't exist
func GetProductPackagingById(id int64) (obj *ProductPackaging, err error) {
	o := orm.NewOrm()
	obj = &ProductPackaging{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetProductPackagingByName retrieves ProductPackaging by Name. Returns error if
// Name doesn't exist
func GetProductPackagingByName(name string) (obj *ProductPackaging, err error) {
	o := orm.NewOrm()
	obj = &ProductPackaging{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductPackaging retrieves all ProductPackaging matches certain condition. Returns empty list if
// no records exist
func GetAllProductPackaging(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductPackaging, error) {
	var (
		objArrs   []ProductPackaging
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductPackaging))
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

// UpdateProductPackaging updates ProductPackaging by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductPackagingById(m *ProductPackaging) (err error) {
	o := orm.NewOrm()
	v := ProductPackaging{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductPackaging deletes ProductPackaging by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductPackaging(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductPackaging{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductPackaging{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
