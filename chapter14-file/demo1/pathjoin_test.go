package demo1

import (
	"fmt"
	"testing"
)

func TestGetWorkPath(t *testing.T) {
	logpath := "/var/lib/systemd/tg-paas"
	unitFileName := "cuk.default.cuktest11-cuk8master1.etcd.service"
	workpathExpect := "/var/lib/systemd/tg-paas/cuk/default/cuktest11-cuk8master1/etcd"
	workpath := getWorkPath(logpath, unitFileName)
	fmt.Println(workpath)

	if workpath == workpathExpect {
		fmt.Println("测试通过")
	} else {
		fmt.Println("测试失败")
	}
}
