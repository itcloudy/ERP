package base

// IndexController home controller
type IndexController struct {
	BaseController
}

// Get home page
func (ctl *IndexController) Get() {

	// 基础布局页面
	ctl.Layout = "base/base.html"
	ctl.TplName = "base/module_dashboard.html"

}
