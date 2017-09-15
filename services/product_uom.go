package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"
	"reflect"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductUom 创建记录
func ServiceCreateProductUom(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductUom"); err == nil {
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
	var obj md.ProductUom
	json.Unmarshal([]byte(requestBody), &obj)
	var requestBodyMap map[string]interface{}
	json.Unmarshal(requestBody, &requestBodyMap)
	var uomcateg md.ProductUomCateg
	if Category, ok := requestBodyMap["Category"]; ok {
		uomcategT := reflect.TypeOf(Category)
		if uomcategT.Kind() == reflect.Map {
			uomcategMap := Category.(map[string]interface{})
			if uomcategID, ok := uomcategMap["ID"]; ok {
				uomcateg.ID, _ = utils.ToInt64(uomcategID)
				obj.Category = &uomcateg
			}
		} else if uomcategT.Kind() == reflect.String {
			uomcateg.ID, _ = utils.ToInt64(Category)
			obj.Category = &uomcateg
		}
	}
	obj.CreateUserID = user.ID
	id, err = md.AddProductUom(&obj, o)

	return
}

// ServiceUpdateProductUom 更新记录
func ServiceUpdateProductUom(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductUom"); err == nil {
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
	var obj md.ProductUom
	var objPtr *md.ProductUom
	if objPtr, err = md.GetProductUomByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductUom(&obj, o)

	return
}

//ServiceGetProductUom 获得城市列表
func ServiceGetProductUom(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductUom"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductUom
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductUom(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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

// ServiceGetProductUomByID get ProductUom by id
func ServiceGetProductUomByID(user *md.User, id int64) (access utils.AccessResult, attrInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductUom"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var attr *md.ProductUom
	if attr, err = md.GetProductUomByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = attr.Name
		objInfo["ID"] = attr.ID
		attrInfo = objInfo
	}
	return
}
