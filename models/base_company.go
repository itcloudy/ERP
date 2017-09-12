package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Company 公司
type Company struct {
	ID           int64            `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64            `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64            `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time        `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time        `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string           `orm:"unique"`                      //公司名称
	Code         string           `orm:"unique"`                      //公司编码
	Children     []*Company       `orm:"reverse(many)"`               //子公司
	Parent       *Company         `orm:"rel(fk);null"`                //上级公司
	Country      *AddressCountry  `orm:"rel(fk);null"`                //国家
	Province     *AddressProvince `orm:"rel(fk);null"`                //省份
	City         *AddressCity     `orm:"rel(fk);null"`                //城市
	District     *AddressDistrict `orm:"rel(fk);null"`                //区县
	Street       string           `orm:"default()"`                   //街道
}

func init() {
	orm.RegisterModel(new(Company))
}

// AddCompany insert a new Company into database and returns last inserted Id on success.
func AddCompany(m *Company, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// DeleteCompanyByID delete  Company by ID
func DeleteCompanyByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &Company{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// BatchAddCompany insert  list of  Company into database and returns  number of  success.
func BatchAddCompany(companies []*Company, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&Company{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, company := range companies {
			if _, err = i.Insert(company); err == nil {
				num = num + 1
			}
		}
	}
	return
}
