package inital_action

import (
	"encoding/xml"
	"fmt"
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
	fmt.Println(filePath)
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
	fmt.Println(filePath)
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
