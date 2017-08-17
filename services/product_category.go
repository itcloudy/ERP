package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductCategory 创建记录
func ServiceCreateProductCategory(user *md.User, obj *md.ProductCategory) (id int64, err error) {
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
	var cateMax md.ProductCategory
	if obj.Parent != nil {
		if parent, err := md.GetProductCategoryByID(obj.Parent.ID, o); err == nil {
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
		}
	} else {
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
	id, err = md.AddProductCategory(obj, o)

	return
}

//ServiceGetProductCategory 获得分类列表
func ServiceGetProductCategory(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var access utils.AccessResult
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
			countryInfo := make(map[string]interface{})
			countryInfo["ID"] = obj.Country.ID
			countryInfo["Name"] = obj.Country.Name
			objInfo["Country"] = countryInfo
			objInfo["ID"] = obj.ID
			results = append(results, objInfo)
		}
	}
	return
}
