package controllers

import (
	"fmt"
	"regexp"
	"strings"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.Data["Website"] = "golangERP"
	c.Data["Email"] = "272685110@qq.com"
	typeMap := map[string]string{"windows": "pc", "android": "mobile", "linux": "pc", "iphone": "mobile"}
	accessType := "pc"
	// 判断访问类型是pc还是移动端，根据不同的终端选择不同的页面，默认为pc端
	userAgent := strings.ToLower(c.Ctx.Request.UserAgent())
	for userAgentReg, atype := range typeMap {
		if ok, _ := regexp.MatchString(userAgentReg, userAgent); ok {
			accessType = atype
			break
		}
	}
	switch accessType {
	case "pc":
		fmt.Println("Pc")
		c.TplName = "index_pc.html"
	case "mobile":
		fmt.Println("mobile")
		c.TplName = "index_mobile.html"
	default:
		c.TplName = "index_pc.html"
	}
}
