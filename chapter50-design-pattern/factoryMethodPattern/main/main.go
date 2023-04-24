package main

import (
	pkg "go-practice/chapter50-design-pattern/factoryMethodPattern/pkg"
)

// 业务逻辑层

func main() {
	// 需求1：需要一个具体的苹果对象
	// 1-先要一个具体的苹果工厂
	var appleFac pkg.AbstractFactory // 分开写是为了方便理解
	appleFac = new(pkg.AppleFactory)

	// 2-生产一个具体的苹果
	var apple pkg.Fruit
	apple = appleFac.CreateFruit()

	apple.Show() // 多态

	// 需求2：需要一个具体的香蕉对象
	// 1-先要一个具体的香蕉工厂
	var bananaFac pkg.AbstractFactory
	bananaFac = new(pkg.BananaFactory)

	// 2-生产一个具体的苹果
	var banana pkg.Fruit
	banana = bananaFac.CreateFruit()

	banana.Show() // 多态

}
