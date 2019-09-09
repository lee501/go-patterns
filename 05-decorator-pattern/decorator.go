package decorator

import "fmt"

/*
	装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能
	*设计思想
		1.声明基础接口Component
		2.定义基础结构体，并实现接口继承
		3.装饰结构体中, 采用匿名组合的方式将接口Component作为结构体的属性，
*/
type Component interface {
	Describe() string
	GetCount() int
}

//基础结构体
type Fruit struct {
	Count 		int
	Description string
}

func (f *Fruit) Describe() string {
	return f.Description
}

func (f *Fruit) GetCount() int {
	return f.Count
}

//装饰结构体
type AppleDecorator struct {
	Component
	Type string
	Num int
}

func (apple *AppleDecorator) Describe() string {
	return fmt.Sprintf("%s, %s", apple.Component.Describe(), apple.Type)
}

func (apple *AppleDecorator) GetCount() int {
	return apple.Component.GetCount() + apple.Num
}

func CreateAppleDecorator(c Component, t string, n int) Component {
	return &AppleDecorator{c, t, n}
}
