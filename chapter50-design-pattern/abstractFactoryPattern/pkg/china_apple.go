package pkg

import "fmt"

type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("我是一个中国苹果")
}
