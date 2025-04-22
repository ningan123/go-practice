package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/godbus/dbus/v5"
	"k8s.io/klog/v2"
)

func getJobType(conn *dbus.Conn, jobID uint32) (string, error) {
	systemd := conn.Object("org.freedesktop.systemd1", dbus.ObjectPath("/org/freedesktop/systemd1"))

	// 获取 job object path
	var jobPath dbus.ObjectPath
	err := systemd.Call("org.freedesktop.systemd1.Manager.GetJob", 0, jobID).Store(&jobPath)
	if err != nil {
		return "", fmt.Errorf("failed to get job path: %v", err)
	}

	jobObj := conn.Object("org.freedesktop.systemd1", jobPath)

	// 读取 Job 的 Type 属性
	var jobTypeVariant dbus.Variant
	err = jobObj.Call("org.freedesktop.DBus.Properties.Get", 0,
		"org.freedesktop.systemd1.Job", "Type").Store(&jobTypeVariant)
	if err != nil {
		return "", fmt.Errorf("failed to get job type: %v", err)
	}

	jobType, ok := jobTypeVariant.Value().(string)
	if !ok {
		return "", fmt.Errorf("unexpected job type format")
	}

	return jobType, nil
}

func listenServiceActions(handler func(unit string, action string)) error {
	conn, err := dbus.SystemBus()
	if err != nil {
		return fmt.Errorf("failed to connect to system bus: %v", err)
	}

	// 匹配 JobRemoved 信号
	matchRule := "type='signal',sender='org.freedesktop.systemd1'," +
		"interface='org.freedesktop.systemd1.Manager',member='JobRemoved'"
	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, matchRule)
	if call.Err != nil {
		return fmt.Errorf("failed to add match: %v", call.Err)
	}

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	klog.InfoS("Listening for actions on all systemd services")

	for sig := range c {
		if len(sig.Body) < 4 {
			continue
		}

		jobID := sig.Body[0].(uint32)
		unit := sig.Body[1].(string)
		result := sig.Body[3].(string)

		jobType, err := getJobType(conn, jobID)
		if err != nil {
			klog.ErrorS(err, "Error getting job type")
			continue
		}

		klog.InfoS("Job detected", "jobID", jobID, "unit", unit, "type", jobType, "result", result)

		if result == "done" {
			handler(unit, strings.ToLower(jobType))
		}
	}

	return nil
}

func main() {

	klog.InitFlags(nil)
	flag.Parse()
	flag.Set("v", "2") // 设置默认日志级别
	defer klog.Flush()

	go func() {
		err := listenServiceActions(func(unit string, action string) {
			klog.InfoS("Action handler invoked", "unit", unit, "action", action)
		})
		if err != nil {
			klog.ErrorS(err, "Failed to listen service actions")
			os.Exit(1)
		}
	}()

	// 优雅退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	klog.InfoS("Shutting down")
}
