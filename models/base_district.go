package models

import (
	"errors"
	"golangERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// AddressDistrict 区县
type AddressDistrict struct {
	ID           int64        `orm:"column(id);pk;auto"`          //主键
	CreateUserID int64        `orm:"column(create_user_id);null"` //创建者
	UpdateUserID int64        `orm:"column(update_user_id);null"` //最后更新者
	CreateDate   time.Time    `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateDate   time.Time    `orm:"auto_now;type(datetime)"`     //最后更新时间
	Name         string       ` `                                 //区县名称
	City         *AddressCity `orm:"rel(fk)"`                     //城市
}

func init() {
	orm.RegisterModel(new(AddressDistrict))
}

// AddAddressDistrict insert a new AddressDistrict into database and returns last inserted Id on success.
func AddAddressDistrict(m *AddressDistrict, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// GetAddressDistrictByID retrieves AddressDistrict by ID. Returns error if ID doesn't exist
func GetAddressDistrictByID(id int64, ormObj orm.Ormer) (obj *AddressDistrict, err error) {
	obj = &AddressDistrict{ID: id}
	err = ormObj.Read(obj)
	ormObj.Read(obj.City)
	ormObj.Read(obj.City.Province)
	ormObj.Read(obj.City.Province.Country)
	return obj, err
}

// DeleteAddressDistrictByID delete  Company by ID
func DeleteAddressDistrictByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &AddressDistrict{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// BatchAddAddressDistrict insert  list of  Country into database and returns  number of  success.
func BatchAddAddressDistrict(districtes []*AddressDistrict, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&AddressDistrict{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, district := range districtes {
			if _, err = i.Insert(district); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateAddressDistrict update AddressDistrict into database and returns id on success
func UpdateAddressDistrict(m *AddressDistrict, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetAllAddressDistrict retrieves all AddressDistrict matches certain condition. Returns empty list if no records exist
func GetAllAddressDistrict(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []AddressDistrict, error) {
	var (
		objArrs   []AddressDistrict
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(AddressDistrict))
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
