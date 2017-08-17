package initalActions

import (
	"encoding/xml"
	md "golangERP/models"
	service "golangERP/services"
	"golangERP/utils"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/orm"
)

// InitModuleModules 模块数据
type InitModuleModules struct {
	XMLName xml.Name     `xml:"Modules"`
	Modules []InitModule `xml:"module"`
}

// InitModule 模块数据解析
type InitModule struct {
	md.ModuleModule
	XMLID    string `xml:"id,attr"`
	Category string `xml:"category"`
}

// InitModuleModule2DB 初始化模块数据
func InitModuleModule2DB(split string) {
	if xmlDir, err := os.Getwd(); err == nil {
		xmlBase := utils.StringsJoin(xmlDir, split, "inital_data", split, "xml", split, "module")
		if dirList, err := ioutil.ReadDir(xmlBase); err == nil {
			var user md.User
			user.ID = 0
			user.IsAdmin = true
			for _, dir := range dirList {
				if dir.IsDir() {
					continue
				}
				if file, err := os.Open(utils.StringsJoin(xmlBase, split, dir.Name())); err == nil {
					defer file.Close()
					if data, err := ioutil.ReadAll(file); err == nil {
						var initModuleModules InitModuleModules
						var moduleName = "ModuleModule"
						ormObj := orm.NewOrm()
						if xml.Unmarshal(data, &initModuleModules) == nil {
							for _, moduleXML := range initModuleModules.Modules {
								var module md.ModuleModule
								var xmlid = utils.StringsJoin(moduleName, ".", moduleXML.XMLID)
								if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {
									module.Name = moduleXML.Name
									module.Description = moduleXML.Description
									if category, err := md.GetModuleCategoryByName(moduleXML.Category, ormObj); err == nil {
										module.Category = category
									}
									if insertID, err := service.ServiceCreateModuleModule(&user, &module); err == nil {
										var moduleData md.ModuleData
										moduleData.InsertID = insertID
										moduleData.XMLID = xmlid
										moduleData.Descrition = module.Name
										moduleData.ModuleName = moduleName
										md.AddModuleData(&moduleData, ormObj)
									}
								}
							}
						}
					}
				}
			}
		}
	}

}
