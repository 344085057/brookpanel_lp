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
	)
}

// // TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

func LpBrookUserTBName() string {
	return TableName("brook_user")
}

func LpBrookServerTBName() string {
	return TableName("brook_server")
}
