package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductUomCateg 创建记录
func ServiceCreateProductUomCateg(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductUomCateg"); err == nil {
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
	var obj md.ProductUomCateg
	json.Unmarshal([]byte(requestBody), &obj)

	obj.CreateUserID = user.ID
	id, err = md.AddProductUomCateg(&obj, o)

	return
}

// ServiceUpdateProductUomCateg 更新记录
func ServiceUpdateProductUomCateg(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductUomCateg"); err == nil {
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
	var obj md.ProductUomCateg
	var objPtr *md.ProductUomCateg
	if objPtr, err = md.GetProductUomCategByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductUomCateg(&obj, o)

	return
}

//ServiceGetProductUomCateg 获得单位类别列表
func ServiceGetProductUomCateg(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductUomCateg"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductUomCateg
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductUomCateg(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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

// ServiceGetProductUomCategByID get ProductUomCateg by id
func ServiceGetProductUomCategByID(user *md.User, id int64) (access utils.AccessResult, attrInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductUomCateg"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var attr *md.ProductUomCateg
	if attr, err = md.GetProductUomCategByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = attr.Name
		objInfo["ID"] = attr.ID
		attrInfo = objInfo
	}
	return
}
