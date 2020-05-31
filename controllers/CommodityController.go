package controllers

type CommodityController struct {
	BaseController
}

// 商品
func (c *CommodityController) Index() {

	c.Data["title"] = "Brook商品列表"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "commodity/css.html"
	c.LayoutSections["footerjs"] = "commodity/js.html"
	c.setTpl("commodity/index.html", "shared/userpanel.html")

}
