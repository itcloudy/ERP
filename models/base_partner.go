package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Partner 合作伙伴，包括客户和供应商，后期会为每个合作伙伴自动创建一个登录帐号
type Partner struct {
	ID         int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string           `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name       string           //合作伙伴名称
	IsCompany  bool             `orm:"default(true)"`              //是公司
	IsSupplier bool             `orm:"default(false)"`             //是供应商
	IsCustomer bool             `orm:"default(true)"`              //是客户
	Active     bool             `orm:"default(true)"`              //有效
	Country    *AddressCountry  `orm:"rel(fk);null"`               //国家
	Province   *AddressProvince `orm:"rel(fk);null"`               //身份
	City       *AddressCity     `orm:"rel(fk);null"`               //城市
	District   *AddressDistrict `orm:"rel(fk);null"`               //区县
	Street     string           `orm:"default(\"\")"`              //街道
	Parent     *Partner         `orm:"rel(fk);null"`               //母公司
	Childs     []*Partner       `orm:"reverse(many)"`              //下级
	Mobile     string           `orm:"default(\"\")"`              //电话号码
	Tel        string           `orm:"default(\"\")"`              //座机
	Email      string           `orm:"default(\"\")"`              //邮箱
	Qq         string           `orm:"default(\"\")" xml:"qq"`     //QQ
	WeChat     string           `orm:"default(\"\")" xml:"wechat"` //微信
	Comment    string           `orm:"type(text)"`                 //备注
}

func init() {
	orm.RegisterModel(new(Partner))
}

// AddPartner insert a new Partner into database and returns
// last inserted ID on success.
func AddPartner(obj *Partner) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(obj)
	return id, err
}

// GetPartnerByID retrieves Partner by ID. Returns error if
// ID doesn't exist
func GetPartnerByID(id int64) (obj *Partner, err error) {
	o := orm.NewOrm()
	obj = &Partner{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllPartner retrieves all Partner matches certain condition. Returns empty list if
// no records exist
func GetAllPartner(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []Partner, error) {
	var (
		objArrs   []Partner
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Partner))
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

// UpdatePartnerByID updates Partner by ID and returns error if
// the record to be updated doesn't exist
func UpdatePartnerByID(m *Partner) (err error) {
	o := orm.NewOrm()
	v := Partner{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetPartnerByName retrieves Partner by Name. Returns error if
// Name doesn't exist
func GetPartnerByName(name string) (obj *Partner, err error) {
	o := orm.NewOrm()
	obj = &Partner{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// DeletePartner deletes Partner by ID and returns error if
// the record to be deleted doesn't exist
func DeletePartner(id int64) (err error) {
	o := orm.NewOrm()
	v := Partner{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Partner{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
