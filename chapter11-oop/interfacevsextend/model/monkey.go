package model

import "fmt"

// Monkey结构体
type Monkey struct {
	Name string
}

func (monkey *Monkey) Climbing() {
	fmt.Println(monkey.Name, "生来会爬树...")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

func (littleMonkey LittleMonkey) Flying() {
	fmt.Println(littleMonkey.Name, "通过学习，会飞翔了...")
}
