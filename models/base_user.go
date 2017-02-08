package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// User table
// 用户表
type User struct {
	ID              int64       `orm:"column(id);pk;auto" json:"id"`                      //主键
	CreateUser      *User       `orm:"rel(fk);null" json:"-"`                             //创建者
	UpdateUser      *User       `orm:"rel(fk);null" json:"-"`                             //最后更新者
	CreateDate      time.Time   `orm:"auto_now_add;type(datetime)" json:"-"`              //创建时间
	UpdateDate      time.Time   `orm:"auto_now;type(datetime)" json:"-"`                  //最后更新时间
	Name            string      `orm:"size(20)" xml:"name" json:"Name"`                   //用户名
	Company         *Company    `orm:"rel(fk);null"`                                      //公司
	NameZh          string      `orm:"size(20)"  json:"NameZh"`                           //中文用户名
	Department      *Department `orm:"rel(fk);null;"  json:"department"`                  //部门
	DepartmentID    int64       `orm:"-" json:"Department"`                               //部门，用于form表单
	Email           string      `orm:"size(20)" xml:"email" json:"Email" json:"email"`    //邮箱
	Mobile          string      `orm:"size(20);default(\"\")" xml:"mobile" json:"Mobile"` //手机号码
	Tel             string      `orm:"size(20);default(\"\")" json:"Tel" json:"tel"`      //固定号码
	Password        string      `xml:"password" json:"Password" json:"password"`          //密码
	ConfirmPassword string      `orm:"-" xml:"ConfirmPassword" json:"ConfirmPassword"`    //确认密码,数据库中不保存
	Roles           []*Role     `orm:"reverse(many)"`                                     //用户拥有的角色
	Teams           []*Team     `orm:"rel(m2m);rel_table(user_teams)"`                    //团队
	TeamIDs         []string    `orm:"-" json:"Team"`                                     //团队，用于form表单
	IsAdmin         bool        `orm:"default(false)" xml:"isAdmin" json:"IsAdmin"`       //是否为超级用户
	Active          bool        `orm:"default(true)" xml:"active" json:"Active"`          //有效
	Qq              string      `orm:"default(\"\")" xml:"qq" json:"Qq"`                  //QQ
	WeChat          string      `orm:"default(\"\")" xml:"wechat" json:"WeChat"`          //微信
	Position        *Position   `orm:"rel(fk);null;" json:"Position"`                     //职位
	PositionID      int64       `orm:"-" json:"Position"`                                 //职位，用于form表单
	FormAction      string      `orm:"-" json:"FormAction"`                               //非数据库字段，用于表示记录的增加，修改

}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted ID on success.
func AddUser(obj *User) (id int64, err error) {
	o := orm.NewOrm()
	password := utils.PasswordMD5(obj.Password, obj.Mobile)
	obj.Password = password
	id, err = o.Insert(obj)
	return id, err
}

// GetUserByID retrieves User by ID. Returns error if
// ID doesn't exist
func GetUserByID(id int64) (obj *User, err error) {
	o := orm.NewOrm()
	obj = &User{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Department != nil {
			o.Read(obj.Department)
		}

		if obj.Position != nil {
			o.Read(obj.Position)
		}
		return obj, nil
	}
	return nil, err
}

// GetUserByName get user
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email", name).Or("name", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	err := qs.One(&user)
	return user, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []User, error) {
	var (
		objArrs   []User
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

// UpdateUserByID updates User by ID and returns error if
// the record to be updated doesn't exist
func UpdateUserByID(m *User) (err error) {
	o := orm.NewOrm()
	v := User{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by ID and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// CheckUserByName  check
func CheckUserByName(name, password string) (User, bool, error) {
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
	return user, ok, err
}
