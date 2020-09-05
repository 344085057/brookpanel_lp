package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//LpBrookGg 文章
type LpBrookGg struct {
	Id         int       `orm:"column(id);auto"`
	Title      string    `orm:"column(title);size(255);" description:"标题"`
	Text       string    `orm:"column(text);size(255);" description:"内容"`
	State      int       `orm:"column(state);" description:"-1:禁用/1:启用 默认启用"`
	GType      int       `orm:"column(g_type);" description:"1:公告 2使用教程"`
	Sx         int       `orm:"column(sx);" description:"顺序"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建日期"`
}

//GetLpBrookAllBygTypeAndState 根据g_type 和 state获取文章
func GetLpBrookAllBygTypeAndState(gType int, state int) (v []LpBrookGg, err error) {
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(LpBrookGGTableName())

	if gType != 0 {
		qs = qs.Filter("g_type", gType)
	}

	if state != 0 {
		qs = qs.Filter("state", state)
	}

	qs.OrderBy("-sx")
	if _, err = qs.All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

// AddLpBrookGg insert a new LpBrookGg into database and returns
// last inserted Id on success.
func AddLpBrookGg(m *LpBrookGg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookGgById retrieves LpBrookGg by Id. Returns error if
// Id doesn't exist
func GetLpBrookGgById(id int) (v *LpBrookGg, err error) {
	o := orm.NewOrm()
	v = &LpBrookGg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	if err == orm.ErrNoRows { //判断是否 是 没有找到的错误
		return nil, nil
	}
	return nil, err
}

// GetAllLpBrookGg retrieves all LpBrookGg matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookGg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookGg))
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

	var l []LpBrookGg
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

// UpdateLpBrookGg updates LpBrookGg by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookGgById(m *LpBrookGg) (err error) {
	o := orm.NewOrm()
	v := LpBrookGg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookGg deletes LpBrookGg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookGg(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookGg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookGg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
