package main

import (
	"fmt"
	"go-practice/chapter11-oop/interfacedemo"
)

func main() {
	// 这里就体现出多态数组 既可以存放Phone类型，又可以存放Camera类型
	var usbArr [3]interfacedemo.Usb
	fmt.Println(usbArr)

	usbArr[0] = interfacedemo.Phone{}
	usbArr[1] = interfacedemo.Phone{}
	usbArr[2] = interfacedemo.Camera{}
	fmt.Println(usbArr)
	fmt.Println()

	for i := 0; i < len(usbArr); i++ {
		fmt.Printf("%T\n", usbArr[i])
	}
	fmt.Println()

	// Phone有一个特有的方法Call()；遍历usb数组，如果是phone，除了调用Usb接口声明的方法外，还需调用Phone特有方法Call() => 类型断言
	var computer interfacedemo.Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

}

/*
[<nil> <nil> <nil>]
[{} {} {}]

interfacedemo.Phone
interfacedemo.Phone
interfacedemo.Camera

手机开始工作
手机 打电话...
手机停止工作

手机开始工作
手机 打电话...
手机停止工作

相机开始工作
相机停止工作
*/
