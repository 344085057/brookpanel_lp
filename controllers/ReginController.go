package controllers

import (
	"encoding/json"
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"myBrookWeb/utils"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/validation"
)

type ReginController struct {
	BaseController
}

// 注册
func (c *ReginController) Index() {
	ip := c.ip
	googleSecret := beego.AppConfig.String("secret")
	googleSitekey := beego.AppConfig.String("sitekey")

	if c.Ctx.Request.Method == "POST" {
		UEmail := strings.TrimSpace(c.GetString("UEmail"))
		UPasswd := strings.TrimSpace(c.GetString("UPasswd"))
		UPasswd = utils.String2md5(UPasswd)
		UName := strings.TrimSpace(c.GetString("UName"))
		token := strings.TrimSpace(c.GetString("token"))
		if token == "" {
			c.jsonResult(enums.JRCodeFailed, "人机身份验证失败", "")
		}
		//人机身份验证
		tokenUrl := `https://www.recaptcha.net/recaptcha/api/siteverify`
		tokenReq := httplib.Post(tokenUrl)
		tokenReq.Param("secret", googleSecret)
		tokenReq.Param("response", token)
		tokenReq.Param("remoteip", ip)
		tokenResponse, err := tokenReq.String()
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "人机身份验证失败，请求服务器错误", err.Error())
		}
		tokenResponseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(tokenResponse), &tokenResponseMap)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "请求远程api成功，但是数据解析失败", err.Error())
		}
		if tokenResponseMap["success"].(bool) == false {
			c.jsonResult(enums.JRCodeFailed, "远程人机身份验证失败", "")
		}

		u := models.LpBrookUserByRegin{UEmail: UEmail, UPasswd: UPasswd, UName: UName}

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

			msgStr := fmt.Sprintf("%s %s", alias, valid.Errors[0].Message)

			c.jsonResult(enums.JRCodeFailed, msgStr, "")
		}

		//userpwd = utils.String2md5(userpwd)

		user, _ := models.GetUserOneByEmail(u.UEmail)

		if user == nil {
			port := models.UPortIsZy() //获取端口

			lpBrookUser := models.LpBrookUser{
				Name:       UName,
				Passwd:     UPasswd,
				Email:      UEmail,
				Port:       port,
				IsAdmin:    0,
				ExpireTime: time.Now(),
			}
			//数据库添加
			// var uid int64
			if _, err := models.AddLpBrookUser(&lpBrookUser); err != nil {
				// uid = id
				c.jsonResult(enums.JRCodeFailed, "注册异常", err.Error())
			}
			//为刚刚注册的用开启服务
			// lpBrookServer, errr := models.GetLpBrookAll(0) //获取可用服务器
			// if errr != nil {
			// 	c.jsonResult(enums.JRCodeSucc, "注册成功，但服务获取失败请联系管理员:(", errr.Error())
			// }else {
			// 	for _, v := range lpBrookServer {
			// 		//http请求
			// 		req := httplib.Get("http://" + v.Ip + ":60001/remote/startservice")

			// 		o := orm.NewOrm()
			// 		sysMap := make(orm.Params)
			// 		o.Raw("SELECT s_name,s_value FROM lp_sys").RowsToMap(&sysMap, "s_name", "s_value")

			// 		req.Param("remote_u", sysMap["remote_u"].(string))
			// 		req.Param("remote_p", sysMap["remote_p"].(string))

			// 		req.Param("uid", strconv.FormatInt(uid, 10))
			// 		fmt.Println(req.String())
			// 	}
			// }

			//获取用户信息
			c.jsonResult(enums.JRCodeSucc, "注册成功", "")
		} else {
			c.jsonResult(enums.JRCodeFailed, "Email已存在", "")
		}
	} else {

		c.Data["googleSitekey"] = googleSitekey
		c.Data["title"] = "注册-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "regin/css.html"
		c.LayoutSections["footerjs"] = "regin/js.html"
		c.setTpl("regin/index.html", "shared/public_page.html")
	}

}
