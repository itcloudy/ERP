package initalActions

import (
	"golangERP/utils"
	"os"
	"runtime"
)

// InitApp 基础数据插入
func InitApp() {
	systemType := runtime.GOOS
	split := "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	if xmlDir, err := os.Getwd(); err == nil {
		xmlBase := utils.StringsJoin(xmlDir, split, "inital_data", split, "xml")
		// 国家信息
		countryXML := utils.StringsJoin(xmlBase, split, "address", split, "Countries.xml")
		InitCountry2DB(countryXML)
		// 省份信息
		provinceXML := utils.StringsJoin(xmlBase, split, "address", split, "Provinces.xml")
		InitProvince2DB(provinceXML)
		// 城市信息
		cityXML := utils.StringsJoin(xmlBase, split, "address", split, "Cities.xml")
		InitCity2DB(cityXML)
		// 地区信息
		districtXML := utils.StringsJoin(xmlBase, split, "address", split, "Districts.xml")
		InitDistrict2DB(districtXML)
		// group初始化要在用户和菜单之前
		groupXML := utils.StringsJoin(xmlBase, split, "Groups.xml")
		InitGroup2DB(groupXML)
		// 模块分类
		moduleCategoryXML := utils.StringsJoin(xmlBase, split, "module_category.xml")
		InitModuleCategory2DB(moduleCategoryXML)
		// 模块信息
		InitModuleModule2DB(split)
		//模块权限信息
		InitModelAccess2DB(split)

		//用户初始化
		userXML := utils.StringsJoin(xmlBase, split, "Users.xml")
		InitUser2DB(userXML)
		//菜单初始化
		InitMenus2DB(split)

	}
}
