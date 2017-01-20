package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//TemplateFile 模版文件
type TemplateFile struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string    `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name       string    //模版名称
	Desc       string    //描述
}

func init() {
	orm.RegisterModel(new(TemplateFile))
}

// AddTemplateFile insert a new TemplateFile into database and returns
// last inserted ID on success.
func AddTemplateFile(obj *TemplateFile) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetTemplateFileByID retrieves TemplateFile by ID. Returns error if
// ID doesn't exist
func GetTemplateFileByID(id int64) (obj *TemplateFile, err error) {
	o := orm.NewOrm()
	obj = &TemplateFile{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetTemplateFileByName retrieves TemplateFile by Name. Returns error if
// Name doesn't exist
func GetTemplateFileByName(name string) (obj *TemplateFile, err error) {
	o := orm.NewOrm()
	obj = &TemplateFile{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

//GetLastTemplateFileByUserID GetLastTemplateFileByUserID
func GetLastTemplateFileByUserID(userID int64) (TemplateFile, error) {
	o := orm.NewOrm()
	var (
		record TemplateFile
		err    error
	)

	o.Using("default")
	err = o.QueryTable(&record).Filter("User", userID).RelatedSel().OrderBy("-id").Limit(1).One(&record)
	return record, err
}

// GetAllTemplateFile retrieves all TemplateFile matches certain condition. Returns empty list if
// no records exist
func GetAllTemplateFile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []TemplateFile, error) {
	var (
		objArrs   []TemplateFile
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(TemplateFile))
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

// UpdateTemplateFileByID updates TemplateFile by ID and returns error if
// the record to be updated doesn't exist
func UpdateTemplateFileByID(m *TemplateFile) error {
	o := orm.NewOrm()
	v := TemplateFile{ID: m.ID}
	var err error
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return err
}

// DeleteTemplateFile deletes TemplateFile by ID and returns error if
// the record to be deleted doesn't exist
func DeleteTemplateFile(id int64) (err error) {
	o := orm.NewOrm()
	v := TemplateFile{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TemplateFile{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
