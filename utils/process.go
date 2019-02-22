package utils

func ProcessExist(pid int) bool {
	if err = syscall.Kill(pid, 0); err == nil {
		err = ErrProcessExists
		return true
	} else {
		return false
	}

}
