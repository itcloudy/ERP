package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductAttribute 创建记录
func ServiceCreateProductAttribute(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttribute"); err == nil {
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
	var obj md.ProductAttribute
	json.Unmarshal([]byte(requestBody), &obj)
	obj.CreateUserID = user.ID
	id, err = md.AddProductAttribute(&obj, o)

	return
}

// ServiceUpdateProductAttribute 更新记录
func ServiceUpdateProductAttribute(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttribute"); err == nil {
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
	var obj md.ProductAttribute
	var objPtr *md.ProductAttribute
	if objPtr, err = md.GetProductAttributeByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)

	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductAttribute(&obj, o)

	return
}

//ServiceGetProductAttribute 获得城市列表
func ServiceGetProductAttribute(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttribute"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductAttribute
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductAttribute(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["Code"] = obj.Code
			objInfo["ID"] = obj.ID
			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetProductAttributeByID get ProductAttribute by id
func ServiceGetProductAttributeByID(user *md.User, id int64) (access utils.AccessResult, attrInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductAttribute"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var attr *md.ProductAttribute
	if attr, err = md.GetProductAttributeByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = attr.Name
		objInfo["ID"] = attr.ID
		objInfo["Code"] = attr.Code
		objInfo["CreatVariant"] = attr.CreatVariant
		lenValues := len(attr.ValueIds)

		if lenValues > 0 {
			var Values []interface{}
			for i := 0; i < lenValues; i++ {
				valueInfo := make(map[string]interface{})
				valueInfo["ID"] = attr.ValueIds[i].ID
				valueInfo["Name"] = attr.ValueIds[i].Name
				Values = append(Values, valueInfo)
			}
			objInfo["Values"] = Values
		}
		attrInfo = objInfo
	}
	return
}
