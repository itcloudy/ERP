package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCity 创建记录
func ServiceCreateAddressCity(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressCity"); err == nil {
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
	var obj md.AddressCity
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddAddressCity(&obj, o)

	return
}

// ServiceDeleteAddressCity 删除记录
func ServiceDeleteAddressCity(user *md.User, id int64) (num int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressCity"); err == nil {
		if !access.Unlink {
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
	var obj md.AddressCity
	obj.ID = id
	num, err = md.DeleteAddressCityByID(id, o)
	return
}

// ServiceUpdateAddressCity 更新记录
func ServiceUpdateAddressCity(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressCity"); err == nil {
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
	var obj md.AddressCity
	var objPtr *md.AddressCity
	if objPtr, err = md.GetAddressCityByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateAddressCity(&obj, o)

	return
}

//ServiceGetAddressCity 获得城市列表
func ServiceGetAddressCity(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "AddressCity"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
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

// ServiceGetAddressCityByID get AddressCity by id
func ServiceGetAddressCityByID(user *md.User, id int64) (access utils.AccessResult, cityInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "AddressCity"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var city *md.AddressCity
	if city, err = md.GetAddressCityByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = city.Name
		objInfo["ID"] = city.ID
		provinceInfo := make(map[string]interface{})
		provinceInfo["ID"] = city.Province.ID
		provinceInfo["Name"] = city.Province.Name
		objInfo["Province"] = provinceInfo

		countryInfo := make(map[string]interface{})
		if city.Province.Country != nil {
			countryInfo["ID"] = city.Province.Country.ID
			countryInfo["Name"] = city.Province.Country.Name
			objInfo["Country"] = countryInfo
		}
		cityInfo = objInfo
	}
	return
}
