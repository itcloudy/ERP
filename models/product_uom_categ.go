package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductUomCateg 产品单位类别
type ProductUomCateg struct {
	ID         int64         `orm:"column(id);pk;auto"`          //主键
	CreateUser *User         `orm:"rel(fk);null"`                //创建者
	UpdateUser *User         `orm:"rel(fk);null"`                //最后更新者
	CreateDate time.Time     `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate time.Time     `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name       string        `orm:"unique"`                      //计量单位分类
	Uoms       []*ProductUom `orm:"reverse(many)"`               //计量单位
}

func init() {
	orm.RegisterModel(new(ProductUomCateg))
}
