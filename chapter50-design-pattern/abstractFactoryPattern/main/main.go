package main

import (
	pkg "go-practice/chapter50-design-pattern/abstractFactoryPattern/pkg"
)

// 业务逻辑层

func main() {
	// // 需求1：需要中国的苹果、香蕉和梨
	// // 1-创建一个具体的中国工厂
	// var cFac pkg.AbstractFactory
	// cFac = new(pkg.ChinaFactory)

	// // 2-生成苹果
	// var cApple pkg.AbstractApple
	// cApple = cFac.CreateApple()
	// cApple.ShowApple()

	// // 2-生成香蕉
	// var cBanana pkg.AbstractBanana
	// cBanana = cFac.CreateBanana()
	// cBanana.ShowBanana()

	// // 2-生成梨
	// var cPear pkg.AbstractPear
	// cPear = cFac.CreatePear()
	// cPear.ShowPear()

	// 需求2：需要中国的苹果、美国的香蕉和日本的梨
	// 1-1 创建一个具体的中国工厂
	var cFac pkg.AbstractFactory
	cFac = new(pkg.ChinaFactory)

	// 1-2 生成苹果
	var cApple pkg.AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()

	// 2-1 创建一个具体的美国工厂
	var aFac pkg.AbstractFactory
	aFac = new(pkg.AmericaFactory)

	// 2-2 生成香蕉
	var aBanana pkg.AbstractBanana
	aBanana = aFac.CreateBanana()
	aBanana.ShowBanana()

	// 3-1 创建一个具体的日本工厂
	var jFac pkg.AbstractFactory
	jFac = new(pkg.JapanFactory)

	// 3-2 生成梨
	var jPear pkg.AbstractPear
	jPear = jFac.CreatePear()
	jPear.ShowPear()

}
