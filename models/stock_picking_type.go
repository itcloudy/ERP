package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockPickingType 分拣类型决定分拣视图
type StockPickingType struct {
	ID         int64             `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User             `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User             `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time         `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time         `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name       string            `orm:"unique" json:"Name"`                   //仓库名称
	Code       string            `json:"Code"`                                //移库类型
	WareHouse  *StockWarehouse   `orm:"rel(fk)"`                              //仓库
	NextStep   *StockPickingType `orm:"null;rel(one)"`                        //下一步
	PrevStep   *StockPickingType `orm:"null;rel(one)"`                        //上一步
}

func init() {
	orm.RegisterModel(new(StockPickingType))
}
