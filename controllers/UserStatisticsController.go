package controllers

//UserStatisticsController 用户使用统计
type UserStatisticsController struct {
	BaseController
}

//Index 首页
func (c *UserStatisticsController) Index() {

	// c.Data["user_info"] = c.curUser
	c.Data["title"] = "统计-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userstatistics/css.html"
	c.LayoutSections["footerjs"] = "userstatistics/js.html"
	c.setTpl("userstatistics/index.html", "shared/userpanel.html")

}
