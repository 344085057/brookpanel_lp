package routers

import (
	"myBrookWeb/controllers"
	"myBrookWeb/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	//未登录的过滤器
	var FilterUser = func(ctx *context.Context) {
		u, _ := ctx.Input.Session("user").(models.LpBrookUser)
		if u.Id == 0 {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)

	//首页
	beego.Router("/", &controllers.IndexController{}, "*:Index")

	//登录
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

	//注册
	beego.Router("/regin", &controllers.ReginController{}, "*:Index")

	//用户面板
	beego.Router("/user/userpanel", &controllers.UserPanelController{}, "*:Index")

	//商店
	beego.Router("/user/commodity", &controllers.CommodityController{}, "*:Index")

	//用户更改个人信息
	beego.Router("/user/userupdate", &controllers.UserUpdateController{}, "*:Index")

	//我的工单
	beego.Router("/user/obj", &controllers.UserObjComtroller{}, "*:Index")

	//用户首页
	// iframe框架中的内容
	beego.Router("/user/userindex", &controllers.UserIndexController{}, "*:Index")

	//用户节点
	// iframe框架中的内容
	beego.Router("/user/nodeindex", &controllers.UserNodeController{}, "*:Index")
	beego.Router("/user/nodegetlist", &controllers.UserNodeController{}, "*:GetList")
}
