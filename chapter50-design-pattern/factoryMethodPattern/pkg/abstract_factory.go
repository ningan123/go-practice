package pkg

// ====抽象层====

// 工厂类(抽象的接口)
type AbstractFactory interface {
	CreateFruit() Fruit // 生产水果类(抽象)的生产者方法
}
