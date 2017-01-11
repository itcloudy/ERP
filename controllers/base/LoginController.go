package base

import md "goERP/models"

type LoginController struct {
	BaseController
}

func (ctl *LoginController) Get() {
	action := ctl.GetString(":action")
	if action == "out" {
		ctl.Logout()
		ctl.Redirect("/login/in", 302)
	} else if action == "in" {
		user := ctl.GetSession("User")
		if user != nil {
			ctl.Redirect("/", 302)
		}
		ctl.TplName = "login.html"
	}

}
func (ctl *LoginController) Post() {

	loginName := ctl.GetString("loginName")
	password := ctl.GetString("password")
	rememberMe := ctl.GetString("remember")

	if loginName == "" && password == "" {
		ctl.Redirect("/login/in", 302)
	}

	var (
		user   md.User
		err    error
		record md.Record
		ok     bool
	)
	if user, err, ok = md.CheckUserByName(loginName, password); ok != true {
		ctl.Redirect("/login/in", 302)
	} else {
		if record, err = md.GetLastRecordByUserID(user.Id); err == nil {

			ctl.SetSession("LastLogin", record.CreateDate)
			ctl.SetSession("LastIp", record.Ip)
		}
		var record md.Record
		record.Ip = ctl.Ctx.Input.IP()
		record.UserAgent = ctl.Ctx.Request.UserAgent()
		record.User = &user
		md.AddRecord(&record)
		ctl.SetSession("User", user)

		ctl.Ctx.SetCookie("Remember", rememberMe, 31536000, "/")
		//通过验证跳转到主界面
		ctl.Redirect("/", 302)
	}
}

//登出
func (ctl *LoginController) Logout() {
	if record, err := md.GetLastRecordByUserID(ctl.User.Id); err == nil {
		record.Ip = ctl.Ctx.Input.IP()
		record.UpdateUser = &ctl.User
		md.UpdateRecordById(&record)
	}
	ctl.SetSession("User", nil)
	ctl.DelSession("User")

}
