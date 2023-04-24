package demo1

import (
	"path/filepath"
	"strings"
)

// 输入参数：
// 1）logpath: /var/lib/systemd/tg-paas
// 2）unitFileName: cuk.default.cuktest11-cuk8master1.etcd.service
//
// 返回：
// 1）workpath: /var/lib/systemd/tg-paas/cuk/default/cuktest11-cuk8master1/etcd
func getWorkPath(logPath, unitFileName string) string {

	array := strings.Split(unitFileName, ".")
	var workPath string
	if array != nil || len(array) > 0 {
		productname := array[0]
		namespace := array[1]
		nodename := array[2] // 包含了clustername
		svcname := array[len(array)-2]
		workPath = filepath.Join(logPath, productname, namespace, nodename, svcname)
	}

	return workPath
}
