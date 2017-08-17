package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateModelAccess 创建记录
func ServiceCreateModelAccess(user *md.User, obj *md.ModelAccess) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModelAccess"); err == nil {
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
	obj.CreateUserID = user.ID
	id, err = md.AddModelAccess(obj, o)

	return
}

// ServiceUpdateModelAccess 更新记录
func ServiceUpdateModelAccess(user *md.User, obj *md.ModelAccess) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModelAccess"); err == nil {
		if !access.Update {
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
	obj.UpdateUserID = user.ID
	id, err = md.UpdateModelAccess(obj, o)

	return
}

// ServiceCheckUserModelAssess 权限检查
func ServiceCheckUserModelAssess(user *md.User, moduleName string) (access utils.AccessResult, err error) {
	var (
		groups  []*md.BaseGroup
		modules []md.ModelAccess
	)
	// 若为系统管理员拥有所有的权限
	if user.IsAdmin {
		access.Create = true
		access.Update = true
		access.Read = true
		access.Unlink = true
		return
	}
	// 获得用户所有的权限组
	if groups, err = ServiceGetUserGroups(user.IsAdmin, user.ID); err == nil {
		// 获得权限组下所有的模块访问权限
		leng := len(groups)
		if leng > 0 {
			groupIDs := make([]int64, leng, leng)
			for index, group := range groups {
				groupIDs[index] = group.ID
			}
			query := make(map[string]interface{})
			exclude := make(map[string]interface{})
			cond := make(map[string]map[string]interface{})
			condAnd := make(map[string]interface{})
			fields := make([]string, 0, 0)
			sortby := make([]string, 0, 1)
			order := make([]string, 0, 1)
			o := orm.NewOrm()
			condAnd["Group__id__in"] = groupIDs
			condAnd["Module__Name"] = moduleName
			if len(condAnd) > 0 {
				cond["and"] = condAnd
			}
			if _, modules, err = md.GetAllModelAccess(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
				for _, module := range modules {
					access.Create = module.PermCreate || access.Create
					access.Update = module.PermWrite || access.Update
					access.Read = module.PermRead || access.Read
					access.Unlink = module.PermUnlink || access.Unlink
				}

			}
		} else {
			err = errors.New("user has no  any permissions")
		}
	}
	return
}
