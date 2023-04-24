package model

type person struct {
	Name   string
	age    int
	salary float64
}

//写一个共产模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age和sal，编写一对Setxxx的方法和Getxxx的方法
func (p *person) SetAge(age int) {
	if age > 0 {
		p.age = age
	}

}

func (p *person) GetAge() int {
	return p.age
}
