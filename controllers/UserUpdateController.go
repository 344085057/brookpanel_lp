package controllers

type UserUpdateController struct {
	BaseController
}

// 更改个人信息　页面
func (c *UserUpdateController) Index() {
	c.Data["title"] = "面板-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userupdate/css.html"
	c.LayoutSections["footerjs"] = "userupdate/js.html"
	c.setTpl("userupdate/index.html", "shared/userpanel.html")

}
