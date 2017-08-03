package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// BaseGroup  权限组
type BaseGroup struct {
	ID            int64          `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID  int64          `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID  int64          `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate    time.Time      `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate    time.Time      `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name          string         `orm:"unique;size(50)" json:"name" form:"Name"`      //权限组名称
	ModelAccesses []*ModelAccess `orm:"reverse(many)"`                                //模块(表)
	Childs        []*BaseGroup   `orm:"reverse(many)" json:"-" form:"-"`              //继承权限
	Parent        *BaseGroup     `orm:"rel(fk);null"`                                 //
	ParentLeft    int64          `orm:"unique"`                                       //左边界
	ParentRight   int64          `orm:"unique"`                                       //右边界
	Category      string         `orm:""`                                             //分类
	Description   string         ``                                                   //说明
}

func init() {
	orm.RegisterModel(new(BaseGroup))
}

// AddBaseGroup insert a new BaseGroup into database and returns last inserted Id on success.
func AddBaseGroup(m *BaseGroup, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddBaseGroup insert  list of  BaseGroup into database and returns  number of  success.
func BatchAddBaseGroup(groups []*BaseGroup, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&BaseGroup{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, group := range groups {
			if _, err = i.Insert(group); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateBaseGroup update BaseGroup into database and returns id on success
func UpdateBaseGroup(m *BaseGroup, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetBaseGroupByID retrieves BaseGroup by ID. Returns error if ID doesn't exist
func GetBaseGroupByID(id int64, ormObj orm.Ormer) (obj *BaseGroup, err error) {
	obj = &BaseGroup{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetBaseGroupByName retrieves BaseGroup by ID. Returns error if ID doesn't exist
func GetBaseGroupByName(name string, ormObj orm.Ormer) (*BaseGroup, error) {
	var obj BaseGroup
	var err error
	qs := ormObj.QueryTable(&obj)
	err = qs.Filter("Name", name).One(&obj)
	return &obj, err
}
