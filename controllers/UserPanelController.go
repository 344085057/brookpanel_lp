package controllers

type UserPanelController struct {
	BaseController
}

// 用户面板首页
func (c *UserPanelController) Index() {
	c.Data["title"] = "面板-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userpanel/css.html"
	c.LayoutSections["footerjs"] = "userpanel/js.html"
	c.setTpl("userpanel/index.html", "shared/userpanel.html")

}
