package pkg

import "fmt"

type AmericaBanana struct {
}

func (ab *AmericaBanana) ShowBanana() {
	fmt.Println("我是一个美国香蕉")
}
