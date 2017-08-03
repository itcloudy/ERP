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
		countryXML := utils.StringsJoin(xmlBase, split, "address", split, "Countries.xml")
		go InitCountry2DB(countryXML)
		provinceXML := utils.StringsJoin(xmlBase, split, "address", split, "Provinces.xml")
		go InitProvince2DB(provinceXML)
		cityXML := utils.StringsJoin(xmlBase, split, "address", split, "Cities.xml")
		go InitCity2DB(cityXML)
		districtXML := utils.StringsJoin(xmlBase, split, "address", split, "Districts.xml")
		go InitDistrict2DB(districtXML)
		groupXML := utils.StringsJoin(xmlBase, split, "Groups.xml")
		// group初始化要在用户和菜单之前
		go InitGroup2DB(groupXML)
		userXML := utils.StringsJoin(xmlBase, split, "Users.xml")
		go InitUser2DB(userXML)
		//菜单初始化
		go InitMenus2DB(split)

	}
}
