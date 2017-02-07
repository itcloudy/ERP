package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockLocation 库位
type StockLocation struct {
	ID               int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser       *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser       *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate       time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate       time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name             string           `orm:"unique"`                               //库位名称
	Company          *Company         `orm:"rel(fk);null"`                         //公司
	Active           bool             `orm:"default(true)"`                        //有效
	Barcode          string           `json:"Barcode"`                             //条码
	Parent           *StockLocation   `orm:"rel(fk);null"`                         //上级库位
	Childs           []*StockLocation `orm:"reverse(many)"`                        //子库位
	ReturnLocation   bool             `orm:"default(false)"`                       //是一个退货库位
	ScrapLocation    bool             `orm:"default(false)"`                       //是一个废料库位
	Posx             int64            `json:"Posx"`                                //通道(X)
	Posy             int64            `json:"Posy"`                                //货架(Y)
	Posz             int64            `json:"Posz"`                                //层
	PutawayStrategy  string           ``                                           //入库策略,需后续改为many2one
	RemovalStrategyd string           ``                                           //出库策略需后续改为many2one
}

func init() {
	orm.RegisterModel(new(StockLocation))
}
