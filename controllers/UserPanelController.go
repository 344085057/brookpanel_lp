package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"myBrookWeb/enums"
	"myBrookWeb/models"
	"time"

	"github.com/astaxie/beego/orm"
)

//UserPanelController 用户面板
type UserPanelController struct {
	BaseController
}

//Index 用户面板首页
func (c *UserPanelController) Index() {

	user, err := models.GetLpBrookUserById(c.curUser.Id)
	if err != nil {
		c.Data["err"] = err.Error()
	}
	c.Data["user_info"] = user

	nowTime := time.Now()
	d := user.ExpireTime.Sub(nowTime)

	if d.Hours() <= 0 {
		c.Data["剩余使用时间"] = fmt.Sprintf("VIP已到期")
	} else if d.Hours()/24 >= 1.00 {
		c.Data["剩余使用时间"] = fmt.Sprintf("%.2f 天", d.Hours()/24)
	} else {
		c.Data["剩余使用时间"] = fmt.Sprintf("%.2f 小时", d.Hours())
	}

	//获取启用的公告
	gg, err := models.GetLpBrookAllBygTypeAndState(1, 1)
	if err != nil {
		c.Data["err"] = err.Error()
	}
	c.Data["公告"] = gg

	//获取启用的教程
	jc, err := models.GetLpBrookAllBygTypeAndState(2, 1)
	if err != nil {
		c.Data["err"] = err.Error()
	}
	c.Data["教程"] = jc

	//流量使用情况
	c.Data["流量使用情况"] = jc

	// c.Data["user_info"] = c.curUser
	c.Data["title"] = "面板-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userpanel/css.html"
	c.LayoutSections["footerjs"] = "userpanel/js.html"
	c.setTpl("userpanel/index.html", "shared/userpanel.html")

}

//SeeGG 查看使用教程和公告
func (c *UserPanelController) SeeGG() {
	ggID, err := c.GetInt("id")
	if err != nil {
		c.Data["err"] = err.Error()
	}

	lpBrookGg, err := models.GetLpBrookGgById(ggID)
	if err != nil {
		c.Data["err"] = err.Error()
	}
	if lpBrookGg == nil {
		c.Data["err"] = errors.New("未知的id")

	}

	c.Data["文章"] = lpBrookGg

	c.Data["title"] = lpBrookGg.Title + "-Brook"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userseegg/css.html"
	c.LayoutSections["footerjs"] = "userseegg/js.html"
	c.setTpl("userseegg/index.html", "shared/userpanel.html")
}

//Recharge 用户使用cdk充值
func (c *UserPanelController) Recharge() {
	obj := make(map[string]string)
	data := c.Ctx.Input.RequestBody

	//json数据封装到map中
	err := json.Unmarshal(data, &obj)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "解析数据错误", err.Error())
	}

	if obj["cdk"] == "" {
		c.jsonResult(enums.JRCodeFailed, "cdk不能为空", "")
	}

	moneycdk, err := models.GetLpBrookMoneycdkByCdk(obj["cdk"])
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	if moneycdk == nil {
		c.jsonResult(enums.JRCodeFailed, "充值码不存在:(", "")
	}
	if moneycdk.UseUid != 0 {
		c.jsonResult(enums.JRCodeFailed, "充值码已经被使用了:(", "")
	}

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	// 事务处理过程
	user := models.LpBrookUser{
		Id:    c.curUser.Id,
		Money: c.curUser.Money + moneycdk.Money,
	}
	//为用户充值
	if _, err = o.Update(&user, "Money"); err != nil {
		o.Rollback()
		c.jsonResult(enums.JRCodeFailed, "充值失败", err.Error())
	}
	//更新充值卡为已使用状态
	updataMoneycdk := models.LpBrookMoneycdk{
		Id:     moneycdk.Id,
		UseUid: c.curUser.Id,
	}
	if _, err = o.Update(&updataMoneycdk, "UseUid"); err != nil {
		o.Rollback()
		c.jsonResult(enums.JRCodeFailed, "充值失败", err.Error())
	}
	err = o.Commit()
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "失败", err.Error())
	}
	c.jsonResult(enums.JRCodeSucc, "充值成功:)", "")
}
