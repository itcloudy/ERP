package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// BaseMenu 城市
type BaseMenu struct {
	ID           int64        `orm:"column(id);pk;auto"`     //主键
	CreateUserID int64        `orm:"column(create_user_id);null"`             //创建者
	UpdateUserID int64        `orm:"column(update_user_id);null"`             //最后更新者
	CreateDate   time.Time    `orm:"auto_now_add;type(datetime)"`             //创建时间
	UpdateDate   time.Time    `orm:"auto_now;type(datetime)"`                 //最后更新时间
	Name         string       `orm:"size(50)"`                 //菜单名称
	Index        string       `orm:"unique"`                                           //唯一标识
	Parent       *BaseMenu    `orm:"rel(fk);null"`              //上级菜单
	Childs       []*BaseMenu  `orm:"reverse(many)"`                      //子菜单
	ParentLeft   int64        `orm:"unique"`                                           //菜单左
	ParentRight  int64        `orm:"unique"`                                           //菜单右
	Sequence     int64        `orm:"default(1)"`                                       //序列号，决定同级菜单显示先后顺序
	Icon         string       `orm:""`                                                 //菜单图标样式
	Groups       []*BaseGroup `orm:"rel(m2m);rel_through(golangERP/models.GroupMenu)"` //权限组
	Path         string       `orm:""`                                                 //菜单路径
	Component    string       `orm:""`                                                 //组件名称
	Meta         string       `orm:""`                                                 //额外参数
	Step         int          `orm:"-"`                                                //用于后台组合菜单使用
}

func init() {
	orm.RegisterModel(new(BaseMenu))
}

// AddBaseMenu insert a new BaseMenu into database and returns last inserted Id on success.
func AddBaseMenu(m *BaseMenu, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddBaseMenu insert  list of  BaseMenu into database and returns  number of  success.
func BatchAddBaseMenu(menus []*BaseMenu, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&BaseMenu{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, menu := range menus {
			if _, err = i.Insert(menu); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdateBaseMenu update BaseMenu into database and returns id on success
func UpdateBaseMenu(m *BaseMenu, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// GetBaseMenuByID retrieves BaseMenu by ID. Returns error if ID doesn't exist
func GetBaseMenuByID(id int64, ormObj orm.Ormer) (obj *BaseMenu, err error) {
	obj = &BaseMenu{ID: id}
	err = ormObj.Read(obj)
	ormObj.LoadRelated(obj, "Childs")
	return obj, err
}

// GetAllBaseMenu retrieves all BaseMenu matches certain condition. Returns empty list if no records exist
func GetAllBaseMenu(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) ([]BaseMenu, error) {
	var (
		objArrs []BaseMenu
		err     error
	)
	if limit == 0 {
		limit = 200
	}

	qs := o.QueryTable(new(BaseMenu))
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	if cnt, err := qs.Count(); err == nil {
		if cnt > 0 {
			_, err = qs.All(&objArrs, fields...)
		}
	}
	for i, _ := range objArrs {
		o.LoadRelated(&objArrs[i], "Groups")
		o.LoadRelated(&objArrs[i], "Childs")
	}
	return objArrs, err
}
