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

	y, ok := x.(float64) // 使用类型断言
	if ok {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", y)
	fmt.Println()

	fmt.Println("okokook~")
}

/*
<nil>
float32

float32
float32

convert fail
float32
float32
float64

okokook~
*/
