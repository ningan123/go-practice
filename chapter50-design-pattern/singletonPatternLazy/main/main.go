package main

import (
	"fmt"
	pkg "go-practice/chapter50-design-pattern/singletonPattern/pkg"
)

func main() {
	s1 := pkg.GetInstance()
	s1.Something()

	s2 := pkg.GetInstance()

	if s1 == s2 {
		fmt.Println("s1==s2")
		fmt.Println(s1)
		fmt.Println(s2)
	}
}
