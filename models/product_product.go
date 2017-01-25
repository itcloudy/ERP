package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

// ProductProduct 产品规格
type ProductProduct struct {
	ID               int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser       *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser       *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate       time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate       time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction       string           `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name             string           `orm:"unique"`                               //产品属性名称
	Company          *Company         `orm:"rel(fk);null"`                         //公司
	IsProductVariant bool             `orm:"default(true)"`                        //是多规格产品
	ProductTags      []*ProductTag    `orm:"reverse(many)"`                        //产品标签
	Categ            *ProductCategory `orm:"rel(fk)"`                              //产品类别
	Active           bool             `orm:"default(true)"`                        //有效
	Barcode          string           `json:"Barcode"`                             //条码,如ean13
	DefaultCode      string           `orm:"unique"`                               //产品编码
	ProductTemplate  *ProductTemplate `orm:"rel(fk)"`                              //产品款式
	// AttributeLine       *ProductAttributeLine    `orm:"rel(fk)"`                              //产品属性明细行，款式含有，规格删除
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
// last inserted ID on success.
func AddProductProduct(obj *ProductProduct) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductProductByID retrieves ProductProduct by ID. Returns error if
// ID doesn't exist
func GetProductProductByID(id int64) (obj *ProductProduct, err error) {
	o := orm.NewOrm()
	obj = &ProductProduct{ID: id}
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
	if limit == 0 {
		limit = 20
	}
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

// UpdateProductProductByID updates ProductProduct by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductProductByID(m *ProductProduct) (err error) {
	o := orm.NewOrm()
	v := ProductProduct{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductProduct deletes ProductProduct by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductProduct(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductProduct{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductProduct{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
