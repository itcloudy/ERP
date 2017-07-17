package inital_action

import (
	"bytes"
	"os"
	"runtime"
)

// InitApp 基础数据插入
func InitApp() {
	systemType := runtime.GOOS
	split := "/"
	switch systemType {
	case "windows":
		split = "\\"
	case "linux":
		split = "/"
	}
	if xmlDir, err := os.Getwd(); err == nil {
		b := bytes.Buffer{}
		b.WriteString(xmlDir)
		b.WriteString(split)
		b.WriteString("inital_data")
		b.WriteString(split)
		b.WriteString("xml")
		xmlBase := b.String()
		countryXML := xmlBase + split + "address" + split + "Countries.xml"
		go InitCountry2DB(countryXML)
		provinceXML := xmlBase + split + "address" + split + "Provinces.xml"
		go InitProvince2DB(provinceXML)
		cityXML := xmlBase + split + "address" + split + "Cities.xml"
		go InitCity2DB(cityXML)
		districtXML := xmlBase + split + "address" + split + "Districts.xml"
		go InitDistrict2DB(districtXML)

	}
}
