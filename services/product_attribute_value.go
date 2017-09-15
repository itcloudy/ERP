package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductAttributeValue 创建记录
func ServiceCreateProductAttributeValue(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeValue"); err == nil {
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
	var obj md.ProductAttributeValue
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddProductAttributeValue(&obj, o)

	return
}

// ServiceDeleteProductAttributeValue 删除记录
func ServiceDeleteProductAttributeValue(user *md.User, id int64) (num int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeValue"); err == nil {
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
	var obj md.ProductAttributeValue

	obj.ID = id
	num, err = md.DeleteProductAttributeValueByID(id, o)
	return
}

// ServiceUpdateProductAttributeValue 更新记录
func ServiceUpdateProductAttributeValue(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeValue"); err == nil {
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
	var obj md.ProductAttributeValue
	var objPtr *md.ProductAttributeValue
	if objPtr, err = md.GetProductAttributeValueByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductAttributeValue(&obj, o)

	return
}

//ServiceGetProductAttributeValue 获得城市列表
func ServiceGetProductAttributeValue(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeValue"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductAttributeValue
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductAttributeValue(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			attrInfo := make(map[string]interface{})
			attrInfo["ID"] = obj.Attribute.ID
			attrInfo["Name"] = obj.Attribute.Name
			objInfo["Attribute"] = attrInfo
			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetProductAttributeValueByID get ProductAttribute by id
func ServiceGetProductAttributeValueByID(user *md.User, id int64) (access utils.AccessResult, valueInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeValue"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var obj *md.ProductAttributeValue
	if obj, err = md.GetProductAttributeValueByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = obj.Name
		objInfo["ID"] = obj.ID
		attributeInfo := make(map[string]interface{})
		attributeInfo["ID"] = obj.Attribute.ID
		attributeInfo["Name"] = obj.Attribute.Name
		objInfo["Attribute"] = attributeInfo
		valueInfo = objInfo
	}
	return
}
