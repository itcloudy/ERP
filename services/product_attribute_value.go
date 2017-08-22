package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductAttributeValue 创建记录
func ServiceCreateProductAttributeValue(user *md.User, obj *md.ProductAttributeValue) (id int64, err error) {

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
	obj.CreateUserID = user.ID
	id, err = md.AddProductAttributeValue(obj, o)

	return
}

// ServiceUpdateProductAttributeValue 更新记录
func ServiceUpdateProductAttributeValue(user *md.User, obj *md.ProductAttributeValue) (id int64, err error) {

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
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductAttributeValue(obj, o)

	return
}

//ServiceGetProductAttributeValue 获得城市列表
func ServiceGetProductAttributeValue(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var access utils.AccessResult
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
