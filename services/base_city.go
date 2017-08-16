package services

import (
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCity 创建记录
func ServiceCreateAddressCity(user *md.User, obj *md.AddressCity) (id int64, err error) {
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
	id, err = md.AddAddressCity(obj, o)

	return
}

// ServiceUpdateAddressCity 更新记录
func ServiceUpdateAddressCity(user *md.User, obj *md.AddressCity) (id int64, err error) {
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
	id, err = md.UpdateAddressCity(obj, o)

	return
}

//ServiceGetAddressCity 获得城市列表
func ServiceGetAddressCity(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var arrs []md.AddressCity
	countryMap := make(map[int64]md.AddressCountry)
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllAddressCity(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			countryInfo := make(map[string]interface{})

			provinceInfo := make(map[string]interface{})
			provinceInfo["ID"] = obj.Province.ID
			provinceInfo["Name"] = obj.Province.Name
			objInfo["Province"] = provinceInfo
			countryID := obj.Province.Country.ID
			if country, ok := countryMap[countryID]; ok {
				countryInfo["Name"] = country.Name
				countryInfo["ID"] = country.ID
			} else {
				if country, err := md.GetAddressCountryByID(countryID, o); err == nil {
					countryMap[countryID] = *country
					countryInfo["Name"] = country.Name
					countryInfo["ID"] = country.ID
				}
			}
			objInfo["Country"] = countryInfo
			results = append(results, objInfo)
		}
	}
	return
}
