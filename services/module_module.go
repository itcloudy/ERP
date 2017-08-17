package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateModuleModule 创建记录
func ServiceCreateModuleModule(user *md.User, obj *md.ModuleModule) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModuleModule"); err == nil {
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
	id, err = md.AddModuleModule(obj, o)

	return
}

// ServiceUpdateModuleModule 更新记录
func ServiceUpdateModuleModule(user *md.User, obj *md.ModuleModule) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ModuleModule"); err == nil {
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
	id, err = md.UpdateModuleModule(obj, o)

	return
}
