package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//  StockPicking 调拨单
type StockPicking struct {
	ID           int64             `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser   *User             `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser   *User             `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate   time.Time         `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate   time.Time         `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name         string            `orm:"unique" json:"Name"`                   //单据名称
	Origin       string            `json:"Origin"`                              //源单据
	Note         string            `orm:"type(text)" json:"Note"`               //备注
	MoveType     string            `orm:"default(\"one\")" json:"MoveType"`     //移动类型:one partial
	State        string            `orm:"default(\"draft\")" json:"-"`          //状态
	Company      *Company          `orm:"rel(fk)"`                              //公司
	LocationDest *StockLocation    `orm:"rel(fk)"`                              //目标库位
	LocationSrc  *StockLocation    `orm:"rel(fk)"`                              //源库位
	Partner      *Partner          `orm:"rel(fk)"`                              //合作伙伴
	Priority     string            `orm:"default(\"normal\")" json:"Priority"`  //优先级
	PickingType  *StockPickingType `orm:"rel(fk)"`                              //分拣类型决定分拣视图
}

func init() {
	orm.RegisterModel(new(StockPicking))
}
