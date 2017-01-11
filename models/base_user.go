package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 用户表
type User struct {
	Base
	Name            string      `orm:"size(20)" xml:"name" form:"Name" json:"name"`                           //用户名
	NameZh          string      `orm:"size(20)"  form:"NameZh" json:"namezh"`                                 //中文用户名
	Department      *Department `orm:"rel(fk);null;" form:"Department" json:"department"`                     //部门
	Email           string      `orm:"size(20)" xml:"email" form:"Email" json:"email"`                        //邮箱
	Mobile          string      `orm:"size(20);default(\"\")" xml:"Mobile" form:"mobile" json:"mobile"`       //手机号码
	Tel             string      `orm:"size(20);default(\"\")" form:"Tel" json:"tel"`                          //固定号码
	Password        string      `xml:"password" form:"Password" json:"password"`                              //密码
	ConfirmPassword string      `orm:"-" xml:"ConfirmPassword" form:"confirmpassword" json:"confirmpassword"` //确认密码,数据库中不保存
	Groups          []*Group    `orm:"rel(m2m);rel_table(user_groups)"`                                       //权限组
	Teams           []*Team     `orm:"rel(m2m);rel_table(user_teams)"`                                        //权限组
	IsAdmin         bool        `orm:"default(false)" xml:"isAdmin" form:"IsAdmin" json:"isadmin"`            //是否为超级用户
	Active          bool        `orm:"default(true)" xml:"active" form:"Active" json:"active"`                //有效
	Qq              string      `orm:"default(\"\")" xml:"qq" form:"Qq" json:"qq"`                            //QQ
	WeChat          string      `orm:"default(\"\")" xml:"wechat" form:"WeChat" json:"wechat"`                //微信
	Position        *Position   `orm:"rel(fk);null;" form:"Position" json:"position"`                         //职位
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(obj *User) (id int64, err error) {
	o := orm.NewOrm()
	password := utils.PasswordMD5(obj.Password, obj.Mobile)
	obj.Password = password
	id, err = o.Insert(obj)
	return id, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (obj *User, err error) {
	o := orm.NewOrm()
	obj = &User{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email", name).Or("name", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	err := qs.One(&user)
	if user.Department != nil {
		o.Read(user.Department)
	}
	if user.Groups != nil {
		o.LoadRelated(user, "Groups")
	}
	if user.Position != nil {
		o.Read(user.Position)
	}
	return user, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []User, error) {
	var (
		objArrs   []User
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func CheckUserByName(name, password string) (User, error, bool) {
	o := orm.NewOrm()
	var (
		user User
		err  error
		ok   bool
	)
	ok = false
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("mobile", name).Or("email", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			ok = true
		}
	}
	return user, err, ok
}
