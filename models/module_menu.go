package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModuleMenu 城市
type ModuleMenu struct {
	ID           int64         `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64         `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64         `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time     `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time     `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string        `orm:"size(50)" json:"name" form:"Name"`             //菜单名称
	Parent       *ModuleMenu   `orm:"rel(fk)" json:"province" form:"-"`             //上级菜单
	Childs       []*ModuleMenu `orm:"reverse(many)" json:"districts"`               //子菜单
	ParenLeft    int64         `orm:"unique"`                                       //菜单左
	ParenRight   int64         `orm:"unique"`                                       //菜单右
	Sequence     int64         `orm:"default(1)"`                                   //序列号，决定同级菜单显示先后顺序
	Icon         string        `orm:""`                                             //菜单图标样式
}

func init() {
	orm.RegisterModel(new(ModuleMenu))
}

// AddModuleMenu insert a new ModuleMenu into database and returns last inserted Id on success.
func AddModuleMenu(m *ModuleMenu, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateModuleMenu update ModuleMenu into database and returns id on success
func UpdateModuleMenu(m *ModuleMenu, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}
