package pubsub

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPublisher_Subscribe(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)
	defer p.Close()

	sub := p.Subscribe()
	p.Publish("hello")

	msg := <-sub
	if msg != "hello" {
		t.Errorf("订阅者应收到 hello, 实际收到: %v", msg)
	}
}

func TestPublisher_SubscribeTopic(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)
	defer p.Close()

	// 只订阅包含 "go" 的消息
	sub := p.SubscribeTopic(func(v interface{}) bool {
		s, ok := v.(string)
		return ok && strings.Contains(s, "go")
	})

	p.Publish("hello world")
	p.Publish("go patterns")

	msg := <-sub
	if msg != "go patterns" {
		t.Errorf("订阅者应收到 go patterns, 实际收到: %v", msg)
	}
}

func TestPublisher_MultipleSubscribers(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)
	defer p.Close()

	sub1 := p.Subscribe()
	sub2 := p.Subscribe()

	p.Publish("broadcast")

	msg1 := <-sub1
	msg2 := <-sub2

	if msg1 != "broadcast" || msg2 != "broadcast" {
		t.Errorf("所有订阅者应收到消息, sub1=%v, sub2=%v", msg1, msg2)
	}
}

func TestPublisher_Exit(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)

	sub := p.Subscribe()
	p.Exit(sub)

	// 退出后 channel 应被关闭
	_, ok := <-sub
	if ok {
		t.Error("退出订阅后 channel 应已关闭")
	}
}

func TestPublisher_Close(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)

	sub := p.Subscribe()
	p.Close()

	// Close后 channel 应被关闭
	_, ok := <-sub
	if ok {
		t.Error("关闭发布者后 channel 应已关闭")
	}
}

func TestPublisher_Timeout(t *testing.T) {
	// 缓冲区为0，超时时间极短，模拟发送超时场景
	p := NewPublisher(0, 1*time.Millisecond)
	defer p.Close()

	_ = p.Subscribe()
	// 发布不应阻塞，超时后直接返回
	done := make(chan struct{})
	go func() {
		p.Publish("timeout test")
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("发布超时处理正常")
	case <-time.After(1 * time.Second):
		t.Error("发布超时后应立即返回，但发生了阻塞")
	}
}
