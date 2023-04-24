package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

var (
	templateFile string
	varsFile     string
	outputFile   string
)

func WriteFile(filePath string, targetFileName, content string) {

	var d1 = []byte(content)
	os.Mkdir(filePath, os.ModePerm)
	f, err3 := os.Create(filePath + "/" + targetFileName) //创建文件
	if err3 != nil {
		fmt.Println("create file fail")
	}
	defer f.Close()
	n2, err3 := f.Write(d1) //写入文件(字节数组)

	fmt.Printf("写入 %d 个字节\n", n2)
	//n3, err3 := f.WriteString("writesn") //写入文件(字节数组)
	//fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()
}

func main() {
	flag.StringVar(&templateFile, "templateFile", "input.txt", "模板文件")
	flag.StringVar(&varsFile, "varsFile", "var.txt", "变量文件")
	flag.StringVar(&outputFile, "outputFile", "output.txt", "输出文件")
	flag.Parse()

	// inputBuf, err := os.ReadFile(*input)
	inputBuf, err := os.ReadFile(templateFile)
	if err != nil {
		log.Println(err)
		return
	}

	tmpl, err := template.New("input").Parse(string(inputBuf))
	if err != nil {
		log.Println(err)
	}
	envs := make(map[string]string, 128)
	fp, err := os.Open(varsFile)
	if err != nil {
		fmt.Println(err) //打开文件错误
		return
	}

	buf := bufio.NewScanner(fp)
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		fmt.Println(line)
		keys := strings.SplitN(line, "=", 2)
		if len(keys) == 2 {
			envs[keys[0]] = keys[1]
		}
	}
	fmt.Println(envs)

	var bufOut bytes.Buffer
	err = tmpl.Execute(&bufOut, envs)
	fmt.Println(bufOut.String())
	// err = tmpl.Execute(os.Stdout, envs)
	if err != nil {
		log.Println(err)
		return
	}

	WriteFile(".", outputFile, bufOut.String())
}

// func main() {
// 	// inputBuf, err := os.ReadFile(*input)
// 	inputBuf, err := os.ReadFile("input.txt")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	tmpl, err := template.New("input").Parse(string(inputBuf))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	envs := make(map[string]string, 128)
// 	for _, kv := range os.Environ() {
// 		keys := strings.SplitN(kv, "=", 2)
// 		if len(keys) == 2 {
// 			envs[keys[0]] = keys[1]
// 		}
// 	}

// 	// var buf bytes.Buffer
// 	// tpl.Execute(os.Stdout, zoo)
// 	err = tmpl.Execute(os.Stdout, envs)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }
