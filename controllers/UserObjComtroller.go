package controllers

type UserObjComtroller struct {
	BaseController
}

// 我的工单　页面
func (c *UserObjComtroller) Index() {
	c.Data["title"] = "俺的工单-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userobj/css.html"
	c.LayoutSections["footerjs"] = "userobj/js.html"
	c.setTpl("userobj/index.html", "shared/userpanel.html")

}
