package template

import "fmt"

/*
	模版方法的核心是父结构体包含接口的引用，同时子类匿名组合父类实现接口继承
	设计思想:
		1. 定义一个接口Shape
		2. 实现父struct, 并接口继承Shape, 同时fu类中包含子类引用，用来调用子类的方法
		3. 实现子struct, 匿名组合父struct， 这样子类也实现接口继承
*/
//定义接口
type Shape interface {
	SetName(name string)
	BeforeAction()
	Exit()
}

//定义父类Person
type Person struct {
	name string
	Concrete Shape  //具体子类的引用， 因为子类继承了接口
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) BeforeAction() {
	//将具体的action延迟到子类中执行
	p.Concrete.BeforeAction()
}

func (p *Person) Exit() {
	p.BeforeAction()
	fmt.Println(p.name + "exit")
}

//定义具体子类，且实现具体的action
type Boy struct {
	Person //匿名组合实现继承
}
//重写BeforeAction
func (b *Boy) BeforeAction() {
	fmt.Println(b.name)
}

type Girl struct {
	Person //匿名组合实现继承
}

func (g *Girl) BeforeAction() {
	fmt.Println(g.name)
}