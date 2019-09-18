package mediator

import "fmt"

/*
	设计思想：
		1. Mediator: 中介者结构体， 实现通信的方法，属性每个具体对象的引用(或对象接口的切片)
			* 也可以为接口，当为接口的时候需要实现具体的对象ConcreteMediator
		2. IDepartment: 对象接口， 邓庄对象struct的的方法
		3. Department: 对象结构体， 继承IDepartment接口， 并包含mediator指针的引用（因为接口实际是一个指针）
	*Department发送消息通过中介者来通知其他Department
*/
//创建Mediator
type Mediator struct {
	Market
	Technical
}

//type ConcreteMediator struct {
//	Market
//	Technical
//}
func (m *Mediator) ForwardMessage(department IDepartment, message string) {
	switch department.(type) {
	case *Technical:
		m.Market.GetMess(message)
	case *Market:
		m.Technical.GetMess(message)
	default:
		fmt.Println("部门不在中介者中")
	}
}

//创建IDepartment接口
type IDepartment interface {
	SendMess(message string)
	GetMess(message string)
}

//创建具体 Technical Department
type Technical struct {
	mediator *Mediator
}

func (t *Technical) SendMess(message string) {
	t.mediator.ForwardMessage(t, message)
}

func (t *Technical) GetMess(message string) {
	fmt.Printf("技术部收到消息: %s\n", message)
}

//创建具体 Market Department
type Market struct {
	mediator *Mediator
}

func (m *Market) SendMess(message string) {
	m.mediator.ForwardMessage(m, message)
}

func (m *Market) GetMess(message string) {
	fmt.Printf("市场部部收到消息: %s\n", message)
}