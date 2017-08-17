package models

import (
	"errors"
	"strings"
	"time"

	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// User 登录用户
type User struct {
	ID              int64        `orm:"column(id);pk;auto"`                               //主键
	CreateUserID    int64        `orm:"column(create_user_id);null"`                      //创建者
	UpdateUserID    int64        `orm:"column(update_user_id);null"`                      //最后更新者
	CreateDate      time.Time    `orm:"auto_now_add;type(datetime)"`                      //创建时间
	UpdateDate      time.Time    `orm:"auto_now;type(datetime)"`                          //最后更新时间
	Name            string       `orm:"size(20);unique" xml:"name"`                       //用户名
	Company         *Company     `orm:"rel(fk);null"`                                     //公司
	NameZh          string       `orm:"size(20)"  xml:"NameZh"`                           //中文用户名
	Email           string       `orm:"size(20);unique" xml:"email"`                      //邮箱
	Mobile          string       `orm:"size(20);unique" xml:"mobile"`                     //手机号码
	Tel             string       `orm:"size(20);default()"`                               //固定号码
	Password        string       `xml:"password"`                                         //密码
	ConfirmPassword string       `orm:"-" xml:"ConfirmPassword"`                          //确认密码,数据库中不保存
	IsAdmin         bool         `orm:"default(false)" xml:"isAdmin"`                     //是否为超级用户
	Active          bool         `orm:"default(true)" xml:"active"`                       //有效
	Qq              string       `orm:"default()" xml:"qq"`                               //QQ
	WeChat          string       `orm:"default()" xml:"wechat"`                           //微信
	Groups          []*BaseGroup `orm:"rel(m2m);rel_through(golangERP/models.GroupUser)"` //权限组
	IsBackground    bool         `orm:"defalut(false)" xml:"isbackground"`                //后台用户可以登录后台
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns last inserted Id on success.
func AddUser(m *User, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateUser update User into database and returns id on success
func UpdateUser(m *User, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetUserByID retrieves User by ID. Returns error if ID doesn't exist
func GetUserByID(id int64, ormObj orm.Ormer) (obj *User, err error) {
	obj = &User{ID: id}
	err = ormObj.Read(obj)
	ormObj.LoadRelated(obj, "Groups")
	return obj, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if no records exist
func GetAllUser(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []User, error) {
	var (
		objArrs   []User
		err       error
		paginator utils.Paginator
		num       int64
	)
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
