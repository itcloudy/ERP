package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductAttributeLine 产品属性明细
type ProductAttributeLine struct {
	ID              int64                    `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser      *User                    `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser      *User                    `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate      time.Time                `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate      time.Time                `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Attribute       *ProductAttribute        `orm:"rel(fk)"`                              //属性
	ProductTemplate *ProductTemplate         `orm:"rel(fk)"`                              //产品模版
	AttributeValues []*ProductAttributeValue `orm:"rel(m2m)"`                             //属性值

	FormAction        string   `orm:"-" json:"FormAction"`        //非数据库字段，用于表示记录的增加，修改
	ActionFields      []string `orm:"-" json:"ActionFields"`      //需要操作的字段,用于update时
	AttributeID       int64    `orm:"-" json:"AttributeId"`       //表单属性
	AttributeValueIds []int64  `orm:"-" json:"AttributeValueIds"` //表单属性值
}

func init() {
	orm.RegisterModel(new(ProductAttributeLine))
}

// AddProductAttributeLine insert a new ProductAttributeLine into database and returns
// last inserted ID on success.
func AddProductAttributeLine(obj *ProductAttributeLine) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductAttributeLineByID retrieves ProductAttributeLine by ID. Returns error if
// ID doesn't exist
func GetProductAttributeLineByID(id int64) (obj *ProductAttributeLine, err error) {
	o := orm.NewOrm()
	obj = &ProductAttributeLine{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductAttributeLine retrieves all ProductAttributeLine matches certain condition. Returns empty list if
// no records exist
func GetAllProductAttributeLine(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductAttributeLine, error) {
	var (
		objArrs   []ProductAttributeLine
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductAttributeLine))
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
				for i, _ := range objArrs {
					o.LoadRelated(&objArrs[i], "AttributeValues")
				}
			}
		}
	}

	return paginator, objArrs, err
}

// UpdateProductAttributeLineByID updates ProductAttributeLine by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductAttributeLineByID(m *ProductAttributeLine) (err error) {
	o := orm.NewOrm()
	v := ProductAttributeLine{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductAttributeLine deletes ProductAttributeLine by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductAttributeLine(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductAttributeLine{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductAttributeLine{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
