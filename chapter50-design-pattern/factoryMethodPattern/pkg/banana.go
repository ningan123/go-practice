package pkg

import "fmt"

// ====基础模块层====
type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是Banana")
}
