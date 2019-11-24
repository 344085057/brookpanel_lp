package controllers

type IndexController struct {
	BaseController
}

// Index
func (c *IndexController) Index() {

	c.Data["title"] = "Brook首页"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "index/css.html"
	c.LayoutSections["footerjs"] = "index/js.html"
	c.setTpl("index/index.html", "shared/layui_page.html")

}
