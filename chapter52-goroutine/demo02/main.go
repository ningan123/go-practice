package main

import (
	"fmt"
	"sync"
)

// // 方法1
// func main() {
// 	for i := 0; i < 100; i++ {
// 		go fmt.Println(i)
// 	}
// 	time.Sleep(time.Second)
// }

// // 方法2
// func main() {
// 	num := 100
// 	c := make(chan bool, num)
// 	for i := 0; i < num; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 			c <- true
// 		}(i)
// 	}

// 	for i := 0; i < num; i++ {
// 		<-c
// 	}
// }

// 方法3
func main() {
	wg := sync.WaitGroup{}
	num := 10
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			// wg.Done() // 位置1  放这会有问题  如果要放在这，需要写成defer wg.Done()
			fmt.Println(i)
			wg.Done() // 位置2
		}(i)
	}
	wg.Wait()
}
