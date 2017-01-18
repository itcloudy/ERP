package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Sequrence  表序列号管理，用于销售订单号，采购订单号等
type Sequrence struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`              //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string    `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name       string    `orm:"unique" xml:"name"`                    //表名称组名称
	Prefix     string    //前缀
	Current    int64     //当前序号
	Padding    int64     //序列位数
}

func init() {
	orm.RegisterModel(new(Sequrence))
}

// AddSequrence insert a new Sequrence into database and returns
// last inserted ID on success.
func AddSequrence(obj *Sequrence) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetSequrenceByID retrieves Sequrence by ID. Returns error if
// ID doesn't exist
func GetSequrenceByID(id int64) (obj *Sequrence, err error) {
	o := orm.NewOrm()
	obj = &Sequrence{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetSequrenceByName retrieves Sequrence by Name. Returns error if
// Name doesn't exist
func GetSequrenceByName(name string) (obj *Sequrence, err error) {
	o := orm.NewOrm()
	obj = &Sequrence{Name: name}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}
