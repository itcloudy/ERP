package models

import (
	"errors"
	"fmt"
	"strings"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// 产品规格
type ProductProduct struct {
	Base
	Name                string                   `orm:"unique"`        //产品属性名称
	IsProductVariant    bool                     `orm:"default(true)"` //是变形产品
	ProductTags         []*ProductTag            `orm:"reverse(many)"` //产品标签
	Categ               *ProductCategory         `orm:"rel(fk)"`       //产品类别
	Active              bool                     `orm:"default(true)"` //有效
	Barcode             string                   //条码,如ean13
	DefaultCode         string                   `orm:"unique"`        //产品编码
	ProductTemplate     *ProductTemplate         `orm:"rel(fk)"`       //产品款式
	AttributeValues     []*ProductAttributeValue `orm:"reverse(many)"` //产品属性
	FirstSaleUom        *ProductUom              `orm:"rel(fk)"`       //第一销售单位
	SecondSaleUom       *ProductUom              `orm:"rel(fk)"`       //第二销售单位
	FirstPurchaseUom    *ProductUom              `orm:"rel(fk)"`       //第一采购单位
	SecondPurchaseUom   *ProductUom              `orm:"rel(fk)"`       //第二采购单位
	ProductPackagings   []*ProductPackaging      `orm:"reverse(many)"` //打包方式
	PackagingDependTemp bool                     `orm:"default(true)"` //根据款式打包
	PurchaseDependTemp  bool                     `orm:"default(true)"` //根据款式采购，ture一个供应商可以供应所有的款式

}

func init() {
	orm.RegisterModel(new(ProductProduct))
}

// AddProductProduct insert a new ProductProduct into database and returns
// last inserted Id on success.
func AddProductProduct(obj *ProductProduct) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductProductById retrieves ProductProduct by Id. Returns error if
// Id doesn't exist
func GetProductProductById(id int64) (obj *ProductProduct, err error) {
	o := orm.NewOrm()
	obj = &ProductProduct{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetProductProductByName retrieves ProductProduct by Name. Returns error if
// Name doesn't exist
func GetProductProductByName(name string) (obj *ProductProduct, err error) {
	o := orm.NewOrm()
	obj = &ProductProduct{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductProduct retrieves all ProductProduct matches certain condition. Returns empty list if
// no records exist
func GetAllProductProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductProduct, error) {
	var (
		objArrs   []ProductProduct
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductProduct))
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

// UpdateProductProduct updates ProductProduct by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductProductById(m *ProductProduct) (err error) {
	o := orm.NewOrm()
	v := ProductProduct{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductProduct deletes ProductProduct by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductProduct(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductProduct{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductProduct{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
