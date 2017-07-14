package routers

import (
	"golangERP/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 首页,返回的为html，其他页面的请求返回的都为json
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginContriller{})
}
