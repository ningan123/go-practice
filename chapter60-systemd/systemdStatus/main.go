package main

import (
    "fmt"
    "strings"

    "github.com/godbus/dbus/v5"
    "k8s.io/klog/v2"
)

// monitorAllServicesState 监听所有 systemd 服务的状态变化
func monitorAllServicesState() error {
    conn, err := dbus.SystemBus()
    if err != nil {
        return fmt.Errorf("failed to connect to system bus: %v", err)
    }
    defer conn.Close()

    // 添加匹配规则，监听 PropertiesChanged 信号
    matchRule := "type='signal',sender='org.freedesktop.systemd1'," +
        "interface='org.freedesktop.DBus.Properties',member='PropertiesChanged'"
    call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, matchRule)
    if call.Err != nil {
        return fmt.Errorf("failed to add match for PropertiesChanged: %v", call.Err)
    }

    klog.InfoS("Monitoring all systemd services state changes")

    // 创建信号通道
    c := make(chan *dbus.Signal, 10)
    conn.Signal(c)

    // 处理收到的信号
    for signal := range c {
        if len(signal.Body) < 3 {
            continue
        }

        // 检查信号是否来自 Unit 对象
        interfaceName, ok := signal.Body[0].(string)
        if !ok || interfaceName != "org.freedesktop.systemd1.Unit" {
            continue
        }

        // 获取服务的属性变化
        changedProps, ok := signal.Body[1].(map[string]dbus.Variant)
        if !ok {
            continue
        }

        // 获取服务名称
        serviceName := extractServiceName(string(signal.Path))

        // 检查 ActiveState 属性
        if activeState, exists := changedProps["ActiveState"]; exists {
            state, ok := activeState.Value().(string)
            if ok {
                klog.InfoS("Service state changed", "service", serviceName, "ActiveState", state)

                // 根据状态变化执行不同的处理逻辑
                switch state {
                case "active":
                    handleStart(serviceName)
                case "inactive":
                    handleStop(serviceName)
                }
            }
        }

        // 检查 SubState 属性
        if subState, exists := changedProps["SubState"]; exists {
            state, ok := subState.Value().(string)
            if ok {
                klog.InfoS("Service sub-state changed", "service", serviceName, "SubState", state)

                // 检查是否是重启状态
                if state == "auto-restart" {
                    handleRestart(serviceName)
                }
            }
        }
    }

    return nil
}

// extractServiceName 从 D-Bus 对象路径中提取服务名称
func extractServiceName(path string) string {
    parts := strings.Split(path, "/")
    if len(parts) > 0 {
        return strings.ReplaceAll(parts[len(parts)-1], "_", ".")
    }
    return ""
}

// handleStart 处理服务启动事件
func handleStart(serviceName string) {
    klog.InfoS("Handling start event", "service", serviceName)
    // 在这里添加启动服务的处理逻辑
}

// handleStop 处理服务停止事件
func handleStop(serviceName string) {
    klog.InfoS("Handling stop event", "service", serviceName)
    // 在这里添加停止服务的处理逻辑
}

// handleRestart 处理服务重启事件
func handleRestart(serviceName string) {
    klog.InfoS("Handling restart event", "service", serviceName)
    // 在这里添加重启服务的处理逻辑
}

func main() {
    klog.InitFlags(nil)
    defer klog.Flush()

    err := monitorAllServicesState()
    if err != nil {
        klog.ErrorS(err, "Failed to monitor all services state")
    }
}