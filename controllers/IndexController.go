package controllers

type IndexController struct {
	BaseController
}

// Index test
func (c *IndexController) Index() {

	c.Data["title"] = "Brook首页"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "index/css.html"
	c.LayoutSections["footerjs"] = "index/js.html"
	c.setTpl("index/index.html", "shared/public_page.html")

}
