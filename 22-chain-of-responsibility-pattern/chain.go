package chain

import "fmt"

/*
	核心是结构体包含下一个结构体的引用
	状态模式和职责链模式区别：
		状态模式下知道自己要处理的状态对象是谁，是属于if、 else if、else操作
		职责链模式是客户端设置请求的类型, 请求直到被具体的某个职责类处理。属于switch-case操作
	设计思想：
		1. 一个Interface接口，用来封装方法集合
		2. 具体struct, 匿名组合接口(对象链中next对象引用)
*/
//定义Interface
type Interface interface {
	SetNext(next Interface) //参数不确定，所以这里使用接口
	HandleEvent(event Event)
}
//定义ObjectA struct
type ObjectA struct {
	Interface
	Level int
	Name string
}

func (ob *ObjectA) SetNext(next Interface) {
	ob.Interface = next
}

func (ob *ObjectA) HandleEvent(event Event) {
	if ob.Level == event.Level {
		fmt.Printf("%s 处理这个事件 %s\n", ob.Name, event.Name)
	} else {
		if ob.Interface != nil {
			ob.Interface.HandleEvent(event)
		} else {
			fmt.Println("无法处理")
		}
	}
}
//定义ObjectB struct
type ObjectB struct {
	Interface
	Level int
	Name string
}

func (ob *ObjectB) SetNext(next Interface) {
	ob.Interface = next
}

func (ob *ObjectB) HandleEvent(event Event) {
	if ob.Level == event.Level {
		fmt.Printf("%s 处理这个事件 %s\n", ob.Name, event.Name)
	} else {
		if ob.Interface != nil {
			ob.Interface.HandleEvent(event)
		} else {
			fmt.Println("无法处理")
		}
	}
}

type Event struct {
	Level int
	Name string
}
