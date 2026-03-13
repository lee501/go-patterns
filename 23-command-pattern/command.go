package command

import "fmt"

/*创建Receiver, 这里使用接口，为了实现多个receiver。(可以创建Receiver struct)*/
type Receiver interface {
	Execute()
}

// 创建具体的接收者，实现接口方法
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

// 创建具体command, 指定接收者
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
	in.cmds = append(in.cmds, c)
}

func (in *Invoker) ExecuteCommand() {
	if in == nil || len(in.cmds) == 0 {
		return
	}
	for _, cmd := range in.cmds {
		cmd.Call()
	}
}

// 使用工厂方法模式来创建ConcreteCommand
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
