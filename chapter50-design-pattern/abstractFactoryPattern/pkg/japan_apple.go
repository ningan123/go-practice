package pkg

import "fmt"

type JapanApple struct {
}

func (ja *JapanApple) ShowApple() {
	fmt.Println("我是一个日本苹果")
}
