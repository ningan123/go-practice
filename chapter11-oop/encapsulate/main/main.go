package main

import (
	"fmt"
	"go-practice/chapter11-oop/encapsulate/model"
)

func main() {
	p := model.NewPerson("test")

	p.SetAge(11)
	fmt.Println(p)
}
