package services

import (
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
	var menuMax md.BaseMenu
	if obj.Parent != nil {
		if parent, err := md.GetBaseMenuByID(obj.Parent.ID, o); err == nil {
			var maxParentRight int64
			if err = o.QueryTable(&parent).Filter("Parent__id", parent.ID).OrderBy("-ParentRight").Limit(1).One(&menuMax); err == nil {
				maxParentRight = menuMax.ParentRight
				obj.ParentLeft = maxParentRight + 1
				obj.ParentRight = maxParentRight + 2
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Exclude("ID", parent.ID).Update(orm.Params{
					"ParentLeft": orm.ColValue(orm.ColAdd, 2),
				})
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Update(orm.Params{
					"ParentRight": orm.ColValue(orm.ColAdd, 2),
				})

			} else {
				maxParentRight = parent.ParentRight
				obj.ParentLeft = parent.ParentLeft + 1
				obj.ParentRight = parent.ParentLeft + 2
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Exclude("ID", parent.ID).Update(orm.Params{
					"ParentLeft": orm.ColValue(orm.ColAdd, 2),
				})
				o.QueryTable(&parent).Filter("ParentRight__gte", maxParentRight).Update(orm.Params{
					"ParentRight": orm.ColValue(orm.ColAdd, 2),
				})
			}
		}
	} else {
		// 判断是否有菜单
		if err = o.QueryTable(&obj).OrderBy("-ParentRight").Limit(1).One(&menuMax); err == nil {
			obj.ParentLeft = menuMax.ParentRight + 1
			obj.ParentRight = menuMax.ParentRight + 2
		} else {
			obj.ParentLeft = 0
			obj.ParentRight = 1
		}
	}
	id, err = md.AddBaseMenu(obj, o)
	err = o.Commit()
	if err != nil {
		return
	}
	return
}
