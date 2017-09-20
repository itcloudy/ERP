package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateSaleOrder 创建记录
func ServiceCreateSaleOrder(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrder"); err == nil {
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
	var obj md.SaleOrder
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddSaleOrder(&obj, o)

	return
}

// ServiceDeleteSaleOrder 删除记录
func ServiceDeleteSaleOrder(user *md.User, id int64) (num int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrder"); err == nil {
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
	var obj md.SaleOrder
	obj.ID = id
	num, err = md.DeleteSaleOrderByID(id, o)
	return
}

// ServiceUpdateSaleOrder 更新记录
func ServiceUpdateSaleOrder(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrder"); err == nil {
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
	var obj md.SaleOrder
	var objPtr *md.SaleOrder
	if objPtr, err = md.GetSaleOrderByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateSaleOrder(&obj, o)

	return
}

//ServiceGetSaleOrder 获得订单列表
func ServiceGetSaleOrder(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrder"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.SaleOrder
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllSaleOrder(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			partnerInfo := make(map[string]interface{})
			partnerInfo["ID"] = obj.Partner.ID
			partnerInfo["Name"] = obj.Partner.Name
			objInfo["Partner"] = partnerInfo
			salemanInfo := make(map[string]interface{})
			salemanInfo["ID"] = obj.SalesMan.ID
			salemanInfo["Name"] = obj.SalesMan.Name
			objInfo["SalesMan"] = salemanInfo
			companyInfo := make(map[string]interface{})
			companyInfo["ID"] = obj.Company.ID
			companyInfo["Name"] = obj.Company.Name
			objInfo["Company"] = companyInfo
			countryInfo := make(map[string]interface{})
			countryInfo["ID"] = obj.Country.ID
			countryInfo["Name"] = obj.Country.Name
			objInfo["Country"] = countryInfo
			provinceInfo := make(map[string]interface{})
			provinceInfo["ID"] = obj.Province.ID
			provinceInfo["Name"] = obj.Province.Name
			objInfo["Province"] = provinceInfo
			cityInfo := make(map[string]interface{})
			cityInfo["ID"] = obj.City.ID
			cityInfo["Name"] = obj.City.Name
			objInfo["City"] = cityInfo
			districtInfo := make(map[string]interface{})
			districtInfo["ID"] = obj.District.ID
			districtInfo["Name"] = obj.District.Name
			objInfo["District"] = districtInfo
			objInfo["Street"] = obj.Street
			objInfo["State"] = obj.State
			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetSaleOrderByID get SaleOrder by id
func ServiceGetSaleOrderByID(user *md.User, id int64) (access utils.AccessResult, cityInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "SaleOrder"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var city *md.SaleOrder
	if city, err = md.GetSaleOrderByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = city.Name
		objInfo["ID"] = city.ID
		cityInfo = objInfo
	}
	return
}
