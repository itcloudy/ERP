package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressDistrict 创建记录
func ServiceCreateAddressDistrict(user *md.User, requestBody []byte) (id int64, err error) {
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
	var obj md.AddressDistrict
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddAddressDistrict(&obj, o)

	return
}

// ServiceUpdateAddressDistrict 更新记录
func ServiceUpdateAddressDistrict(user *md.User, requestBody []byte, id int64) (err error) {
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
	var obj md.AddressDistrict
	var objPtr *md.AddressDistrict
	if objPtr, err = md.GetAddressDistrictByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)

	obj.UpdateUserID = user.ID
	id, err = md.UpdateAddressDistrict(&obj, o)
	return
}

//ServiceGetAddressDistrict 获得区县列表
func ServiceGetAddressDistrict(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {

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

// ServiceGetAddressDistrictByID get AddressDistrict by id
func ServiceGetAddressDistrictByID(user *md.User, id int64) (access utils.AccessResult, districtInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "AddressDistrict"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var district *md.AddressDistrict

	if district, err = md.GetAddressDistrictByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = district.Name
		objInfo["ID"] = district.ID
		cityInfo := make(map[string]interface{})
		cityInfo["ID"] = district.City.ID
		cityInfo["Name"] = district.City.Name
		objInfo["City"] = cityInfo
		provinceInfo := make(map[string]interface{})
		if district.City.Province != nil {
			provinceInfo["ID"] = district.City.Province.ID
			provinceInfo["Name"] = district.City.Province.Name
			objInfo["Province"] = provinceInfo
			countryInfo := make(map[string]interface{})
			if district.City.Province.Country != nil {
				countryInfo["ID"] = district.City.Province.Country.ID
				countryInfo["Name"] = district.City.Province.Country.Name
				objInfo["Country"] = countryInfo
			}
		}

		districtInfo = objInfo
	}
	return
}
