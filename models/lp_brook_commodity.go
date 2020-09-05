package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

//LpBrookCommodity 商品
type LpBrookCommodity struct {
	Id       int     `orm:"column(id);auto" description:"商品"`
	Sort     string  `orm:"column(sort);size(255)" description:"商品类别"`
	Title    string  `orm:"column(title);size(255)" description:"商品名称"`
	Describe string  `orm:"column(describe);size(255)" description:"商品描述"`
	Money    int     `orm:"column(money);digits(20);decimals(2)" description:"商品价格"`
	Time     int     `orm:"column(time)" description:"时长（天数）"`
	Cover    int     `orm:"column(cover)" description:"-1:覆盖/1:叠加/ 默认覆盖"`
	State    int     `orm:"column(state)" description:"-1:禁用/1:启用 默认启用"`
	Sx       int     `orm:"column(sx)" description:"顺序"`
	Ll       float64 `orm:"column(ll);digits(40);decimals(5)" description:"流量  mb"`
	// TableTime time.Time `orm:"column(table_time);type(datetime);auto_now" description:"直接修改表的日期"`
}

//GetSortAllByState1 获取已启用的分组
func GetSortAllByState1() (sortArr []string, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select sort from lp_brook_commodity where state = 1 group by sort ORDER BY sx desc").QueryRows(&sortArr)
	if err != nil {
		return nil, err
	}
	return sortArr, nil
}

//GetSortAll 获取所有分组
func GetSortAll() (sortArr []string, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select sort from lp_brook_commodity group by sort ORDER BY sx desc").QueryRows(&sortArr)
	if err != nil {
		return nil, err
	}
	return sortArr, nil
}

//GetCommodityArrBySort 根据sort获取已启用的商品
func GetCommodityArrBySort(sort string) (lpBrookCommodityArr []LpBrookCommodity, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from lp_brook_commodity where sort = ? and state = 1 ORDER BY sx desc", sort).QueryRows(&lpBrookCommodityArr)
	if err != nil {
		return nil, err
	}
	return lpBrookCommodityArr, nil
}

//GetCommodityAllBySort 根据sort获取所有商品
func GetCommodityAllBySort(sort string) (lpBrookCommodityArr []LpBrookCommodity, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from lp_brook_commodity where sort = ? ORDER BY sx desc", sort).QueryRows(&lpBrookCommodityArr)
	if err != nil {
		return nil, err
	}
	return lpBrookCommodityArr, nil
}

// AddLpBrookCommodity insert a new LpBrookCommodity into database and returns
// last inserted Id on success.
func AddLpBrookCommodity(m *LpBrookCommodity) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookCommodityById retrieves LpBrookCommodity by Id. Returns error if
// Id doesn't exist
func GetLpBrookCommodityById(id int) (v *LpBrookCommodity, err error) {
	o := orm.NewOrm()
	v = &LpBrookCommodity{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	if err == orm.ErrNoRows { //判断是否 是 没有找到的错误
		return nil, nil
	}
	return nil, err
}

// GetAllLpBrookCommodity retrieves all LpBrookCommodity matches certain condition. Returns empty list if
// no records exist
func GetAllLpBrookCommodity(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LpBrookCommodity))
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

	var l []LpBrookCommodity
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

// UpdateLpBrookCommodity updates LpBrookCommodity by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookCommodityById(m *LpBrookCommodity) (err error) {
	o := orm.NewOrm()
	v := LpBrookCommodity{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookCommodity deletes LpBrookCommodity by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookCommodity(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookCommodity{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookCommodity{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
