package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(
		new(LpBrookUser),
		new(LpBrookServer),
		new(LpBrookGg),
		new(LpBrookCommodity),
		new(LpBrookMoneycdk),
		new(LpBrookUserLoginLog),
	)
}

//TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

//LpBrookUserTBName 用户
func LpBrookUserTBName() string {
	return TableName("brook_user")
}

//LpBrookServerTBName 服务器
func LpBrookServerTBName() string {
	return TableName("brook_server")
}

//LpBrookGGTableName 公告
func LpBrookGGTableName() string {
	return TableName("brook_gg")
}

//LpBrookCommodityTableName 商品
func LpBrookCommodityTableName() string {
	return TableName("brook_commodity")
}

//LpBrookMoneycdkTableName 商品
func LpBrookMoneycdkTableName() string {
	return TableName("brook_moneycdk")
}

//LpBrookUserLoginLogTableName 登录日志
func LpBrookUserLoginLogTableName() string {
	return TableName("brook_user_login_log")
}
