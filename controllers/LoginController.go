package controllers

import (
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"myBrookWeb/utils"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

// 登录
func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		ip := c.ip
		ipStrArr := make([]string, 0)

		if err := utils.GetCache(ip, &ipStrArr); err != nil { //获取ip数组
			utils.SetCache(ip, ipStrArr, 0) // 设置缓存
		}
		loginErrorNum, err := beego.AppConfig.Int("sys_config::login_error_num") // 获取登录错误限制次数
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "系统设置错误！", err)
		}
		if len(ipStrArr) == loginErrorNum {
			c.jsonResult(enums.JRCodeFailed, "您被限制登录！", "")
		}

		loginErrorTimeout, err := beego.AppConfig.Int("sys_config::login_error_timeout") // 获取登录错误限制时间
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "系统设置错误！", err)
		}

		userEmail := strings.TrimSpace(c.GetString("UEmail"))
		userPasswd := strings.TrimSpace(c.GetString("UPasswd"))
		userPasswd = utils.String2md5(userPasswd)
		// if len(username) == 0 || len(userpwd) == 0 {
		// 	c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", s)
		// }
		//userpwd = utils.String2md5(userpwd)
		u := models.LpBrookUserByLogin{UEmail: userEmail, UPasswd: userPasswd}

		valid := validation.Validation{}
		b, _ := valid.Valid(&u)
		if !b {
			// validation does not pass
			// blabla...
			//表示获取验证的结构体
			st := reflect.TypeOf(u)
			// for _, err := range valid.Errors {
			// 	filed, _ := st.FieldByName(err.Field)
			// 	var alias = filed.Tag.Get("alias")
			// 	log.Println(alias, err.Message)
			// }
			filed, _ := st.FieldByName(valid.Errors[0].Field)
			var alias = filed.Tag.Get("description")
			// log.Println(alias, valid.Errors[0].Message)
			ipStrArr := append(ipStrArr, ip)
			utils.SetCache(ip, ipStrArr, loginErrorTimeout)
			fmt.Println(u)
			msgStr := fmt.Sprintf("%s %s,您还有%s次机会", alias, valid.Errors[0].Message, strconv.Itoa(loginErrorNum-len(ipStrArr)))
			if loginErrorNum-len(ipStrArr) == 0 {
				msgStr = "你已经被限制登录！"
			}
			c.jsonResult(enums.JRCodeFailed, msgStr, "")
		}

		user, err := models.GetUserOneByEmailAndPasswd(userEmail, userPasswd)
		if user != nil && err == nil {
			if user.IsAdmin == -1 {
				ipStrArr := append(ipStrArr, ip)
				utils.SetCache(ip, ipStrArr, loginErrorTimeout)
				msgStr := fmt.Sprintf("用户被禁用，请联系管理员;您还有%s次机会", strconv.Itoa(loginErrorNum-len(ipStrArr)))
				if loginErrorNum-len(ipStrArr) == 0 {
					msgStr = "你已经被限制登录！"
				}
				c.jsonResult(enums.JRCodeFailed, msgStr, "")
			}
			//保存用户信息到session
			c.SetSession("user", *user)

			//插入登入日志
			userLoginLog := models.LpBrookUserLoginLog{
				UId:       user.Id,
				LoginTime: time.Now(),
				LoginIp:   ip,
			}
			models.AddLpBrookUserLoginLog(&userLoginLog)

			//获取用户信息
			c.jsonResult(enums.JRCodeSucc, "登录成功", "")
		} else {
			ipStrArr := append(ipStrArr, ip)
			utils.SetCache(ip, ipStrArr, loginErrorTimeout)

			msgStr := fmt.Sprintf("用户名或者密码错误,您还有%s次机会", strconv.Itoa(loginErrorNum-len(ipStrArr)))
			if loginErrorNum-len(ipStrArr) == 0 {
				msgStr = "你已经被限制登录！"
			}
			c.jsonResult(enums.JRCodeFailed, msgStr, "")
		}
	} else {
		c.Data["title"] = "登录-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "login/css.html"
		c.LayoutSections["footerjs"] = "login/js.html"
		c.setTpl("login/index.html", "shared/public_page.html")
	}

}
