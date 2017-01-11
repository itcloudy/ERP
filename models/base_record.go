package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Record struct {
	Base
	User      *User     `orm:"rel(fk)"`
	Logout    time.Time `orm:"type(datetime);null"` //登出时间
	UserAgent string    `orm:"null"`                //用户代理
	Ip        string    //上次登录IP
}

// 用户登录记录
func init() {
	orm.RegisterModel(new(Record))
}

// AddRecord insert a new Record into database and returns
// last inserted Id on success.
func AddRecord(obj *Record) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetRecordById retrieves Record by Id. Returns error if
// Id doesn't exist
func GetRecordById(id int64) (obj *Record, err error) {
	o := orm.NewOrm()
	obj = &Record{Base: Base{Id: id}}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}
func GetLastRecordByUserID(userId int64) (Record, error) {
	o := orm.NewOrm()
	var (
		record Record
		err    error
	)

	o.Using("default")
	err = o.QueryTable(&record).Filter("User", userId).RelatedSel().OrderBy("-id").Limit(1).One(&record)
	return record, err
}

// GetAllRecord retrieves all Record matches certain condition. Returns empty list if
// no records exist
func GetAllRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []Record, error) {
	var (
		objArrs   []Record
		paginator utils.Paginator
		num       int64
		err       error
	)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Record))
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

// UpdateRecord updates Record by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecordById(m *Record) error {
	o := orm.NewOrm()
	v := Record{Base: Base{Id: m.Id}}
	var err error
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return err
}

// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecord(id int64) (err error) {
	o := orm.NewOrm()
	v := Record{Base: Base{Id: id}}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Record{Base: Base{Id: id}}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
