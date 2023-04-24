package main

import "fmt"

func main() {
	var x interface{}
	var f float32 = 1.1
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", f)
	fmt.Println()

	x = f //空接口，可以接收任意类型
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", f)
	fmt.Println()

	y := x.(float32) // 使用类型断言
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", y)
}

/*
<nil>
float32

float32
float32

float32
float32
float32
*/

/*
若注释掉  y := x.(float32)  输出如下：
<nil>
float32

<nil>
float32

panic: interface conversion: interface {} is nil, not float32

goroutine 1 [running]:
main.main()
        /root/go-practice/chapter11-oop/assert/demo2/main.go:17 +0x17a
exit status 2
*/
