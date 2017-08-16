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
	id, err = md.AddUser(obj, o)
	return
}

// ServiceUpdateUser 更新记录
func ServiceUpdateUser(obj *md.User) (id int64, err error) {
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
	id, err = md.UpdateUser(obj, o)

	return
}

// ServiceUpdateUserPassWord 更新密码
func ServiceUpdateUserPassWord(obj *md.User) (id int64, err error) {
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
