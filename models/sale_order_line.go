package models

import (
	"errors"
	"golangERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// SaleOrderLine 销售订单明细
type SaleOrderLine struct {
	ID            int64           `orm:"column(id);pk;auto"`              //主键
	CreateUserID  int64           `orm:"column(create_user_id);null"`     //创建者
	UpdateUserID  int64           `orm:"column(update_user_id);null"`     //最后更新者
	CreateDate    time.Time       `orm:"auto_now_add;type(datetime)"`     //创建时间
	UpdateDate    time.Time       `orm:"auto_now;type(datetime)"`         //最后更新时间
	Name          string          `orm:"default()" json:"Name"`           //订单明细号
	Company       *Company        `orm:"rel(fk)"`                         //公司
	SaleOrder     *SaleOrder      `orm:"rel(fk)"`                         //销售订单
	Partner       *Partner        `orm:"rel(fk)"`                         //客户
	Product       *ProductProduct `orm:"rel(fk)"`                         //产品
	ProductName   string          `json:"ProductName"`                    //产品名称
	ProductCode   string          `json:"ProductCode"`                    //产品编码
	FirstSaleUom  *ProductUom     `orm:"rel(fk)"`                         //第一销售单位
	SecondSaleUom *ProductUom     `orm:"rel(fk);null"`                    //第二销售单位
	FirstSaleQty  float32         `orm:"default(1)" json:"FirstSaleQty"`  //第一销售单位
	SecondSaleQty float32         `orm:"default(0)" json:"SecondSaleQty"` //第二销售单位
	State         string          `orm:"default(draft)"`                  //订单明细状态:draft/confirm/process/done/cancel
	PriceUnit     float32         `orm:"default(0)" json:"PriceUnit"`     //单价
	Total         float32         `orm:"default(0)" json:"Total"`         //小计
}

func init() {
	orm.RegisterModel(new(SaleOrderLine))
}

// AddSaleOrderLine insert a new SaleOrderLine into database and returns last inserted Id on success.
func AddSaleOrderLine(m *SaleOrderLine, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddSaleOrderLine insert  list of  SaleOrderLine into database and returns  number of  success.
func BatchAddSaleOrderLine(cities []*SaleOrderLine, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&SaleOrderLine{})
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

// UpdateSaleOrderLine update SaleOrderLine into database and returns id on success
func UpdateSaleOrderLine(m *SaleOrderLine, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// DeleteSaleOrderLineByID delete  ProductAttributeValue by ID
func DeleteSaleOrderLineByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &SaleOrderLine{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// GetSaleOrderLineByID retrieves SaleOrderLine by ID. Returns error if ID doesn't exist
func GetSaleOrderLineByID(id int64, ormObj orm.Ormer) (obj *SaleOrderLine, err error) {
	obj = &SaleOrderLine{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetAllSaleOrderLine retrieves all SaleOrderLine matches certain condition. Returns empty list if no records exist
func GetAllSaleOrderLine(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []SaleOrderLine, error) {
	var (
		objArrs   []SaleOrderLine
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(SaleOrderLine))
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
