package models

import (
	"errors"
	"strings"
	"time"

	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// AddressCountry 国家
type AddressCountry struct {
	ID           int64              `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64              `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64              `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time          `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time          `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string             `orm:"unique;size(50)" xml:"name"`  //国家名称
	Provinces    []*AddressProvince `orm:"reverse(many)"`               //省份
}

func init() {
	orm.RegisterModel(new(AddressCountry))

}

// AddAddressCountry insert a new AddressCountry into database and returns last inserted Id on success.
func AddAddressCountry(m *AddressCountry, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// UpdateAddressCountry update AddressCountry into database and returns id on success
func UpdateAddressCountry(m *AddressCountry, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetAddressCountryByID retrieves AddressCountry by ID. Returns error if ID doesn't exist
func GetAddressCountryByID(id int64, ormObj orm.Ormer) (obj *AddressCountry, err error) {
	obj = &AddressCountry{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// DeleteAddressCountryByID delete  Company by ID
func DeleteAddressCountryByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &AddressCountry{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// BatchAddAddressCountry insert  list of  Country into database and returns  number of  success.
func BatchAddAddressCountry(countries []*AddressCountry, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&AddressCountry{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, country := range countries {
			if _, err = i.Insert(country); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// GetAllAddressCountry retrieves all AddressCountry matches certain condition. Returns empty list if no records exist
func GetAllAddressCountry(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []AddressCountry, error) {
	var (
		objArrs   []AddressCountry
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(AddressCountry))
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
