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
			return
		}

		user, err := models.GetLpBrookUserById(u.Id)
		if err != nil {
			ctx.Redirect(302, "/login")
		}
		ctx.Input.CruSession.Set("user", *user)
	}
	beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)

	//后台管理
	var AdminFilterUser = func(ctx *context.Context) {
		u, _ := ctx.Input.Session("user").(models.LpBrookUser)
		if u.Id == 0 {
			ctx.Redirect(302, "/login")
			return
		}

		user, err := models.GetLpBrookUserById(u.Id)
		if err != nil {
			ctx.Redirect(302, "/login")
		}
		if user.IsAdmin != 1 {
			ctx.Redirect(302, "/login")
		}
		ctx.Input.CruSession.Set("user", *user)
	}
	beego.InsertFilter("/admin/*", beego.BeforeRouter, AdminFilterUser)

	//首页
	beego.Router("/", &controllers.IndexController{}, "*:Index")

	//登录
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

	//注册
	beego.Router("/regin", &controllers.ReginController{}, "*:Index")

	//用户面板
	beego.Router("/user/userpanel", &controllers.UserPanelController{}, "*:Index")
	//查看文章
	beego.Router("/user/seegg", &controllers.UserPanelController{}, "*:SeeGG")
	//用户充值
	beego.Router("/user/recharge", &controllers.UserPanelController{}, "Post:Recharge")

	//用户 节点
	beego.Router("/user/usernode", &controllers.UserNodeController{}, "*:Index")

	//用户 使用统计
	beego.Router("/user/userstatistics", &controllers.UserStatisticsController{}, "*:Index")

	//商店
	beego.Router("/user/commodity", &controllers.UserCommodityController{}, "*:Index")
	//用户购买商品api
	beego.Router("/user/shopping", &controllers.UserCommodityController{}, "Post:Shopping")

	//用户更改个人信息
	beego.Router("/user/userupdate", &controllers.UserUpdateController{}, "*:Index")
	//用户更新密码api
	beego.Router("/user/updatepasswd", &controllers.UserUpdateController{}, "Post:UpdataPasswd")
	//用户更新连接密码api
	beego.Router("/user/updateporxypasswd", &controllers.UserUpdateController{}, "Post:UpdataPorxyPasswd")
	//用户修改端口api
	// beego.Router("/user/updateport", &controllers.UserUpdateController{}, "Post:UpdataPort")
	//用户退出登录
	beego.Router("/user/logout", &controllers.UserUpdateController{}, "*:Logout")

	//节点管理
	beego.Router("/admin/usernode", &controllers.AdminUserNodeController{}, "*:Index")
	beego.Router("/admin/usernode/ae", &controllers.AdminUserNodeController{}, "Post:AE")
	beego.Router("/admin/usernode/del", &controllers.AdminUserNodeController{}, "Post:Del")

	//商品管理
	beego.Router("/admin/commodity", &controllers.AdminCommodityController{}, "*:Index")
	beego.Router("/admin/commodity/ae", &controllers.AdminCommodityController{}, "Post:AE")
	beego.Router("/admin/commodity/del", &controllers.AdminCommodityController{}, "Post:Del")

	//用户管理
	beego.Router("/admin/user", &controllers.AdminUserController{}, "*:Index")
	beego.Router("/admin/user/ae", &controllers.AdminUserController{}, "Post:AE")
	beego.Router("/admin/user/del", &controllers.AdminUserController{}, "Post:Del")

	//CDK管理
	beego.Router("/admin/cdk", &controllers.AdminMoneycdkController{}, "*:Index")
	beego.Router("/admin/cdk/ae", &controllers.AdminMoneycdkController{}, "Post:AE")
	beego.Router("/admin/cdk/del", &controllers.AdminMoneycdkController{}, "Post:Del")

	//登录日志
	beego.Router("/admin/loginlog", &controllers.AdminLoginLogController{}, "*:Index")

}
