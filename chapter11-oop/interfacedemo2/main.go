package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

// 实现AInterface接口
func (m Monster) Say() {
	fmt.Println("Monster Say()")
}

// 实现BInterface接口
func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}

func main() {
	var monster Monster
	var a AInterface = monster
	var b BInterface = monster

	a.Say()
	b.Hello()
}
