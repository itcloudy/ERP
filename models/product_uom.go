package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductUom 产品单位
type ProductUom struct {
	ID         int64            `orm:"column(id);pk;auto"`          //主键
	CreateUser *User            `orm:"rel(fk);null"`                //创建者
	UpdateUser *User            `orm:"rel(fk);null"`                //最后更新者
	CreateDate time.Time        `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate time.Time        `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name       string           `orm:"unique"`                      //计量单位名称
	Active     bool             `orm:"default(true)"`               //有效
	Category   *ProductUomCateg `orm:"rel(fk)"`                     //计量单位类别
	Factor     float64          ``                                  //比率
	FactorInv  float64          ``                                  //更大比率
	Rounding   float64          ``                                  //舍入精度
	Type       int64            ``                                  //类型
	Symbol     string           ``                                  //符号，后置

}

func init() {
	orm.RegisterModel(new(ProductUom))
}
