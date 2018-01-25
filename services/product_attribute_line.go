package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"
	"reflect"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductAttributeLine 创建记录
func ServiceCreateProductAttributeLine(user *md.User, requestBody []byte) (id int64, err error) {

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
	json.Unmarshal([]byte(requestBody), &obj)

	obj.CreateUserID = user.ID
	id, err = md.AddProductAttributeLine(&obj, o)
	if err != nil {
		return
	}
	var objPtr *md.ProductAttributeLine
	objPtr, _ = md.GetProductAttributeLineByID(id, o)
	var requestBodyMap map[string]interface{}
	json.Unmarshal(requestBody, &requestBodyMap)
	if AttributeValues, ok := requestBodyMap["AttributeValues"]; ok {
		s := reflect.ValueOf(AttributeValues)
		if s.Kind() == reflect.Slice {
			m2m := o.QueryM2M(objPtr, "AttributeValues")
			m2m.Clear()
			for i := 0; i < s.Len(); i++ {
				valueID := s.Index(i).Interface()
				var valueObj md.ProductAttributeValue
				if valueObj.ID, _ = utils.ToInt64(valueID); valueObj.ID > 0 {
					m2m.Add(valueObj)
				}
			}

		}
	}

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
func ServiceUpdateProductAttributeLine(user *md.User, requestBody []byte, id int64) (err error) {

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
	objPtr, _ = md.GetProductAttributeLineByID(id, o)
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	var requestBodyMap map[string]interface{}
	json.Unmarshal(requestBody, &requestBodyMap)
	if AttributeValues, ok := requestBodyMap["AttributeValues"]; ok {
		s := reflect.ValueOf(AttributeValues)
		if s.Kind() == reflect.Slice {
			m2m := o.QueryM2M(objPtr, "AttributeValues")
			m2m.Clear()
			for i := 0; i < s.Len(); i++ {
				valueID := s.Index(i).Interface()
				var valueObj md.ProductAttributeValue
				if valueObj.ID, _ = utils.ToInt64(valueID); valueObj.ID > 0 {
					m2m.Add(valueObj)
				}
			}

		}
	}

	id, err = md.UpdateProductAttributeLine(&obj, o)
	return
}

//ServiceGetProductAttributeLine 获得属性明细列表
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
			attrInfo := make(map[string]interface{})
			attrInfo["ID"] = obj.Attribute.ID
			attrInfo["Name"] = obj.Attribute.Name
			objInfo["Attribute"] = attrInfo
			tempInfo := make(map[string]interface{})
			tempInfo["ID"] = obj.ProductTemplate.ID
			tempInfo["Name"] = obj.ProductTemplate.Name
			objInfo["ProductTemplate"] = tempInfo
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
		attrInfo := make(map[string]interface{})
		attrInfo["ID"] = obj.Attribute.ID
		attrInfo["Name"] = obj.Attribute.Name
		objInfo["Attribute"] = attrInfo
		tempInfo := make(map[string]interface{})
		tempInfo["ID"] = obj.ProductTemplate.ID
		tempInfo["Name"] = obj.ProductTemplate.Name
		objInfo["ProductTemplate"] = tempInfo
		var attributeValues []map[string]interface{}
		valuesLen := len(obj.AttributeValues)
		if valuesLen > 0 {
			for i := 0; i < valuesLen; i++ {
				lineInfo := make(map[string]interface{})
				lineInfo["ID"] = obj.AttributeValues[i].ID
				lineInfo["Name"] = obj.AttributeValues[i].Name
				attributeValues = append(attributeValues, lineInfo)
			}
		}
		objInfo["AttributeValues"] = attributeValues
		valueInfo = objInfo
	}
	return
}
