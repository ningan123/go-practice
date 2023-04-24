package pkg

// ====具体的工厂模块====
// 具体的梨工厂

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	fruit = new(Pear)

	return fruit
}
