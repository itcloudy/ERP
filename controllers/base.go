package controllers

import (
	"fmt"
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

}
func (ctl *BaseController) Post() {
	fmt.Println(12312313)
}
