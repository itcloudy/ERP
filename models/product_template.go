package models

import (
	"errors"
	"strings"
	"time"

	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ProductTemplate 产品款式
type ProductTemplate struct {
	ID                  int64                   `orm:"column(id);pk;auto"`          //主键
	CreateUserID        int64                   `orm:"column(create_user_id);null"` //创建者
	UpdateUserID        int64                   `orm:"column(update_user_id);null"` //最后更新者
	CreateDate          time.Time               `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate          time.Time               `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name                string                  `orm:""`                            //款式名称
	Description         string                  `orm:"type(text);null"`             //描述
	DescriptionSale     string                  `orm:"type(text);null"`             //销售描述
	DescriptionPurchase string                  `orm:"type(text);null"`             //采购描述
	Rental              bool                    `orm:"default(false)"`              //代售品
	Category            *ProductCategory        `orm:"rel(fk)"`                     //产品类别
	Price               float64                 ``                                  //款式价格
	StandardPrice       float64                 ``                                  //成本价格
	StandardWeight      float64                 ``                                  //标准重量
	SaleOk              bool                    `orm:"default(true)"`               //可销售
	Active              bool                    `orm:"default(true)"`               //有效
	IsProductVariant    bool                    `orm:"default(true)"`               //是规格产品
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
	ProductType         string                  `orm:"default(stock)"`              //产品类型 stock consume service
	ProductMethod       string                  `orm:"default(hand)"`               //产品规格创建方式 auto hand
	// TemplatePackagings  []*ProductPackaging     `orm:"reverse(many)"`               //打包方式
}

func init() {
	orm.RegisterModel(new(ProductTemplate))
}

// AddProductTemplate insert a new ProductTemplate into database and returns last inserted Id on success.
func AddProductTemplate(m *ProductTemplate, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddProductTemplate insert  list of  ProductTemplate into database and returns  number of  success.
func BatchAddProductTemplate(cities []*ProductTemplate, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&ProductTemplate{})
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

// UpdateProductTemplate update ProductTemplate into database and returns id on success
func UpdateProductTemplate(m *ProductTemplate, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetProductTemplateByID retrieves ProductTemplate by ID. Returns error if ID doesn't exist
func GetProductTemplateByID(id int64, ormObj orm.Ormer) (obj *ProductTemplate, err error) {
	obj = &ProductTemplate{ID: id}
	err = ormObj.Read(obj)
	ormObj.Read(obj.Category)
	ormObj.Read(obj.FirstPurchaseUom)
	ormObj.Read(obj.SecondPurchaseUom)
	ormObj.Read(obj.FirstSaleUom)
	ormObj.Read(obj.SecondSaleUom)
	ormObj.LoadRelated(obj, "AttributeLines")
	lenA := len(obj.AttributeLines)
	for i := 0; i < lenA; i++ {
		ormObj.Read(obj.AttributeLines[i].Attribute)
		ormObj.LoadRelated(obj.AttributeLines[i], "AttributeValues")
	}
	return obj, err
}

// DeleteProductTemplateByID delete  Company by ID
func DeleteProductTemplateByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &ProductTemplate{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// GetAllProductTemplate retrieves all ProductTemplate matches certain condition. Returns empty list if no records exist
func GetAllProductTemplate(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductTemplate, error) {
	var (
		objArrs   []ProductTemplate
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(ProductTemplate))
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
		// rewrite dot-notation to Object__Template
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Template
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
