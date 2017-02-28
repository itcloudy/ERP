package models

import (
	"errors"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// User table
// 用户表
type User struct {
	ID              int64       `orm:"column(id);pk;auto" json:"id"`                      //主键
	CreateUser      *User       `orm:"rel(fk);null" json:"-"`                             //创建者
	UpdateUser      *User       `orm:"rel(fk);null" json:"-"`                             //最后更新者
	CreateDate      time.Time   `orm:"auto_now_add;type(datetime)" json:"-"`              //创建时间
	UpdateDate      time.Time   `orm:"auto_now;type(datetime)" json:"-"`                  //最后更新时间
	Name            string      `orm:"size(20)" xml:"name" json:"Name"`                   //用户名
	Company         *Company    `orm:"rel(fk);null" json:"-"`                             //公司
	NameZh          string      `orm:"size(20)"  xml:"NameZh" json:"NameZh"`              //中文用户名
	Department      *Department `orm:"rel(fk);null;" json:"-"`                            //部门
	Email           string      `orm:"size(20)" xml:"email" json:"Email"`                 //邮箱
	Mobile          string      `orm:"size(20);default(\"\")" xml:"mobile" json:"Mobile"` //手机号码
	Tel             string      `orm:"size(20);default(\"\")" json:"Tel" `                //固定号码
	Password        string      `xml:"password" json:"Password"`                          //密码
	ConfirmPassword string      `orm:"-" xml:"ConfirmPassword" json:"ConfirmPassword"`    //确认密码,数据库中不保存
	Roles           []*Role     `orm:"rel(m2m)"`                                          //用户拥有的角色
	Teams           []*Team     `orm:"rel(m2m)"`                                          //团队
	Groups          []*Group    `orm:"rel(m2m)"`                                          //用户组
	IsAdmin         bool        `orm:"default(false)" xml:"isAdmin" json:"IsAdmin"`       //是否为超级用户
	Active          bool        `orm:"default(true)" xml:"active" json:"Active"`          //有效
	Qq              string      `orm:"default()" xml:"qq" json:"Qq"`                      //QQ
	WeChat          string      `orm:"default()" xml:"wechat" json:"WeChat"`              //微信
	Position        *Position   `orm:"rel(fk);null;" json:"-" `                           //职位

	FormAction   string             `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string           `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
	DepartmentID int64              `orm:"-" json:"Department"`
	CompanyID    int64              `orm:"-" json:"Company"`
	PositionID   int64              `orm:"-" json:"Position"`
	TeamIDs      map[string][]int64 `orm:"-" json:"TeamIds"`
	RoleIDs      map[string][]int64 `orm:"-" json:"RoleIds"`
}

func init() {
	orm.RegisterModel(new(User))
}

// TableName 表名
func (u *User) TableName() string {
	return "base_user"
}

// AddUser insert a new User into database and returns
// last inserted ID on success.
func AddUser(obj *User, addUser *User) (id int64, err error) {

	o := orm.NewOrm()
	errBegin := o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return 0, errBegin
	}
	obj.CreateUser = addUser
	obj.UpdateUser = addUser
	password := utils.PasswordMD5(obj.Password, obj.Mobile)
	obj.Password = password
	if obj.CompanyID > 0 {
		obj.Company, _ = GetCompanyByID(obj.CompanyID)
	}
	if obj.DepartmentID > 0 {
		obj.Department, _ = GetDepartmentByID(obj.DepartmentID)
	}
	if obj.PositionID > 0 {
		obj.Position, _ = GetPositionByID(obj.PositionID)
	}
	if id, err = o.Insert(obj); err == nil {
		obj.ID = id
		if createTeamRecords, ok := obj.TeamIDs["create"]; ok {
			m2mTeams := o.QueryM2M(obj, "Teams")
			for _, teamID := range createTeamRecords {
				if team, err := GetTeamByID(int64(teamID)); err == nil {
					m2mTeams.Add(team)
				} else {
					utils.LogOut("error", "add user teams failed:"+err.Error())
				}
			}
		}
		if createRoleRecords, ok := obj.RoleIDs["create"]; ok {
			m2mRoles := o.QueryM2M(obj, "Roles")
			for _, RoleID := range createRoleRecords {
				if role, err := GetRoleByID(int64(RoleID)); err == nil {
					m2mRoles.Add(role)
				} else {
					utils.LogOut("error", "add user roles failed:"+err.Error())
				}
			}
		}
	}
	if err != nil {
		return 0, err
	}
	errCommit := o.Commit()
	if errCommit != nil {
		return 0, errCommit
	}

	return id, err
}

// GetUserByID retrieves User by ID. Returns error if
// ID doesn't exist
func GetUserByID(id int64) (obj *User, err error) {
	o := orm.NewOrm()
	obj = &User{ID: id}
	if err = o.Read(obj); err == nil {
		if obj.Company != nil {
			o.Read(obj.Company)
		}
		if obj.Department != nil {
			o.Read(obj.Department)
		}
		if obj.Position != nil {
			o.Read(obj.Position)
		}
		o.LoadRelated(obj, "Teams")
		o.LoadRelated(obj, "Roles")

		return obj, nil
	}
	return nil, err
}

// GetUserByName get user
func GetUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("mobile", name).Or("email__icontains", name).Or("name", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	qs = qs.RelatedSel()
	err := qs.One(&user)
	return user, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []User, error) {
	var (
		objArrs   []User
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))

	//cond k=v cond必须放到Filter和Exclude前面
	cond := orm.NewCondition()
	if _, ok := condMap["and"]; ok {
		andMap := condMap["and"]
		for k, v := range andMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.And(k, v)
		}
	}
	if _, ok := condMap["or"]; ok {
		orMap := condMap["or"]
		for k, v := range orMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.Or(k, v)
		}
	}
	qs = qs.SetCond(cond)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Exclude(k, v)
	}

	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[i] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[0] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return paginator, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return paginator, nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	qs = qs.RelatedSel()
	if cnt, err := qs.Count(); err == nil {
		if cnt > 0 {
			paginator = utils.GenPaginator(limit, offset, cnt)
			if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
				paginator.CurrentPageSize = num
				for obj := range objArrs {
					o.LoadRelated(&obj, "Roles")
					o.LoadRelated(&obj, "Teams")
				}
			}
		}
	}
	return paginator, objArrs, err
}

// UpdateUser updates User by ID and returns error if
// the record to be updated doesn't exist
func UpdateUser(obj *User, updateUser *User) (err error) {
	o := orm.NewOrm()
	v := User{ID: obj.ID}
	errBegin := o.Begin()
	defer func() {
		if err != nil {
			if errRollback := o.Rollback(); errRollback != nil {
				err = errRollback
			}
		}
	}()
	if errBegin != nil {
		return errBegin
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if obj.CompanyID > 0 {
			obj.Company, _ = GetCompanyByID(obj.CompanyID)
		}
		if obj.DepartmentID > 0 {
			obj.Department, _ = GetDepartmentByID(obj.DepartmentID)
		}
		if obj.PositionID > 0 {
			obj.Position, _ = GetPositionByID(obj.PositionID)
		}
		if createTeamRecords, ok := obj.TeamIDs["create"]; ok {
			m2mTeams := o.QueryM2M(obj, "Teams")
			for _, teamID := range createTeamRecords {
				if team, err := GetTeamByID(int64(teamID)); err == nil {
					m2mTeams.Add(team)
				} else {
					utils.LogOut("error", "add user teams failed:"+err.Error())
				}
			}
		}
		if deleteTeamRecords, ok := obj.TeamIDs["delete"]; ok {
			m2mTeams := o.QueryM2M(obj, "Teams")
			for _, teamID := range deleteTeamRecords {
				if team, err := GetTeamByID(int64(teamID)); err == nil {
					m2mTeams.Remove(team)
				} else {
					utils.LogOut("error", "delete user teams failed:"+err.Error())

				}
			}
		}
		if createRoleRecords, ok := obj.RoleIDs["create"]; ok {
			m2mRoles := o.QueryM2M(obj, "Roles")
			for _, RoleID := range createRoleRecords {
				if role, err := GetRoleByID(int64(RoleID)); err == nil {
					m2mRoles.Add(role)
				} else {
					utils.LogOut("error", "add user roles failed:"+err.Error())
				}
			}
		}
		if deleteRoleRecords, ok := obj.RoleIDs["delete"]; ok {
			m2mRoles := o.QueryM2M(obj, "Roles")
			for _, RoleID := range deleteRoleRecords {
				if role, err := GetRoleByID(int64(RoleID)); err == nil {
					m2mRoles.Remove(role)
				} else {
					utils.LogOut("error", "delete user roles failed:"+err.Error())
				}
			}
		}
		if _, err = o.Update(obj, append(obj.ActionFields, "UpdateUser", "UpdateDate")...); err != nil {
			utils.LogOut("error", "update user fields failed:"+err.Error())
		}
	}
	if err != nil {
		return err
	}
	errCommit := o.Commit()
	if errCommit != nil {
		return errCommit
	}
	return nil
}

// DeleteUser deletes User by ID and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&User{ID: id})
	}
	return
}

// CheckUserByName  check
func CheckUserByName(name, password string) (User, bool, error) {
	o := orm.NewOrm()
	var (
		user User
		err  error
		ok   bool
	)
	ok = false
	//7LR8ZC-855575-64657756081974692
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("Name", name).Or("Email", name).Or("Mobile", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Password == utils.PasswordMD5(password, user.Mobile) {
			ok = true
			if user.Company != nil {
				o.Read(user.Company)
			}
			if user.Department != nil {
				o.Read(user.Department)
			}
			if user.Position != nil {
				o.Read(user.Position)
			}
		}
	}
	return user, ok, err
}
