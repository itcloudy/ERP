package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// AddressCity 城市
type AddressCity struct {
	ID           int64              `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64              `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64              `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time          `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time          `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string             `orm:"size(50)" json:"name" form:"Name"`             //城市名称
	Province     *AddressProvince   `orm:"rel(fk)" json:"province" form:"-"`             //省份
	Districts    []*AddressDistrict `orm:"reverse(many)" json:"districts"`               //区县
}

func init() {
	orm.RegisterModel(new(AddressCity))
}
