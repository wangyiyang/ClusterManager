package task

import (
	"ClusterManager/models"

	"github.com/astaxie/beego/toolbox"
)

func Run() {
	checkPid := toolbox.NewTask("checkPid", "0/10 * * * * *", func() error {
		models.CheckProcess()
		return nil
	})
	checkMaster := toolbox.NewTask("checkMaster", "0/10 * * * * *", func() error {
		models.CheckMaster()
		return nil
	})

	toolbox.AddTask("checkPid", checkPid)
	toolbox.AddTask("checkMaster", checkMaster)

	toolbox.StartTask()
}
