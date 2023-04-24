package main

import "fmt"

type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println("test01(): ", p.Name)
}

func test02(p *Person) {
	fmt.Println("test02(): ", p.Name)
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03(): ", p.Name)
}

func (p *Person) test04() {
	p.Name = "marry"
	fmt.Println("test04(): ", p.Name)
}

func main() {
	p := Person{"tom"}

	test01(p)
	// test01(&p) // cannot use &p (type *Person) as type Person in argument to test01
	fmt.Println()

	// test02(p)  // cannot use p (type Person) as type *Person in argument to test02
	test02(&p)
	fmt.Println()

	p.test03()
	fmt.Println("main(): ", p.Name)
	(&p).test03() //形式上看是传入地址，但是本质仍然是值拷贝
	fmt.Println("main(): ", p.Name)
	fmt.Println()

	p.test04() // 等价于(&p).test04()
	fmt.Println("main(): ", p.Name)
	(&p).test04()
	fmt.Println("main(): ", p.Name)
	fmt.Println()
}

/*
test01():  tom

test02():  tom

test03():  jack
main():  tom
test03():  jack
main():  tom

test04():  marry
main():  marry
test04():  marry
main():  marry
*/
