package pkg

import "fmt"

// 实现层
type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是Banana")
}
