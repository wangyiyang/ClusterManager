package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func PathExists(filePth string) (bool, error) {
	_, err := os.Stat(filePth)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func ReadAll(filePth string) (string, error) {
	f, err := os.Open(filePth)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	pid_byte, err := ioutil.ReadAll(f)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	pid_str := strings.Replace(string(pid_byte), "\n", "", -1)
	return pid_str, nil
}

func GetPid() (int, error) {
	pidPath := beego.AppConfig.String("cluster::pid_path")
	pid_str, _ := ReadAll(pidPath)
	pid, err := strconv.Atoi(pid_str)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	return pid, nil
}
