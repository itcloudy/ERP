package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// StockQuant  	库存分析
type StockQuant struct {
	ID                   int64               `orm:"column(id);pk;auto" json:"id"`                  //主键
	CreateUser           *User               `orm:"rel(fk);null" json:"-"`                         //创建者
	UpdateUser           *User               `orm:"rel(fk);null" json:"-"`                         //最后更新者
	CreateDate           time.Time           `orm:"auto_now_add;type(datetime)" json:"-"`          //创建时间
	UpdateDate           time.Time           `orm:"auto_now;type(datetime)" json:"-"`              //最后更新时间
	Name                 string              `json:"Name"`                                         //编号
	Product              *ProductProduct     `orm:"rel(fk)"`                                       //产品
	Location             *StockLocation      `orm:"rel(fk);null"`                                  //库位
	FirstUomQty          float64             `orm:"default(0)"`                                    //第一单位数量
	SecondUomQty         float64             `orm:"default(0)"`                                    //第二单位数量
	FirstUom             *ProductUom         `orm:"rel(fk)"`                                       //第一单位
	SecondUom            *ProductUom         `orm:"rel(fk);null"`                                  //第二单位
	Package              *StockQuantPackage  `orm:"rel(fk);null"`                                  //物理包装
	PackagingType        *ProductPackaging   `orm:"rel(fk)"`                                       //包装
	Reservation          *StockMove          `orm:"rel(fk);null"`                                  //调拨保留
	Lot                  *StockProductionLot `orm:"rel(fk)"`                                       // 批次
	Cost                 float64             `orm:"default(0)"`                                    //成本
	InDate               time.Time           `orm:"auto_now_add;type(datetime)"`                   //接收时间
	Historys             []*StockMove        `orm:"reverse(many);rel_table(stock_quant_move_rel)"` //调拨
	Company              *Company            `orm:"rel(fk)"`                                       //公司
	PropagatedFrom       *StockQuant         `orm:"rel(fk);null"`                                  //连接份数
	NegativeDestLocation *StockLocation      `orm:"rel(fk);null"`                                  //负值目标库位
	NegativeMove         *StockMove          `orm:"rel(fk);null"`                                  //调拨负数分析

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	CompanyID    int64    `orm:"-" json:"Company"`
}

func init() {
	orm.RegisterModel(new(StockQuant))
}
