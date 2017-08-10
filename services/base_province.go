package services

import (
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressProvince 创建记录
func ServiceCreateAddressProvince(obj *md.AddressProvince) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	id, err = md.AddAddressProvince(obj, o)

	return
}

// ServiceUpdateAddressProvince 更新记录
func ServiceUpdateAddressProvince(obj *md.AddressProvince) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if err != nil {
		return
	}
	id, err = md.UpdateAddressProvince(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}

//ServiceGetAddressProvince 获得省份列表
func ServiceGetAddressProvince(userID int64, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var arrs []md.AddressProvince

	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllAddressProvince(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			countryInfo := make(map[string]interface{})
			countryInfo["ID"] = obj.Country.ID
			countryInfo["Name"] = obj.Country.Name
			objInfo["Country"] = countryInfo
			objInfo["ID"] = obj.ID
			results = append(results, objInfo)
		}
	}
	return
}
