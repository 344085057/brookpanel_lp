package controllers

import (
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"myBrookWeb/utils"
)

//AdminMoneycdkController CDK管理
type AdminMoneycdkController struct {
	BaseController
}

// Index
func (c *AdminMoneycdkController) Index() {
	if c.Ctx.Request.Method == "POST" {

	} else {

		page, _ := c.GetInt("p") //当前页

		var obj models.LpBrookMoneycdk
		if err := c.ParseForm(&obj); err != nil {
			c.Data["err"] = err.Error()
		}
		objInfoArr, totle, err := models.GetCDKPage(page, 10, obj)
		if err != nil {
			c.Data["err"] = err.Error()
		}

		p := utils.NewPaginator(c.Ctx.Request, 10, totle)
		c.Data["paginator"] = p
		c.Data["data"] = objInfoArr
		c.Data["sodata"] = obj

		c.Data["title"] = "CDK管理-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "admin/cdk/css.html"
		c.LayoutSections["footerjs"] = "admin/cdk/js.html"
		c.setTpl("admin/cdk/index.html", "shared/userpanel.html")
	}

}

//AE 添加/修改
func (c *AdminMoneycdkController) AE() {

	var obj models.LpBrookMoneycdk

	//数据封装到对象中
	err := c.ParseForm(&obj)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	if obj.Id == 0 {
		cdk, err := models.GetLpBrookMoneycdkByCdk(obj.Cdk)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", err.Error())
		}
		if cdk != nil {
			c.jsonResult(enums.JRCodeFailed, "CDK已经存在", "")
		}

		//添加
		id, err := models.AddLpBrookMoneycdk(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "添加成功", id)
	} else {
		cdk, err := models.GetLpBrookMoneycdkById(obj.Id)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
		}
		if cdk == nil {
			c.jsonResult(enums.JRCodeFailed, "CDK已经不存在", "")
		}
		if cdk.Cdk != obj.Cdk {
			cdk, err := models.GetLpBrookMoneycdkByCdk(obj.Cdk)
			if err != nil {
				c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
			}
			if cdk != nil {
				c.jsonResult(enums.JRCodeFailed, "CDK已经存在", "")
			}
		}
		//修改
		models.UpdateLpBrookMoneycdkById(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "修改成功", obj)

	}

}

//Del 删除
func (c *AdminMoneycdkController) Del() {

	id, err := c.GetInt("id")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if id == 0 {
		c.jsonResult(enums.JRCodeFailed, "数据错误", "id is not 0")
	}
	err = models.DeleteLpBrookMoneycdk(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "删除失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "删除成功", id)

}
