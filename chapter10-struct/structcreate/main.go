package main

import "fmt"

/*
创建结构体变量和访问结构体字段
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式1
	var p1 Person
	p1.Name = "test1"
	p1.Age = 10
	fmt.Println(p1)

	// 方式2 用的比较多
	p2 := Person{"test2", 20}
	fmt.Println(p2)

	// 方式3   创建一个结构体指针
	// 下面两种方式都可以，go语言的设计者为了程序员方便，底层做了处理
	var p3 *Person = new(Person)
	(*p3).Name = "test3"
	(*p3).Age = 30
	fmt.Println(p3)
	p3.Name = "test3-new"
	p3.Age = 32
	fmt.Println(p3)

	// 方式4   创建一个结构体指针
	// 下面两种方式都可以，go语言的设计者为了程序员方便，底层做了处理
	var p4 *Person = &Person{}
	(*p4).Name = "test4"
	(*p4).Age = 40
	fmt.Println(p4)
	p4.Name = "test4-new"
	p4.Age = 43
	fmt.Println(p4)

	var p5 *Person = &Person{"test5", 50}
	fmt.Println(p5)

}

/*
{test1 10}
{test2 20}
&{test3 30}
&{test3-new 32}
&{test4 40}
&{test4-new 43}
&{test5 50}
*/
