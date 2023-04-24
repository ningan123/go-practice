package main

import (
	pkg "go-practice/chapter50-design-pattern/simpleFactoryPattern/pkg"
)

// 业务逻辑层
// 业务逻辑层只会和工厂模块进行依赖，业务逻辑层不再关心Fruit类具体是怎么创建基础对象的

func main() {
	factory := new(pkg.Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
