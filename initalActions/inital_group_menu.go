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
func groupMenu(group md.BaseGroup) {

}
