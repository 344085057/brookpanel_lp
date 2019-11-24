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

	//用户面板 iframe框架
	beego.Router("/user/userpanel", &controllers.UserPanelController{}, "*:Index")

	//用户首页
	// iframe框架中的内容
	beego.Router("/user/userindex", &controllers.UserIndexController{}, "*:Index")

	//用户节点
	// iframe框架中的内容
	beego.Router("/user/nodeindex", &controllers.UserNodeController{}, "*:Index")
	beego.Router("/user/nodegetlist", &controllers.UserNodeController{}, "*:GetList")
}
