package observer


/*
	allow an instance to publish events to other observers when a particular evert occur
	核心思想
		1. Event struct
		2. Observer interface
			OnNotify(Event) 处理事件
		3. 被观察者Notifier interface 实现以下三个方法
			实现Register ObServer
			取消DeRegister Observer
			通知 Notify(Event)
*/
//定义事件结构体
type Event struct {
	Info string
}
//设计观察者和被观察者接口
type Observer interface {
	Update()
}

type Notifier interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event Event)
}

//通过具体对象来实现接口
//投资人观察者
type ConcreteInvestor struct {
	Name string
}

func (invester *ConcreteInvestor) Update() {

}
//股票被观察者
type ConcreteShare struct {
	Price float64
	oblist []Observer //注册链表
}
