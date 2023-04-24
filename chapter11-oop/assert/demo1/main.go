package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{} // 空接口
	var p Point = Point{
		x: 1,
		y: 2,
	}
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Println()

	// 将自定义类型变量 赋给 接口变量  没有问题
	// a = p
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Println()

	var p2 Point
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p2)
	fmt.Println()

	// 将接口变量 赋给 自定义类型的变量 不能字节赋
	// p2 = a // cannot use a (variable of type interface{}) as Point value in assignment

	p2 = a.(Point)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p2)
	fmt.Println()
}

/*
<nil>
main.Point

main.Point
main.Point

main.Point
main.Point
main.Point

main.Point
main.Point
main.Point
*/

/*

如果将21行a = p注释掉，就会报错
<nil>
main.Point

<nil>
main.Point

<nil>
main.Point
main.Point

panic: interface conversion: interface {} is nil, not main.Point

goroutine 1 [running]:
main.main()
        /root/go-practice/chapter11-oop/assert/main.go:35 +0x2e5
exit status 2
*/
