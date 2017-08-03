package models

import (
	"errors"
	"goCMS/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// BaseGroup  权限组
type BaseGroup struct {
	ID            int64          `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID  int64          `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID  int64          `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate    time.Time      `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate    time.Time      `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name          string         `orm:"unique;size(50)" json:"name" form:"Name"`      //权限组名称
	ModelAccesses []*ModelAccess `orm:"reverse(many)"`                                //模块(表)
	Childs        []*BaseGroup   `orm:"reverse(many)" json:"-" form:"-"`              //继承权限
	Parent        *BaseGroup     `orm:"rel(fk);null"`                                 //
	ParentLeft    int64          `orm:"unique"`                                       //左边界
	ParentRight   int64          `orm:"unique"`                                       //右边界
	Category      string         `orm:""`                                             //分类
	Description   string         ``                                                   //说明
}

func init() {
	orm.RegisterModel(new(BaseGroup))
}

// AddBaseGroup insert a new BaseGroup into database and returns last inserted Id on success.
func AddBaseGroup(m *BaseGroup, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddBaseGroup insert  list of  BaseGroup into database and returns  number of  success.
func BatchAddBaseGroup(groups []*BaseGroup, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&BaseGroup{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, group := range groups {
			if _, err = i.Insert(group); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateBaseGroup update BaseGroup into database and returns id on success
func UpdateBaseGroup(m *BaseGroup, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetBaseGroupByID retrieves BaseGroup by ID. Returns error if ID doesn't exist
func GetBaseGroupByID(id int64, ormObj orm.Ormer) (obj *BaseGroup, err error) {
	obj = &BaseGroup{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetBaseGroupByName retrieves BaseGroup by ID. Returns error if ID doesn't exist
func GetBaseGroupByName(name string, ormObj orm.Ormer) (*BaseGroup, error) {
	var obj BaseGroup
	var err error
	qs := ormObj.QueryTable(&obj)
	err = qs.Filter("Name", name).One(&obj)
	return &obj, err
}

// GetAllBaseGroup retrieves all BaseGroup matches certain condition. Returns empty list if no records exist
func GetAllBaseGroup(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []BaseGroup, error) {
	var (
		objArrs   []BaseGroup
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 200
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(BaseGroup))
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
