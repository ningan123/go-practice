package pkg

import "fmt"

type JapanPear struct {
}

func (jp *JapanPear) ShowPear() {
	fmt.Println("我是一个日本梨")
}
