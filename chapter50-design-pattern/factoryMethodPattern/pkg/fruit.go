package pkg

// ====抽象层====

// 水果类(抽象的接口)
type Fruit interface {
	Show() // 疑问：此处不继承可不可以
}
