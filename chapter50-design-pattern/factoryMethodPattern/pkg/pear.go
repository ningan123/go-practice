package pkg

import "fmt"

// ====基础模块层====

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是Pear")
}
