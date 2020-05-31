package controllers

import (
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"strings"
)

type LoginController struct {
	BaseController
}

// Index
func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(c.GetString("UName"))
		userpwd := strings.TrimSpace(c.GetString("UPasswd"))

		if len(username) == 0 || len(userpwd) == 0 {
			c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
		}
		//userpwd = utils.String2md5(userpwd)

		user, err := models.BackendUserOneByUserName(username, userpwd)
		if user != nil && err == nil {
			if user.UIsAdmin == -1 {
				c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
			}
			//保存用户信息到session
			c.SetSession("user", *user)

			//获取用户信息
			c.jsonResult(enums.JRCodeSucc, "登录成功", "")
		} else {
			c.jsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
		}
	} else {
		c.Data["title"] = "登录-Brook"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "login/css.html"
		c.LayoutSections["footerjs"] = "login/js.html"
		c.setTpl("login/index.html", "shared/public_page.html")
	}

}
