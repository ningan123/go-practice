package pkg

import "fmt"

type ChinaPear struct {
}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("我是一个中国梨")
}
