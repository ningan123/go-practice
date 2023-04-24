package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", p.Name, p.Age)
	return str
}

func main() {
	p := Person{
		Name: "test",
		Age:  10,
	}

	fmt.Println(p)
	fmt.Println(&p)

}

/*
{test 10}
Name=[test] Age=[10]
*/
