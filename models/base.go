package models

import "time"

// Base struct，数据库中的表基于此表，该struct不会生成数据库中对应的表结构
type Base struct {
	ID         int64     `orm:"column(id);pk" json:"id"`              //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string    `orm:"-" form:"formAction"`                  //非数据库字段，用于表示记录的增加，修改
}

// QueryMap struct
type QueryMap struct {
	Type    string                 //用户数据库查询的类型:filter,exclude,cond等
	Context map[string]interface{} //数据库查询的条件
}
