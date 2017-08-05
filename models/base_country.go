package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// AddressCountry 国家
type AddressCountry struct {
	ID           int64              `orm:"column(id);pk;auto"`                    //主键
	CreateUserID int64              `orm:"column(create_user_id);null"`           //创建者
	UpdateUserID int64              `orm:"column(update_user_id);null"`           //最后更新者
	CreateDate   time.Time          `orm:"auto_now_add;type(datetime)"`           //创建时间
	UpdateDate   time.Time          `orm:"auto_now;type(datetime)"`               //最后更新时间
	Name         string             `orm:"unique;size(50)" xml:"name"` //国家名称
	Provinces    []*AddressProvince `orm:"reverse(many)"`                          //省份
}

func init() {
	orm.RegisterModel(new(AddressCountry))

}

// AddAddressCountry insert a new AddressCountry into database and returns last inserted Id on success.
func AddAddressCountry(m *AddressCountry, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateAddressCountry update AddressCountry into database and returns id on success
func UpdateAddressCountry(m *AddressCountry, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// BatchAddAddressCountry insert  list of  Country into database and returns  number of  success.
func BatchAddAddressCountry(countries []*AddressCountry, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&AddressCountry{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, country := range countries {
			if _, err = i.Insert(country); err == nil {
				num = num + 1
			}
		}
	}
	return
}
