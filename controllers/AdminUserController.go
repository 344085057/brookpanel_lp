package controllers

import (
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"myBrookWeb/utils"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

//AdminUserController 用户管理
type AdminUserController struct {
	BaseController
}

// Index
func (c *AdminUserController) Index() {
	if c.Ctx.Request.Method == "POST" {

	} else {

		page, _ := c.GetInt("p") //当前页

		var user models.LpBrookUser
		if err := c.ParseForm(&user); err != nil {
			c.Data["err"] = err.Error()
		}
		userInfoArr, totle, err := models.GetUserPage(page, 10, user)
		if err != nil {
			c.Data["err"] = err.Error()
		}

		p := utils.NewPaginator(c.Ctx.Request, 10, totle)
		c.Data["paginator"] = p
		c.Data["data"] = userInfoArr
		c.Data["sodata"] = user

		c.Data["title"] = "用户管理-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "admin/user/css.html"
		c.LayoutSections["footerjs"] = "admin/user/js.html"
		c.setTpl("admin/user/index.html", "shared/userpanel.html")
	}

}

//AE 添加/修改
func (c *AdminUserController) AE() {

	var obj models.LpBrookUser

	//数据封装到对象中
	err := c.ParseForm(&obj)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	if obj.Id == 0 {
		//添加
		id, err := models.AddLpBrookUser(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "添加成功", id)
	} else {
		//修改
		models.UpdateLpBrookUserById(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "修改成功", obj)
	}

}

//Del 删除
func (c *AdminUserController) Del() {

	id, err := c.GetInt("id")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if id == 0 {
		c.jsonResult(enums.JRCodeFailed, "数据错误", "id is not 0")
	}
	err = models.DeleteLpBrookUser(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "删除失败", err.Error())
	}

	lpBrookServer, err := models.GetLpBrookAll(0) //获取可用服务器
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "关闭服务时出错", err.Error())
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
	c.jsonResult(enums.JRCodeSucc, "删除成功", id)

}
