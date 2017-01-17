package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

//ProductAttributeValue 产品属性值
type ProductAttributeValue struct {
	Base
	Name       string            `orm:"unique" form:"name"` //产品属性名称
	Attribute  *ProductAttribute `orm:"rel(fk)"`            //属性
	Products   []*ProductProduct `orm:"rel(m2m)"`           //产品规格
	PriceExtra float64           `orm:"default(0)"`         //额外价格
	// Prices     *ProductAttributePrice `orm:"reverse(many)"`
	Sequence int32 //序列
}

func init() {
	orm.RegisterModel(new(ProductAttributeValue))
}

// AddProductAttributeValue insert a new ProductAttributeValue into database and returns
// last inserted ID on success.
func AddProductAttributeValue(obj *ProductAttributeValue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductAttributeValueByID retrieves ProductAttributeValue by ID. Returns error if
// ID doesn't exist
func GetProductAttributeValueByID(id int64) (obj *ProductAttributeValue, err error) {
	o := orm.NewOrm()
	obj = &ProductAttributeValue{Base: Base{ID: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetProductAttributeValueByName retrieves ProductAttributeValue by Name. Returns error if
// Name doesn't exist
func GetProductAttributeValueByName(name string) (obj *ProductAttributeValue, err error) {
	o := orm.NewOrm()
	obj = &ProductAttributeValue{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductAttributeValue retrieves all ProductAttributeValue matches certain condition. Returns empty list if
// no records exist
func GetAllProductAttributeValue(query map[string]interface{}, exclude map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductAttributeValue, error) {
	var (
		objArrs   []ProductAttributeValue
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductAttributeValue))
	qs = qs.RelatedSel()
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

// UpdateProductAttributeValueByID updates ProductAttributeValue by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductAttributeValueByID(m *ProductAttributeValue) (err error) {
	o := orm.NewOrm()
	v := ProductAttributeValue{Base: Base{ID: m.ID}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductAttributeValue deletes ProductAttributeValue by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductAttributeValue(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductAttributeValue{Base: Base{ID: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductAttributeValue{Base: Base{ID: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
