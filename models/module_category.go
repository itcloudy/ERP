package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModuleCategory 模块(表)分类
type ModuleCategory struct {
	ID           int64     `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64     `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64     `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string    `orm:"size(50)" json:"name" form:"Name"`             //模块分类名称
	XMLID        string    `orm:"-"`                                            //xml初始化数据的ID，数据库中不保存

}

func init() {
	orm.RegisterModel(new(ModuleCategory))
}
