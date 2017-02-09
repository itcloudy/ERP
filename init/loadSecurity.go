package init

import (
	"encoding/xml"
	md "goERP/models"
	"io/ioutil"
	"os"
	"runtime"
)

// LoadSecurity 权限控制加载
func LoadSecurity() {
	// 加载系统资源
	systemType := runtime.GOOS
	split := "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	if xmDir, err := os.Getwd(); err == nil {
		xmDir += split + "security" + split
		loadSources(xmDir + "sources.xml")
		loadMenus(xmDir + "menus.xml")

	}
}
func loadSources(filename string) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var sources InitSources
			if xml.Unmarshal(data, &sources) == nil {
				for _, k := range sources.Sources {
					user := new(md.User)
					user.ID = 1
					if obj, err := md.GetSourceByModelName(k.ModelName); err != nil {
						if obj.ID == 0 {
							md.AddSource(&k, user)
						}
					}
				}
			}
		}
	}
}
func loadMenus(filename string) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var menus InitMenus
			if xml.Unmarshal(data, &menus) == nil {
				for _, k := range menus.Menus {
					user := new(md.User)
					user.ID = 1
					if obj, err := md.GetMenuByIdentity(k.Identity); err != nil {
						if obj.ID == 0 {
							md.AddMenu(&k, user)
						}
					}
				}
			}
		}
	}
}
