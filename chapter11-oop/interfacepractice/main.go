package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明Hero结构体切片类型
type HeroSlice []Hero

// 3.实现Interface接口 有三个方法： Len Less Swap
func (hs HeroSlice) Len() int {
	return len(hs)
}

// 该方法决定使用什么标准排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// 实现方式1
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp

	// 实现方式2
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	// 1. 定义数组切片，对数组切片进行排序
	var intSlice = []int{0, -1, 10, 7, 90}
	fmt.Println(intSlice)

	sort.Ints(intSlice) // 切片是引用类型
	fmt.Println(intSlice)

	// 2. 定义结构体切片，对结构体切片进行排序
	var heroSlices HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄*%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlices = append(heroSlices, hero)
	}
	fmt.Println("===结构体切片 排序前===")
	for _, hero := range heroSlices {
		fmt.Println(hero)
	}

	sort.Sort(heroSlices)
	fmt.Println("===结构体切片 排序后===")
	for _, hero := range heroSlices {
		fmt.Println(hero)
	}

}

/*
[0 -1 10 7 90]
[-1 0 7 10 90]
===结构体切片 排序前===
{英雄*81 87}
{英雄*47 59}
{英雄*81 18}
{英雄*25 40}
{英雄*56 0}
{英雄*94 11}
{英雄*62 89}
{英雄*28 74}
{英雄*11 45}
{英雄*37 6}
===结构体切片 排序后===
{英雄*56 0}
{英雄*37 6}
{英雄*94 11}
{英雄*81 18}
{英雄*25 40}
{英雄*11 45}
{英雄*47 59}
{英雄*28 74}
{英雄*81 87}
{英雄*62 89}
*/
