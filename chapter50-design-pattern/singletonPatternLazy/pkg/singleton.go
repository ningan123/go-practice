package pkg

import (
	"fmt"
	"sync"
)

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
var instance *singleton

// 3. 如果全部都是私有化，那么外部模块将永远无法访问到这个对象
//    所以需要对外提供一个方法来获取到这个对象
// 讨论：GetInstance是否可以定义为singleton的一个成员方法呢？  不行

// 定义一个锁
var lock sync.Mutex

// 标记
// var initialized uint32

var once sync.Once

// 3. once
func GetInstance() *singleton {

	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

/*
// 2. 原子 + 锁
func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 如果没有，则加锁申请
	lock.Lock()
	defer lock.Unlock()

	// 只有首次GetInstance()方法被调用，才会生成这个单例的对象
	if instance == nil {
		instance = new(singleton)
		// 设置标记为1
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}
*/

/*  1. 锁   每次调用这个函数都要加锁，太慢了
func GetInstance() *singleton {
	// 为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	// 只有首次GetInstance()方法被调用，才会生成这个单例的对象
	if instance == nil {
		instance = new(singleton)
		return instance
	}
	return instance
}
*/

func (s *singleton) Something() {
	fmt.Println("单例的某方法")
}
