package services

import (
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateUser 创建记录
func ServiceCreateUser(user *md.User, obj *md.User) (id int64, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "User"); err == nil {
		if !access.Create {
			err = errors.New("has no create permission ")
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
	obj.Password = utils.PasswordMD5(obj.Password, obj.Mobile)
	obj.CreateUserID = user.ID
	id, err = md.AddUser(obj, o)
	return
}

// ServiceUpdateUser 更新记录
func ServiceUpdateUser(user *md.User, obj *md.User) (id int64, err error) {
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
	obj.UpdateUserID = user.ID
	id, err = md.UpdateUser(obj, o)

	return
}

// ServiceUpdateUserPassWord 更新密码
func ServiceUpdateUserPassWord(user *md.User, obj *md.User) (id int64, err error) {
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
	obj.UpdateUserID = user.ID
	id, err = md.UpdateUser(obj, o)

	return
}

// ServiceUserLogin 用户登录
func ServiceUserLogin(username string, password string) (*md.User, bool) {
	o := orm.NewOrm()
	var (
		user md.User
		err  error
	)
	ok := false
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("Name", username).Or("Email", username).Or("Mobile", username)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			ok = true

		}
	}
	return &user, ok
}

// ServiceUserLogout 用户登出
func ServiceUserLogout(id int64) (ok bool, err error) {
	return
}

// ServiceGetUser 获得用户列表
func ServiceGetUser(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (paginator utils.Paginator, results []map[string]interface{}, err error) {
	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "User"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.User
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllUser(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)

		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			objInfo["NameZh"] = obj.NameZh
			objInfo["Email"] = obj.Email
			objInfo["Mobile"] = obj.Mobile
			objInfo["Tel"] = obj.Tel
			objInfo["IsAdmin"] = obj.IsAdmin
			objInfo["Active"] = obj.Active
			objInfo["Qq"] = obj.Qq
			objInfo["WeChat"] = obj.WeChat
			objInfo["IsBackground"] = obj.IsBackground
			results = append(results, objInfo)
		}
	}
	return
}
