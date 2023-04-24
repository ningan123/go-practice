package pkg

// ====具体的工厂模块====
// 具体的香蕉工厂

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Banana)

	return fruit
}
