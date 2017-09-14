package models

import (
	"errors"
	"golangERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductCategory 产品分类
type ProductCategory struct {
	ID           int64              `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64              `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64              `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time          `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time          `orm:"auto_now;type(datetime)"`     //最后更新时间
	ParentLeft   int64              `orm:"unique"`                      //分类左
	ParentRight  int64              `orm:"unique"`                      //分类右
	Name         string             `orm:"size(50)" json:"name"`        //分类名称
	Parent       *ProductCategory   `orm:"rel(fk);null"`                //上级分类
	Childs       []*ProductCategory `orm:"reverse(many)" `              //下级分类
}

func init() {
	orm.RegisterModel(new(ProductCategory))
}

// AddProductCategory insert a new ProductCategory into database and returns last inserted Id on success.
func AddProductCategory(m *ProductCategory, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddProductCategory insert  list of  ProductCategory into database and returns  number of  success.
func BatchAddProductCategory(menus []*ProductCategory, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&ProductCategory{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, menu := range menus {
			if _, err = i.Insert(menu); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateProductCategory update ProductCategory into database and returns id on success
func UpdateProductCategory(m *ProductCategory, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetProductCategoryByID retrieves ProductCategory by ID. Returns error if ID doesn't exist
func GetProductCategoryByID(id int64, ormObj orm.Ormer) (obj *ProductCategory, err error) {
	obj = &ProductCategory{ID: id}
	err = ormObj.Read(obj)
	ormObj.LoadRelated(obj, "Childs")
	ormObj.Read(obj.Parent)
	return obj, err
}

// GetAllProductCategory retrieves all ProductCategory matches certain condition. Returns empty list if no records exist
func GetAllProductCategory(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductCategory, error) {
	var (
		objArrs   []ProductCategory
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(ProductCategory))
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
