package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
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
 
func serverArchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Fprintf(w, "GOARCH: %s\n", runtime.GOARCH)  
}
func main() {  	
	// 第一个HTTP服务  
	mux1 := http.NewServeMux()  
	mux1.HandleFunc("/", serverIPHandler)  
	go func() {  
		log.Println("Starting first server on port 8080")  
		if err := http.ListenAndServe(":8080", mux1); err != nil {  
			log.Fatal(err)  
		}  
	}()  
  
	// 第二个HTTP服务  
	mux2 := http.NewServeMux()  
	mux2.HandleFunc("/", serverArchHandler)  
	log.Println("Starting second server on port 8081")  
	if err := http.ListenAndServe(":8081", mux2); err != nil {  
		log.Fatal(err)  
	} 
}