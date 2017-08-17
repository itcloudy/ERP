package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ProductTemplate 产品款式
type ProductTemplate struct {
	ID                  int64                   `orm:"column(id);pk;auto"`          //主键
	CreateUserID        int64                   `orm:"column(create_user_id);null"` //创建者
	UpdateUserID        int64                   `orm:"column(update_user_id);null"` //最后更新者
	CreateDate          time.Time               `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate          time.Time               `orm:"auto_now;type(datetime)"`     //最后更新时间
	Description         string                  `orm:"type(text);null"`             //描述
	DescriptionSale     string                  `orm:"type(text);null"`             //销售描述
	DescriptionPurchase string                  `orm:"type(text);null"`             //采购描述
	Rental              bool                    `orm:"default(false)"`              //代售品
	Category            *ProductCategory        `orm:"rel(fk)"`                     //产品类别
	Price               float64                 ``                                  //模版产品价格
	StandardPrice       float64                 ``                                  //成本价格
	StandardWeight      float64                 ``                                  //标准重量
	SaleOk              bool                    `orm:"default(true)"`               //可销售
	Active              bool                    `orm:"default(true)"`               //有效
	IsProductVariant    bool                    `orm:"default(true)"`               //是变形产品
	FirstSaleUom        *ProductUom             `orm:"rel(fk)"`                     //第一销售单位
	SecondSaleUom       *ProductUom             `orm:"rel(fk);null"`                //第二销售单位
	FirstPurchaseUom    *ProductUom             `orm:"rel(fk)"`                     //第一采购单位
	SecondPurchaseUom   *ProductUom             `orm:"rel(fk);null"`                //第二采购单位
	AttributeLines      []*ProductAttributeLine `orm:"reverse(many)"`               //属性明细
	ProductVariants     []*ProductProduct       `orm:"reverse(many)"`               //产品规格明细
	VariantCount        int32                   ``                                  //产品规格数量
	Barcode             string                  ``                                  //条码,如ean13
	DefaultCode         string                  ``                                  //产品编码
	BigImages           []*ProductImage         `orm:"reverse(many)"`               //产品款式图片
	MidImages           []*ProductImage         `orm:"reverse(many)"`               //产品款式图片
	SmallImages         []*ProductImage         `orm:"reverse(many)"`               //产品款式图片
	ProductType         string                  `orm:"default()"`                   //产品类型 stock consume service
	ProductMethod       string                  `orm:"default()"`                   //产品规格创建方式 auto hand
	// TemplatePackagings  []*ProductPackaging     `orm:"reverse(many)"`                                       //打包方式
}

func init() {
	orm.RegisterModel(new(ProductTemplate))
}
