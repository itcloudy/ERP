package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateSaleOrderLine 创建记录
func ServiceCreateSaleOrderLine(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrderLine"); err == nil {
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
	var obj md.SaleOrderLine
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddSaleOrderLine(&obj, o)

	return
}

// ServiceDeleteSaleOrderLine 删除记录
func ServiceDeleteSaleOrderLine(user *md.User, id int64) (num int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrderLine"); err == nil {
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
	var obj md.SaleOrderLine
	obj.ID = id
	num, err = md.DeleteSaleOrderLineByID(id, o)
	return
}

// ServiceUpdateSaleOrderLine 更新记录
func ServiceUpdateSaleOrderLine(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrderLine"); err == nil {
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
	var obj md.SaleOrderLine
	var objPtr *md.SaleOrderLine
	if objPtr, err = md.GetSaleOrderLineByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateSaleOrderLine(&obj, o)

	return
}

//ServiceGetSaleOrderLine 获得城市列表
func ServiceGetSaleOrderLine(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "SaleOrderLine"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.SaleOrderLine
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllSaleOrderLine(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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

// ServiceGetSaleOrderLineByID get SaleOrderLine by id
func ServiceGetSaleOrderLineByID(user *md.User, id int64) (access utils.AccessResult, orderlineInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "SaleOrderLine"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var orderline *md.SaleOrderLine
	if orderline, err = md.GetSaleOrderLineByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = orderline.Name
		objInfo["ID"] = orderline.ID
		orderlineInfo = objInfo
	}
	return
}
