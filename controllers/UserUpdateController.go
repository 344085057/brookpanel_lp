package controllers

import (
	"encoding/json"
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"myBrookWeb/utils"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserUpdateController struct {
	BaseController
}

//Logout 用户退出登录
func (c *UserUpdateController) Logout() {
	c.DelSession("user")
	c.Ctx.Redirect(302, "/login")
}

//Index 更改个人信息　页面
func (c *UserUpdateController) Index() {

	updata_port_money, err := beego.AppConfig.Int("sys_config::updata_port_money")
	if err != nil {
		c.Data["err"] = err.Error()
	}
	c.Data["updata_port_money"] = updata_port_money
	c.Data["title"] = "更改信息-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userupdate/css.html"
	c.LayoutSections["footerjs"] = "userupdate/js.html"
	c.setTpl("userupdate/index.html", "shared/userpanel.html")

}

//UpdataPasswd 用户修改密码
func (c *UserUpdateController) UpdataPasswd() {
	var updataPasswd models.LpBrookUserByUpdataPasswd

	//json数据封装到user对象中
	err := c.ParseForm(&updataPasswd)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if updataPasswd.NewPasswd == c.curUser.ProxyPasswd {
		c.jsonResult(enums.JRCodeFailed, "新的登录密码不能和连接密码一致", "")
	}
	updataPasswd.Passwd = utils.String2md5(updataPasswd.Passwd)
	updataPasswd.NewPasswd = utils.String2md5(updataPasswd.NewPasswd)

	valid := validation.Validation{}
	b, _ := valid.Valid(&updataPasswd)
	if !b {
		// validation does not pass
		// blabla...
		//表示获取验证的结构体
		st := reflect.TypeOf(updataPasswd)
		// for _, err := range valid.Errors {
		// 	filed, _ := st.FieldByName(err.Field)
		// 	var alias = filed.Tag.Get("alias")
		// 	log.Println(alias, err.Message)
		// }
		filed, _ := st.FieldByName(valid.Errors[0].Field)
		var alias = filed.Tag.Get("description")
		// log.Println(alias, valid.Errors[0].Message)
		msgStr := fmt.Sprintf("%s %s", alias, valid.Errors[0].Message)

		c.jsonResult(enums.JRCodeFailed, msgStr, "")
	}

	if updataPasswd.Passwd != c.curUser.Passwd {
		c.jsonResult(enums.JRCodeFailed, "当前登录密码错误", "")
	}

	if updataPasswd.ProxyPasswd != c.curUser.ProxyPasswd {
		c.jsonResult(enums.JRCodeFailed, "当前连接密码错误", "")
	}

	err = models.UpdataNewPasswdByUserID(c.curUser.Id, updataPasswd.NewPasswd)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "修改成功:)", "")

}

//UpdataPorxyPasswd 用户连接密码
func (c *UserUpdateController) UpdataPorxyPasswd() {
	var updataPasswd models.LpBrookUserByUpdataProxyPasswd

	//json数据封装到user对象中
	err := c.ParseForm(&updataPasswd)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if updataPasswd.Passwd == updataPasswd.NewProxyPasswd {
		c.jsonResult(enums.JRCodeFailed, "新的连接密码不能和登录密码一致", "")
	}
	updataPasswd.Passwd = utils.String2md5(updataPasswd.Passwd)

	valid := validation.Validation{}
	b, _ := valid.Valid(&updataPasswd)
	if !b {
		// validation does not pass
		// blabla...
		//表示获取验证的结构体
		st := reflect.TypeOf(updataPasswd)
		// for _, err := range valid.Errors {
		// 	filed, _ := st.FieldByName(err.Field)
		// 	var alias = filed.Tag.Get("alias")
		// 	log.Println(alias, err.Message)
		// }
		filed, _ := st.FieldByName(valid.Errors[0].Field)
		var alias = filed.Tag.Get("description")
		// log.Println(alias, valid.Errors[0].Message)
		msgStr := fmt.Sprintf("%s %s", alias, valid.Errors[0].Message)

		c.jsonResult(enums.JRCodeFailed, msgStr, "")
	}

	if updataPasswd.Passwd != c.curUser.Passwd {
		c.jsonResult(enums.JRCodeFailed, "当前登录密码错误", "")
	}

	if updataPasswd.ProxyPasswd != c.curUser.ProxyPasswd {
		c.jsonResult(enums.JRCodeFailed, "当前连接密码错误", "")
	}

	err = models.UpdataNewProxyPasswdByUserID(c.curUser.Id, updataPasswd.NewProxyPasswd)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "修改成功:)", "")

}

//UpdataPort 用户修改端口
func (c *UserUpdateController) UpdataPort() {
	var updataPort models.LpBrookUserByUpdataPort
	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &updataPort)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	valid := validation.Validation{}
	b, _ := valid.Valid(&updataPort)
	if !b {
		// validation does not pass
		// blabla...
		//表示获取验证的结构体
		st := reflect.TypeOf(updataPort)
		// for _, err := range valid.Errors {
		// 	filed, _ := st.FieldByName(err.Field)
		// 	var alias = filed.Tag.Get("alias")
		// 	log.Println(alias, err.Message)
		// }
		filed, _ := st.FieldByName(valid.Errors[0].Field)
		var alias = filed.Tag.Get("description")
		// log.Println(alias, valid.Errors[0].Message)
		msgStr := fmt.Sprintf("%s %s", alias, valid.Errors[0].Message)

		c.jsonResult(enums.JRCodeFailed, msgStr, "")
	}

	updata_port_money, err := beego.AppConfig.Int("sys_config::updata_port_money")
	if err != nil {
		c.Data["err"] = err.Error()
	}
	if updata_port_money > c.curUser.Money {
		c.jsonResult(enums.JRCodeFailed, "金币不足", "")
	}
	err = models.UpdataNewPortByUserID(c.curUser.Id, updataPort.Port, updata_port_money)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "修改成功:)", "")

}
