package pkg

import "fmt"

/*
三个要点：
	一是某个类只能有一个实例
	二是它必须自行创建这个实例
	三个它必须自行向整个系统提供这个实例
*/

/*
保证一个类永远只能有一个对象，这个对象还能被系统的其他模块调用
*/

// 1. 保证这个类非公有化，外界不能通过这个类直接创建一个对象，
//    那么这个类就应该变得非共有访问，类名称是字母要小写
type singleton struct {
}

// 2. 还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
//    golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
var instance *singleton = new(singleton)

// 3. 如果全部都是私有化，那么外部模块将永远无法访问到这个对象
//    所以需要对外提供一个方法来获取到这个对象
// 讨论：GetInstance是否可以定义为singleton的一个成员方法呢？  不行

func GetInstance() *singleton {
	return instance
}

func (s *singleton) Something() {
	fmt.Println("单例的某方法")
}
