package main

import "fmt"

type MethodUtils struct {
}

func (methodUtils *MethodUtils) print() {
	a := 10
	b := 8

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (methodUtils *MethodUtils) print2(a, b int) {

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func (methodUtils *MethodUtils) getArea(a, b float64) float64 {

	return a * b
}

func main() {
	var methodUtils *MethodUtils
	methodUtils.print()
	methodUtils.print2(3, 4)

	fmt.Println(methodUtils.getArea(3, 4))
}
