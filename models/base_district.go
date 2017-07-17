package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// AddressDistrict 区县
type AddressDistrict struct {
	ID           int64        `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64        `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64        `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time    `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time    `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string       `json:"Name" form:"Name"`                            //区县名称
	City         *AddressCity `orm:"rel(fk)" form:"-"`                             //城市
}

func init() {
	orm.RegisterModel(new(AddressDistrict))
}

// AddAddressDistrict insert a new AddressDistrict into database and returns last inserted Id on success.
func AddAddressDistrict(m *AddressDistrict, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}
