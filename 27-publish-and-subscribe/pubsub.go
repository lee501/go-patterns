package pubsub

import (
	"sync"
	"time"
)

/*
	将消息发送给主题
	设计思想：
		publisher结构体中包含属性为subscribers的map结构，该map key为channel， 通过Publish方法，将消息发送到channel中
*/
type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m sync.RWMutex //读写锁
	buffer int  //订阅队列的大小
	timeout time.Duration  //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

//构建发布者对象，设置订阅队列的大小和超时时间
func NewPublisher(buf int, t time.Duration) *Publisher {
	return &Publisher{
		buffer: buf,
		timeout: t,
		subscribers: make(map[subscriber]topicFunc),
	}
}
//添加一个订阅者，订阅所有的主题
func (p *Publisher) Subscribe() subscriber {
	return p.SubscribeTopic(nil)
}

//添加一个订阅者，订阅指定的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
	ch := make(subscriber, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[ch] = topic
	return ch
}

//退出订阅, 同时关闭chan
func (p *Publisher) Exit(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

//关闭发布者，同时关闭所有订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

//发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//发送主题， 使用select语句处理channel， 允许超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub<-v:
	case <-time.After(p.timeout):
	}
}
