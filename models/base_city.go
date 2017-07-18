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
	XMLID        string             `orm:"-"`                                            //xml初始化数据的ID，数据库中不保存
}

func init() {
	orm.RegisterModel(new(AddressCity))
}

// AddAddressCity insert a new AddressCity into database and returns last inserted Id on success.
func AddAddressCity(m *AddressCity, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddAddressCity insert  list of  AddressCity into database and returns  number of  success.
func BatchAddAddressCity(cities []*AddressCity, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&AddressCity{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, city := range cities {
			if _, err = i.Insert(city); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateAddressCity update AddressCity into database and returns id on success
func UpdateAddressCity(m *AddressCity, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}
