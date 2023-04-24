package pkg

// ====具体的工厂模块====
// 具体的苹果工厂

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Apple)

	return fruit
}
