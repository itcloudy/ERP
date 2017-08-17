package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateModuleCategory 创建记录
func ServiceCreateModuleCategory(user *md.User, obj *md.ModuleCategory) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModuleCategory"); err == nil {
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
	id, err = md.AddModuleCategory(obj, o)

	return
}

// ServiceUpdateModuleCategory 更新记录
func ServiceUpdateModuleCategory(user *md.User, obj *md.ModuleCategory) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModuleCategory"); err == nil {
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
	obj.UpdateUserID = user.ID
	id, err = md.UpdateModuleCategory(obj, o)
	return
}
