package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateBaseGroup 创建记录
func ServiceCreateBaseGroup(user *md.User, obj *md.BaseGroup) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "BaseGroup"); err == nil {
		if !access.Create {
			err = errors.New("has no create permission ")
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
	var groupMax md.BaseGroup
	if obj.Parent != nil {
		if parent, err := md.GetBaseGroupByID(obj.Parent.ID, o); err == nil {
			var maxParentRight int64
			// 获得同级最右的group
			if err = o.QueryTable(&parent).Filter("Parent__id", parent.ID).OrderBy("-ParentRight").Limit(1).One(&groupMax); err == nil {

				maxParentRight = groupMax.ParentRight
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
		// 判断是否有菜单
		if err = o.QueryTable(&obj).OrderBy("-ParentRight").Limit(1).One(&groupMax); err == nil {
			obj.ParentLeft = groupMax.ParentRight + 1
			obj.ParentRight = groupMax.ParentRight + 2
		} else {
			obj.ParentLeft = 0
			obj.ParentRight = 1
		}
	}
	obj.CreateUserID = user.ID
	id, err = md.AddBaseGroup(obj, o)

	return
}

// ServiceUpdateBaseGroup 更新记录
func ServiceUpdateBaseGroup(user *md.User, obj *md.BaseGroup) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "BaseGroup"); err == nil {
		if !access.Update {
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
	id, err = md.UpdateBaseGroup(obj, o)
	return
}

// ServiceGetUserGroups 获得用户的权限组信息
func ServiceGetUserGroups(isAdmin bool, userID int64) (groups []*md.BaseGroup, err error) {
	var (
		tGroups []md.BaseGroup
	)
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	condAnd := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	o := orm.NewOrm()
	if isAdmin {
		condAnd["Parent__isnull"] = true
		if len(condAnd) > 0 {
			cond["and"] = condAnd
		}
		if tGroups, err = md.GetAllBaseGroup(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
			for _, group := range tGroups {
				groups = append(groups, &group)
			}
		}
	} else {
		if user, err := md.GetUserByID(userID, o); err == nil {
			groups = user.Groups
		}
	}
	return
}

// ServiceGetGroup 获得用户列表
func ServiceGetGroup(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "Group"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.BaseGroup
	o := orm.NewOrm()
	if arrs, err = md.GetAllBaseGroup(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			objInfo["Category"] = obj.Category
			objInfo["Description"] = obj.Description
			results = append(results, objInfo)
		}
	}
	return
}
