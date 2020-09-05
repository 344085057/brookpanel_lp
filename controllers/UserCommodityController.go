package controllers

import (
	"encoding/json"
	"myBrookWeb/enums"
	"myBrookWeb/models"
)

type UserCommodityController struct {
	BaseController
}

//Index 商品首页
func (c *UserCommodityController) Index() {

	commodity := make([]map[string][]models.LpBrookCommodity, 0)

	//获取所有分类
	sortArr, err := models.GetSortAllByState1()
	if err != nil {
		c.Data["err"] = err.Error()
	}

	//获取分类下的 商品
	for _, v := range sortArr {
		m := make(map[string][]models.LpBrookCommodity)
		lpBrookCommodityArr, err := models.GetCommodityArrBySort(v)
		if err != nil {
			c.Data["err"] = err.Error()
		}
		m[v] = lpBrookCommodityArr
		commodity = append(commodity, m)
	}

	c.Data["商品"] = commodity

	c.Data["title"] = "商品列表-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "commodity/css.html"
	c.LayoutSections["footerjs"] = "commodity/js.html"
	c.setTpl("commodity/index.html", "shared/userpanel.html")

}

//Shopping 购买商品
func (c *UserCommodityController) Shopping() {
	//commodityId 商品id
	obj := make(map[string]int)

	data := c.Ctx.Input.RequestBody

	//json数据封装到user对象中
	err := json.Unmarshal(data, &obj)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if obj["id"] == 0 {
		c.jsonResult(enums.JRCodeFailed, "id不能为0", "")
	}
	lpBrookCommodity, err := models.GetLpBrookCommodityById(obj["id"])
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, err.Error(), "")
	}
	if lpBrookCommodity == nil {
		c.jsonResult(enums.JRCodeFailed, "未知商品", "")
	}
	if lpBrookCommodity.State != 1 {
		c.jsonResult(enums.JRCodeFailed, "商品未启用", "")
	}

	err = models.UserShopping(c.curUser.Id, *lpBrookCommodity)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, err.Error(), "")
	}

	c.jsonResult(enums.JRCodeSucc, "购买成功", "")
}
