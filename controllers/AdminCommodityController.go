package controllers

import (
	"myBrookWeb/enums"
	"myBrookWeb/models"
)

//AdminCommodityController 商品管理
type AdminCommodityController struct {
	BaseController
}

//Index 商品首页
func (c *AdminCommodityController) Index() {

	commodity := make([]map[string][]models.LpBrookCommodity, 0)

	//获取所有分类
	sortArr, err := models.GetSortAll()
	if err != nil {
		c.Data["err"] = err.Error()
	}

	//获取分类下的 商品
	for _, v := range sortArr {
		m := make(map[string][]models.LpBrookCommodity)
		lpBrookCommodityArr, err := models.GetCommodityAllBySort(v)
		if err != nil {
			c.Data["err"] = err.Error()
		}
		m[v] = lpBrookCommodityArr
		commodity = append(commodity, m)
	}

	c.Data["商品"] = commodity

	c.Data["title"] = "商品管理-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "admin/commodity/css.html"
	c.LayoutSections["footerjs"] = "admin/commodity/js.html"
	c.setTpl("admin/commodity/index.html", "shared/userpanel.html")

}

//AE 添加/修改
func (c *AdminCommodityController) AE() {

	var brookCommodity models.LpBrookCommodity

	//数据封装到对象中
	err := c.ParseForm(&brookCommodity)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	if brookCommodity.Id == 0 {
		//添加
		id, err := models.AddLpBrookCommodity(&brookCommodity)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "添加成功", id)
	} else {
		//修改
		models.UpdateLpBrookCommodityById(&brookCommodity)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "修改成功", brookCommodity)
	}

}

//Del 删除
func (c *AdminCommodityController) Del() {

	id, err := c.GetInt("id")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if id == 0 {
		c.jsonResult(enums.JRCodeFailed, "数据错误", "id is not 0")
	}
	err = models.DeleteLpBrookCommodity(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "删除失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "删除成功", id)

}
