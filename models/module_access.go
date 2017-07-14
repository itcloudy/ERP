package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModelAccess 模块(表)操作权限
type ModelAccess struct {
	ID           int64             `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64             `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64             `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time         `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time         `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Module       *ModuleTable      `orm:"rel(fk)"`                                      //模块(表)
	Permission   *ModulePermission `orm:"rel(fk)"`                                      //权限组
	PermCreate   bool              `orm:"default(true)"`                                //创建权限
	PermUnlink   bool              `orm:"default(false)"`                               //删除权限
	PermWrite    bool              `orm:"default(true)"`                                //修改权限
	PermRead     bool              `orm:"default(true)"`                                //读权限
}

func init() {
	orm.RegisterModel(new(ModelAccess))
}
