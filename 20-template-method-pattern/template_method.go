package template

import "fmt"

// 定义接口
type Shape interface {
	SetName(name string)
	BeforeAction()
	Exit()
}

// 定义父类Person
type Person struct {
	name     string
	Concrete Shape //具体子类的引用， 因为子类继承了接口
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

// 定义具体子类，且实现具体的action
type Boy struct {
	Person //匿名组合实现继承
}

// 重写BeforeAction
func (b *Boy) BeforeAction() {
	fmt.Println(b.name)
}

type Girl struct {
	Person //匿名组合实现继承
}

func (g *Girl) BeforeAction() {
	fmt.Println(g.name)
}
