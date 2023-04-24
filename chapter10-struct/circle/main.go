package main

import "fmt"

type Circle struct {
	redius float64
}

func (c *Circle) getArea() float64 {
	return 3.14 * c.redius * c.redius
}

func main() {
	var c Circle
	c.redius = 1
	res := c.getArea()
	fmt.Println("面积为：", res)
}
