package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"goERP/utils"

	"github.com/astaxie/beego/orm"
)

//ProductTemplate 产品款式
type ProductTemplate struct {
	ID                  int64                   `orm:"column(id);pk;auto" json:"id" form:"-"`  //主键
	CreateUser          *User                   `orm:"rel(fk);null" json:"-"`                  //创建者
	UpdateUser          *User                   `orm:"rel(fk);null" json:"-"`                  //最后更新者
	CreateDate          time.Time               `orm:"auto_now_add;type(datetime)" json:"-"`   //创建时间
	UpdateDate          time.Time               `orm:"auto_now;type(datetime)" json:"-"`       //最后更新时间
	Name                string                  `orm:"unique;index" json:"Name"`               //产品款式名称
	Company             *Company                `orm:"rel(fk);null"`                           //公司
	Sequence            int32                   `json:"Sequence"`                              //序列号
	Description         string                  `orm:"type(text);null"`                        //描述
	DescriptionSale     string                  `orm:"type(text);null"`                        //销售描述
	DescriptionPurchase string                  `orm:"type(text);null"`                        //采购描述
	Rental              bool                    `orm:"default(false)"`                         //代售品
	Category            *ProductCategory        `orm:"rel(fk)"`                                //产品类别
	Price               float64                 `json:"Price"`                                 //模版产品价格
	StandardPrice       float64                 `json:"StandardPrice"`                         //成本价格
	SaleOk              bool                    `orm:"default(true)" json:"SaleOk"`            //可销售
	Active              bool                    `orm:"default(true)" json:"Active"`            //有效
	IsProductVariant    bool                    `orm:"default(true)"`                          //是变形产品
	FirstSaleUom        *ProductUom             `orm:"rel(fk)"`                                //第一销售单位
	SecondSaleUom       *ProductUom             `orm:"rel(fk);null"`                           //第二销售单位
	FirstPurchaseUom    *ProductUom             `orm:"rel(fk)"`                                //第一采购单位
	SecondPurchaseUom   *ProductUom             `orm:"rel(fk);null"`                           //第二采购单位
	AttributeLines      []*ProductAttributeLine `orm:"reverse(many)"`                          //属性明细
	ProductVariants     []*ProductProduct       `orm:"reverse(many)"`                          //产品规格明细
	TemplatePackagings  []*ProductPackaging     `orm:"reverse(many)"`                          //打包方式
	VariantCount        int32                   `json:"VariantCount"`                          //产品规格数量
	Barcode             string                  `json:"Barcode"`                               //条码,如ean13
	DefaultCode         string                  `json:"DefaultCode"`                           //产品编码
	BigImages           []*ProductImage         `orm:"reverse(many)"`                          //产品款式图片
	MidImages           []*ProductImage         `orm:"reverse(many)"`                          //产品款式图片
	SmallImages         []*ProductImage         `orm:"reverse(many)"`                          //产品款式图片
	ProductType         string                  `orm:"default()"`                     //产品类型
	ProductMethod       string                  `orm:"default()" json:"ProductMethod"` //产品规格创建方式
	PackagingDependTemp bool                    `orm:"default(true)"`                          //根据款式打包
	PurchaseDependTemp  bool                    `orm:"default(true)"`                          //根据款式采购，ture一个供应商可以供应所有的款式

	FormAction            string                 `orm:"-" json:"FormAction"`        //非数据库字段，用于表示记录的增加，修改
	ActionFields          []string               `orm:"-" json:"ActionFields"`      //需要操作的字段,用于update时
	CategoryID            int64                  `orm:"-" json:"Category"`          //产品类别
	FirstSaleUomID        int64                  `orm:"-" json:"FirstSaleUom"`      //第一销售单位form
	SecondSaleUomID       int64                  `orm:"-" json:"SecondSaleUom"`     //第二销售单位form
	FirstPurchaseUomID    int64                  `orm:"-" json:"FirstPurchaseUom"`  //第一采购单位form
	SecondPurchaseUomID   int64                  `orm:"-" json:"SecondPurchaseUom"` //第二采购单位form
	ProductCounterID      int64                  `orm:"-" json:"ProductCounter"`    //产品柜台
	ProductAttributeLines []ProductAttributeLine `orm:"-" json:"ProductAttributes"`
}

func init() {
	orm.RegisterModel(new(ProductTemplate))
}

func GetVariantCount(obj *ProductTemplate) (count int64) {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	condMap := make(map[string]map[string]interface{})
	query["ProductTemplate.id"] = obj.ID
	if paginaotor, _, err := GetAllProductProduct(query, exclude, condMap, fields, sortby, order, 0, 5); err == nil {
		return paginaotor.TotalCount
	} else {
		return 0
	}
}

// AddProductTemplate insert a new ProductTemplate into database and returns
// last inserted ID on success.
func AddProductTemplate(obj *ProductTemplate, addUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.CreateUser = addUser
	obj.UpdateUser = addUser
	errBegin := o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return 0, errBegin
	}
	if obj.CategoryID > 0 {
		obj.Category, _ = GetProductCategoryByID(obj.CategoryID)
	}
	if obj.FirstSaleUomID > 0 {
		obj.FirstSaleUom, _ = GetProductUomByID(obj.FirstSaleUomID)
	}
	if obj.SecondSaleUomID > 0 {
		obj.SecondSaleUom, _ = GetProductUomByID(obj.SecondSaleUomID)
	}
	if obj.FirstPurchaseUomID > 0 {
		obj.FirstPurchaseUom, _ = GetProductUomByID(obj.FirstPurchaseUomID)
	}
	if obj.SecondPurchaseUomID > 0 {
		obj.SecondPurchaseUom, _ = GetProductUomByID(obj.SecondPurchaseUomID)
	}
	// 获得款式产品编码
	obj.DefaultCode, _ = GetNextSequece(reflect.Indirect(reflect.ValueOf(obj)).Type().Name(), addUser.Company.ID)
	if id, err = o.Insert(obj); err == nil {
		obj.ID = id
		if len(obj.ProductAttributeLines) > 0 {
			for _, item := range obj.ProductAttributeLines {
				if attribute, err := GetProductAttributeByID(item.AttributeID); err == nil {
					productAttributeLine := new(ProductAttributeLine)
					productAttributeLine.Attribute = attribute
					productAttributeLine.CreateUser = addUser
					productAttributeLine.UpdateUser = addUser
					productAttributeLine.ProductTemplate = obj
					if lineID, err := o.Insert(productAttributeLine); err == nil {
						productAttributeLine.ID = lineID
						m2m := o.QueryM2M(productAttributeLine, "AttributeValues")
						attributeValueIDArr := item.AttributeValueIds
						for _, attrValueID := range attributeValueIDArr {
							if valueObj, err := GetProductAttributeValueByID(attrValueID); err == nil {
								productAttributeLine.AttributeValues = append(productAttributeLine.AttributeValues, valueObj)
								UpdateProductAttributeTemplatesCount(productAttributeLine.Attribute, addUser)
							} else {
								return 0, err
							}
						}
						// for _, attrVal := range productAttributeLine.AttributeValues {
						// 	if !m2m.Exist(attrVal) {
						// 		m2m.Add(attrVal)
						// 	}
						// }
						// 创建直接添加
						m2m.Add(productAttributeLine.AttributeValues)
						obj.AttributeLines = append(obj.AttributeLines, productAttributeLine)
					} else {
						return 0, err
					}

				} else {
					return 0, err
				}
			}
		}
	}
	if err != nil {
		return 0, err
	} else {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetProductTemplateByID retrieves ProductTemplate by ID. Returns error if
// ID doesn't exist
func GetProductTemplateByID(id int64) (obj *ProductTemplate, err error) {
	o := orm.NewOrm()
	obj = &ProductTemplate{ID: id}
	if err = o.Read(obj); err == nil {
		if _, err := o.LoadRelated(obj, "AttributeLines"); err == nil {
			for index, _ := range obj.AttributeLines {
				o.LoadRelated(obj.AttributeLines[index], "Attribute")
				o.LoadRelated(obj.AttributeLines[index], "AttributeValues")
			}
		}
		if obj.Category != nil {
			o.Read(obj.Category)
		}
		if obj.FirstSaleUom != nil {
			o.Read(obj.FirstSaleUom)
		}
		if obj.FirstPurchaseUom != nil {
			o.Read(obj.FirstPurchaseUom)
		}
		if obj.SecondSaleUom != nil {
			o.Read(obj.SecondSaleUom)
		}
		if obj.SecondPurchaseUom != nil {
			o.Read(obj.SecondPurchaseUom)
		}
		return obj, nil
	}
	return nil, err
}

// GetProductTemplateByName retrieves ProductTemplate by Name. Returns error if
// Name doesn't exist
func GetProductTemplateByName(name string) (*ProductTemplate, error) {
	o := orm.NewOrm()
	var obj ProductTemplate
	cond := orm.NewCondition()
	cond = cond.And("Name", name)
	qs := o.QueryTable(&obj)
	qs = qs.SetCond(cond)
	err := qs.One(&obj)
	return &obj, err
}

// GetAllProductTemplate retrieves all ProductTemplate matches certain condition. Returns empty list if
// no records exist
func GetAllProductTemplate(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []ProductTemplate, error) {
	var (
		objArrs   []ProductTemplate
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}

	o := orm.NewOrm()
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

// UpdateProductTemplate updates ProductTemplate by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductTemplate(obj *ProductTemplate, updateUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.UpdateUser = updateUser
	var num int64
	if num, err = o.Update(obj); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return obj.ID, err
}

// DeleteProductTemplate deletes ProductTemplate by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductTemplate(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductTemplate{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductTemplate{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
