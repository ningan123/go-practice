package main

import (
	"go-practice/chapter11-oop/interfacevsextend/model"
)

func main() {
	// littleMonkey := LittleMonkey{
	// 	Name: "悟空",
	// } //这样子写会报错  unknown field Name in struct literal

	littleMonkey := model.LittleMonkey{
		model.Monkey{
			Name: "悟空",
		},
	} // go-practice/chapter11-oop/interfacevsextend/model.LittleMonkey struct literal uses unkeyed fields
	littleMonkey.Climbing()
	littleMonkey.Flying()
}

/*
悟空 生来会爬树...
悟空 通过学习，会飞翔了...
*/
