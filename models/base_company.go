package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Company 公司
type Company struct {
	ID           int64            `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64            `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64            `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time        `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time        `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string           `orm:"unique" json:"Name" form:"Name"`               //公司名称
	Code         string           `orm:"unique" json:"Code" form:"Code"`               //公司编码
	Children     []*Company       `orm:"reverse(many)" json:"-"`                       //子公司
	Parent       *Company         `orm:"rel(fk);null" json:"-" form:"-"`               //上级公司
	Country      *AddressCountry  `orm:"rel(fk);null" json:"-" form:"-"`               //国家
	Province     *AddressProvince `orm:"rel(fk);null" json:"-" form:"-"`               //省份
	City         *AddressCity     `orm:"rel(fk);null" json:"-" form:"-"`               //城市
	District     *AddressDistrict `orm:"rel(fk);null" json:"-" form:"-"`               //区县
	Street       string           `orm:"default()" json:"Street" form:"Street"`        //街道
	XMLID        string           `orm:"-"`                                            //xml初始化数据的ID，数据库中不保存
}

func init() {
	orm.RegisterModel(new(Company))
}

// AddCompany insert a new Company into database and returns last inserted Id on success.
func AddCompany(m *Company, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
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
