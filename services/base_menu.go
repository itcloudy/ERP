package services

import (
	"fmt"
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateBaseMenu 创建表
func ServiceCreateBaseMenu(obj *md.BaseMenu) (id int64, err error) {
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
	if obj.Parent != nil {
		if menu, err := md.GetBaseMenuByID(obj.Parent.ID, o); err == nil {
			parentRight := menu.ParenRight
			o.QueryTable(&menu).Filter("ParentRight__gt", parentRight).Update(orm.Params{
				"ParentRight": orm.ColValue(orm.ColAdd, 2),
				"ParentLeft":  orm.ColValue(orm.ColAdd, 2),
			})
			obj.ParenLeft = menu.ParenLeft + 1
			obj.ParenRight = menu.ParenLeft + 2
			o.Update(obj)
		}
	} else {
		obj.ParenLeft = 0
		obj.ParenRight = 1
	}
	fmt.Printf("%+v\n", obj)
	id, err = md.AddBaseMenu(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
