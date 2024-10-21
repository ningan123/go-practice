
go run hEqual1/main.go -a="0.0.0.0:6010" -b="0.0.0.0:6011" 
curl http://192.168.20.201:6010/debug/pprof/goroutine?debug=1
curl http://192.168.20.201:6011/debug/pprof/goroutine?debug=1

go run hEqual0/main.go -a="0.0.0.0:6020" -b="0.0.0.0:6021" 
curl http://192.168.20.201:6020/debug/pprof/goroutine?debug=1
curl http://192.168.20.201:6021/debug/pprof/goroutine?debug=1