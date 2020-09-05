package controllers

import "myBrookWeb/models"

//UserNodeController 节点
type UserNodeController struct {
	BaseController
}

// Index
func (c *UserNodeController) Index() {
	if c.Ctx.Request.Method == "POST" {

	} else {

		//获取启用的节点
		brookServerArr, err := models.GetLpBrookAll(1)
		if err != nil {
			c.Data["err"] = err.Error()
		}

		c.Data["节点"] = brookServerArr
		c.Data["端口"] = c.curUser.Port

		c.Data["title"] = "节点-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "usernode/css.html"
		c.LayoutSections["footerjs"] = "usernode/js.html"
		c.setTpl("usernode/index.html", "shared/userpanel.html")
	}

}
