package models

import (
	"errors"
	"fmt"
	"myBrookWeb/utils"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LpBrookUser struct {
	Id       int    `orm:"column(u_id);auto"`
	UName    string `orm:"column(u_name);size(255)" description:"用户名"`
	UPasswd  string `orm:"column(u_passwd);size(255)" description:"密码"`
	UProt    string `orm:"column(u_prot);size(255)" description:"端口"`
	UFlow    string `orm:"column(u_flow);size(255)" description:"流量"`
	UIsAdmin int    `orm:"column(u_is_admin)" description:"是否是管理员"`
}

// AddLpBrookUser insert a new LpBrookUser into database and returns
// last inserted Id on success.
func AddLpBrookUser(m *LpBrookUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookUserById retrieves LpBrookUser by Id. Returns error if
// Id doesn't exist
func GetLpBrookUserById(id int) (v *LpBrookUser, err error) {
	o := orm.NewOrm()
	v = &LpBrookUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLpBrookUser retrieves all LpBrookUser matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []LpBrookUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLpBrookUser updates LpBrookUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookUserById(m *LpBrookUser) (err error) {
	o := orm.NewOrm()
	v := LpBrookUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookUser deletes LpBrookUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookUser(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据用户名和密码 查询用户
func BackendUserOneByUserName(user, passwd string) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("UName", user).Filter("UPasswd", passwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//根据用户名 查询用户
func BackendUserOneByUser(user string) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("UName", user).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//根据端口 查询用户
func BackendUserOneByUProt(prot string) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("UProt", prot).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//获取没用到的端口

func UProtIsZy() (prot string) {

	pro := strconv.Itoa(utils.GenerateRangeNum(10000, 60000))

	m := LpBrookUser{}
	orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("UProt", pro).One(&m)
	fmt.Println(pro)
	fmt.Println(&m)
	if m.Id != 0 {
		UProtIsZy()
		return ""
	} else {
		return pro
	}

}
