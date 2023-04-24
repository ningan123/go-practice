package main

import (
	"fmt"
	"go-practice/chapter10-struct/factory/model"
)

func main() {
	// var stu = model.Student{
	// 	Name:  "test",
	// 	Score: 100,
	// }
	// fmt.Println(stu)

	var stu = model.NewStudent("tom", 17)
	fmt.Println(stu)
	fmt.Println(*stu)
	// fmt.Println(stu.Name, stu.score) // stu.score undefined (type *model.student has no field or method score)
	fmt.Println(stu.Name, stu.GetScore())
}
