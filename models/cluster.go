package models

import (
	"ClusterManager/utils"
	"github.com/astaxie/beego/logs"
)

var CLUSTER_OBJ Cluster = Cluster{}

type Cluster struct {
	Role          string
	BindURL       string
	Hosts         []string
	LocalPid      int
	ProcessStatus bool
}

func FindMaster() []string {
	result := []string{}
	hosts := utils.GetHosts()

	for _, host := range hosts {
		if host != CLUSTER_OBJ.BindURL {
			url := "http://" + host + "/v1/cluster/role"
			role, _ := utils.HttpGet(url)
			if role == "\"master\"" {
				result = append(result, host)
			}
		}
	}
	if CLUSTER_OBJ.Role == "master" {
		result = append(result, CLUSTER_OBJ.BindURL)
	}
	return result
}

func CheckMaster() {
	if CLUSTER_OBJ.ProcessStatus != false {
		masters := FindMaster()
		if len(masters) > 1 && utils.In(masters, CLUSTER_OBJ.BindURL) {
			logs.Error("master count:", len(masters))
			SetRole("slave")
		}
		if len(masters) == 0 {
			logs.Error("no master")
			SetRole("master")
		}
	}
}

func CheckProcess() {
	logs.Info("start CheckProcess()")
	pid, err := utils.GetPid()
	if !utils.ProcessExist(pid) || err != nil {
		logs.Error("Process Not Exist")
		CLUSTER_OBJ.ProcessStatus = false
		SetRole("error")
	}
	if utils.ProcessExist(pid) && CLUSTER_OBJ.ProcessStatus == false {
		CLUSTER_OBJ.ProcessStatus = true
		SetRole("slave")
	}
}

func SetRole(role string) string {
	logs.Info("set role", role)
	CLUSTER_OBJ.Role = role
	return CLUSTER_OBJ.Role
}

// TODO: 多线程单例
func init() {
	logs.Info("init cluster role")
	SetRole("slave")
	CLUSTER_OBJ.BindURL = utils.GetBindURL()
	CLUSTER_OBJ.Hosts = utils.GetHosts()
}
