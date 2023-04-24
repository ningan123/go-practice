package pkg

import "fmt"

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是Pear")
}
