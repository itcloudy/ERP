package initalActions

import (
	md "golangERP/models"

	"github.com/astaxie/beego/orm"
)

// InitGroupMenu 权限菜单初始化
func InitGroupMenu() {
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	o := orm.NewOrm()
	var (
		groups []md.BaseGroup
		err    error
	)
	if groups, err = md.GetAllBaseGroup(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
		for _, group := range groups {
			go groupMenu(group)
		}
	}
}

//某个权限的菜单
func groupMenu(group md.BaseGroup) map[string]interface{} {
	// 最终结果，按照菜单的层级关系
	var resultMenus = make(map[string]interface{})
	// 临时保存没有找到上级菜单的菜单信息
	var tempMenus = make(map[string]interface{})
	o := orm.NewOrm()
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 0)
	order := make([]string, 0, 0)
	condAnd := make(map[string]interface{})
	// 过滤条件使用增加 等于(e)，包括本身
	condAnd["ParentRight__lte"] = group.ParentRight
	condAnd["ParentLeft__gte"] = group.ParentLeft
	groupIDs := make([]int64, 0, 0)
	if len(condAnd) > 0 {
		cond["and"] = condAnd
	}
	if childs, err := md.GetAllBaseGroup(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
		for _, child := range childs {
			groupIDs = append(groupIDs, child.ID)
		}
	}
	// 获得menus
	menuCondAnd := make(map[string]interface{})
	menuCondAnd["Group__id__in"] = groupIDs
	if len(menuCondAnd) > 0 {
		cond["and"] = menuCondAnd
	}

	// 按ParentRight排序，则第一个为最顶级的菜单，如果顶级菜单上面还有菜单，则忽略后面的子菜单，即子菜单不自动升级为上级菜单
	if groupMenus, err := md.GetAllGroupMenu(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
		menuLen := len(groupMenus)
		for i := 0; i < menuLen; i++ {
			menuID := groupMenus[i].Menu.ID
			if menu, err := md.GetBaseMenuByID(menuID, o); err == nil {
				var menuInfo = make(map[string]interface{})

				if len(menu.Childs) > 0 {
					menuInfo["HasChild"] = true
				} else {
					menuInfo["HasChild"] = false
				}
				menuInfo["ID"] = menu.ID
				menuInfo["Name"] = menu.Name
				menuInfo["Index"] = menu.Index
				menuInfo["Icon"] = menu.Icon
				menuInfo["Path"] = menu.Path
				menuInfo["ComponentPath"] = menu.ComponentPath
				menuInfo["Mete"] = menu.Meta
				menuInfo["Child"] = make(map[string]interface{})
				tempMenus[menu.Index] = menuInfo
				step := menu.ParentRight - menu.ParentLeft
				menuInfo["Step"] = step
			}

		}
	}
	return resultMenus
}
