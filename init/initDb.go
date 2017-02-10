//数据库部分数据初始化
package init

import (
	"encoding/xml"
	md "goERP/models"
	"io/ioutil"
	"os"
	"runtime"
)

//初始化数据库
func InitDb() {
	systemType := runtime.GOOS
	split := "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	if xmDir, err := os.Getwd(); err == nil {
		if _, err := md.GetUserByID(1); err != nil {

			xmDir += split + "init_xml" + split
			initUser(xmDir + "Users.xml")
			if user, err := md.GetUserByID(1); err == nil {
				initCountry(xmDir+"Countries.xml", user)
				initProvince(xmDir+"Provinces.xml", user)
				initCity(xmDir+"Cities.xml", user)
				initDistrict(xmDir+"Districts.xml", user)
				initDistrict(xmDir+"Sequence.xml", user)
			}
		}
	}

}
func initSequence(filename string, user *md.User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initSequences InitSequences
			if xml.Unmarshal(data, &initSequences) == nil {
				for _, k := range initSequences.Sequence {
					md.AddSequence(&k, user)
				}
			}
		}
	}
}
func initUser(filename string) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initUsers InitUsers
			var user md.User
			user.ID = 1
			if xml.Unmarshal(data, &initUsers) == nil {
				for _, k := range initUsers.Users {
					//admin系统管理员
					md.AddUser(&k, &user)
				}
			}
		}
	}

}

func initCountry(filename string, user *md.User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCountries InitCountries
			if xml.Unmarshal(data, &initCountries) == nil {
				for _, k := range initCountries.Countries {
					md.AddAddressCountry(&k, user)
				}
			}
		}
	}
}
func initProvince(filename string, user *md.User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initProvinces InitProvinces
			if xml.Unmarshal(data, &initProvinces) == nil {
				for _, k := range initProvinces.Provinces {
					var province md.AddressProvince
					pid := int64(k.PID)
					if country, err := md.GetAddressCountryByID(pid); err == nil {
						province.Country = country
						province.Name = k.Name
						md.AddAddressProvince(&province, user)
					}
				}
			}
		}
	}
}
func initCity(filename string, user *md.User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCities InitCities
			if xml.Unmarshal(data, &initCities) == nil {
				for _, k := range initCities.Cities {
					var city md.AddressCity
					pid := int64(k.PID)
					if province, e := md.GetAddressProvinceByID(pid); e == nil {
						city.Province = province
						city.Name = k.Name
						md.AddAddressCity(&city, user)
					}
				}
			}
		}
	}
}
func initDistrict(filename string, user *md.User) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initDistricts InitDistricts
			if xml.Unmarshal(data, &initDistricts) == nil {
				for _, k := range initDistricts.Districts {

					var (
						district md.AddressDistrict
					)
					pid := int64(k.PID)
					if city, e := md.GetAddressCityByID(pid); e == nil {
						district.City = city
						district.Name = k.Name
						md.AddAddressDistrict(&district, user)
					}
				}
			}
		}
	}
}
