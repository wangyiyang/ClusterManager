package controllers

import (
	"ClusterManager/models"

	"github.com/astaxie/beego"
)

type ClusterController struct {
	beego.Controller
}

// @Title Cluster
// @Description find object by objectid
// @Param   objectId        path    string  true        "the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /status [get]
func (c *ClusterController) Status() {
	models.CLUSTER_OBJ.Role = "master"
	c.Data["json"] = map[string]string{"status": models.CLUSTER_OBJ.Role}
	c.ServeJSON()
}
