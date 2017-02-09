package init

import (
	"encoding/xml"
	md "goERP/models"
)

// InitSources 资源标识符
type InitSources struct {
	XMLName xml.Name    `xml:"Sources"`
	Sources []md.Source `xml:"source"`
}

// InitMenus 资源标识符
type InitMenus struct {
	XMLName xml.Name  `xml:"Menus"`
	Menus   []md.Menu `xml:"menu"`
}
type InitUsers struct {
	XMLName xml.Name  `xml:"Users"`
	Users   []md.User `xml:"user"`
}
type InitCountries struct {
	XMLName   xml.Name            `xml:"Countries"`
	Countries []md.AddressCountry `xml:"country"`
}

type InitProvince struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"ProvinceName,attr"`
	PID  uint   `xml:"PID,attr"`
}
type InitProvinces struct {
	XMLName   xml.Name       `xml:"Provinces"`
	Provinces []InitProvince `xml:"Province"`
}
type InitCity struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"CityName,attr"`
	PID  uint   `xml:"PID,attr"`
}
type InitCities struct {
	XMLName xml.Name   `xml:"Cities"`
	Cities  []InitCity `xml:"City"`
}

type InitDistrict struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"DistrictName,attr"`
	PID  uint   `xml:"CID,attr"`
}
type InitDistricts struct {
	XMLName   xml.Name       `xml:"Districts"`
	Districts []InitDistrict `xml:"District"`
}
type InitSequences struct {
	XMLName  xml.Name      `xml:"Sequences"`
	Sequence []md.Sequence `xml:"Sequence"`
}
