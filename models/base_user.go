package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User 登录用户
type User struct {
	ID              int64        `orm:"column(id);pk;auto"`                                                    //主键
	CreateUserID    int64        `orm:"column(create_user_id);null" json:"-"`                                  //创建者
	UpdateUserID    int64        `orm:"column(update_user_id);null" json:"-"`                                  //最后更新者
	CreateDate      time.Time    `orm:"auto_now_add;type(datetime)" json:"-"`                                  //创建时间
	UpdateDate      time.Time    `orm:"auto_now;type(datetime)" json:"-"`                                      //最后更新时间
	Name            string       `orm:"size(20);unique" xml:"name" json:"Name" form:"Name"`                    //用户名
	Company         *Company     `orm:"rel(fk);null" json:"-"`                                                 //公司
	NameZh          string       `orm:"size(20)"  xml:"NameZh" json:"NameZh" form:"NameZh"`                    //中文用户名
	Email           string       `orm:"size(20);unique" xml:"email" json:"Email" form:"Email"`                 //邮箱
	Mobile          string       `orm:"size(20);unique" xml:"mobile" json:"Mobile" form:"Mobile"`              //手机号码
	Tel             string       `orm:"size(20);default()" json:"Tel" form:"Tel"`                              //固定号码
	Password        string       `xml:"password" json:"Password" form:"Password"`                              //密码
	ConfirmPassword string       `orm:"-" xml:"ConfirmPassword" json:"ConfirmPassword" form:"ConfirmPassword"` //确认密码,数据库中不保存
	IsAdmin         bool         `orm:"default(false)" xml:"isAdmin" json:"IsAdmin" form:"IsAdmin"`            //是否为超级用户
	Active          bool         `orm:"default(true)" xml:"active" json:"Active" form:"Active"`                //有效
	Qq              string       `orm:"default()" xml:"qq" json:"Qq" form:"Qq"`                                //QQ
	WeChat          string       `orm:"default()" xml:"wechat" json:"WeChat" form:"WeChat"`                    //微信
	Groups          []*BaseGroup `orm:"rel(m2m);rel_through(golangERP/models.GroupUser)"`                      //权限组
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
