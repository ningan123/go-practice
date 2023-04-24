package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
	// time.Sleep(1 * time.Second)
}

func main() {
	go sayHello()
	// time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
