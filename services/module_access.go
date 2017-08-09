package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateModelAccess 创建记录
func ServiceCreateModelAccess(obj *md.ModelAccess) (id int64, err error) {
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
	id, err = md.AddModelAccess(obj, o)

	return
}

// ServiceUpdateModelAccess 更新记录
func ServiceUpdateModelAccess(obj *md.ModelAccess) (id int64, err error) {
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
	id, err = md.UpdateModelAccess(obj, o)

	return
}

// ServiceCheckUserModelAssess 权限检查
func ServiceCheckUserModelAssess() {

}
