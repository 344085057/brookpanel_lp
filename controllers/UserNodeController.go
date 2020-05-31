package controllers

import (
	"myBrookWeb/models"

	"github.com/astaxie/beego/orm"
)

type UserNodeController struct {
	BaseController
}

// 节点列表
func (c *UserNodeController) Index() {
	c.Data["title"] = "节点列表-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "usernode/css.html"
	c.LayoutSections["footerjs"] = "usernode/js.html"
	c.setTpl("usernode/index.html", "shared/public_page.html")
}

func (c *UserNodeController) GetList() {

	//page=1&limit=10
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	reportSo := c.GetString("reportso")
	reportSo = "%" + reportSo + "%"

	page = (page - 1) * limit

	o := orm.NewOrm()
	var lpBrookServer []models.LpBrookServer
	o.Raw("select * from lp_brook_server where (s_ip like ? or s_title like ? ) limit ?,?", reportSo, reportSo, page, limit).QueryRows(&lpBrookServer)

	var dataNum int64
	o.Raw("select count(s_id) from lp_brook_server where  (s_ip like ? or s_title like ? )", reportSo, reportSo).QueryRow(&dataNum)

	returnJson := make(map[string]interface{})
	returnJson["data"] = lpBrookServer
	returnJson["code"] = 0
	returnJson["msg"] = ""
	returnJson["count"] = dataNum

	c.Data["json"] = returnJson
	c.ServeJSON()
}
