package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// AddressProvince 省份
type AddressProvince struct {
	ID           int64           `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64           `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64           `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time       `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time       `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string          `xml:"ProvinceName,attr" json:"Name" form:"Name"`    //省份名称
	Country      *AddressCountry `orm:"rel(fk)" form:"-"`                             //国家
	Citys        []*AddressCity  `orm:"reverse(many)"`                                //城市
}

func init() {
	orm.RegisterModel(new(AddressProvince))
}
