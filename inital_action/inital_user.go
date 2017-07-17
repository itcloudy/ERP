package inital_action

import (
	"encoding/xml"
	md "golangERP/models"
	service "golangERP/services"
	"io/ioutil"
	"os"
)

type InitUsers struct {
	XMLName xml.Name  `xml:"Users"`
	Users   []md.User `xml:"user"`
}

// InitUser2DB 初始化用户数据
func InitUser2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initUsers InitUsers
			if xml.Unmarshal(data, &initUsers) == nil {
				for _, user := range initUsers.Users {
					service.ServiceCreateUser(&user)
				}
			}
		}
	}
}
