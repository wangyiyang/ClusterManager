package models

var CLUSTER_OBJ Cluster = Cluster{}

type Cluster struct {
	Role string
}

func findMaster() {

}

func checkMaster() {

}

func setRole(role string) string {
	CLUSTER_OBJ.Role = role
	return CLUSTER_OBJ.Role
}

func init() {
	CLUSTER_OBJ.Role = "master"
}
