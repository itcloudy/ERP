package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.Data["Website"] = "golangERP"
	c.Data["Email"] = "272685110@qq.com"
	c.TplName = "index.html"
}
