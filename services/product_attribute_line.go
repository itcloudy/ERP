package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"
	"reflect"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductAttributeLine 创建记录
func ServiceCreateProductAttributeLine(user *md.User, requestBody map[string]interface{}) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeLine"); err == nil {
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
	var obj md.ProductAttributeLine

	obj.CreateUserID = user.ID
	id, err = md.AddProductAttributeLine(&obj, o)

	return
}

// ServiceDeleteProductAttributeLine 删除记录
func ServiceDeleteProductAttributeLine(user *md.User, id int64) (num int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeLine"); err == nil {
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
	var obj md.ProductAttributeLine

	obj.ID = id
	num, err = md.DeleteProductAttributeLineByID(id, o)
	return
}

// ServiceUpdateProductAttributeLine 更新记录
func ServiceUpdateProductAttributeLine(user *md.User, requestBody map[string]interface{}, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeLine"); err == nil {
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
	var obj md.ProductAttributeLine
	var objPtr *md.ProductAttributeLine
	if objPtr, err = md.GetProductAttributeLineByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	var attribute md.ProductAttribute
	if Attribute, ok := requestBody["Attribute"]; ok {
		attributeT := reflect.TypeOf(Attribute)
		if attributeT.Kind() == reflect.Map {
			attributeMap := Attribute.(map[string]interface{})
			if attributeID, ok := attributeMap["ID"]; ok {
				attribute.ID, _ = utils.ToInt64(attributeID)
				obj.Attribute = &attribute
			}
		} else if attributeT.Kind() == reflect.String {
			attribute.ID, _ = utils.ToInt64(Attribute)
			obj.Attribute = &attribute
		}
	}

	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductAttributeLine(&obj, o)

	return
}

//ServiceGetProductAttributeLine 获得城市列表
func ServiceGetProductAttributeLine(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeLine"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductAttributeLine
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductAttributeLine(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["ID"] = obj.ID

			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetProductAttributeLineByID get ProductAttribute by id
func ServiceGetProductAttributeLineByID(user *md.User, id int64) (access utils.AccessResult, valueInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductAttributeLine"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var obj *md.ProductAttributeLine
	if obj, err = md.GetProductAttributeLineByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["ID"] = obj.ID
		valueInfo = objInfo
	}
	return
}
