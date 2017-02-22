package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockMoveOperationLink 将库存调拨与包装业务联系在一起
type StockMoveOperationLink struct {
	ID         int64      `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User      `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User      `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time  `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time  `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Move       *StockMove `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(StockMoveOperationLink))
}
