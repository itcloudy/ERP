package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 城市
type AddressCity struct {
	Base
	Name      string             `orm:"size(50)" json:"name"`           //城市名称
	Province  *AddressProvince   `orm:"rel(fk)" json:"province"`        //国家
	Districts []*AddressDistrict `orm:"reverse(many)" json:"districts"` //城市
}

func init() {
	orm.RegisterModel(new(AddressCity))
}

// AddAddressCity insert a new AddressCity into database and returns
// last inserted Id on success.
func AddAddressCity(obj *AddressCity) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetAddressCityById retrieves AddressCity by Id. Returns error if
// Id doesn't exist
func GetAddressCityById(id int64) (obj *AddressCity, err error) {
	o := orm.NewOrm()
	obj = &AddressCity{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAddressCityByName retrieves AddressCity by Name. Returns error if
// Name doesn't exist
func GetAddressCityByName(name string) (obj *AddressCity, err error) {
	o := orm.NewOrm()
	obj = &AddressCity{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllAddressCity retrieves all AddressCity matches certain condition. Returns empty list if
// no records exist
func GetAllAddressCity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []AddressCity, error) {
	var (
		objArrs   []AddressCity
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(AddressCity))
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

// UpdateAddressCity updates AddressCity by Id and returns error if
// the record to be updated doesn't exist
func UpdateAddressCityById(m *AddressCity) error {
	o := orm.NewOrm()
	v := AddressCity{Base: Base{Id: m.Id}}
	var err error
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return err
}

// DeleteAddressCity deletes AddressCity by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAddressCity(id int64) (err error) {
	o := orm.NewOrm()
	v := AddressCity{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AddressCity{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
