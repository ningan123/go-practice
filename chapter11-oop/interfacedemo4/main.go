package main

import (
	"fmt"
	"reflect"
)

// 定义空接口
type T interface {
}

type Student struct {
}

func main() {
	// test1
	var stu Student
	var t T = stu
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t))
	fmt.Println()

	// test2
	var t2 interface{} = stu //也是一个空接口
	fmt.Println(t2)
	fmt.Println(reflect.TypeOf(t2))
	fmt.Println()

	// test3
	var num1 float64 = 8.9
	t2 = num1
	fmt.Println(t2)
	fmt.Println(reflect.TypeOf(t2))
	fmt.Printf("%T\n", t2)
	fmt.Println()

}

/*
{}
main.Student

{}
main.Student

8.9
float64
float64
*/
