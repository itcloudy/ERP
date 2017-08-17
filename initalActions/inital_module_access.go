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

// InitModuleAccesses 模块权限数据
type InitModuleAccesses struct {
	XMLName     xml.Name           `xml:"Permissions"`
	Permissions []InitModuleAccess `xml:"permission"`
}

// InitModuleAccess 模块数据解析
type InitModuleAccess struct {
	XMLID      string `xml:"id,attr"`
	Module     string `xml:"module"`
	PermCreate bool   `xml:"create"` //创建权限
	PermUnlink bool   `xml:"unlink"` //删除权限
	PermWrite  bool   `xml:"write"`  //修改权限
	PermRead   bool   `xml:"read"`   //读权限
}

// InitModelAccess2DB 初始化模块数据
func InitModelAccess2DB(split string) {
	if xmlDir, err := os.Getwd(); err == nil {
		xmlBase := utils.StringsJoin(xmlDir, split, "inital_data", split, "xml", split, "permission")

		var moduleName = "ModuleAccess"
		if dirList, err := ioutil.ReadDir(xmlBase); err == nil {
			var user md.User
			user.ID = 0
			user.IsAdmin = true
			for _, dir := range dirList {
				if dir.IsDir() {
					var groupStr = dir.Name()
					if subDirList, err := ioutil.ReadDir(utils.StringsJoin(xmlBase, split, groupStr)); err == nil {
						for _, subDir := range subDirList {
							if subDir.IsDir() {
								continue
							}
							if file, err := os.Open(utils.StringsJoin(xmlBase, split, dir.Name(), split, subDir.Name())); err == nil {
								defer file.Close()
								if data, err := ioutil.ReadAll(file); err == nil {

									ormObj := orm.NewOrm()
									var initModuleAccesses InitModuleAccesses
									if xml.Unmarshal(data, &initModuleAccesses) == nil {
										for _, moduleXML := range initModuleAccesses.Permissions {
											var module md.ModelAccess
											var xmlid = utils.StringsJoin(moduleName, ".", groupStr, ".", moduleXML.XMLID)
											if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {
												module.Module, _ = md.GetModuleModuleByName(moduleXML.Module, ormObj)
												module.Group, _ = md.GetBaseGroupByName(groupStr, ormObj)
												module.PermCreate = moduleXML.PermCreate
												module.PermRead = moduleXML.PermRead
												module.PermUnlink = moduleXML.PermUnlink
												module.PermWrite = moduleXML.PermWrite
												if insertID, err := service.ServiceCreateModelAccess(&user, &module); err == nil {
													var moduleData md.ModuleData
													moduleData.InsertID = insertID
													moduleData.XMLID = xmlid
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
		}
	}

}
