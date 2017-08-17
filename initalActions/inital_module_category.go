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

// InitModuleCategories 模块分类数据
type InitModuleCategories struct {
	XMLName    xml.Name       `xml:"Categories"`
	Categories []InitCategory `xml:"category"`
}

// InitCategory 用户数据解析
type InitCategory struct {
	md.ModuleCategory
	XMLID string `xml:"id,attr"`
}

// InitModuleCategory2DB 初始化用户数据
func InitModuleCategory2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initModuleCategories InitModuleCategories
			var moduleName = "ModuleCategory"
			ormObj := orm.NewOrm()
			if xml.Unmarshal(data, &initModuleCategories) == nil {
				var user md.User
				user.ID = 0
				user.IsAdmin = true
				for _, categoryXML := range initModuleCategories.Categories {
					var category md.ModuleCategory
					var xmlid = utils.StringsJoin(moduleName, ".", categoryXML.XMLID)
					if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {
						category.Name = categoryXML.Name
						category.Description = categoryXML.Description
						if insertID, err := service.ServiceCreateModuleCategory(&user, &category); err == nil {
							var moduleData md.ModuleData
							moduleData.InsertID = insertID
							moduleData.XMLID = xmlid
							moduleData.Descrition = category.Name
							moduleData.ModuleName = moduleName
							md.AddModuleData(&moduleData, ormObj)
						}
					}
				}
			}
		}
	}
}
