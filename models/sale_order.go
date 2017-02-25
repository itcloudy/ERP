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

// SaleOrder 产品分类
type SaleOrder struct {
	ID             int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser     *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser     *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate     time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate     time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name           string           `orm:"unique" json:"name"`                   //订单号
	Partner        *Partner         `orm:"rel(fk)"`                              //客户
	SalesMan       *User            `orm:"rel(fk)"`                              //业务员
	Company        *Company         `orm:"rel(fk)"`                              //公司
	Country        *AddressCountry  `orm:"rel(fk);null" json:"-"`                //国家
	Province       *AddressProvince `orm:"rel(fk);null" json:"-"`                //省份
	City           *AddressCity     `orm:"rel(fk);null" json:"-"`                //城市
	District       *AddressDistrict `orm:"rel(fk);null" json:"-"`                //区县
	Street         string           `orm:"default()" json:"Street"`              //街道
	OrderLine      []*SaleOrderLine `orm:"reverse(many)"`                        //订单明细
	State          *SaleOrderState  `orm:"rel(fk)"`                              //订单状态
	StockWarehouse *StockWarehouse  `orm:"rel(fk)"`                              //仓库
	PickingPolicy  string           `orm:"default(one)" json:"PickingPolicy"`    //发货策略one/mult

	FormAction       string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields     []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	CompanyID        int64    `orm:"-" json:"Company"`
	PartnerID        int64    `orm:"-" json:"Partner"`
	SalesManID       int64    `orm:"-" json:"SalesMan"`
	CountryID        int64    `orm:"-" json:"Country"`
	ProvinceID       int64    `orm:"-" json:"Province"`
	CityID           int64    `orm:"-" json:"City"`
	DistrictID       int64    `orm:"-" json:"District"`
	StockWarehouseID int64    `orm:"-" json:"StockWarehouse"`
}

func init() {
	orm.RegisterModel(new(SaleOrder))
}

// AddSaleOrder insert a new SaleOrder into database and returns
// last inserted ID on success.
func AddSaleOrder(obj *SaleOrder, addUser *User) (id int64, err error) {
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
	if obj.SalesManID > 0 {
		obj.SalesMan, _ = GetUserByID(obj.SalesManID)

	}

	if obj.CompanyID > 0 {
		obj.Company, _ = GetCompanyByID(obj.CompanyID)
	}
	if obj.CountryID > 0 {
		obj.Country, _ = GetAddressCountryByID(obj.CountryID)
	}
	if obj.ProvinceID > 0 {
		obj.Province, _ = GetAddressProvinceByID(obj.ProvinceID)
	}
	if obj.CityID > 0 {
		obj.City, _ = GetAddressCityByID(obj.CityID)
	}
	if obj.DistrictID > 0 {
		obj.District, _ = GetAddressDistrictByID(obj.DistrictID)
	}
	if obj.StockWarehouseID > 0 {
		obj.StockWarehouse, _ = GetStockWarehouseByID(obj.StockWarehouseID)
	}
	if obj.StockWarehouse != nil && obj.Company != nil {
		obj.State, _ = GetSaleOrderStateByCompanyStock(obj.Company, obj.StockWarehouse, nil)
	}
	if obj.PartnerID > 0 {
		obj.Partner, _ = GetPartnerByID(obj.PartnerID)
	}
	// 获得款式产品编码
	obj.Name, _ = GetNextSequece(reflect.Indirect(reflect.ValueOf(obj)).Type().Name(), obj.Company.ID)
	id, err = o.Insert(obj)
	if err == nil {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetSaleOrderByID retrieves SaleOrder by ID. Returns error if
// ID doesn't exist
func GetSaleOrderByID(id int64) (obj *SaleOrder, err error) {
	o := orm.NewOrm()
	obj = &SaleOrder{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Partner != nil {
			o.Read(obj.Partner)
		}
		if obj.SalesMan != nil {
			o.Read(obj.SalesMan)
		}
		if obj.Company != nil {
			o.Read(obj.Company)
		}
		if obj.Country != nil {
			o.Read(obj.Country)
		}
		if obj.Province != nil {
			o.Read(obj.Province)
		}
		if obj.City != nil {
			o.Read(obj.City)
		}
		if obj.StockWarehouse != nil {
			o.Read(obj.StockWarehouse)
		}
		if obj.State != nil {
			o.Read(obj.State)
		}
		return obj, nil
	}
	return nil, err
}

// GetAllSaleOrder retrieves all SaleOrder matches certain condition. Returns empty list if
// no records exist
func GetAllSaleOrder(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []SaleOrder, error) {
	var (
		objArrs   []SaleOrder
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(SaleOrder))
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

// UpdateSaleOrderByID updates SaleOrder by ID and returns error if
// the record to be updated doesn't exist
func UpdateSaleOrderByID(m *SaleOrder) (err error) {
	o := orm.NewOrm()
	v := SaleOrder{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetSaleOrderByName retrieves SaleOrder by Name. Returns error if
// Name doesn't exist
func GetSaleOrderByName(name string) (obj *SaleOrder, err error) {
	o := orm.NewOrm()
	obj = &SaleOrder{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeleteSaleOrder deletes SaleOrder by ID and returns error if
// the record to be deleted doesn't exist
func DeleteSaleOrder(id int64) (err error) {
	o := orm.NewOrm()
	v := SaleOrder{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SaleOrder{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
