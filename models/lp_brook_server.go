package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LpBrookServer struct {
	Id         int     `orm:"column(s_id);auto"`
	SIp        string  `orm:"column(s_ip);size(255)" description:"ip地址"`
	SFlowRatio float64 `orm:"column(s_flow_ratio)" description:"流量比例"`
	STitle     string  `orm:"column(s_title);size(255)" description:"服务标题"`
	SType      int     `orm:"column(s_type);size(255)" description:"服务器类型"`
	SDelay     string  `orm:"column(s_delay);size(255)" description:"服务器延迟"`
	SOrder     int     `orm:"column(s_order);size(255)" description:"顺序"`
}

// AddLpBrookServer insert a new LpBrookServer into database and returns
// last inserted Id on success.
func AddLpBrookServer(m *LpBrookServer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookServerById retrieves LpBrookServer by Id. Returns error if
// Id doesn't exist
func GetLpBrookServerById(id int) (v *LpBrookServer, err error) {
	o := orm.NewOrm()
	v = &LpBrookServer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLpBrookServer retrieves all LpBrookServer matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookServer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookServer))
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

	var l []LpBrookServer
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

// UpdateLpBrookServer updates LpBrookServer by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookServerById(m *LpBrookServer) (err error) {
	o := orm.NewOrm()
	v := LpBrookServer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookServer deletes LpBrookServer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookServer(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookServer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookServer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取所有可用服务器
func GetLpBrookAll() (v *[]LpBrookServer, err error) {
	o := orm.NewOrm()
	var lpBrookServer []LpBrookServer

	_, err = o.QueryTable(LpBrookServerTBName()).Exclude("s_type__exact", -1).All(&lpBrookServer)
	if err == nil {
		return &lpBrookServer, nil
	}
	return nil, err
}
