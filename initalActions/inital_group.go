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

// InitGroup  权限组数据解析
type InitGroup struct {
	XMLID       string `xml:"id,attr"`
	Childs      string `xml:"childs"`
	Parent      string `xml:"parent"`
	Category    string `xml:"category"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
}

// InitGroups 权限组数据列表
type InitGroups struct {
	XMLName xml.Name    `xml:"Groups"`
	Groups  []InitGroup `xml:"group"`
}

// InitGroup2DB 菜单初始化，数据库创建记录
func InitGroup2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initGroups InitGroups
			if xml.Unmarshal(data, &initGroups) == nil {
				ormObj := orm.NewOrm()
				var moduleName = "BaseGroup"
				var user md.User
				user.ID = 0
				user.IsAdmin = true
				for _, groupXML := range initGroups.Groups {
					var xmlid = utils.StringsJoin(moduleName, ".", groupXML.XMLID)
					// 检查在系统中是否已经存在
					if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {
						var group md.BaseGroup
						var parent md.BaseGroup
						group.Name = groupXML.Name
						group.Category = groupXML.Category
						group.Description = groupXML.Description
						parentIDStr := groupXML.Parent
						if parentIDStr != "" {
							if mobuleData, err := md.GetModuleDataByXMLID(utils.StringsJoin(moduleName, ".", parentIDStr), ormObj); err == nil {
								parent.ID = mobuleData.InsertID
								group.Parent = &parent
							}
						}
						if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {

							if insertID, err := service.ServiceCreateBaseGroup(&user, &group); err == nil {
								var moduleData md.ModuleData
								moduleData.InsertID = insertID
								moduleData.XMLID = xmlid
								moduleData.Descrition = group.Name
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
