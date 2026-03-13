package chain

import "fmt"

// 定义Interface
type Interface interface {
	SetNext(next Interface) //参数不确定，所以这里使用接口
	HandleEvent(event Event)
}

// 定义ObjectA struct
type ObjectA struct {
	Interface
	Level int
	Name  string
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

// 定义ObjectB struct
type ObjectB struct {
	Interface
	Level int
	Name  string
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
	Name  string
}
