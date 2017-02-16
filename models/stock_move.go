package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockMove  	移动明细
type StockMove struct {
	ID           int64           `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser   *User           `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser   *User           `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate   time.Time       `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate   time.Time       `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name         string          `json:"Name"`                                //明细产品名称
	Picking      *StockPicking   `orm:"rel(fk)"`                              //调拨单
	Sequence     int64           `orm:"default(0)" json:"Sequence"`           //序列号
	Priority     string          `orm:"default(\"normal\")" json:"Priority"`  //优先级
	Product      *ProductProduct `orm:"rel(fk)"`                              //产品规格
	FirstUomQty  float64         `orm:"default(0)"`                           //第一单位数量
	SecondUomQty float64         `orm:"default(0)"`                           //第二单位数量
	FirstUom     *ProductUom     `orm:"rel(fk)"`                              //第一单位
	SecondUom    *ProductUom     `orm:"rel(fk);null"`                         //第二单位
	State        string          `orm:"default(\"draft\")" json:"State"`      //状态

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
}

func init() {
	orm.RegisterModel(new(StockMove))
}
