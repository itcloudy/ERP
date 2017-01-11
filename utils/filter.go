package utils

import "github.com/astaxie/beego/context"

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login/in")
	}
}
