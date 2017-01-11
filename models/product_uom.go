package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// 产品单位
type ProductUom struct {
	Base
	Name      string           `orm:"unique" form:"name"`          //计量单位名称
	Active    bool             `orm:"default(true)" form:"active"` //有效
	Category  *ProductUomCateg `orm:"rel(fk)"`                     //计量单位类别
	Factor    float64          `form:"factor"`                     //比率
	FactorInv float64          `form:"factorInv"`                  //更大比率
	Rounding  float64          `form:"rounding"`                   //舍入精度
	Type      int64            `form:"type"`                       //类型
	Symbol    string           `form:"symbol"`                     //符号，后置
}

func init() {
	orm.RegisterModel(new(ProductUom))
}

// AddProductUom insert a new ProductUom into database and returns
// last inserted Id on success.
func AddProductUom(obj *ProductUom) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetProductUomById retrieves ProductUom by Id. Returns error if
// Id doesn't exist
func GetProductUomById(id int64) (obj *ProductUom, err error) {
	o := orm.NewOrm()
	obj = &ProductUom{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetProductUomByName retrieves ProductUom by Name. Returns error if
// Name doesn't exist
func GetProductUomByName(name string) (obj *ProductUom, err error) {
	o := orm.NewOrm()
	obj = &ProductUom{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductUom retrieves all ProductUom matches certain condition. Returns empty list if
// no records exist
func GetAllProductUom(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductUom, error) {
	var (
		objArrs   []ProductUom
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductUom))
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

// UpdateProductUom updates ProductUom by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductUomById(m *ProductUom) (err error) {
	o := orm.NewOrm()
	v := ProductUom{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductUom deletes ProductUom by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductUom(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductUom{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductUom{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
