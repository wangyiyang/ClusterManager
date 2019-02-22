package main

import (
	_ "ClusterManager/routers"
	"ClusterManager/task"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	task.Run()
	beego.Run()
}
