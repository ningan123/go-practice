package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"github.com/spf13/pflag"
)

type runOpts struct {
	addr1      string
	addr2      string
	handleFunc string
}

func main() {
	opts := &runOpts{}
	pflag.StringVarP(&opts.addr1, "addr1", "a", "", "addr1")
	pflag.StringVarP(&opts.addr2, "addr2", "b", "", "addr2")
	pflag.StringVarP(&opts.handleFunc, "handleFunc", "h", "1", "handleFunc")
	pflag.Parse()

	go func() {
		if opts.handleFunc == "1" {
			http.HandleFunc("/debug/pprof/block", pprof.Index)
			http.HandleFunc("/debug/pprof/goroutine", pprof.Index)
			http.HandleFunc("/debug/pprof/heap", pprof.Index)
		}

		log.Println("Starting server on ", opts.addr2)
		if err := http.ListenAndServe(opts.addr2, nil); err != nil {
			log.Fatal(err)
		}
	}()

	//这里实现了远程获取pprof数据的接口
	log.Println("Starting server on ", opts.addr1)
	if err := http.ListenAndServe(opts.addr1, nil); err != nil {
	  log.Fatal(err)
	}
}
