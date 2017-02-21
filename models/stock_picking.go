package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//  StockPicking 调拨单
type StockPicking struct {
	ID           int64             `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser   *User             `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser   *User             `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate   time.Time         `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate   time.Time         `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name         string            `orm:"unique" json:"Name"`                   //单据名称
	Origin       string            `json:"Origin"`                              //源单据
	Note         string            `orm:"type(text)" json:"Note"`               //备注
	MoveType     string            `orm:"default(one)" json:"MoveType"`     //移动类型:one partial
	State        string            `orm:"default(draft)" json:"-"`          //状态 draft confirm process done cancel
	Company      *Company          `orm:"rel(fk)"`                              //公司
	LocationDest *StockLocation    `orm:"rel(fk)"`                              //目标库位
	LocationSrc  *StockLocation    `orm:"rel(fk)"`                              //源库位
	Partner      *Partner          `orm:"rel(fk)"`                              //合作伙伴
	Priority     int64             `orm:"default(1)" json:"Priority"`           //优先级,值越大优先级越高
	PickingType  *StockPickingType `orm:"rel(fk)"`                              //分拣类型决定分拣视图

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	CompanyID    int64    `orm:"-" json:"Company"`
}

func init() {
	orm.RegisterModel(new(StockPicking))
}

// AddStockPicking insert a new StockPicking into database and returns
// last inserted ID on success.
func AddStockPicking(obj *StockPicking, addUser *User) (id int64, err error) {
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

	// 获得款式产品编码

	id, err = o.Insert(obj)
	if err != nil {
		return 0, err
	} else {
		errCommit := o.Commit()
		if errCommit != nil {
			return 0, errCommit
		}
	}
	return id, err
}

// GetStockPickingByID retrieves StockPicking by ID. Returns error if
// ID doesn't exist
func GetStockPickingByID(id int64) (obj *StockPicking, err error) {
	o := orm.NewOrm()
	obj = &StockPicking{ID: id}
	if err = o.Read(obj); err == nil {

		return obj, nil
	}
	return nil, err
}

// GetStockPickingByName retrieves StockPicking by Name. Returns error if
// Name doesn't exist
func GetStockPickingByName(name string) (*StockPicking, error) {
	o := orm.NewOrm()
	var obj StockPicking
	cond := orm.NewCondition()
	cond = cond.And("Name", name)
	qs := o.QueryTable(&obj)
	qs = qs.SetCond(cond)
	err := qs.One(&obj)
	return &obj, err
}

// GetAllStockPicking retrieves all StockPicking matches certain condition. Returns empty list if
// no records exist
func GetAllStockPicking(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []StockPicking, error) {
	var (
		objArrs   []StockPicking
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(StockPicking))
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

// UpdateStockPicking updates StockPicking by ID and returns error if
// the record to be updated doesn't exist
func UpdateStockPicking(obj *StockPicking, updateUser *User) error {
	o := orm.NewOrm()
	v := User{ID: obj.ID}
	errBegin := o.Begin()
	var err error
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return errBegin
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Update(obj, append(obj.ActionFields, "UpdateUser", "UpdateDate")...); err != nil {
			utils.LogOut("error", "update user fields failed:"+err.Error())
		}
	}
	if err != nil {
		return err
	}
	errCommit := o.Commit()
	if errCommit != nil {
		return errCommit
	}
	return nil
}

// DeleteStockPicking deletes StockPicking by ID and returns error if
// the record to be deleted doesn't exist
func DeleteStockPicking(id int64) (err error) {
	o := orm.NewOrm()
	v := StockPicking{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StockPicking{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
