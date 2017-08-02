package services

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateAddressDistrict 创建表
func ServiceCreateAddressDistrict(obj *md.AddressDistrict) (id int64, err error) {
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
	id, err = md.AddAddressDistrict(obj, o)

	return
}

// ServiceUpdateAddressDistrict 更新表
func ServiceUpdateAddressDistrict(obj *md.AddressDistrict) (id int64, err error) {
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
	id, err = md.UpdateAddressDistrict(obj, o)
	return
}
