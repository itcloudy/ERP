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
	Name       string           `orm:"unique" json:"Name"`                   //合作伙伴名称
	IsCompany  bool             `orm:"default(true)" json:"IsCompany"`       //是公司
	IsSupplier bool             `orm:"default(false)" json:"IsSupplier"`     //是供应商
	IsCustomer bool             `orm:"default(true)" json:"IsCustomer"`      //是客户
	Active     bool             `orm:"default(true)" json:"Active"`          //有效
	Country    *AddressCountry  `orm:"rel(fk);null"`                         //国家
	Province   *AddressProvince `orm:"rel(fk);null"`                         //省份
	City       *AddressCity     `orm:"rel(fk);null"`                         //城市
	District   *AddressDistrict `orm:"rel(fk);null"`                         //区县
	Street     string           `orm:"default(\"\")" json:"Street"`          //街道
	Parent     *Partner         `orm:"rel(fk);null"`                         //母公司
	Childs     []*Partner       `orm:"reverse(many)"`                        //下级
	Mobile     string           `orm:"default(\"\")" json:"Mobile"`          //电话号码
	Tel        string           `orm:"default(\"\")" json:"Tel"`             //座机
	Email      string           `orm:"default(\"\")" json:"Email"`           //邮箱
	Qq         string           `orm:"default(\"\")" json:"Qq"`              //QQ
	WeChat     string           `orm:"default(\"\")" json:"WeChat"`          //微信
	Comment    string           `orm:"type(text)" json:"Comment"`            //备注

	FormAction string `orm:"-" json:"FormAction"` //非数据库字段，用于表示记录的增加，修改
	ParentID   int64  `orm:"-" json:"Parent"`     //母公司
	CountryID  int64  `orm:"-" json:"Country"`    //国家
	ProvinceID int64  `orm:"-" json:"Province"`   //省份
	CityID     int64  `orm:"-" json:"City"`       //城市
	DistrictID int64  `orm:"-" json:"District"`   //区县

}

func init() {
	orm.RegisterModel(new(Partner))
}
func (u *Partner) TableName() string {
	return "base_partner"
}

// AddPartner insert a new Partner into database and returns
// last inserted ID on success.
func AddPartner(obj *Partner, addUser *User) (id int64, err error) {
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
	if obj.ParentID > 0 {
		obj.Parent, _ = GetPartnerByID(obj.ParentID)
	}
	if obj.CountryID > 0 {
		obj.Country, _ = GetAddressCountryByID(obj.CountryID)
	}
	if obj.ProvinceID > 0 {
		obj.Province, _ = GetAddressProvinceByID(obj.ProvinceID)
	}
	if obj.CityID > 0 {
		obj.City, _ = GetAddressCityByID(obj.CityID)
	}
	if obj.DistrictID > 0 {
		obj.District, _ = GetAddressDistrictByID(obj.DistrictID)
	}
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

// GetPartnerByID retrieves Partner by ID. Returns error if
// ID doesn't exist
func GetPartnerByID(id int64) (obj *Partner, err error) {
	o := orm.NewOrm()
	obj = &Partner{ID: id}
	if err = o.Read(obj); err == nil {
		o.LoadRelated(obj, "Parent")
		o.LoadRelated(obj, "Country")
		o.LoadRelated(obj, "Province")
		o.LoadRelated(obj, "City")
		o.LoadRelated(obj, "District")
		return obj, nil
	}
	return nil, err
}

// GetPartnerByName retrieves Partner by Name. Returns error if
// Name doesn't exist
func GetPartnerByName(name string) (*Partner, error) {
	o := orm.NewOrm()
	var obj Partner
	cond := orm.NewCondition()
	cond = cond.And("Name", name)
	qs := o.QueryTable(&obj)
	qs = qs.SetCond(cond)
	err := qs.One(&obj)
	return &obj, err
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

// UpdatePartner updates Partner by ID and returns error if
// the record to be updated doesn't exist
func UpdatePartner(obj *Partner, updateUser *User) (id int64, err error) {
	o := orm.NewOrm()
	obj.UpdateUser = updateUser
	var num int64
	if num, err = o.Update(obj); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return obj.ID, err
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
