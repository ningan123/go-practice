package pkg

import "fmt"

type AmericaPear struct {
}

func (ap *AmericaPear) ShowPear() {
	fmt.Println("我是一个美国梨")
}
