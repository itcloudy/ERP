package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCountry 创建记录
func ServiceCreateAddressCountry(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressCountry"); err == nil {
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

	var obj md.AddressCountry

	json.Unmarshal([]byte(requestBody), &obj)

	obj.CreateUserID = user.ID
	id, err = md.AddAddressCountry(&obj, o)

	return
}

// ServiceUpdateAddressCountry 更新记录
func ServiceUpdateAddressCountry(user *md.User, requestBody []byte, id int64) (err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "AddressCountry"); err == nil {
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
	var obj md.AddressCountry
	var objPtr *md.AddressCountry
	if objPtr, err = md.GetAddressCountryByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateAddressCountry(&obj, o)

	return
}

//ServiceGetAddressCountry 获得国家列表
func ServiceGetAddressCountry(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "AddressCountry"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.AddressCountry
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllAddressCountry(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetAddressCountryByID get AddressCountry by id
func ServiceGetAddressCountryByID(user *md.User, id int64) (access utils.AccessResult, cityInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "AddressCountry"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var country *md.AddressCountry
	if country, err = md.GetAddressCountryByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = country.Name
		objInfo["ID"] = country.ID
		cityInfo = objInfo
	}
	return
}
