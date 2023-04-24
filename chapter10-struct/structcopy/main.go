package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"test1", 10}
	fmt.Println("p1", p1)
	fmt.Println()

	// p2改了  p1不会改    说明go语言结构体是值类型
	p2 := p1
	fmt.Println("p2", p2)
	p2.Name = "test2"
	p2.Age = 20
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println()

	// p3改了 p1也会改
	p3 := &p1
	fmt.Println("p3", p3)
	p3.Name = "test3"
	p3.Age = 30
	fmt.Println("p1", p1)
	fmt.Println("p3", p3)
}
