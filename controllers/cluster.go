package controllers

import (
	"ClusterManager/models"

	"github.com/astaxie/beego"
)

type ClusterController struct {
	beego.Controller
}

// @Title Status
// @router /status [get]
func (c *ClusterController) Status() {
	c.Data["json"] = map[string]string{"role": models.CLUSTER_OBJ.Role}
	c.ServeJSON()
}

// @Title Role
// @router /role [get]
func (c *ClusterController) Role() {
	c.Data["json"] = models.CLUSTER_OBJ.Role
	c.ServeJSON()
}
