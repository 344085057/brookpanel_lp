package models

import (
	"errors"
	"fmt"
	"myBrookWeb/utils"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

//LpBrookUser 用户
type LpBrookUser struct {
	Id         int       `orm:"column(u_id);auto"`
	Email      string    `orm:"column(u_email);size(255)" description:"邮箱" valid:"Email; MaxSize(50)"`
	Name       string    `orm:"column(u_name);size(255)" description:"用户名" valid:"Range(2, 20)"`
	Passwd     string    `orm:"column(u_passwd);size(255)" description:"密码" valid:"Range(6, 20)"`
	Port       int       `orm:"column(u_port);size(255)" description:"端口"`
	Flow       float64   `orm:"column(u_flow);digits(40);decimals(5)" description:"剩余流量"`
	IsAdmin    int       `orm:"column(u_is_admin)" description:"是否是管理员 0普通用户/1管理员/-1停用"`
	ExpireTime time.Time `orm:"column(expire_time);type(timestamp);" description:"vip到期时间"`
	FlowTotal  float64   `orm:"column(u_flow_total);digits(40);decimals(5)" description:"总使用流量"`
	Money      int       `orm:"column(u_money)" description:"金币 100 = 1元"`
	// TableTime  time.Time `orm:"column(table_time);type(datetime);auto_now" description:"直接修改表的日期"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建日期"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新日期"`
}

//LpBrookUserByLogin 用户登录
type LpBrookUserByLogin struct {
	UEmail  string `orm:"column(u_email);size(255)" description:"邮箱" valid:"Email; MaxSize(50)"`
	UPasswd string `orm:"column(u_passwd);size(255)" description:"密码" valid:"MinSize(6);MaxSize(20)"`
}

//LpBrookUserByRegin 用户注册
type LpBrookUserByRegin struct {
	UEmail  string `orm:"column(u_email);size(255)" description:"邮箱" valid:"Email; MaxSize(50)"`
	UPasswd string `orm:"column(u_passwd);size(255)" description:"密码" valid:"MinSize(6);MaxSize(20)"`
	UName   string `orm:"column(u_name);size(255)" description:"名称" valid:"MinSize(2);MaxSize(20)"`
}

//LpBrookUserByUpdataPasswd 用户修改密码
type LpBrookUserByUpdataPasswd struct {
	Passwd    string `json:"passwd" description:"旧密码" valid:"MinSize(6);MaxSize(20)"`
	NewPasswd string `json:"newPasswd"  description:"新密码" valid:"MinSize(6);MaxSize(20)"`
}

//LpBrookUserByUpdataPasswd 用户修改密码
type LpBrookUserByUpdataPort struct {
	Port int `json:"port" description:"端口" valid:"Min(1024);Max(60000)"`
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
	if err == orm.ErrNoRows { //判断是否 是 没有找到的错误
		return nil, nil
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

//GetUserOneByEmailAndPasswd 根据email和密码 查询用户
func GetUserOneByEmailAndPasswd(email, passwd string) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("Email", email).Filter("Passwd", passwd).One(&m)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//UpdataNewPasswdByUserID 根据id 更新密码
func UpdataNewPasswdByUserID(id int, newPasswd string) (err error) {
	o := orm.NewOrm()
	user := LpBrookUser{
		Id:     id,
		Passwd: newPasswd,
	}

	if _, err = o.Update(&user, "Passwd"); err != nil {
		return err
	}

	lpBrookServer, errr := GetLpBrookAll(0) //获取可用服务器
	if errr != nil {
		return err
	} else {
		for _, v := range lpBrookServer {
			//http请求
			req := httplib.Get("http://" + v.Ip + ":60001/remote/UpdataServicePasswd")

			o := orm.NewOrm()
			sysMap := make(orm.Params)
			o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

			req.Param("remote_u", sysMap["remote_u"].(string))
			req.Param("remote_p", sysMap["remote_p"].(string))

			req.Param("user_id", fmt.Sprintf("%v", id))
			fmt.Println(req.String())
		}
	}

	return nil
}

//UpdataNewPortByUserID 根据id 更新端口
func UpdataNewPortByUserID(id, port, money int) (err error) {

	BrookUser, err := BackendUserOneByUPort(port)
	if err != nil {
		return err
	}
	if BrookUser != nil {
		return errors.New("端口已被使用:(")
	}

	user, err := GetLpBrookUserById(id)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	upUser := LpBrookUser{
		Id:    id,
		Port:  port,
		Money: user.Money - money,
	}

	if _, err = o.Update(&upUser, "Port", "Money"); err != nil {
		return err
	}

	return nil
}

//GetUserOneByEmail 根据email 查询用户
func GetUserOneByEmail(email string) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("Email", email).One(&m)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//BackendUserOneByUPort 根据端口 查询用户
func BackendUserOneByUPort(port int) (*LpBrookUser, error) {
	m := LpBrookUser{}
	err := orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("Port", port).One(&m)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//UPortIsZy 获取没用到的端口
func UPortIsZy() int {

	pro := utils.GenerateRangeNum(1024, 60000)

	m := LpBrookUser{}
	orm.NewOrm().QueryTable(LpBrookUserTBName()).Filter("Port", pro).One(&m)
	if m.Id != 0 {
		UPortIsZy()
		return 0
	} else {
		return pro
	}

}

//用户购买商品
func UserShopping(userid int, v LpBrookCommodity) error {
	o := orm.NewOrm()

	user, err := GetLpBrookUserById(userid)
	if err != nil {
		return err
	}
	if user.Money < v.Money {
		return errors.New("账户金额不足,请充值:(")
	}

	if v.Cover == -1 { //覆盖
		//当前系统时间
		now := time.Now()

		dd, err := time.ParseDuration(fmt.Sprintf("%vh", v.Time*24))
		if err != nil {
			return err
		}
		dd1 := now.Add(dd)

		user := LpBrookUser{
			Id:         userid,
			Flow:       v.Ll,                 //   `orm:"column(u_flow);digits(40);decimals(5)" description:"剩余流量"`
			ExpireTime: dd1,                  // `orm:"column(expire_time);type(timestamp);auto_now_add" description:"vip到期时间"`
			Money:      user.Money - v.Money, //       `orm:"column(u_money)" description:"金币 100 = 1元"`
		}
		if _, err = o.Update(&user, "Flow", "ExpireTime", "Money"); err != nil {
			return err
		}
	} else if v.Cover == 1 { //叠加

		dd, err := time.ParseDuration(fmt.Sprintf("%vh", v.Time*24))
		if err != nil {
			return err
		}
		dd1 := user.ExpireTime.Add(dd)

		user := LpBrookUser{
			Id:         userid,
			Flow:       user.Flow + v.Ll,     //   `orm:"column(u_flow);digits(40);decimals(5)" description:"剩余流量"`
			ExpireTime: dd1,                  // `orm:"column(expire_time);type(timestamp);auto_now_add" description:"vip到期时间"`
			Money:      user.Money - v.Money, //       `orm:"column(u_money)" description:"金币 100 = 1元"`
		}
		if _, err = o.Update(&user, "Flow", "ExpireTime", "Money"); err != nil {
			return err
		}

	}

	return nil
}

//获取所有用户
func GetUserPage(page, num int, user LpBrookUser) (v []LpBrookUser, totle int64, err error) {
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable(LpBrookUserTBName())
	qsCount := o.QueryTable(LpBrookUserTBName())
	if user.Email != "" {
		qs = qs.Filter("Email__icontains", user.Email)
		qsCount = qsCount.Filter("Email__icontains", user.Email)
	}
	if user.Name != "" {
		qs = qs.Filter("Name__icontains", user.Name)
		qsCount = qsCount.Filter("Name__icontains", user.Name)
	}
	if user.Port != 0 {
		qs = qs.Filter("Port__icontains", user.Port)
		qsCount = qsCount.Filter("Port__icontains", user.Port)
	}
	if user.IsAdmin != 1 {
		qs = qs.Filter("IsAdmin", user.IsAdmin)
		qsCount = qsCount.Filter("IsAdmin", user.IsAdmin)
	}
	qs = qs.Exclude("IsAdmin", 1)
	qsCount = qsCount.Exclude("IsAdmin", 1)

	if page == 0 {
		page = 1
	}
	qs = qs.Limit(num, (page-1)*num)
	userArr := make([]LpBrookUser, 0)
	_, err = qs.All(&userArr)
	if err != nil {
		return nil, 0, err
	}

	cnt, err := qsCount.Count()
	if err != nil {
		return nil, 0, err
	}

	return userArr, cnt, nil

}
