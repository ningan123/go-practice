package pkg

import "fmt"

// 实现层
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("我是Apple")
}
