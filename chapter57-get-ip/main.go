package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)  
  
func serverIPHandler(w http.ResponseWriter, r *http.Request) {  
	// 获取本地网络接口列表
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces:", err)
		return
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 忽略回环接口（localhost）
		if iface.Flags&net.FlagUp == 0 {
			continue // 接口未启用
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // 忽略回环接口
		}

		// 获取接口上的地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error fetching addresses for interface:", iface.Name, err)
			continue
		}

		// 遍历地址列表
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 确保IP不是IPv6的未指定地址（::）或IPv4的0.0.0.0
			if ip == nil || ip.IsLoopback() || ip.IsUnspecified()  {
				continue
			}

			
			// 打印非回环和非未指定的IP地址
			fmt.Printf("Interface Name: %s, IP Address: %s\n", iface.Name, ip)
			fmt.Fprintf(w, "Interface: %s, IP: %s\n", iface.Name, ip)  
		}
	}

}  
  
func main() {  
	http.HandleFunc("/", serverIPHandler)  
  
	// 监听并启动HTTP服务器  
	port := ":8080"  
	log.Printf("Starting HTTP server on port %s", port)  
	if err := http.ListenAndServe(port, nil); err != nil {  
		log.Fatal(err)  
	}  
}