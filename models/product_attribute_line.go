package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductAttributeLine 产品属性明细
type ProductAttributeLine struct {
	ID              int64                    `orm:"column(id);pk;auto"`          //主键
	CreateUser      *User                    `orm:"rel(fk);null"`                //创建者
	UpdateUser      *User                    `orm:"rel(fk);null"`                //最后更新者
	CreateDate      time.Time                `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate      time.Time                `orm:"auto_now;type(datetime)"`     //最后更新时间
	Attribute       *ProductAttribute        `orm:"rel(fk)"`                     //属性
	ProductTemplate *ProductTemplate         `orm:"rel(fk)"`                     //产品模版
	AttributeValues []*ProductAttributeValue `orm:"rel(m2m)"`                    //属性值
}

func init() {
	orm.RegisterModel(new(ProductAttributeLine))
}

// AddProductAttributeLine insert a new ProductAttributeLine into database and returns
