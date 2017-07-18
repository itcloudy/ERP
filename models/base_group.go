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
	Inherits      []*BaseGroup   `orm:"reverse(many)" json:"-" form:"-"`              //继承权限
	Child         *BaseGroup     `orm:"rel(fk);null"`                                 //
	ParenLeft     int64          `orm:"unique"`                                       //左边界
	ParenRight    int64          `orm:"unique"`                                       //右边界
	XMLID         string         `orm:"-"`                                            //xml初始化数据的ID，数据库中不保存
}

func init() {
	orm.RegisterModel(new(BaseGroup))
}
