package main

import "fmt"

/*
不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个；结构体是值类型
*/

type Monster struct {
	Name string
	Age  int
}

func main() {
	var mon1 Monster
	mon1.Name = "test1"
	mon1.Age = 10

	mon2 := mon1 //结构体是值类型，默认是值拷贝
	fmt.Println("mon1: ", mon1)
	fmt.Println("mon2: ", mon2)

	mon2.Name = "test2"
	fmt.Println("mon1: ", mon1)
	fmt.Println("mon2: ", mon2)

	//可以让mon1的变化影响到mon2，让mon2等于mon1的地址就可以了，具体后面再讨论
}

/*
mon1:  {test1 10}
mon2:  {test1 10}
mon1:  {test1 10}
mon2:  {test2 10}
*/
