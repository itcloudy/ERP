package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Department 部门
type Department struct {
	ID         int64       `orm:"column(id);pk;auto" json:"id"`              //主键
	CreateUser *User       `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User       `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time   `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time   `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string      `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name       string      `orm:"unique"`                               //团队名称
	Leader     *User       `orm:"rel(fk);null"`                         //团队领导者
	Parent     *Department `orm:"rel(fk);null"`                         //上级分类
	Members    []*User     `orm:"reverse(many)"`                        //组员
	Company    *Company    `orm:"rel(fk);null"`                         //公司
}

func init() {
	orm.RegisterModel(new(Department))
}

// AddDepartment insert a new Department into database and returns
// last inserted ID on success.
func AddDepartment(obj *Department) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetDepartmentByID retrieves Department by ID. Returns error if
// ID doesn't exist
func GetDepartmentByID(id int64) (obj *Department, err error) {
	o := orm.NewOrm()
	obj = &Department{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetDepartmentByName retrieves Department by ID. Returns error if
// Name doesn't exist
func GetDepartmentByName(name string) (obj *Department, err error) {
	o := orm.NewOrm()
	obj = &Department{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllDepartment retrieves all Department matches certain condition. Returns empty list if
// no records exist
func GetAllDepartment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []Department, error) {
	var (
		objArrs   []Department
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Department))
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

// UpdateDepartmentByID updates Department by ID and returns error if
// the record to be updated doesn't exist
func UpdateDepartmentByID(m *Department) (err error) {
	o := orm.NewOrm()
	v := Department{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDepartment deletes Department by ID and returns error if
// the record to be deleted doesn't exist
func DeleteDepartment(id int64) (err error) {
	o := orm.NewOrm()
	v := Department{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Department{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
