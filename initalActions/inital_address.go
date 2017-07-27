package initalActions

import (
	"encoding/xml"
	md "golangERP/models"
	"io/ioutil"
	"os"

	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// InitCountry  国家数据解析
type InitCountry struct {
	ID    uint   `xml:"ID,attr"`
	Name  string `xml:"name"`
	XMLID string `xml:"xml_id,attr"`
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
					if insertID, err := md.AddAddressCountry(&country, ormObj); err == nil {
						var moduleData md.ModuleData
						moduleData.InsertID = insertID
						moduleData.XMLID = utils.StringsJoin("AddressCountry.", countryXML.XMLID)
						moduleData.Descrition = country.Name
						// moduleData.ModuleName = reflect.Indirect(reflect.ValueOf(country)).Type().Name()
						moduleData.ModuleName = "AddressCountry"
						md.AddModuleData(&moduleData, ormObj)
					}
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
	XMLID string `xml:"xml_id,attr"`
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
					var province md.AddressProvince
					var country md.AddressCountry
					pid := int64(provinceXML.PID)
					country.ID = pid
					province.Country = &country
					province.Name = provinceXML.Name
					if insertID, err := md.AddAddressProvince(&province, ormObj); err == nil {
						var moduleData md.ModuleData
						moduleData.InsertID = insertID
						moduleData.XMLID = utils.StringsJoin("AddressProvince.", provinceXML.XMLID)
						moduleData.Descrition = province.Name
						// moduleData.ModuleName = reflect.Indirect(reflect.ValueOf(province)).Type().Name()
						moduleData.ModuleName = "AddressProvince"
						md.AddModuleData(&moduleData, ormObj)
					}
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
	XMLID string `xml:"xml_id,attr"`
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
				for _, cityXML := range initCities.Cities {
					var city md.AddressCity
					var province md.AddressProvince
					pid := int64(cityXML.PID)
					province.ID = pid
					city.Province = &province
					city.Name = cityXML.Name
					if insertID, err := md.AddAddressCity(&city, ormObj); err == nil {
						var moduleData md.ModuleData
						moduleData.InsertID = insertID
						moduleData.XMLID = utils.StringsJoin("AddressCity.", cityXML.XMLID)
						moduleData.Descrition = city.Name
						// moduleData.ModuleName = reflect.Indirect(reflect.ValueOf(city)).Type().Name()
						moduleData.ModuleName = "AddressCity"
						md.AddModuleData(&moduleData, ormObj)
					}
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
	XMLID string `xml:"xml_id,attr"`
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
				for _, districtXML := range initDistricts.Districts {
					var district md.AddressDistrict
					var city md.AddressCity
					pid := int64(districtXML.PID)
					city.ID = pid
					district.City = &city
					district.Name = districtXML.Name
					if insertID, err := md.AddAddressDistrict(&district, ormObj); err == nil {
						var moduleData md.ModuleData
						moduleData.InsertID = insertID
						moduleData.XMLID = utils.StringsJoin("AddressDistrict.", districtXML.XMLID)
						moduleData.Descrition = district.Name
						// moduleData.ModuleName = reflect.Indirect(reflect.ValueOf(district)).Type().Name()
						moduleData.ModuleName = "AddressDistrict"
						md.AddModuleData(&moduleData, ormObj)
					}
				}
			}
		}
	}
}
