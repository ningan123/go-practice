package pkg

import "fmt"

type AmericaApple struct {
}

func (aa *AmericaApple) ShowApple() {
	fmt.Println("我是一个美国苹果")
}
