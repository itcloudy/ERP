package models

import (
	"errors"
	"golangERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Partner 合作伙伴，包括客户和供应商，后期会为每个合作伙伴自动创建一个登录帐号
type Partner struct {
	ID           int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUserID int64            `orm:"column(create_user_id);null"`          //创建者
	UpdateUserID int64            `orm:"column(update_user_id);null"`          //最后更新者
	CreateDate   time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate   time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	Name         string           `orm:"unique" json:"Name"`                   //合作伙伴名称
	IsCompany    bool             `orm:"default(true)" json:"IsCompany"`       //是公司
	IsSupplier   bool             `orm:"default(false)" json:"IsSupplier"`     //是供应商
	IsCustomer   bool             `orm:"default(true)" json:"IsCustomer"`      //是客户
	Active       bool             `orm:"default(true)" json:"Active"`          //有效
	Country      *AddressCountry  `orm:"rel(fk);null"`                         //国家
	Province     *AddressProvince `orm:"rel(fk);null"`                         //省份
	City         *AddressCity     `orm:"rel(fk);null"`                         //城市
	District     *AddressDistrict `orm:"rel(fk);null"`                         //区县
	Street       string           `orm:"default(\"\")" json:"Street"`          //街道
	Parent       *Partner         `orm:"rel(fk);null"`                         //母公司
	Childs       []*Partner       `orm:"reverse(many)"`                        //下级
	Mobile       string           `orm:"default(\"\")" json:"Mobile"`          //电话号码
	Tel          string           `orm:"default(\"\")" json:"Tel"`             //座机
	Email        string           `orm:"default(\"\")" json:"Email"`           //邮箱
	Qq           string           `orm:"default(\"\")" json:"Qq"`              //QQ
	WeChat       string           `orm:"default(\"\")" json:"WeChat"`          //微信
	Comment      string           `orm:"type(text)" json:"Comment"`            //备注

}

func init() {
	orm.RegisterModel(new(Partner))
}

// AddPartner insert a new Partner into database and returns last inserted Id on success.
func AddPartner(m *Partner, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// BatchAddPartner insert  list of  Partner into database and returns  number of  success.
func BatchAddPartner(cities []*Partner, ormObj orm.Ormer) (num int64, err error) {
	qs := ormObj.QueryTable(&Partner{})
	if i, err := qs.PrepareInsert(); err == nil {
		defer i.Close()
		for _, partner := range cities {
			if _, err = i.Insert(partner); err == nil {
				num = num + 1
			}
		}
	}
	return
}

// UpdatePartner update Partner into database and returns id on success
func UpdatePartner(m *Partner, ormObj orm.Ormer) (id int64, err error) {
	if _, err = ormObj.Update(m); err == nil {
		id = m.ID
	}
	return
}

// DeletePartnerByID delete  ProductAttributeValue by ID
func DeletePartnerByID(id int64, ormObj orm.Ormer) (num int64, err error) {
	obj := &Partner{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// GetPartnerByID retrieves Partner by ID. Returns error if ID doesn't exist
func GetPartnerByID(id int64, ormObj orm.Ormer) (obj *Partner, err error) {
	obj = &Partner{ID: id}
	err = ormObj.Read(obj)
	ormObj.Read(obj.Province)
	ormObj.Read(obj.Province.Country)
	return obj, err
}

// GetAllPartner retrieves all Partner matches certain condition. Returns empty list if no records exist
func GetAllPartner(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []Partner, error) {
	var (
		objArrs   []Partner
		err       error
		paginator utils.Paginator
		num       int64
	)
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
