package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"
	"reflect"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductCategory 创建记录
func ServiceCreateProductCategory(user *md.User, requestBody []byte) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductCategory"); err == nil {
		if !access.Create {
			err = errors.New("has no update permission ")
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
	var obj md.ProductCategory
	var cateMax md.ProductCategory
	json.Unmarshal([]byte(requestBody), &obj)
	var requestBodyMap map[string]interface{}
	json.Unmarshal(requestBody, &requestBodyMap)
	parentNotOK := false
	var parent md.ProductCategory
	if Parent, ok := requestBodyMap["Parent"]; ok {

		parentT := reflect.TypeOf(Parent)
		if parentT.Kind() == reflect.Map {
			parentMap := Parent.(map[string]interface{})
			if pID, ok := parentMap["ID"]; ok {
				parent.ID, _ = utils.ToInt64(pID)
				obj.Parent = &parent
			}
		} else if parentT.Kind() == reflect.String {
			parent.ID, _ = utils.ToInt64(Parent)
			obj.Parent = &parent
		}
		if parent.ID > 0 {
			if pt, err := md.GetProductCategoryByID(parent.ID, o); err == nil {
				parent = *pt
			}
			var maxParentRight int64
			if err = o.QueryTable(&parent).Filter("Parent__id", parent.ID).OrderBy("-ParentRight").Limit(1).One(&cateMax); err == nil {
				maxParentRight = cateMax.ParentRight
				obj.ParentLeft = maxParentRight + 1
				obj.ParentRight = maxParentRight + 2
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Exclude("ID", parent.ID).Update(orm.Params{
					"ParentLeft": orm.ColValue(orm.ColAdd, 2),
				})
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Update(orm.Params{
					"ParentRight": orm.ColValue(orm.ColAdd, 2),
				})

			} else {
				maxParentRight = parent.ParentRight
				obj.ParentLeft = parent.ParentLeft + 1
				obj.ParentRight = parent.ParentLeft + 2
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Filter("ParentLeft__gte", maxParentRight).Exclude("ID", parent.ID).Update(orm.Params{
					"ParentLeft": orm.ColValue(orm.ColAdd, 2),
				})
				o.QueryTable(&parent).Filter("ParentRight__gte", maxParentRight).Update(orm.Params{
					"ParentRight": orm.ColValue(orm.ColAdd, 2),
				})
			}
		} else {
			parentNotOK = true
		}
	}
	if parentNotOK {
		// 判断是否有分类
		if err = o.QueryTable(&obj).OrderBy("-ParentRight").Limit(1).One(&cateMax); err == nil {
			obj.ParentLeft = cateMax.ParentRight + 1
			obj.ParentRight = cateMax.ParentRight + 2
		} else {
			obj.ParentLeft = 0
			obj.ParentRight = 1
		}
	}
	obj.CreateUserID = user.ID
	id, err = md.AddProductCategory(&obj, o)

	return
}

// ServiceUpdateProductCategory 更新记录
func ServiceUpdateProductCategory(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductCategory"); err == nil {
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
	var obj md.ProductCategory
	var objPtr *md.ProductCategory
	if objPtr, err = md.GetProductCategoryByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)
	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductCategory(&obj, o)

	return
}

//ServiceGetProductCategory 获得分类列表
func ServiceGetProductCategory(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductCategory"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission ")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductCategory
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductCategory(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			parentInfo := make(map[string]interface{})
			parentInfo["ID"] = obj.Parent.ID
			parentInfo["Name"] = obj.Parent.Name
			objInfo["Parent"] = parentInfo
			results = append(results, objInfo)
		}
	}
	return
}

// ServiceGetProductCategoryByID get ProductCategory by id
func ServiceGetProductCategoryByID(user *md.User, id int64) (access utils.AccessResult, catetoryInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductCategory"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var catetory *md.ProductCategory
	if catetory, err = md.GetProductCategoryByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = catetory.Name
		objInfo["ID"] = catetory.ID
		parentInfo := make(map[string]interface{})
		parentInfo["ID"] = catetory.Parent.ID
		parentInfo["Name"] = catetory.Parent.Name
		objInfo["Parent"] = parentInfo

		catetoryInfo = objInfo
	}
	return
}
