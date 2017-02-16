package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Role  角色
type Role struct {
	ID          int64         `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser  *User         `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser  *User         `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate  time.Time     `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate  time.Time     `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name        string        `orm:"unique" json:"Name"`                   //角色名称
	Users       []*User       `orm:"reverse(many)"`                        //角色所对应的用户
	Permissions []*Permission `orm:"reverse(many)"`                        //权限列表
	Menus       []*Menu       `orm:"rel(m2m)"`                             //用户可见的菜单

	FormAction    string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields  []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	UserIDs       []int64  `orm:"-" json:"UserIds"`
	PermissionIDs []int64  `orm:"-" json:"PermissionIds"`
	MenuIDs       []int64  `orm:"-" json:"MenuIds"`
}

func init() {
	orm.RegisterModel(new(Role))
}
func (u *Role) TableName() string {
	return "base_role"
}

// Role insert a new Role into database and returns
// last inserted ID on success.
func AddRole(obj *Role, addUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.CreateUser = addUser
	obj.UpdateUser = addUser
	errBegin := o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return 0, errBegin
	}
	if id, err = o.Insert(obj); err == nil {
		obj.ID = id
		for _, item := range obj.UserIDs {
			m2m := o.QueryM2M(obj, "Users")
			if user, err := GetUserByID(item); err == nil {
				m2m.Add(user)
			}
		}
		for _, item := range obj.PermissionIDs {
			m2m := o.QueryM2M(obj, "Permissions")
			if permission, err := GetPermissionByID(item); err == nil {
				m2m.Add(permission)
			}
		}
		for _, item := range obj.MenuIDs {
			m2m := o.QueryM2M(obj, "Menus")
			if menu, err := GetMenuByID(item); err == nil {
				m2m.Add(menu)
			}
		}
	}
	if err == nil {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetRoleByID retrieves Role by ID. Returns error if
// ID doesn't exist
func GetRoleByID(id int64) (obj *Role, err error) {
	o := orm.NewOrm()
	obj = &Role{ID: id}
	if err = o.Read(obj); err == nil {
		o.LoadRelated(obj, "Users")
		o.LoadRelated(obj, "Permissions")
		o.LoadRelated(obj, "Menus")
		return obj, err
	}
	return nil, err
}

// GetRoleByName retrieves Role by Name. Returns error if
// Name doesn't exist
func GetRoleByName(name string) (*Role, error) {
	o := orm.NewOrm()
	var obj Role
	cond := orm.NewCondition()
	cond = cond.And("Name", name)
	qs := o.QueryTable(&obj)
	qs = qs.SetCond(cond)
	err := qs.One(&obj)
	return &obj, err
}

// GetAllRole retrieves all Role matches certain condition. Returns empty list if
// no records exist
func GetAllRole(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []Role, error) {
	var (
		objArrs   []Role
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
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
					o.LoadRelated(&objArrs[i], "Permissions")
					o.LoadRelated(&objArrs[i], "Users")
					o.LoadRelated(&objArrs[i], "Menus")
				}
			}
		}
	}

	return paginator, objArrs, err
}

// UpdateRole updates Role by ID and returns error if
// the record to be updated doesn't exist
func UpdateRole(obj *Role, updateUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.UpdateUser = updateUser
	var num int64
	if num, err = o.Update(obj); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return obj.ID, err
}

// DeleteRole deletes Role by ID and returns error if
// the record to be deleted doesn't exist
func DeleteRole(id int64) (err error) {
	o := orm.NewOrm()
	v := Role{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Role{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
