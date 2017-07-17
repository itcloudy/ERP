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

// AddAddressProvince insert a new AddressProvince into database and returns last inserted Id on success.
func AddAddressProvince(m *AddressProvince, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateAddressProvince update AddressProvince into database and returns id on success
func UpdateAddressProvince(m *AddressProvince, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}
