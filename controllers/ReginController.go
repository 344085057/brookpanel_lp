package controllers

import (
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

type ReginController struct {
	BaseController
}

// 注册
func (c *ReginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(c.GetString("UName"))
		userpwd := strings.TrimSpace(c.GetString("UPasswd"))

		if len(username) == 0 || len(userpwd) == 0 {
			c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
		}
		//userpwd = utils.String2md5(userpwd)

		user, _ := models.BackendUserOneByUser(username)

		if user == nil {
			p := models.UProtIsZy() //获取端口

			lpBrookUser := models.LpBrookUser{
				UName:    username,
				UPasswd:  userpwd,
				UProt:    p,
				UIsAdmin: 0,
			}
			//数据库添加
			var uid int64
			if id, err := models.AddLpBrookUser(&lpBrookUser); err != nil {
				uid = id
				c.jsonResult(enums.JRCodeSucc, "注册异常", id)
			}
			//为刚刚注册的用开启服务
			lpBrookServer, errr := models.GetLpBrookAll() //获取可用服务器
			if errr != nil {
				c.jsonResult(enums.JRCodeSucc, "注册成功，但服务获取失败请联系管理员:(", "")
			} else {
				for _, v := range *lpBrookServer {
					//http请求
					req := httplib.Get("http://" + v.SIp + ":8080/remote/startservice")

					o := orm.NewOrm()
					sysMap := make(orm.Params)
					o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

					req.Param("remote_u", sysMap["remote_u"].(string))
					req.Param("remote_p", sysMap["remote_p"].(string))

					req.Param("uid", strconv.FormatInt(uid, 10))
					fmt.Println(req.String())
				}
			}

			//获取用户信息
			c.jsonResult(enums.JRCodeSucc, "注册成功", "")
		} else {
			c.jsonResult(enums.JRCodeFailed, "用户已存在", "")
		}
	} else {
		c.Data["title"] = "注册-Brook"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "regin/css.html"
		c.LayoutSections["footerjs"] = "regin/js.html"
		c.setTpl("regin/index.html", "shared/public_page.html")
	}

}
