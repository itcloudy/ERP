package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockWarehouse  仓库
type StockWarehouse struct {
	ID         int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string           `orm:"unique" json:"Name"`                   //仓库名称
	Code       string           `orm:"unique" json:"	"`                      //仓库编码
	Company    *Company         `orm:"rel(fk)"`                              //所属公司
	Country    *AddressCountry  `orm:"rel(fk);null" json:"-"`                //国家
	Province   *AddressProvince `orm:"rel(fk);null" json:"-"`                //省份
	City       *AddressCity     `orm:"rel(fk);null" json:"-"`                //城市
	District   *AddressDistrict `orm:"rel(fk);null" json:"-"`                //区县
	Street     string           `orm:"default(\"\")" json:"Street"`          //街道
}

func init() {
	orm.RegisterModel(new(StockWarehouse))
}
