package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateModuleModule 创建表
func ServiceCreateModuleModule(obj *md.ModuleModule) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if err != nil {
		return
	}
	id, err = md.AddModuleModule(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}

// ServiceUpdateModuleModule 更新表
func ServiceUpdateModuleModule(obj *md.ModuleModule) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if err != nil {
		return
	}
	id, err = md.UpdateModuleModule(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
