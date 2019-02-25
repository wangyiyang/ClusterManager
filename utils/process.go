package utils

import (
	"github.com/astaxie/beego/logs"
	"syscall"
)

func ProcessExist(pid int) bool {
	if err := syscall.Kill(pid, 0); err == nil {
		logs.Error(err)
		return true
	} else {
		return false
	}

}
