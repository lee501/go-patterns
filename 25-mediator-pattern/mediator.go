package mediator

import "fmt"

// 创建Mediator
type Mediator struct {
	Market
	Technical
}

//	type ConcreteMediator struct {
//		Market
//		Technical
//	}
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

// 创建IDepartment接口
type IDepartment interface {
	SendMess(message string)
	GetMess(message string)
}

// 创建具体 Technical Department
type Technical struct {
	mediator *Mediator
}

func (t *Technical) SendMess(message string) {
	t.mediator.ForwardMessage(t, message)
}

func (t *Technical) GetMess(message string) {
	fmt.Printf("技术部收到消息: %s\n", message)
}

// 创建具体 Market Department
type Market struct {
	mediator *Mediator
}

func (m *Market) SendMess(message string) {
	m.mediator.ForwardMessage(m, message)
}

func (m *Market) GetMess(message string) {
	fmt.Printf("市场部部收到消息: %s\n", message)
}
