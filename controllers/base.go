package controllers

import (
	md "golangERP/models"
	"time"

	"github.com/astaxie/beego"
)

// BaseController 基础controller
type BaseController struct {
	beego.Controller
	IsAdmin   bool
	UserName  string
	URL       string
	LastLogin time.Time
	User      md.User
}

// Prepare implemented Prepare method for baseRouter.
func (ctl *BaseController) Prepare() {
	// flash := beego.NewFlash()
	// Setting properties.
	ctl.StartSession()
	ctl.Data["PageStartTime"] = time.Now()
	user := ctl.GetSession("User")
	if user != nil {
		ctl.User = user.(md.User)
		if ctl.User.IsAdmin {
			ctl.IsAdmin = true
		}
		ctl.Data["LoginUser"] = user
		ctl.Data["LastLogin"] = ctl.GetSession("LastLogin")
	}

}
