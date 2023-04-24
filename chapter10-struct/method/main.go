package main

import "fmt"

type Person struct {
	Name string
}

// 给Person类型绑定一个方法
func (p Person) test() {
	p.Name = "5678"
	fmt.Println("方法内：", p)
}

func main() {
	var p Person
	p.Name = "1234"
	p.test()
	fmt.Println("main内: ", p)

}

/*  说明是  值传递
方法内： {5678}
main内:  {1234}
*/
