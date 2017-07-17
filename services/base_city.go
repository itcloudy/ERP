package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCity 创建表
func ServiceCreateAddressCity(obj *md.AddressCity) (id int64, err error) {
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
	id, err = md.AddAddressCity(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}

// ServiceUpdateAddressCity 更新表
func ServiceUpdateAddressCity(obj *md.AddressCity) (id int64, err error) {
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
	id, err = md.UpdateAddressCity(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
