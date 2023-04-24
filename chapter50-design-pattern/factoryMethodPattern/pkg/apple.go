package pkg

import "fmt"

// ====基础模块层====
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("我是Apple")
}
