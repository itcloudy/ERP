package routers

import (
	"golangERP/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 首页,返回的为html，其他页面的请求返回的都为json
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login/?:id", &controllers.LoginController{})
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/address/country/?:id", &controllers.AddressCountryController{})
	beego.Router("/address/province/?:id", &controllers.AddressProvinceController{})
	beego.Router("/address/city/?:id", &controllers.AddressCityController{})
	beego.Router("/address/district/?:id", &controllers.AddressDistrictController{})
	beego.Router("/setting/user/?:id", &controllers.UserController{})
	beego.Router("/setting/group/?:id", &controllers.GroupController{})
	beego.Router("/partner/?:id", &controllers.PartnerController{})
	// 产品管理
	beego.Router("/product/attribute/?:id", &controllers.ProductAttributeController{})
	beego.Router("/product/attribute/line/?:id", &controllers.ProductAttributeLineController{})
	beego.Router("/product/attributevalue/?:id", &controllers.ProductAttributeValueController{})
	beego.Router("/product/template/?:id", &controllers.ProductTemplateController{})
	beego.Router("/product/product/?:id", &controllers.ProductProductController{})
	beego.Router("/product/uom/?:id", &controllers.ProductUomController{})
	beego.Router("/product/uomcateg/?:id", &controllers.ProductUomCategController{})
	beego.Router("/product/category/?:id", &controllers.ProductCategoryController{})
	// 销售管理
	beego.Router("/sale/order/?:id", &controllers.SaleOrderController{})
	beego.Router("/sale/order/line/?:id", &controllers.SaleOrderLineController{})
}
