package inital_action

import (
	"encoding/xml"
	md "golangERP/models"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/orm"
)

// InitCountries 国家数据解析
type InitCountries struct {
	XMLName   xml.Name            `xml:"Countries"`
	Countries []md.AddressCountry `xml:"country"`
}

// InitCountry2DB 初始化国家数据
func InitCountry2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCountries InitCountries
			if xml.Unmarshal(data, &initCountries) == nil {
				ormObj := orm.NewOrm()
				countries := make([]*md.AddressCountry, 5)
				var lastIndex int
				for i, country := range initCountries.Countries {
					countries[i] = &country
					lastIndex = i
				}
				countries = countries[0:lastIndex]
				md.BatchAddAddressCountry(countries, ormObj)
			}
		}
	}
}

// InitProvince  省份数据解析
type InitProvince struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"ProvinceName,attr"`
	PID  uint   `xml:"PID,attr"`
}

// InitProvinces 省份数据列表
type InitProvinces struct {
	XMLName   xml.Name       `xml:"Provinces"`
	Provinces []InitProvince `xml:"Province"`
}

// InitProvince2DB 初始化省份数据
func InitProvince2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initProvinces InitProvinces
			if xml.Unmarshal(data, &initProvinces) == nil {
				ormObj := orm.NewOrm()
				provinces := make([]*md.AddressProvince, 50)
				var lastIndex int
				for i, provinceXML := range initProvinces.Provinces {
					lastIndex = i
					var province md.AddressProvince
					var country md.AddressCountry
					pid := int64(provinceXML.PID)
					country.ID = pid
					province.Country = &country
					province.Name = provinceXML.Name
					provinces[i] = &province
				}

				provinces = provinces[0:lastIndex]
				md.BatchAddAddressProvince(provinces, ormObj)
			}
		}
	}
}

// InitCity 初始化城市数据
type InitCity struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"CityName,attr"`
	PID  uint   `xml:"PID,attr"`
}

// InitCities 初始化城市数据列表
type InitCities struct {
	XMLName xml.Name   `xml:"Cities"`
	Cities  []InitCity `xml:"City"`
}

// InitCity2DB 初始化省份城市数据
func InitCity2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCities InitCities
			if xml.Unmarshal(data, &initCities) == nil {
				ormObj := orm.NewOrm()
				cities := make([]*md.AddressCity, 1000)
				var lastIndex int
				for i, cityXML := range initCities.Cities {
					var city md.AddressCity
					var province md.AddressProvince
					pid := int64(cityXML.PID)
					province.ID = pid
					city.Province = &province
					city.Name = cityXML.Name
					cities[i] = &city
					lastIndex = i
				}
				cities = cities[0:lastIndex]
				md.BatchAddAddressCity(cities, ormObj)
			}
		}
	}
}

// InitDistrict 初始化区县数据
type InitDistrict struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"DistrictName,attr"`
	PID  uint   `xml:"CID,attr"`
}

// InitDistricts 初始化区县数据列表
type InitDistricts struct {
	XMLName   xml.Name       `xml:"Districts"`
	Districts []InitDistrict `xml:"District"`
}

// InitDistrict2DB 初始化省份城市数据
func InitDistrict2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initDistricts InitDistricts
			if xml.Unmarshal(data, &initDistricts) == nil {
				ormObj := orm.NewOrm()
				districtes := make([]*md.AddressDistrict, 10000)
				var lastIndex int
				for i, districtXML := range initDistricts.Districts {
					var district md.AddressDistrict
					var city md.AddressCity
					pid := int64(districtXML.PID)
					city.ID = pid
					district.City = &city
					district.Name = districtXML.Name
					districtes[i] = &district
					lastIndex = i
				}
				districtes = districtes[0:lastIndex]
				md.BatchAddAddressDistrict(districtes, ormObj)
			}
		}
	}
}
