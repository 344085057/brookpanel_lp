package controllers

import (
	"fmt"
	"myBrookWeb/utils"
)

//UserStatisticsController 用户使用统计
type UserStatisticsController struct {
	BaseController
}

//Index 首页
func (c *UserStatisticsController) Index() {
	zFlowLog := make([]map[string]float64, 0)

	flowLog := make([]map[string][]float64, 0)
	utils.GetCache("flow_log_"+fmt.Sprintf("%v", c.curUser.Id), &flowLog)
	if len(flowLog) != 0 {
		for _, v := range flowLog {
			var totle float64
			var month_day string
			for k := range v {
				month_day = k
				for _, vv := range v[k] {
					totle = totle + vv
				}
			}
			zFlowLog = append(zFlowLog, map[string]float64{month_day: totle})
		}
	}

	// c.Data["user_info"] = c.curUser
	c.Data["周使用统计"] = zFlowLog
	c.Data["title"] = "统计-" + c.appname
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "userstatistics/css.html"
	c.LayoutSections["footerjs"] = "userstatistics/js.html"
	c.setTpl("userstatistics/index.html", "shared/userpanel.html")

}
