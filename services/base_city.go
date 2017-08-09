package services

import (
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCity 创建记录
func ServiceCreateAddressCity(obj *md.AddressCity) (id int64, err error) {
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
func ServiceUpdateAddressCity(obj *md.AddressCity) (id int64, err error) {
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
func ServiceGetAddressCity(userID int64, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var arrs []md.AddressCity
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllAddressCity(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			objInfo["Province"] = obj.Province.Name
			results = append(results, objInfo)
		}
	}
	return
}
