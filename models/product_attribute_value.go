package models

import (
	"errors"
	"golangERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductAttributeValue 产品属性
type ProductAttributeValue struct {
	ID           int64             `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64             `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64             `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time         `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time         `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string            `orm:"size(50)"`                    //属性值
	Attribute    *ProductAttribute `orm:"rel(fk)"`                     //属性
	Products     []*ProductProduct `orm:"rel(m2m)"`                    //产品规格
}

func init() {
	orm.RegisterModel(new(ProductAttributeValue))
}

// AddProductAttributeValue insert a new ProductAttributeValue into database and returns last inserted Id on success.
func AddProductAttributeValue(m *ProductAttributeValue, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddProductAttributeValue insert  list of  ProductAttributeValue into database and returns  number of  success.
func BatchAddProductAttributeValue(cities []*ProductAttributeValue, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&ProductAttributeValue{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, city := range cities {
			if _, err = i.Insert(city); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateProductAttributeValue update ProductAttributeValue into database and returns id on success
func UpdateProductAttributeValue(m *ProductAttributeValue, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetProductAttributeValueByID retrieves ProductAttributeValue by ID. Returns error if ID doesn't exist
func GetProductAttributeValueByID(id int64, ormObj orm.Ormer) (obj *ProductAttributeValue, err error) {
	obj = &ProductAttributeValue{ID: id}
	err = ormObj.Read(obj)
	ormObj.Read(obj.Attribute)
	return obj, err
}

// DeleteProductAttributeValueByID delete  ProductAttributeValue by ID
func DeleteProductAttributeValueByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &ProductAttributeValue{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// GetAllProductAttributeValue retrieves all ProductAttributeValue matches certain condition. Returns empty list if no records exist
func GetAllProductAttributeValue(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductAttributeValue, error) {
	var (
		objArrs   []ProductAttributeValue
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(ProductAttributeValue))
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
