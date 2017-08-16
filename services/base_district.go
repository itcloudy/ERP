package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressDistrict 创建记录
func ServiceCreateAddressDistrict(user *md.User, obj *md.AddressDistrict) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressDistrict"); err == nil {
		if !access.Create {
			err = errors.New("has no create permission")
			return
		}
	} else {
		return
	}
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
	id, err = md.AddAddressDistrict(obj, o)

	return
}

// ServiceUpdateAddressDistrict 更新记录
func ServiceUpdateAddressDistrict(user *md.User, obj *md.AddressDistrict) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressDistrict"); err == nil {
		if !access.Update {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
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
	id, err = md.UpdateAddressDistrict(obj, o)
	return
}

//ServiceGetAddressDistrict 获得区县列表
func ServiceGetAddressDistrict(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressDistrict"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.AddressDistrict
	provinceMap := make(map[int64]md.AddressProvince)
	countryMap := make(map[int64]md.AddressCountry)
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllAddressDistrict(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			cityInfo := make(map[string]interface{})
			cityInfo["ID"] = obj.City.ID
			cityInfo["Name"] = obj.City.Name
			objInfo["City"] = cityInfo
			provinceInfo := make(map[string]interface{})
			countryInfo := make(map[string]interface{})
			provinceID := obj.City.Province.ID
			if province, ok := provinceMap[provinceID]; ok {
				provinceInfo["Name"] = province.Name
				provinceInfo["ID"] = province.ID
				countryID := province.Country.ID
				if country, ok := countryMap[countryID]; ok {
					countryInfo["Name"] = country.Name
					countryInfo["ID"] = country.ID
				}
			} else {
				if province, err := md.GetAddressProvinceByID(provinceID, o); err == nil {
					provinceMap[provinceID] = *province
					provinceInfo["Name"] = province.Name
					provinceInfo["ID"] = province.ID
					countryID := province.Country.ID
					if country, err := md.GetAddressCountryByID(countryID, o); err == nil {
						countryInfo["Name"] = country.Name
						countryInfo["ID"] = country.ID
						countryMap[countryID] = *country
					}
				}
			}
			objInfo["Country"] = countryInfo
			objInfo["Province"] = provinceInfo
			results = append(results, objInfo)
		}
	}
	return
}
