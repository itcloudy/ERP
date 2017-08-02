package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressProvince 创建表
func ServiceCreateAddressProvince(obj *md.AddressProvince) (id int64, err error) {
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
	id, err = md.AddAddressProvince(obj, o)

	return
}

// ServiceUpdateAddressProvince 更新表
func ServiceUpdateAddressProvince(obj *md.AddressProvince) (id int64, err error) {
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
	id, err = md.UpdateAddressProvince(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
