package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductTag  产品标签
type ProductTag struct {
	ID         int64             `orm:"column(id);pk;auto"`          //主键
	CreateUser *User             `orm:"rel(fk);null"`                //创建者
	UpdateUser *User             `orm:"rel(fk);null"`                //最后更新者
	CreateDate time.Time         `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate time.Time         `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name       string            `orm:"size(20);unique"`             //产品标签名称
	Type       string            `orm:"size(20);default()"`          //标签类型，
	Products   []*ProductProduct `orm:"rel(m2m)"`                    //产品规格

}

func init() {
	orm.RegisterModel(new(ProductTag))
}
