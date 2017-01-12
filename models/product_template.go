package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

//ProductTemplate 产品款式
type ProductTemplate struct {
	Base
	Name               string                  `orm:"unique"` //产品属性名称
	Sequence           int32                   //序列号
	Description        string                  `orm:"type(text);null"` //描述
	DescriptioSale     string                  `orm:"type(text);null"` //销售描述
	DescriptioPurchase string                  `orm:"type(text);null"` //采购描述
	Rental             bool                    `orm:"default(false)"`  //代售品
	Categ              *ProductCategory        `orm:"rel(fk)"`         //产品类别
	Price              float64                 //模版产品价格
	StandardPrice      float64                 //成本价格
	SaleOk             bool                    `orm:"default(true)"` //可销售
	Active             bool                    `orm:"default(true)"` //有效
	IsProductVariant   bool                    `orm:"default(true)"` //是变形产品
	FirstSaleUom       *ProductUom             `orm:"rel(fk)"`       //第一销售单位
	SecondSaleUom      *ProductUom             `orm:"rel(fk)"`       //第二销售单位
	FirstPurchaseUom   *ProductUom             `orm:"rel(fk)"`       //第一采购单位
	SecondPurchaseUom  *ProductUom             `orm:"rel(fk)"`       //第二采购单位
	AttributeLines     []*ProductAttributeLine `orm:"reverse(many)"` //属性明细
	ProductVariants    []*ProductProduct       `orm:"reverse(many)"` //产品规格明细
	TemplatePackagings []*ProductPackaging     `orm:"reverse(many)"` //打包方式
	VariantCount       int32                   //产品规格数量
	Barcode            string                  //条码,如ean13
	DefaultCode        string                  //产品编码
	ProductType        string                  `orm:"default(\"stock\")"` //产品类型
	ProductMethod      string                  `orm:"default(\"hand\")"`  //产品规格创建方式
	// ProductPricelistItems []*ProductPricelistItem `orm:"reverse(many)"`
	PackagingDependTemp bool `orm:"default(true)"` //根据款式打包
	PurchaseDependTemp  bool `orm:"default(true)"` //根据款式采购，ture一个供应商可以供应所有的款式
}

func init() {
	orm.RegisterModel(new(ProductTemplate))
}

// AddProductTemplate insert a new ProductTemplate into database and returns
// last inserted ID on success.
func AddProductTemplate(obj *ProductTemplate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductTemplateByID retrieves ProductTemplate by ID. Returns error if
// ID doesn't exist
func GetProductTemplateByID(id int64) (obj *ProductTemplate, err error) {
	o := orm.NewOrm()
	obj = &ProductTemplate{Base: Base{ID: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetProductTemplateByName retrieves ProductTemplate by Name. Returns error if
// Name doesn't exist
func GetProductTemplateByName(name string) (obj *ProductTemplate, err error) {
	o := orm.NewOrm()
	obj = &ProductTemplate{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductTemplate retrieves all ProductTemplate matches certain condition. Returns empty list if
// no records exist
func GetAllProductTemplate(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductTemplate, error) {
	var (
		objArrs   []ProductTemplate
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductTemplate))
	qs = qs.RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return paginator, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return paginator, nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(limit, offset, cnt)
	}
	if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
		paginator.CurrentPageSize = num
	}
	return paginator, objArrs, err
}

// UpdateProductTemplateByID updates ProductTemplate by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductTemplateByID(m *ProductTemplate) (err error) {
	o := orm.NewOrm()
	v := ProductTemplate{Base: Base{ID: m.ID}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductTemplate deletes ProductTemplate by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductTemplate(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductTemplate{Base: Base{ID: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductTemplate{Base: Base{ID: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
