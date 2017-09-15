package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressProvince 创建记录
func ServiceCreateAddressProvince(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressProvince"); err == nil {
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
	var obj md.AddressProvince
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddAddressProvince(&obj, o)
	return
}

// ServiceUpdateAddressProvince 更新记录
func ServiceUpdateAddressProvince(user *md.User, requestBody []byte, id int64) (err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressProvince"); err == nil {
		if !access.Update {
			err = errors.New("has no update permission ")
			return
		}
	} else {
		return
	}
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
	var obj md.AddressProvince
	var objPtr *md.AddressProvince
	if objPtr, err = md.GetAddressProvinceByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateAddressProvince(&obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}

//ServiceGetAddressProvince 获得省份列表
func ServiceGetAddressProvince(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "AddressProvince"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission ")
			return
		}
	} else {
		return
	}
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

// ServiceGetAddressProvinceByID get AddressProvince by id
func ServiceGetAddressProvinceByID(user *md.User, id int64) (access utils.AccessResult, provinceInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "AddressProvince"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var province *md.AddressProvince

	if province, err = md.GetAddressProvinceByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = province.Name
		objInfo["ID"] = province.ID
		countryInfo := make(map[string]interface{})
		countryInfo["ID"] = province.Country.ID
		countryInfo["Name"] = province.Country.Name
		objInfo["Country"] = countryInfo
		provinceInfo = objInfo
	}
	return
}
