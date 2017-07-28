package initalActions

import (
	"encoding/xml"
	"fmt"
	md "golangERP/models"
	service "golangERP/services"
	"golangERP/utils"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/orm"
)

// InitMenu  菜单数据解析
type InitMenu struct {
	Name      string `xml:"name"`
	XMLID     string `xml:"id,attr"`
	Path      string `xml:"path"`
	Icon      string `xml:"icon"`
	Component string `xml:"component"`
	Sequence  int64  `xml:"sequence"`
	ParentID  string `xml:"parent_id,attr"`
}

// InitMenus 国家数据列表
type InitMenus struct {
	XMLName xml.Name   `xml:"Menus"`
	Menus   []InitMenu `xml:"menu"`
}

// InitMenus2DB 菜单初始化，数据库创建记录
func InitMenus2DB(split string) {
	if xmlDir, err := os.Getwd(); err == nil {
		xmlBase := utils.StringsJoin(xmlDir, split, "inital_data", split, "xml", split, "menu")
		if dirList, err := ioutil.ReadDir(xmlBase); err == nil {
			for _, dir := range dirList {
				if dir.IsDir() {
					continue
				}
				if file, err := os.Open(utils.StringsJoin(xmlBase, split, dir.Name())); err == nil {
					defer file.Close()
					if data, err := ioutil.ReadAll(file); err == nil {
						var initMenus InitMenus
						if xml.Unmarshal(data, &initMenus) == nil {
							ormObj := orm.NewOrm()
							for _, menuXML := range initMenus.Menus {
								var menu md.BaseMenu
								var parent md.BaseMenu
								menu.Name = menuXML.Name
								menu.Path = menuXML.Path
								menu.Component = menuXML.Component
								menu.Icon = menuXML.Icon
								menu.Sequence = menuXML.Sequence
								parentIDStr := menuXML.ParentID
								if parentIDStr != "" {
									if mobuleData, err := md.GetModuleDataByXMLID(utils.StringsJoin("BaseMenu.", parentIDStr), ormObj); err == nil {
										parent.ID = mobuleData.InsertID
										fmt.Print(mobuleData)
										menu.Parent = &parent
									}
								}
								if insertID, err := service.ServiceCreateBaseMenu(&menu); err == nil {
									var moduleData md.ModuleData
									moduleData.InsertID = insertID
									moduleData.XMLID = utils.StringsJoin("BaseMenu.", menuXML.XMLID)
									moduleData.Descrition = menu.Name
									moduleData.ModuleName = "BaseMenu"
									md.AddModuleData(&moduleData, ormObj)
								} else {
									fmt.Println("service create result ")
									fmt.Println(err)
								}

							}
						}
					}

				}
			}
		}
	}
}
