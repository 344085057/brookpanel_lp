package controllers

import (
	"myBrookWeb/enums"
	"myBrookWeb/models"
)

//UserNodeController 节点管理
type AdminUserNodeController struct {
	BaseController
}

// Index
func (c *AdminUserNodeController) Index() {
	if c.Ctx.Request.Method == "POST" {

	} else {

		//获取所有的节点
		brookServerArr, err := models.GetLpBrookAll(0)
		if err != nil {
			c.Data["err"] = err.Error()
		}

		c.Data["节点"] = brookServerArr
		// c.Data["端口"] = c.curUser.Port

		c.Data["title"] = "节点管理-" + c.appname
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["headcssjs"] = "admin/usernode/css.html"
		c.LayoutSections["footerjs"] = "admin/usernode/js.html"
		c.setTpl("admin/usernode/index.html", "shared/userpanel.html")
	}

}

//AE 添加/修改
func (c *AdminUserNodeController) AE() {

	var obj models.LpBrookServer

	//数据封装到对象中
	err := c.ParseForm(&obj)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	if obj.Id == 0 {
		//添加
		id, err := models.AddLpBrookServer(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "添加成功", id)
	} else {
		//修改
		models.UpdateLpBrookServerById(&obj)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "修改失败", err.Error())
		}
		c.jsonResult(enums.JRCodeSucc, "修改成功", obj)
	}

}

//Del 删除
func (c *AdminUserNodeController) Del() {

	id, err := c.GetInt("id")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}
	if id == 0 {
		c.jsonResult(enums.JRCodeFailed, "数据错误", "id is not 0")
	}
	err = models.DeleteLpBrookServer(id)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "删除失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "删除成功", id)

}
