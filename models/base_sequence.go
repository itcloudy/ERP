package models

import (
	"errors"
	"fmt"
	"pms/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Sequence  表序列号管理，用于销售订单号，采购订单号等
type Sequence struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`              //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                     //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                     //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"`      //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`          //最后更新时间
	Name       string    `orm:"unique" xml:"Name"`                         //表名称组名称
	Prefix     string    `xml:"Prefix" json:"Prefix"`                      //前缀
	Current    int64     `xml:"Current" json:"Current"`                    //当前序号
	Padding    int64     `orm:"default(8)" xml:"Padding" json:"Padding"`   //序列位数,默认位数8
	StructName string    `orm:"unique" xml:"StructName" json:"StructName"` //表结构体名称
	// form表单字段
	FormAction string `orm:"-" json:"FormAction"` //非数据库字段，用于表示记录的增加，修改
}

func init() {
	orm.RegisterModel(new(Sequence))
}

// AddSequence insert a new Sequence into database and returns
// last inserted ID on success.
func AddSequence(obj *Sequence) (id int64, errs []error) {
	o := orm.NewOrm()
	var err error
	err = o.Begin()
	if err != nil {
		errs = append(errs, err)
	}
	id, err = o.Insert(obj)
	if err != nil {
		errs = append(errs, err)
		err = o.Rollback()
		if err != nil {
			errs = append(errs, err)
		}
	} else {
		err = o.Commit()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return id, errs
}

// GetSequenceByID retrieves Sequence by ID. Returns error if
// ID doesn't exist
func GetSequenceByID(id int64) (obj *Sequence, err error) {
	o := orm.NewOrm()
	obj = &Sequence{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetSequenceByName retrieves Sequence by Name. Returns error if
// Name doesn't exist
func GetSequenceByName(name string) (obj *Sequence, err error) {
	o := orm.NewOrm()
	obj = &Sequence{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllSequence retrieves all Sequence matches certain condition. Returns empty list if
// no records exist
func GetAllSequence(query map[string]interface{}, exclude map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []Sequence, error) {
	var (
		objArrs   []Sequence
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sequence))
	qs = qs.RelatedSel()
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
	for i, _ := range objArrs {
		o.LoadRelated(&objArrs[i], "ValueIDs")
	}

	return paginator, objArrs, err
}

// UpdateSequenceByID updates Sequence by ID and returns error if
// the record to be updated doesn't exist
func UpdateSequenceByID(m *Sequence) (err error) {
	o := orm.NewOrm()
	v := Sequence{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSequence deletes Sequence by ID and returns error if
// the record to be deleted doesn't exist
func DeleteSequence(id int64) (err error) {
	o := orm.NewOrm()
	v := Sequence{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sequence{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
