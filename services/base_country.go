package services

import (
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCountry 创建记录
func ServiceCreateAddressCountry(obj *md.AddressCountry) (id int64, err error) {
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
	id, err = md.AddAddressCountry(obj, o)

	return
}

// ServiceUpdateAddressCountry 更新记录
func ServiceUpdateAddressCountry(obj *md.AddressCountry) (id int64, err error) {
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
	id, err = md.UpdateAddressCountry(obj, o)

	return
}

//ServiceGetAddressCountry 获得国家列表
func ServiceGetAddressCountry(userID int64, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
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
