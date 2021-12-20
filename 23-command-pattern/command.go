package command

import "fmt"

/*
	命令模式是一种数据驱动模式，将请求封装成一个对象，从而可以用不同的请求对客户进行参数化，实现调用者和接收者的解藕
	设计思想：
		*接收者(Receiver): 执行请求相关的操作Execute()
		*请求对象(Invoker):
		*命令接口(Command)
		*具体命令的结构体(ConcreteCommand)
	Invoker负责维护Command队列
	ConcreteCommand匿名组合Receiver
	
	协作过程：
		1.创建一个ConcreteCommand对象并指定它的Receiver对象
		2.某Invoker对象存储该ConcreteCommand对象
		3.该Invoker通过调用Command的Excute操作来提交一个请求。
		4.具体的ConcreteCommand对象对调用它的Receiver的一些操作以执行该请求
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
//创建具体command, 指定接收者
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


/*创建请求对象， 维护请求cmd*/
type Invoker struct {
	cmds []Command
}

func (in *Invoker) AddCommand(c Command) {
	if in == nil {
		return
	}
	in.cmds = append(in.list, c)
}

func (in *Invoker) ExecuteCommand() {
	if in  == nil || len(in.cmds) == 0 {
		return
	}
	for _, cmd := range in.cmds {
		cmd.Call()
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
