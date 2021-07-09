package mediator

import "testing"

func TestMediator_ForwardMessage(t *testing.T) {
	//创建中介者
	mediator := &Mediator{}
	//创建部门
	technical := Technical{mediator}
	market := Market{mediator}
	//添加对象到中介者中
	mediator.Market = market
	mediator.Technical = technical
	//测试发送消息
	technical.SendMess("开发已经完成")
	market.SendMess("市场推广中")
}
