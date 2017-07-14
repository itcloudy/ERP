package controllers

import (
	md "golangERP/models"
	"html/template"
	"time"

	"github.com/astaxie/beego"
)

var (
	//AppVer 版本
	AppVer string
	//IsPro 生产还是开发环境
	IsPro bool
)

// BaseController 基础controller
type BaseController struct {
	beego.Controller
	IsAdmin    bool
	UserName   string
	URL        string
	LastLogin  time.Time
	User       md.User
	PageName   string //页面名称，用于提示用户
	PageAction string //页面动作
}

// Prepare implemented Prepare method for baseRouter.
func (ctl *BaseController) Prepare() {
	// flash := beego.NewFlash()
	// Setting properties.
	ctl.StartSession()
	ctl.Data["AppVer"] = AppVer
	ctl.Data["IsPro"] = IsPro
	ctl.Data["xsrf"] = template.HTML(ctl.XSRFFormHTML())
	ctl.Data["PageStartTime"] = time.Now()

}
