package main

import "fmt"

type BInterface interface {
	testB()
}

type CInterface interface {
	testC()
}

type AInterface interface {
	BInterface
	CInterface
	testA()
}

type Student struct {
}

func (stu Student) testA() {
	fmt.Println("stu testA()")
}

func (stu Student) testB() {
	fmt.Println("stu testB()")
}

func (stu Student) testC() {
	fmt.Println("stu testC()")
}

func main() {
	var stu Student
	var a AInterface = stu
	var b BInterface = stu
	var c CInterface = stu

	a.testA()
	a.testB()
	a.testC()

	// b.testA()
	b.testB()
	// b.testC()

	// c.testA()
	// c.testB()
	c.testC()

}
