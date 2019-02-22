package controllers

import (
	"ClusterManager/models"

	"github.com/astaxie/beego"
)

type ClusterController struct {
	beego.Controller
}

// @Title Status
// @Description find object by objectid
// @Param   objectId        path    string  true        "the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /status [get]
func (c *ClusterController) Status() {
	c.Data["json"] = map[string]string{"status": models.CLUSTER_OBJ.Role}
	c.ServeJSON()
}

// @Title Role
// @Description find object by objectid
// @Param   objectId        path    string  true        "the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /role [get]
func (c *ClusterController) Role() {
	c.Data["json"] = models.CLUSTER_OBJ.Role
	c.ServeJSON()
}
