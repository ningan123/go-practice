package pkg

// 工厂
type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit { // 传一个参数，告诉我你想要哪种水果
	var fruit Fruit

	if kind == "apple" {
		// xxx 一堆操作
		fruit = new(Apple) // 满足多台条件的赋值，父类指针指向子类对象
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}
