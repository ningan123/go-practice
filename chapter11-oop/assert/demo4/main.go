package main

import "fmt"

type Student struct {
	Name string
}

// 编写一个函数，可以判断传入的参数是什么类型
func TypeJudge(items ...interface{}) { // 这种写法可以传入任意多个任意类型的参数

	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%d个参数是bool类型,值是：%v\n", index, x)
		case float32:
			fmt.Printf("第%d个参数是float32类型,值是：%v\n", index, x)
		case float64:
			fmt.Printf("第%d个参数是float64类型,值是：%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%d个参数是int类型,值是：%v\n", index, x)
		case string:
			fmt.Printf("第%d个参数是string类型,值是：%v\n", index, x)
		case Student:
			fmt.Printf("第%d个参数是Student类型,值是：%v\n", index, x)
		case *Student:
			fmt.Printf("第%d个参数是*Student类型,值是：%v\n", index, x)
		default:
			fmt.Printf("第%d个参数类型不确定,值是：%v\n", index, x)
		}
	}

}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 3.2
	var n3 int = 10
	var n4 string = "test"
	n5 := 300
	n6 := Student{"test"}
	n7 := &Student{"test2"}
	TypeJudge(n1, n2, n3, n4, n5, n6, n7)
}

/*
第0个参数是float32类型,值是：1.1
第1个参数是float64类型,值是：3.2
第2个参数是int类型,值是：10
第3个参数是string类型,值是：test
第4个参数是int类型,值是：300
第5个参数是Student类型,值是：{test}
第6个参数是*Student类型,值是：&{test2}
*/
