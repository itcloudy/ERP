package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateUser 创建表
func ServiceCreateUser(obj *md.User) (id int64, err error) {
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
	id, err = md.AddUser(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}

// ServiceUpdateUser 更新表
func ServiceUpdateUser(obj *md.User) (id int64, err error) {
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
	id, err = md.UpdateUser(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
