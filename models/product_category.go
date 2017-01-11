package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// 产品分类
type ProductCategory struct {
	Base
	Name           string             `orm:"unique" form:"name" json:"name"` //产品属性名称
	Parent         *ProductCategory   `orm:"rel(fk);null"`                   //上级分类
	Childs         []*ProductCategory `orm:"reverse(many)"`                  //下级分类
	Sequence       int64              //序列
	ParentFullPath string             //上级全路径
}

func init() {
	orm.RegisterModel(new(ProductCategory))
}

// AddProductCategory insert a new ProductCategory into database and returns
// last inserted Id on success.
func AddProductCategory(obj *ProductCategory) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetProductCategoryById retrieves ProductCategory by Id. Returns error if
// Id doesn't exist
func GetProductCategoryById(id int64) (obj *ProductCategory, err error) {
	o := orm.NewOrm()
	obj = &ProductCategory{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductCategory retrieves all ProductCategory matches certain condition. Returns empty list if
// no records exist
func GetAllProductCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductCategory, error) {
	var (
		objArrs   []ProductCategory
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductCategory))
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

// UpdateProductCategory updates ProductCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductCategoryById(m *ProductCategory) (err error) {
	o := orm.NewOrm()
	v := ProductCategory{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetProductCategoryByName retrieves ProductCategory by Name. Returns error if
// Name doesn't exist
func GetProductCategoryByName(name string) (obj *ProductCategory, err error) {
	o := orm.NewOrm()
	obj = &ProductCategory{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeleteProductCategory deletes ProductCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductCategory(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductCategory{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductCategory{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
