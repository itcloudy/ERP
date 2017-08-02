package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressCountry 创建表
func ServiceCreateAddressCountry(obj *md.AddressCountry) (id int64, err error) {
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
	id, err = md.AddAddressCountry(obj, o)

	return
}

// ServiceUpdateAddressCountry 更新表
func ServiceUpdateAddressCountry(obj *md.AddressCountry) (id int64, err error) {
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
	id, err = md.UpdateAddressCountry(obj, o)

	return
}
