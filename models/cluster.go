package models

import (
	"sync"
)

var CLUSTER_OBJ *Cluster
var once sync.Once

func GetIns() *Cluster {
	once.Do(func() {
		CLUSTER_OBJ = &Cluster{}
	})
	return CLUSTER_OBJ
}

type Cluster struct {
	Role string
}

func init() {
	CLUSTER_OBJ.Role = "master"
}
