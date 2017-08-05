package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModuleCategory 模块(表)分类
type ModuleCategory struct {
	ID           int64     `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64     `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64     `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string    `orm:"size(50)" xml:"name"`        //模块分类名称
	Description  string    `xml:"description"`                  //说明

}

func init() {
	orm.RegisterModel(new(ModuleCategory))
}

// AddModuleCategory insert a new ModuleCategory into database and returns last inserted Id on success.
func AddModuleCategory(m *ModuleCategory, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddModuleCategory insert  list of  ModuleCategory into database and returns  number of  success.
func BatchAddModuleCategory(groups []*ModuleCategory, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&ModuleCategory{})
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

// UpdateModuleCategory update ModuleCategory into database and returns id on success
func UpdateModuleCategory(m *ModuleCategory, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetModuleCategoryByID retrieves ModuleCategory by ID. Returns error if ID doesn't exist
func GetModuleCategoryByID(id int64, ormObj orm.Ormer) (obj *ModuleCategory, err error) {
	obj = &ModuleCategory{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetModuleCategoryByName retrieves ModuleCategory by ID. Returns error if ID doesn't exist
func GetModuleCategoryByName(name string, ormObj orm.Ormer) (*ModuleCategory, error) {
	var obj ModuleCategory
	var err error
	qs := ormObj.QueryTable(&obj)
	err = qs.Filter("Name", name).One(&obj)
	return &obj, err
}
