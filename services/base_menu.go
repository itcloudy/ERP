package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateBaseMenu 创建记录
func ServiceCreateBaseMenu(user *md.User, obj *md.BaseMenu) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "BaseMenu"); err == nil {
		if !access.Create {
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
				o.QueryTable(&parent).Filter("ParentRight__gt", maxParentRight).Filter("ParentLeft__gte", maxParentRight).Exclude("ID", parent.ID).Update(orm.Params{
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
	obj.CreateUserID = user.ID
	id, err = md.AddBaseMenu(obj, o)

	return
}

// ServiceGetMenus 获得菜单
func ServiceGetMenus(isAdmin bool, groupIDs []int64) (menus []md.BaseMenu, err error) {

	o := orm.NewOrm()
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	// 如果为管理员则获得所有的菜单
	if isAdmin {
		menus, _ = md.GetAllBaseMenu(o, query, exclude, cond, fields, sortby, order, 0, 0)
	} else {
		// 获得权限组下所有的下级权限组
		var allGroupIDs []int64
		for _, groupID := range groupIDs {
			if group, err := md.GetBaseGroupByID(groupID, o); err == nil {
				condFor := make(map[string]map[string]interface{})
				condAndFor := make(map[string]interface{})
				condAndFor["ParentRight__lte"] = group.ParentRight
				condAndFor["ParentLeft__gte"] = group.ParentLeft
				condFor["and"] = condAndFor
				if childs, err := md.GetAllBaseGroup(o, query, exclude, condFor, fields, sortby, order, 0, 0); err == nil {
					for _, child := range childs {
						allGroupIDs = append(allGroupIDs, child.ID)
					}
				}
			}
		}
		// 获得权限组所有的菜单
		menuCondAnd := make(map[string]interface{})
		menuCondAnd["Group__id__in"] = allGroupIDs
		if len(menuCondAnd) > 0 {
			cond["and"] = menuCondAnd
		}
		if groupMenus, err := md.GetAllGroupMenu(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
			groupMenuLen := len(groupMenus)
			// 临时保存没有找到上级菜单的菜单信息,后面是否为int类型不重要，只为tempMenus中有该index
			tempMenus := make(map[string]int)
			for i := 0; i < groupMenuLen; i++ {
				menu := groupMenus[i].Menu
				index := menu.Index
				if _, ok := tempMenus[index]; ok {
					continue
				}
				menu.ViewType = strings.Replace(menu.ViewType, "\n", "", 0)
				menus = append(menus, *menu)
				tempMenus[index] = 1
			}
		}
	}

	return
}
