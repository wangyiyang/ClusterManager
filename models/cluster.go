package models

import (
	"ClusterManager/utils"
	"fmt"
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
		fmt.Println(masters)
		if len(masters) > 1 && utils.In(masters, CLUSTER_OBJ.BindURL) {
			fmt.Println("len(masters) > 1 && utils.In(masters, CLUSTER_OBJ.BindURL)")
			SetRole("slave")
		}
		if len(masters) == 0 {
			fmt.Println("len(masters) == 0")
			SetRole("master")
		}
	}
}

func CheckProcess() {
	pid, err := utils.GetPid()
	fmt.Println(pid)
	fmt.Println(err)
	if !utils.ProcessExist(pid) || err != nil {
		CLUSTER_OBJ.ProcessStatus = false
		fmt.Println("!utils.ProcessExist")
		SetRole("error")
	}
	if utils.ProcessExist(pid) && CLUSTER_OBJ.ProcessStatus == false {
		CLUSTER_OBJ.ProcessStatus = true
		SetRole("slave")
	}
}

func SetRole(role string) string {
	CLUSTER_OBJ.Role = role
	return CLUSTER_OBJ.Role
}

func init() {
	CLUSTER_OBJ.Role = "slave"
	CLUSTER_OBJ.BindURL = utils.GetBindURL()
	CLUSTER_OBJ.Hosts = utils.GetHosts()
}
