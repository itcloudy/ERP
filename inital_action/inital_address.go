package inital_action

import (
	"encoding/xml"
	md "golangERP/models"
	service "golangERP/services"
	"io/ioutil"
	"os"
)

// InitCountries 国家数据解析
type InitCountries struct {
	XMLName   xml.Name            `xml:"Countries"`
	Countries []md.AddressCountry `xml:"country"`
}

// InitCountry 初始化国家数据
func InitCountry2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCountries InitCountries
			if xml.Unmarshal(data, &initCountries) == nil {
				for _, country := range initCountries.Countries {
					service.ServiceCreateAddressCountry(&country)
				}
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
				for _, provinceXML := range initProvinces.Provinces {
					var province md.AddressProvince
					var country md.AddressCountry
					pid := int64(provinceXML.PID)
					country.ID = pid
					province.Country = &country
					province.Name = provinceXML.Name
					service.ServiceCreateAddressProvince(&province)
				}
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
				for _, cityXML := range initCities.Cities {
					var city md.AddressCity
					var province md.AddressProvince
					pid := int64(cityXML.PID)
					province.ID = pid
					city.Province = &province
					city.Name = cityXML.Name
					service.ServiceCreateAddressCity(&city)
				}
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
				for _, districtXML := range initDistricts.Districts {
					var district md.AddressDistrict
					var city md.AddressCity
					pid := int64(districtXML.PID)
					city.ID = pid
					district.City = &city
					district.Name = districtXML.Name
					service.ServiceCreateAddressDistrict(&district)
				}
			}
		}
	}
}
