package main

import "fmt"

/*

 */
type Cat struct {
	Name  string // 大写 说明这个字段是可以公开的
	Age   int
	Color string
}

func main() {
	var cat1 Cat // var a int
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "white"
	fmt.Println("cat1: ", cat1)
	fmt.Printf("cat1的地址: %p\n", &cat1)
}
