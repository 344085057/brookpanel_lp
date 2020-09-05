package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//LpBrookMoneycdk 充值码
type LpBrookMoneycdk struct {
	Id          int          `orm:"column(id);pk" description:"充值码"`
	Cdk         string       `orm:"column(cdk);size(255)" description:"cdk"`
	Money       int          `orm:"column(money);" description:"金额"`
	CreateTime  time.Time    `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	UseTime     time.Time    `orm:"column(use_time);type(datetime);auto_now" description:"使用时间"`
	UseUid      int          `orm:"column(use_uid)" description:"使用者id"`
	LpBrookUser *LpBrookUser `orm:"-"`
}

// AddLpBrookMoneycdk insert a new LpBrookMoneycdk into database and returns
// last inserted Id on success.
func AddLpBrookMoneycdk(m *LpBrookMoneycdk) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookMoneycdkById retrieves LpBrookMoneycdk by Id. Returns error if
// Id doesn't exist
func GetLpBrookMoneycdkById(id int) (v *LpBrookMoneycdk, err error) {
	o := orm.NewOrm()
	v = &LpBrookMoneycdk{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	if err == orm.ErrNoRows { //判断是否 是 没有找到的错误
		return nil, nil
	}
	return nil, err
}

// GetLpBrookMoneycdkByCdk 根据cdk获取对象
func GetLpBrookMoneycdkByCdk(cdk string) (v *LpBrookMoneycdk, err error) {
	o := orm.NewOrm()
	m := LpBrookMoneycdk{}
	err = o.QueryTable(LpBrookMoneycdkTableName()).Filter("Cdk", cdk).One(&m)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// GetAllLpBrookMoneycdk retrieves all LpBrookMoneycdk matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookMoneycdk(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookMoneycdk))
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

	var l []LpBrookMoneycdk
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

// UpdateLpBrookMoneycdk updates LpBrookMoneycdk by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookMoneycdkById(m *LpBrookMoneycdk) (err error) {
	o := orm.NewOrm()
	v := LpBrookMoneycdk{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookMoneycdk deletes LpBrookMoneycdk by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookMoneycdk(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookMoneycdk{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookMoneycdk{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取CDK
func GetCDKPage(page, num int, cdk LpBrookMoneycdk) (v []LpBrookMoneycdk, totle int64, err error) {
	o := orm.NewOrm()

	// Cdk        string    `orm:"column(cdk);size(255)" description:"cdk"`
	// Money      int       `orm:"column(money);" description:"金额"`
	// CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now_add" description:"创建时间"`
	// UseTime    time.Time `orm:"column(use_time);type(datetime);auto_now" description:"使用时间"`
	// UseUid     int       `orm:"column(use_uid)" description:"使用者id"`

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(LpBrookMoneycdkTableName())
	qsCount := o.QueryTable(LpBrookMoneycdkTableName())
	if cdk.Cdk != "" {
		qs = qs.Filter("Cdk__icontains", cdk.Cdk)
		qsCount = qsCount.Filter("Cdk__icontains", cdk.Cdk)
	}
	if cdk.Money != 0 {
		qs = qs.Filter("Money", cdk.Money)
		qsCount = qsCount.Filter("Money", cdk.Money)
	}
	if cdk.UseUid == -1 {
		qs = qs.Filter("UseUid", 0)
		qsCount = qsCount.Filter("UseUid", 0)
	}
	if cdk.UseUid == 1 {
		qs = qs.Filter("UseUid__gt", 0)
		qsCount = qsCount.Filter("UseUid__gt", 0)
	}

	if page == 0 {
		page = 1
	}
	qs = qs.Limit(num, (page-1)*num)
	data := make([]LpBrookMoneycdk, 0)
	_, err = qs.All(&data)
	if err != nil {
		return nil, 0, err
	}

	cnt, err := qsCount.Count()
	if err != nil {
		return nil, 0, err
	}

	for i, cdkPage := range data {
		userInfo, err := GetLpBrookUserById(cdkPage.UseUid)
		if err != nil {
			return nil, 0, err
		}
		data[i].LpBrookUser = userInfo
	}

	return data, cnt, nil

}
