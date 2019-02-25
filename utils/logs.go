package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLogs() {
	log_path := beego.AppConfig.String("logs::path")
	level := beego.AppConfig.String("logs::level")
	beego.SetLogger(logs.AdapterMultiFile, `{"filename":"`+log_path+`/root.log",
                                                    "separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
                                                            "level":`+level+`,
                                                                    "daily":true,
                                                                            "maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Async()
}
