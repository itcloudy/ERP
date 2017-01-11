package models

import "time"

// base struct，数据库中的表基于此表，该struct不会生成数据库中对应的表结构
type Base struct {
	Id         int64     `orm:"pk;auto" json:"id"`                    //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string    `orm:"-" form:"formAction"`                  //非数据库字段，用于表示记录的增加，修改
}
