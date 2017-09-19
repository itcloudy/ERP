package services

import (
	"encoding/json"
	"errors"
	md "golangERP/models"
	"golangERP/utils"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateProductTemplate 创建记录
func ServiceCreateProductTemplate(user *md.User, requestBody []byte) (id int64, err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductTemplate"); err == nil {
		if !access.Create {
			err = errors.New("has no create permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	var obj md.ProductTemplate
	json.Unmarshal([]byte(requestBody), &obj)

	obj.CreateUserID = user.ID
	id, err = md.AddProductTemplate(&obj, o)

	return
}

// ServiceUpdateProductTemplate 更新记录
func ServiceUpdateProductTemplate(user *md.User, requestBody []byte, id int64) (err error) {

	var access utils.AccessResult
	if access, err = ServiceCheckUserModelAssess(user, "ProductTemplate"); err == nil {
		if !access.Update {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	var obj md.ProductTemplate
	var objPtr *md.ProductTemplate
	if objPtr, err = md.GetProductTemplateByID(id, o); err != nil {
		return
	}
	obj = *objPtr
	json.Unmarshal([]byte(requestBody), &obj)

	obj.UpdateUserID = user.ID
	id, err = md.UpdateProductTemplate(&obj, o)

	return
}

//ServiceGetProductTemplate 获得城市列表
func ServiceGetProductTemplate(user *md.User, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckUserModelAssess(user, "ProductTemplate"); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.ProductTemplate
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllProductTemplate(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			objInfo["Name"] = obj.Name
			objInfo["ID"] = obj.ID
			objInfo["Description"] = obj.Description
			objInfo["DescriptionSale"] = obj.DescriptionSale
			objInfo["DescriptionPurchase"] = obj.DescriptionPurchase
			objInfo["Rental"] = obj.Rental
			categoryInfo := make(map[string]interface{})
			categoryInfo["ID"] = obj.Category.ID
			categoryInfo["Name"] = obj.Category.Name
			objInfo["Category"] = categoryInfo
			objInfo["Price"] = obj.Price
			objInfo["StandardPrice"] = obj.StandardPrice
			objInfo["StandardWeight"] = obj.StandardWeight
			objInfo["SaleOk"] = obj.SaleOk
			objInfo["Active"] = obj.Active
			objInfo["IsProductVariant"] = obj.IsProductVariant
			firstSaleUomInfo := make(map[string]interface{})
			firstSaleUomInfo["ID"] = obj.FirstSaleUom.ID
			firstSaleUomInfo["Name"] = obj.FirstSaleUom.Name
			objInfo["FirstSaleUom"] = firstSaleUomInfo
			SecondSaleUomInfo := make(map[string]interface{})
			SecondSaleUomInfo["ID"] = obj.SecondSaleUom.ID
			SecondSaleUomInfo["Name"] = obj.SecondSaleUom.Name
			objInfo["SecondSaleUom"] = SecondSaleUomInfo
			firstPurchaseUomInfo := make(map[string]interface{})
			firstPurchaseUomInfo["ID"] = obj.FirstPurchaseUom.ID
			firstPurchaseUomInfo["Name"] = obj.FirstPurchaseUom.Name
			objInfo["FirstPurchaseUom"] = firstPurchaseUomInfo
			secondPurchaseUomInfo := make(map[string]interface{})
			secondPurchaseUomInfo["ID"] = obj.SecondPurchaseUom.ID
			secondPurchaseUomInfo["Name"] = obj.SecondPurchaseUom.Name
			objInfo["SecondPurchaseUom"] = secondPurchaseUomInfo
			objInfo["VariantCount"] = obj.VariantCount
			objInfo["Barcode"] = obj.Barcode
			objInfo["DefaultCode"] = obj.DefaultCode
			objInfo["ProductType"] = obj.ProductType
			objInfo["ProductMethod"] = obj.ProductMethod
			results = append(results, objInfo)

		}
	}
	return
}

// ServiceGetProductTemplateByID get ProductTemplate by id
func ServiceGetProductTemplateByID(user *md.User, id int64) (access utils.AccessResult, templateInfo map[string]interface{}, err error) {

	if access, err = ServiceCheckUserModelAssess(user, "ProductTemplate"); err == nil {
		if !access.Read {
			err = errors.New("has no update permission")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	var obj *md.ProductTemplate
	if obj, err = md.GetProductTemplateByID(id, o); err == nil {
		objInfo := make(map[string]interface{})
		objInfo["Name"] = obj.Name
		objInfo["ID"] = obj.ID
		objInfo["Description"] = obj.Description
		objInfo["DescriptionSale"] = obj.DescriptionSale
		objInfo["DescriptionPurchase"] = obj.DescriptionPurchase
		objInfo["Rental"] = obj.Rental
		categoryInfo := make(map[string]interface{})
		categoryInfo["ID"] = obj.Category.ID
		categoryInfo["Name"] = obj.Category.Name
		objInfo["Category"] = categoryInfo
		objInfo["Price"] = obj.Price
		objInfo["StandardPrice"] = obj.StandardPrice
		objInfo["StandardWeight"] = obj.StandardWeight
		objInfo["SaleOk"] = obj.SaleOk
		objInfo["Active"] = obj.Active
		objInfo["IsProductVariant"] = obj.IsProductVariant
		firstSaleUomInfo := make(map[string]interface{})
		firstSaleUomInfo["ID"] = obj.FirstSaleUom.ID
		firstSaleUomInfo["Name"] = obj.FirstSaleUom.Name
		objInfo["FirstSaleUom"] = firstSaleUomInfo
		SecondSaleUomInfo := make(map[string]interface{})
		SecondSaleUomInfo["ID"] = obj.SecondSaleUom.ID
		SecondSaleUomInfo["Name"] = obj.SecondSaleUom.Name
		objInfo["SecondSaleUom"] = SecondSaleUomInfo
		firstPurchaseUomInfo := make(map[string]interface{})
		firstPurchaseUomInfo["ID"] = obj.FirstPurchaseUom.ID
		firstPurchaseUomInfo["Name"] = obj.FirstPurchaseUom.Name
		objInfo["FirstPurchaseUom"] = firstPurchaseUomInfo
		secondPurchaseUomInfo := make(map[string]interface{})
		secondPurchaseUomInfo["ID"] = obj.SecondPurchaseUom.ID
		secondPurchaseUomInfo["Name"] = obj.SecondPurchaseUom.Name
		objInfo["SecondPurchaseUom"] = secondPurchaseUomInfo
		objInfo["VariantCount"] = obj.VariantCount
		objInfo["Barcode"] = obj.Barcode
		objInfo["DefaultCode"] = obj.DefaultCode
		objInfo["ProductType"] = obj.ProductType
		objInfo["ProductMethod"] = obj.ProductMethod
		lenAttLine := len(obj.AttributeLines)
		attributeLines := make([]map[string]interface{}, lenAttLine, lenAttLine)
		for i := 0; i < lenAttLine; i++ {
			attLine := obj.AttributeLines[i]
			lineInfo := make(map[string]interface{})
			lineInfo["ID"] = attLine.ID
			valueInfo := make(map[string]interface{})
			valueInfo["ID"] = attLine.Attribute.ID
			valueInfo["Name"] = attLine.Attribute.Name
			lineInfo["Attribute"] = valueInfo
			lenV := len(attLine.AttributeValues)
			values := make([]map[string]interface{}, lenV, lenV)
			for k := 0; k < lenV; k++ {
				attValueInfo := make(map[string]interface{})
				attValueInfo["ID"] = attLine.AttributeValues[k].ID
				attValueInfo["Name"] = attLine.AttributeValues[k].Name
				values[k] = attValueInfo
			}
			lineInfo["AttributeValues"] = values
			attributeLines[i] = lineInfo
		}
		objInfo["attributeLines"] = attributeLines
		templateInfo = objInfo
	}
	return
}
