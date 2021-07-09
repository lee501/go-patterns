package command

import "fmt"

/*
	命令模式是一种数据驱动模式，将请求封装成一个对象，从而可以用不同的请求对客户进行参数化，实现调用者和接收者的解藕
	设计思想：
		*接收者(Receiver): 执行请求相关的操作Execute()
		*调用者(Invoker):
		*命令接口(Command)
		*具体命令的结构体(ConcreteCommand)
	Invoker负责维护Command队列
	ConcreteCommand匿名组合Receiver
*/
/*创建Receiver, 这里使用接口，为了实现多个receiver。(可以创建Receiver struct)*/
type Receiver interface {
	Execute()
}
//创建具体的接收者，实现接口方法
type ReceiverA struct {

}
func (a *ReceiverA) Execute() {
	fmt.Println("接收者A处理请求")
}

type ReceiverB struct {

}
func (b *ReceiverB) Execute() {
	fmt.Println("接收者B处理请求")
}


/*创建Command接口*/
type Command interface {
	Call()
}
//创建具体command struct
type ConcreteCommandA struct {
	Receiver
}
func (ca *ConcreteCommandA) Call() {
	ca.Receiver.Execute()
}

type ConcreteCommandB struct {
	Receiver
}
func (cb *ConcreteCommandB) Call() {
	cb.Receiver.Execute()
}


/*创建调用者， 实现添加，执行命令*/
type Invoker struct {
	list []Command
}

func (in *Invoker) AddCommand(c Command) {
	if in == nil {
		return
	}
	in.list = append(in.list, c)
}

func (in *Invoker) ExecuteCommand() {
	if in  == nil || len(in.list) == 0 {
		return
	}
	for _, item := range in.list {
		item.Call()
	}
}

//使用工厂方法模式来创建ConcreteCommand
type TYPE string

const (
	Acommand TYPE = "a"
	Bcommand TYPE = "b"
)
func CreateCommand(kind TYPE, receiver Receiver) Command {
	switch kind {
	case Acommand:
		return &ConcreteCommandA{receiver}
	case Bcommand:
		return &ConcreteCommandB{receiver}
	default:
		return nil
	}
}
