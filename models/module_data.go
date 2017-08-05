package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// ModuleData xml初始化数据记录
type ModuleData struct {
	ID           int64     `orm:"column(id);pk;auto"` //主键
	CreateUserID int64     `orm:"column(create_user_id);null"`         //创建者
	UpdateUserID int64     `orm:"column(update_user_id);null"`         //最后更新者
	CreateDate   time.Time `orm:"auto_now_add;type(datetime)"`         //创建时间
	UpdateDate   time.Time `orm:"auto_now;type(datetime)"`             //最后更新时间
	XMLID        string    `orm:"column(xml_id);unique;index"`                  //xml文件中的id
	Data         string    `orm:"null"`                                         //数据内容
	Descrition   string    `orm:"null"`                                         //记录描述
	InsertID     int64     `orm:"column(insert_id)"`                            //插入记录的ID
	ModuleName   string    `orm:""`                                             //模块(表)的名称

}

func init() {
	orm.RegisterModel(new(ModuleData))
}

// AddModuleData insert a new ModuleData into database and returns last inserted Id on success.
func AddModuleData(m *ModuleData, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// GetModuleDataByXMLID get moduledata by xmlid
func GetModuleDataByXMLID(xmlid string, ormObj orm.Ormer) (*ModuleData, error) {
	var obj ModuleData
	var err error
	qs := ormObj.QueryTable(&obj)
	err = qs.Filter("xml_id", xmlid).One(&obj)
	return &obj, err

}
