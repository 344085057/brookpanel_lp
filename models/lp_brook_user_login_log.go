package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type LpBrookUserLoginLog struct {
	Id             int          `orm:"column(l_id);auto" description:"登录日志"`
	UId            int          `orm:"column(u_id)" description:"用户id"`
	LoginTime      time.Time    `orm:"column(login_time);type(timestamp);auto_now_add" description:"登录时间"`
	LoginIp        string       `orm:"column(login_ip);size(255)" description:"登录ip"`
	LoginIpAddress string       `orm:"column(login_ip_address);size(255);null" description:"登录归属地"`
	LpBrookUser    *LpBrookUser `orm:"-"`
}

// AddLpBrookUserLoginLog insert a new LpBrookUserLoginLog into database and returns
// last inserted Id on success.
func AddLpBrookUserLoginLog(m *LpBrookUserLoginLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookUserLoginLogById retrieves LpBrookUserLoginLog by Id. Returns error if
// Id doesn't exist
func GetLpBrookUserLoginLogById(id int) (v *LpBrookUserLoginLog, err error) {
	o := orm.NewOrm()
	v = &LpBrookUserLoginLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	if err == orm.ErrNoRows { //判断是否 是 没有找到的错误
		return nil, nil
	}
	return nil, err
}

// GetAllLpBrookUserLoginLog retrieves all LpBrookUserLoginLog matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookUserLoginLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookUserLoginLog))
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

	var l []LpBrookUserLoginLog
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

// UpdateLpBrookUserLoginLog updates LpBrookUserLoginLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookUserLoginLogById(m *LpBrookUserLoginLog) (err error) {
	o := orm.NewOrm()
	v := LpBrookUserLoginLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookUserLoginLog deletes LpBrookUserLoginLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookUserLoginLog(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookUserLoginLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookUserLoginLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取登录日志
func GetLoginLogPage(page, num int, loginLog LpBrookUserLoginLog) (v []LpBrookUserLoginLog, totle int64, err error) {
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(LpBrookUserLoginLogTableName())
	qsCount := o.QueryTable(LpBrookUserLoginLogTableName())
	if loginLog.LoginIp != "" {
		qs = qs.Filter("LoginIp__icontains", loginLog.LoginIp)
		qsCount = qsCount.Filter("LoginIp__icontains", loginLog.LoginIp)
	}
	if loginLog.LoginIpAddress != "" {
		qs = qs.Filter("LoginIpAddress__icontains", loginLog.LoginIpAddress)
		qsCount = qsCount.Filter("LoginIpAddress__icontains", loginLog.LoginIpAddress)
	}

	if page == 0 {
		page = 1
	}
	qs = qs.Limit(num, (page-1)*num)
	data := make([]LpBrookUserLoginLog, 0)
	_, err = qs.All(&data)
	if err != nil {
		return nil, 0, err
	}

	cnt, err := qsCount.Count()
	if err != nil {
		return nil, 0, err
	}

	for i, cdkPage := range data {
		userInfo, err := GetLpBrookUserById(cdkPage.UId)
		if err != nil {
			return nil, 0, err
		}
		data[i].LpBrookUser = userInfo
	}

	return data, cnt, nil

}
