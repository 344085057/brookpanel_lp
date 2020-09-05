package controllers

import (
	"myBrookWeb/models"
	"myBrookWeb/utils"
)

//AdminLoginLogController 登录日志
type AdminLoginLogController struct {
	BaseController
}

// Index
func (c *AdminLoginLogController) Index() {
	if c.Ctx.Request.Method == "POST" {

	} else {

		page, _ := c.GetInt("p") //当前页

		var obj models.LpBrookUserLoginLog
		if err := c.ParseForm(&obj); err != nil {
			c.Data["err"] = err.Error()
		}
		objInfoArr, totle, err := models.GetLoginLogPage(page, 10, obj)
		if err != nil {
			c.Data["err"] = err.Error()
		}

		p := utils.NewPaginator(c.Ctx.Request, 10, totle)
		c.Data["paginator"] = p
		c.Data["data"] = objInfoArr
		c.Data["sodata"] = obj

		c.Data["title"] = "用户登录日志-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "admin/loginlog/css.html"
		c.LayoutSections["footerjs"] = "admin/loginlog/js.html"
		c.setTpl("admin/loginlog/index.html", "shared/userpanel.html")
	}

}
