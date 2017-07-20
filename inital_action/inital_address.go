package inital_action

import (
	"encoding/xml"
	md "golangERP/models"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/orm"
)

// InitCountry  国家数据解析
type InitCountry struct {
	ID    uint   `xml:"ID,attr"`
	Name  string `xml:"name"`
	XMLID uint   `xml:"xml_id,attr"`
}

// InitCountries 国家数据列表
type InitCountries struct {
	XMLName   xml.Name      `xml:"Countries"`
	Countries []InitCountry `xml:"country"`
}

// InitCountry2DB 初始化国家数据
func InitCountry2DB(filePath string) {
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err == nil {
			var initCountries InitCountries
			if xml.Unmarshal(data, &initCountries) == nil {
				ormObj := orm.NewOrm()
				for _, countryXML := range initCountries.Countries {
					var country md.AddressCountry
					country.Name = countryXML.Name
					md.AddAddressCountry(&country, ormObj)
				}
			}
		}
	}
}

// InitProvince  省份数据解析
type InitProvince struct {
	ID    uint   `xml:"ID,attr"`
	Name  string `xml:"ProvinceName,attr"`
	PID   uint   `xml:"PID,attr"`
	XMLID uint   `xml:"xml_id,attr"`
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
				for _, provinceXML := range initProvinces.Provinces {
					fmt.println(provinceXML)
					var province md.AddressProvince
					var country md.AddressCountry
					pid := int64(provinceXML.PID)
					country.ID = pid
					province.Country = &country
					province.Name = provinceXML.Name
					md.AddAddressProvince(&province, ormObj)
				}

			}
		}
	}
}

// InitCity 初始化城市数据
type InitCity struct {
	ID    uint   `xml:"ID,attr"`
	Name  string `xml:"CityName,attr"`
	PID   uint   `xml:"PID,attr"`
	XMLID uint   `xml:"xml_id,attr"`
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
				for i, cityXML := range initCities.Cities {
					var city md.AddressCity
					var province md.AddressProvince
					pid := int64(cityXML.PID)
					province.ID = pid
					city.Province = &province
					city.Name = cityXML.Name
					md.AddAddressCity(&city, ormObj)
				}
			}
		}
	}
}

// InitDistrict 初始化区县数据
type InitDistrict struct {
	ID    uint   `xml:"ID,attr"`
	Name  string `xml:"DistrictName,attr"`
	PID   uint   `xml:"CID,attr"`
	XMLID uint   `xml:"xml_id,attr"`
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
				for i, districtXML := range initDistricts.Districts {
					var district md.AddressDistrict
					var city md.AddressCity
					pid := int64(districtXML.PID)
					city.ID = pid
					district.City = &city
					district.Name = districtXML.Name
					md.AddAddressDistrict(district, ormObj)
				}
			}
		}
	}
}
