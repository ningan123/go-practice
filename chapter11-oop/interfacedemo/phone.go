package interfacedemo

import "fmt"

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机 打电话...")
}
