package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductProduct 创建记录
func ServiceCreateProductProduct(user *md.User, requestBody map[string]interface{}) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductProduct"); err == nil {
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
	var obj md.ProductProduct
	if Name, ok := requestBody["Name"]; ok {
		obj.Name = utils.ToString(Name)
	}
	obj.CreateUserID = user.ID
	id, err = md.AddProductProduct(&obj, o)

	return
}

// ServiceUpdateProductProduct 更新记录
func ServiceUpdateProductProduct(user *md.User, requestBody map[string]interface{}, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductProduct"); err == nil {
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
	var obj md.ProductProduct
	var objPtr *md.ProductProduct
	if objPtr, err = md.GetProductProductByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	if Name, ok := requestBody["Name"]; ok {
		obj.Name = utils.ToString(Name)
	}
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductProduct(&obj, o)

	return
}

//ServiceGetProductProduct 获得城市列表
func ServiceGetProductProduct(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductProduct"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductProduct
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductProduct(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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

// ServiceGetProductProductByID get ProductProduct by id
func ServiceGetProductProductByID(user *md.User, id int64) (access utils.AccessResult, attrInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductProduct"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var attr *md.ProductProduct
	if attr, err = md.GetProductProductByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = attr.Name
		objInfo["ID"] = attr.ID

		attrInfo = objInfo
	}
	return
}
