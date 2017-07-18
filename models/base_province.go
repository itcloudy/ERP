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
	XMLID        string          `orm:"-"`                                            //xml初始化数据的ID，数据库中不保存
}

func init() {
	orm.RegisterModel(new(AddressProvince))
}

// AddAddressProvince insert a new AddressProvince into database and returns last inserted Id on success.
func AddAddressProvince(m *AddressProvince, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddAddressProvince insert  list of  Country into database and returns  number of  success.
func BatchAddAddressProvince(privinces []*AddressProvince, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&AddressProvince{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, province := range privinces {
			if _, err = i.Insert(province); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateAddressProvince update AddressProvince into database and returns id on success
func UpdateAddressProvince(m *AddressProvince, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}
