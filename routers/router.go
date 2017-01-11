package routers

import (
	"goERP/controllers/address"
	"goERP/controllers/base"
	"goERP/controllers/product"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &base.IndexController{})
	//=======================================基本操作===========================================
	//登录
	beego.Router("/login/:action([A-Za-z]+)/", &base.LoginController{})
	//用户
	beego.Router("/user/?:id", &base.UserController{})
	//公司
	beego.Router("/company/?:id", &base.CompanyController{})
	//部门
	beego.Router("/department/?:id", &base.DepartmentController{})
	//职位
	beego.Router("/position/?:id", &base.TemplateFileController{})

	//权限
	beego.Router("/group/?:id", &base.GroupController{})
	//团队
	beego.Router("/team/?:id", &base.TeamController{})
	//登录日志
	beego.Router("/record/", &base.RecordController{})
	//文件模版
	beego.Router("/templatefile/?:id", &base.TemplateFileController{})
	// ===============================地址===========================================
	//国家
	beego.Router("/address/country/?:id", &address.CountryController{})
	//省份
	beego.Router("/address/province/?:id", &address.ProvinceController{})
	//城市
	beego.Router("/address/city/?:id", &address.CityController{})
	//区县
	beego.Router("/address/district/?:id", &address.DistrictController{})
	//=======================================产品管理===========================================
	//产品类别
	beego.Router("/product/category/?:id", &product.ProductCategoryController{})
	//属性
	beego.Router("/product/attribute/?:id", &product.ProductAttributeController{})
	//属性值
	beego.Router("/product/attributevalue/?:id", &product.ProductAttributeValueController{})
	//属性值明细
	beego.Router("/product/attributeline/:action([A-Za-z]+)/?:id", &product.ProductAttributeLineController{})
	//产品款式
	beego.Router("/product/template/?:id", &product.ProductTemplateController{})
	//产品规格
	beego.Router("/product/product/?:id", &product.ProductProductController{})

	//产品标签
	beego.Router("/product/tag/:action([A-Za-z]+)/?:id", &product.ProductTagController{})
	//产品包装
	beego.Router("/product/packaging/:action([A-Za-z]+)/?:id", &product.ProductPackagingController{})
	//产品属性价格
	beego.Router("/product/attributeprice/:action([A-Za-z]+)/?:id", &product.ProductAttributePriceController{})
	//产品计量单位
	beego.Router("/product/uom/?:id", &product.ProductUomController{})
	//产品计量单位类别
	beego.Router("/product/uomcateg/?:id", &product.ProductUomCategController{})

}
