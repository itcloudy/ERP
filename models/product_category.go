package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// ProductCategory 产品分类
type ProductCategory struct {
	ID             int64              `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser     *User              `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser     *User              `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate     time.Time          `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate     time.Time          `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name           string             `orm:"unique"  json:"Name"`                  //产品属性名称
	Parent         *ProductCategory   `orm:"rel(fk);null"`                         //上级分类
	Childs         []*ProductCategory `orm:"reverse(many)"`                        //下级分类
	Sequence       int64              //序列
	ParentFullPath string             //上级全路径
	// form表单使用字段
	ParentID   int64  `orm:"-" json:"Parent"`     //上级类别
	FormAction string `orm:"-" json:"FormAction"` //非数据库字段，用于表示记录的增加，修改

}

func init() {
	orm.RegisterModel(new(ProductCategory))
}

// AddProductCategory insert a new ProductCategory into database and returns
// last inserted ID on success.
func AddProductCategory(obj *ProductCategory, addUser *User) (id int64, errs []error) {
	o := orm.NewOrm()
	obj.CreateUser = addUser
	obj.UpdateUser = addUser
	var err error
	err = o.Begin()
	if err != nil {
		errs = append(errs, err)
	}
	if obj.ParentID != 0 {
		obj.Parent, _ = GetProductCategoryByID(obj.ParentID)
	}
	id, err = o.Insert(obj)
	if err != nil {
		errs = append(errs, err)
		err = o.Rollback()
		if err != nil {
			errs = append(errs, err)
		}
	} else {
		err = o.Commit()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return id, errs
}

// GetProductCategoryByID retrieves ProductCategory by ID. Returns error if
// ID doesn't exist
func GetProductCategoryByID(id int64) (obj *ProductCategory, err error) {
	o := orm.NewOrm()
	obj = &ProductCategory{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Parent != nil {
			o.Read(obj.Parent)
		}
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
	if limit == 0 {
		limit = 20
	}
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

// UpdateProductCategoryByID updates ProductCategory by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductCategoryByID(m *ProductCategory) (err error) {
	o := orm.NewOrm()
	v := ProductCategory{ID: m.ID}
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

// DeleteProductCategory deletes ProductCategory by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductCategory(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductCategory{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductCategory{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
