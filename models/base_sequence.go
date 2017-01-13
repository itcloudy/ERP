package models

import "github.com/astaxie/beego/orm"

//Sequrence  表序列号管理，用于销售订单号，采购订单号等
type Sequrence struct {
	Base
	Name    string `orm:"unique" xml:"name"` //表名称组名称
	Prefix  string //前缀
	Current int64  //当前序号
	Padding int64  //序列位数
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
	obj = &Sequrence{Base: Base{ID: id}}
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
