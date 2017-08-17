package initalActions

import (
	"encoding/xml"
	md "golangERP/models"
	service "golangERP/services"
	"golangERP/utils"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego/orm"
)

// InitUsers 用户数据列表
type InitUsers struct {
	XMLName xml.Name   `xml:"Users"`
	Users   []InitUser `xml:"user"`
}

// InitUser 用户数据解析
type InitUser struct {
	md.User
	XMLID  string `xml:"id,attr"`
	Groups string `xml:"group"`
}

// InitUser2DB 初始化用户数据
func InitUser2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initUsers InitUsers
			var moduleName = "User"
			ormObj := orm.NewOrm()
			if xml.Unmarshal(data, &initUsers) == nil {
				var createU md.User
				createU.ID = 0
				createU.IsAdmin = true
				for _, userXML := range initUsers.Users {
					var user md.User
					var xmlid = utils.StringsJoin(moduleName, ".", userXML.XMLID)
					if _, err = md.GetModuleDataByXMLID(xmlid, ormObj); err != nil {
						user.Name = userXML.Name
						user.NameZh = userXML.NameZh
						user.Email = userXML.Email
						user.Mobile = userXML.Mobile
						user.Password = userXML.Password
						user.IsAdmin = userXML.IsAdmin
						user.Active = userXML.Active
						user.Qq = userXML.Qq
						user.WeChat = userXML.WeChat
						user.IsBackground = userXML.IsBackground
						if insertID, err := service.ServiceCreateUser(&createU, &user); err == nil {
							var moduleData md.ModuleData
							moduleData.InsertID = insertID
							moduleData.XMLID = xmlid
							moduleData.Descrition = user.Name
							moduleData.ModuleName = moduleName
							md.AddModuleData(&moduleData, ormObj)
							//权限设置
							groupStr := userXML.Groups
							if groupStr != "" {
								groups := strings.Split(groupStr, ",")
								for _, groupName := range groups {
									if groupName != "" {
										if group, err := md.GetBaseGroupByName(groupName, ormObj); err == nil {
											var groupUser md.GroupUser
											var us md.User
											us.ID = insertID
											groupUser.User = &us
											groupUser.Group = group
											md.AddGroupUser(&groupUser, ormObj)
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
