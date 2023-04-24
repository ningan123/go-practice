package main

import "fmt"

/*
如果结构体的字段类型是：指针、slice、和map的零值都是nil，即还没有分配空间
如果需要使用这样的字段，需要先make，才能使用
*/

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)

	// 确切判断一下 ptr slice map1是否是nil类型
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	// 显示各个字段的地址
	fmt.Println()
	fmt.Printf("p1的地址: %p\n", &p1)
	fmt.Printf("p1.Name的地址: %p\n", &p1.Name)
	fmt.Printf("p1.Age的地址: %p\n", &p1.Age)
	fmt.Printf("p1.Scores的地址: %p\n", &p1.Scores)
	fmt.Printf("p1.ptr的地址: %p\n", &p1.ptr)
	fmt.Printf("p1.slice的地址: %p\n", &p1.slice)
	fmt.Printf("p1.map1的地址: %p\n", &p1.map1)

	// 使用slice，一定要先make
	// p1.slice[0] = 100 // panic: runtime error: index out of range [0] with length 0
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println(p1)

	// 使用map，一定要先make   map可以自动增长
	// p1.map1["k1"] = "v1" // panic: assignment to entry in nil map
	p1.map1 = make(map[string]string)
	p1.map1["k1"] = "v1"
	fmt.Println(p1)

	// 显示各个字段的地址
	fmt.Println()
	fmt.Printf("p1的地址: %p\n", &p1)
	fmt.Printf("p1.Name的地址: %p\n", &p1.Name)
	fmt.Printf("p1.Age的地址: %p\n", &p1.Age)
	fmt.Printf("p1.Scores的地址: %p\n", &p1.Scores)
	fmt.Printf("p1.ptr的地址: %p\n", &p1.ptr)
	fmt.Printf("p1.slice的地址: %p\n", &p1.slice)
	fmt.Printf("p1.map1的地址: %p\n", &p1.map1)
}

/*
{ 0 [0 0 0 0 0] <nil> [] map[]}
ok1
ok2
ok3

p1的地址: 0xc000132000
p1.Name的地址: 0xc000132000
p1.Age的地址: 0xc000132010
p1.Scores的地址: 0xc000132018
p1.ptr的地址: 0xc000132040
p1.slice的地址: 0xc000132048
p1.map1的地址: 0xc000132060
{ 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[]}
{ 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[k1:v1]}

p1的地址: 0xc000132000
p1.Name的地址: 0xc000132000
p1.Age的地址: 0xc000132010
p1.Scores的地址: 0xc000132018
p1.ptr的地址: 0xc000132040
p1.slice的地址: 0xc000132048
p1.map1的地址: 0xc000132060
*/
