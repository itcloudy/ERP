package utils

import (
	"reflect"
	"strconv"
	"strings"
)

var sliceOfInts = reflect.TypeOf([]int(nil))
var sliceOfStrings = reflect.TypeOf([]string(nil))

// JSONreflectStruct json数据反射为结构体
func JSONreflectStruct(jsonData map[string]interface{}, obj interface{}) (err error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	objT = objT.Elem()
	objV = objV.Elem()
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		fieldV := objV.Field(i)
		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}

		jsonValue := jsonData[tag]
		if jsonValue == nil {
			continue
		}

		if fieldT.Type.Kind() != reflect.Ptr {
			switch fieldT.Type.Kind() {
			case reflect.Bool:
				value := jsonValue.(string)
				if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
					fieldV.SetBool(true)
					continue
				}
				if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
					fieldV.SetBool(false)
					continue
				}
				b, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				fieldV.SetBool(b)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				x, err := ToInt64(jsonValue)
				if err != nil {
					return err
				}
				fieldV.SetInt(x)
			case reflect.Float32, reflect.Float64:
				value := jsonValue.(string)
				x, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				fieldV.SetFloat(x)
			case reflect.Interface:
				fieldV.Set(reflect.ValueOf(jsonValue))
			case reflect.String:
				value := jsonValue.(string)
				fieldV.SetString(value)
			}

		}
	}
	return
}
