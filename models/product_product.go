package models

import (
	"errors"
	"strings"
	"time"

	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ProductProduct 产品规格
type ProductProduct struct {
	ID                    int64                    `orm:"column(id);pk;auto" json:"id"` //主键
	CreateUserID          int64                    `orm:"column(create_user_id);null"`  //创建者
	UpdateUserID          int64                    `orm:"column(update_user_id);null"`  //最后更新者
	CreateDate            time.Time                `orm:"auto_now_add;type(datetime)"`  //创建时间
	UpdateDate            time.Time                `orm:"auto_now;type(datetime)"`      //最后更新时间
	Name                  string                   `orm:"index"`                        //产品属性名称
	Company               *Company                 `orm:"rel(fk);null"`                 //公司
	Category              *ProductCategory         `orm:"rel(fk)"`                      //产品类别
	IsProductVariant      bool                     `orm:"default(true)"`                //是多规格产品
	ProductTags           []*ProductTag            `orm:"reverse(many)"`                //产品标签
	SaleOk                bool                     `orm:"default(true)" json:"SaleOk"`  //可销售
	Active                bool                     `orm:"default(true)"`                //有效
	Barcode               string                   `orm:"null" json:"Barcode"`          //条码,如ean13
	StandardPrice         float64                  `json:"StandardPrice"`               //成本价格
	DefaultCode           string                   `orm:"unique"`                       //产品编码
	ProductTemplate       *ProductTemplate         `orm:"rel(fk)"`                      //产品款式
	AttributeValues       []*ProductAttributeValue `orm:"reverse(many)"`                //产品属性值
	ProductType           string                   `orm:"default(stock)"`               //产品类型
	AttributeValuesString string                   `orm:"index;default()"`              //产品属性值ID编码，用于修改和增加时对应的产品是否已经存在
	FirstSaleUom          *ProductUom              `orm:"rel(fk)"`                      //第一销售单位
	SecondSaleUom         *ProductUom              `orm:"rel(fk);null"`                 //第二销售单位
	FirstPurchaseUom      *ProductUom              `orm:"rel(fk)"`                      //第一采购单位
	SecondPurchaseUom     *ProductUom              `orm:"rel(fk);null"`                 //第二采购单位
	PackagingDependTemp   bool                     `orm:"default(true)"`                //根据款式打包
	BigImages             []*ProductImage          `orm:"reverse(many)"`                //产品款式图片
	MidImages             []*ProductImage          `orm:"reverse(many)"`                //产品款式图片
	SmallImages           []*ProductImage          `orm:"reverse(many)"`                //产品款式图片
	PurchaseDependTemp    bool                     `orm:"default(true)"`                //根据款式采购，ture一个供应商可以供应所有的款式
	// ProductPackagings     []*ProductPackaging      `orm:"reverse(many)"`                        //打包方式

}

func init() {
	orm.RegisterModel(new(ProductProduct))
}

// AddProductProduct insert a new ProductProduct into database and returns last inserted Id on success.
func AddProductProduct(m *ProductProduct, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddProductProduct insert  list of  ProductProduct into database and returns  number of  success.
func BatchAddProductProduct(cities []*ProductProduct, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&ProductProduct{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, city := range cities {
			if _, err = i.Insert(city); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateProductProduct update ProductProduct into database and returns id on success
func UpdateProductProduct(m *ProductProduct, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// DeleteProductProductByID delete  ProductAttributeValue by ID
func DeleteProductProductByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &ProductProduct{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// GetProductProductByID retrieves ProductProduct by ID. Returns error if ID doesn't exist
func GetProductProductByID(id int64, ormObj orm.Ormer) (obj *ProductProduct, err error) {
	obj = &ProductProduct{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetAllProductProduct retrieves all ProductProduct matches certain condition. Returns empty list if no records exist
func GetAllProductProduct(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductProduct, error) {
	var (
		objArrs   []ProductProduct
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(ProductProduct))
	qs = qs.RelatedSel()

	//cond k=v cond必须放到Filter和Exclude前面
	cond := orm.NewCondition()
	if _, ok := condMap["and"]; ok {
		andMap := condMap["and"]
		for k, v := range andMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.And(k, v)
		}
	}
	if _, ok := condMap["or"]; ok {
		orMap := condMap["or"]
		for k, v := range orMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.Or(k, v)
		}
	}
	qs = qs.SetCond(cond)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Exclude(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[i] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
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
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[0] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
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
		if cnt > 0 {
			paginator = utils.GenPaginator(limit, offset, cnt)
			if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
				paginator.CurrentPageSize = num
			}
		}
	}
	return paginator, objArrs, err
}
