package utils

import (
	"github.com/astaxie/beego"
	"strings"
)

func GetHosts() []string {
	hosts_str := beego.AppConfig.String("cluster::hosts")
	hosts_list := strings.Split(hosts_str, ",")
	return hosts_list
}

func GetBindURL() string {
	host := beego.AppConfig.String("cluster::bind")
	port := beego.AppConfig.String("httpport")
	return host + ":" + port
}

func In(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
