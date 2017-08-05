package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModuleModule 模块(表)名称
type ModuleModule struct {
	ID           int64           `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64           `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64           `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time       `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time       `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string          `orm:"unique;size(50)" xml:"name"`  //表名称
	Description  string          `xml:"description"`                 //说明
	Category     *ModuleCategory `orm:"rel(fk);null"`                //模块分类

}

func init() {
	orm.RegisterModel(new(ModuleModule))
}

// AddModuleModule insert a new ModuleModule into database and returns last inserted Id on success.
func AddModuleModule(m *ModuleModule, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateModuleModule update ModuleModule into database and returns id on success
func UpdateModuleModule(m *ModuleModule, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetModuleModuleByName retrieves ModuleModule by ID. Returns error if ID doesn't exist
func GetModuleModuleByName(name string, ormObj orm.Ormer) (*ModuleModule, error) {
	var obj ModuleModule
	var err error
	qs := ormObj.QueryTable(&obj)
	err = qs.Filter("Name", name).One(&obj)
	return &obj, err
}
