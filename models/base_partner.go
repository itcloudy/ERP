package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Partner 合作伙伴，包括客户和供应商，后期会为每个合作伙伴自动创建一个登录帐号
type Partner struct {
	ID         int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string           `orm:"unique" json:"Name"`                   //合作伙伴名称
	IsCompany  bool             `orm:"default(true)" json:"IsCompany"`       //是公司
	IsSupplier bool             `orm:"default(false)" json:"IsSupplier"`     //是供应商
	IsCustomer bool             `orm:"default(true)" json:"IsCustomer"`      //是客户
	Active     bool             `orm:"default(true)" json:"Active"`          //有效
	Country    *AddressCountry  `orm:"rel(fk);null"`                         //国家
	Province   *AddressProvince `orm:"rel(fk);null"`                         //省份
	City       *AddressCity     `orm:"rel(fk);null"`                         //城市
	District   *AddressDistrict `orm:"rel(fk);null"`                         //区县
	Street     string           `orm:"default(\"\")" json:"Street"`          //街道
	Parent     *Partner         `orm:"rel(fk);null"`                         //母公司
	Childs     []*Partner       `orm:"reverse(many)"`                        //下级
	Mobile     string           `orm:"default(\"\")" json:"Mobile"`          //电话号码
	Tel        string           `orm:"default(\"\")" json:"Tel"`             //座机
	Email      string           `orm:"default(\"\")" json:"Email"`           //邮箱
	Qq         string           `orm:"default(\"\")" json:"Qq"`              //QQ
	WeChat     string           `orm:"default(\"\")" json:"WeChat"`          //微信
	Comment    string           `orm:"type(text)" json:"Comment"`            //备注

}

func init() {
	orm.RegisterModel(new(Partner))
}
