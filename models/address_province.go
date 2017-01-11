package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 省份
type AddressProvince struct {
	Base
	Name    string          `xml:"ProvinceName,attr"` //省份名称
	Country *AddressCountry `orm:"rel(fk)"`           //国家
	Citys   []*AddressCity  `orm:"reverse(many)"`     //城市

}

func init() {
	orm.RegisterModel(new(AddressProvince))
}

// AddAddressProvince insert a new AddressProvince into database and returns
// last inserted Id on success.
func AddAddressProvince(obj *AddressProvince) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetAddressProvinceById retrieves AddressProvince by Id. Returns error if
// Id doesn't exist
func GetAddressProvinceById(id int64) (obj *AddressProvince, err error) {
	o := orm.NewOrm()
	obj = &AddressProvince{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAddressProvinceByName retrieves AddressProvince by Name. Returns error if
// Name doesn't exist
func GetAddressProvinceByName(name string) (obj *AddressProvince, err error) {
	o := orm.NewOrm()
	obj = &AddressProvince{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllAddressProvince retrieves all AddressProvince matches certain condition. Returns empty list if
// no records exist
func GetAllAddressProvince(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []AddressProvince, error) {
	var (
		objArrs   []AddressProvince
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(AddressProvince))
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

// UpdateAddressProvince updates AddressProvince by Id and returns error if
// the record to be updated doesn't exist
func UpdateAddressProvinceById(m *AddressProvince) (err error) {
	o := orm.NewOrm()
	v := AddressProvince{Base: Base{Id: m.Id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAddressProvince deletes AddressProvince by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAddressProvince(id int64) (err error) {
	o := orm.NewOrm()
	v := AddressProvince{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AddressProvince{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
