package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductProduct 产品规格
type ProductProduct struct {
	ID                    int64                    `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser            *User                    `orm:"rel(fk);null"`                //创建者
	UpdateUser            *User                    `orm:"rel(fk);null"`                //最后更新者
	CreateDate            time.Time                `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate            time.Time                `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name                  string                   `orm:"index"`                                //产品属性名称
	Company               *Company                 `orm:"rel(fk);null"`                         //公司
	Category              *ProductCategory         `orm:"rel(fk)"`                              //产品类别
	IsProductVariant      bool                     `orm:"default(true)"`                        //是多规格产品
	ProductTags           []*ProductTag            `orm:"reverse(many)"`                        //产品标签
	SaleOk                bool                     `orm:"default(true)" json:"SaleOk"`          //可销售
	Active                bool                     `orm:"default(true)"`                        //有效
	Barcode               string                   `orm:"null" json:"Barcode"`                  //条码,如ean13
	StandardPrice         float64                  `json:"StandardPrice"`                       //成本价格
	DefaultCode           string                   `orm:"unique"`                               //产品编码
	ProductTemplate       *ProductTemplate         `orm:"rel(fk)"`                              //产品款式
	AttributeValues       []*ProductAttributeValue `orm:"reverse(many)"`                        //产品属性值
	ProductType           string                   `orm:"default(stock)"`                       //产品类型
	AttributeValuesString string                   `orm:"index;default()"`                      //产品属性值ID编码，用于修改和增加时对应的产品是否已经存在
	FirstSaleUom          *ProductUom              `orm:"rel(fk)"`                              //第一销售单位
	SecondSaleUom         *ProductUom              `orm:"rel(fk);null"`                         //第二销售单位
	FirstPurchaseUom      *ProductUom              `orm:"rel(fk)"`                              //第一采购单位
	SecondPurchaseUom     *ProductUom              `orm:"rel(fk);null"`                         //第二采购单位
	PackagingDependTemp   bool                     `orm:"default(true)"`                        //根据款式打包
	BigImages             []*ProductImage          `orm:"reverse(many)"`                        //产品款式图片
	MidImages             []*ProductImage          `orm:"reverse(many)"`                        //产品款式图片
	SmallImages           []*ProductImage          `orm:"reverse(many)"`                        //产品款式图片
	PurchaseDependTemp    bool                     `orm:"default(true)"`                        //根据款式采购，ture一个供应商可以供应所有的款式
	// ProductPackagings     []*ProductPackaging      `orm:"reverse(many)"`                        //打包方式

}

func init() {
	orm.RegisterModel(new(ProductProduct))
}
