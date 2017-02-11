package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// SaleCounterProducct 柜台产品
type SaleCounterProducct struct {
	ID               int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser       *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser       *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate       time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate       time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	SaleCounter      *SaleCounter     `orm:"rel(fk)"`                              //柜台
	ProductProducts  *ProductProduct  `orm:"rel(fk);null"`                         //产品规格
	ProductTemplates *ProductTemplate `orm:"rel(fk)"`                              //产品款式
}

func init() {
	orm.RegisterModel(new(SaleCounterProducct))
}
