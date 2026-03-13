# Publish-Subscribe Pattern（发布订阅模式）

## 概述

发布订阅模式是一种**消息传递**设计模式，发布者将消息发布到特定主题（Topic），订阅者订阅感兴趣的主题，消息通过**消息队列/channel** 异步传递。与观察者模式的核心区别在于，发布者和订阅者**完全解耦**，互不感知。

> Publishers send messages to topics without knowing who will receive them. Subscribers receive messages from topics without knowing who sent them.

---

## 适用场景

- 需要在**完全解耦**的组件之间传递消息
- 消息生产者和消费者的处理速度不一致，需要**异步缓冲**
- 需要**按主题过滤**消息，不同订阅者接收不同类型的消息
- 典型场景：事件总线、日志收集系统、微服务消息通知、实时数据推送

---

## 结构

```
                   ┌──────────────────────────────┐
                   │          Publisher            │
                   │  subscribers map[chan]topicFn │
                   │                               │
                   │  Subscribe()   → 订阅所有    │
                   │  SubscribeTopic() → 订阅指定  │
                   │  Publish(v)    → 广播消息     │
                   │  Exit(sub)     → 退订         │
                   │  Close()       → 关闭发布者   │
                   └──────────────┬────────────────┘
                                  │ sendTopic()（goroutine）
               ┌──────────────────┼──────────────────┐
               ▼                  ▼                  ▼
         ┌──────────┐       ┌──────────┐       ┌──────────┐
         │ sub chan │       │ sub chan │       │ sub chan │
         │(订阅所有) │       │(过滤主题) │       │(过滤主题) │
         └──────────┘       └──────────┘       └──────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Publisher** | 发布者，维护订阅者 map，负责消息广播 |
| **subscriber** | 订阅者 channel（`chan interface{}`），消息通过它传递 |
| **topicFunc** | 主题过滤函数 `func(v interface{}) bool`，返回 false 则跳过该消息 |
| **Publish()** | 并发向所有订阅者发送消息（每个订阅者一个 goroutine） |
| **sendTopic()** | 向单个订阅者发送，使用 `select` 支持超时 |

---

## Go 实现

### 类型定义

```go
type (
    subscriber chan interface{}
    topicFunc  func(v interface{}) bool
)

type Publisher struct {
    m           sync.RWMutex
    buffer      int
    timeout     time.Duration
    subscribers map[subscriber]topicFunc
}
```

### 创建发布者

```go
func NewPublisher(buf int, t time.Duration) *Publisher {
    return &Publisher{
        buffer:      buf,
        timeout:     t,
        subscribers: make(map[subscriber]topicFunc),
    }
}
```

### 订阅

```go
// 订阅所有主题（topicFunc = nil）
func (p *Publisher) Subscribe() subscriber {
    return p.SubscribeTopic(nil)
}

// 订阅指定主题（传入过滤函数）
func (p *Publisher) SubscribeTopic(topic topicFunc) subscriber {
    ch := make(subscriber, p.buffer)
    p.m.Lock()
    defer p.m.Unlock()
    p.subscribers[ch] = topic
    return ch
}
```

### 发布消息

```go
func (p *Publisher) Publish(v interface{}) {
    p.m.Lock()
    defer p.m.Unlock()
    var wg sync.WaitGroup
    for sub, topic := range p.subscribers {
        wg.Add(1)
        go p.sendTopic(sub, topic, v, &wg) // 并发发送
    }
    wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
    defer wg.Done()
    if topic != nil && !topic(v) { // 主题过滤
        return
    }
    select {
    case sub <- v:                      // 发送消息
    case <-time.After(p.timeout):       // 超时丢弃
    }
}
```

### 使用示例

```go
pub := NewPublisher(10, 100*time.Millisecond)
defer pub.Close()

// 订阅所有消息
allSub := pub.Subscribe()

// 只订阅包含 "Go" 的消息
goSub := pub.SubscribeTopic(func(v interface{}) bool {
    if msg, ok := v.(string); ok {
        return strings.Contains(msg, "Go")
    }
    return false
})

// 发布消息
pub.Publish("Hello Go!")
pub.Publish("Hello Python!")
pub.Publish("Go is awesome!")

// 读取 allSub（接收所有消息）
// allSub 收到: "Hello Go!", "Hello Python!", "Go is awesome!"

// 读取 goSub（只接收含 "Go" 的消息）
// goSub 收到: "Hello Go!", "Go is awesome!"
```

---

## 发布订阅 vs 观察者模式

| 维度 | 观察者模式 | 发布订阅模式 |
|------|-----------|-------------|
| **耦合程度** | Subject 持有 Observer 列表，有一定耦合 | 发布者和订阅者完全解耦，通过 channel 传递 |
| **消息过滤** | 所有 Observer 都收到通知 | 可以通过 `topicFunc` 按主题过滤 |
| **通信方式** | 同步（直接调用 Receive） | 异步（通过 channel，支持缓冲） |
| **超时处理** | 无内置超时 | `sendTopic` 中 `select + time.After` 支持超时 |
| **适用场景** | 对象间紧密协作，实时响应 | 组件解耦，异步处理，消息队列 |

---

## 优缺点

### 优点

- **完全解耦**：发布者无需知道订阅者，订阅者无需知道发布者
- **异步传递**：基于 channel，生产和消费异步进行
- **灵活过滤**：`topicFunc` 支持按内容过滤消息
- **并发安全**：使用 `sync.RWMutex` 保护订阅者 map

### 缺点

- 消息投递不保证顺序（并发 goroutine 发送）
- 超时情况下消息可能被丢弃
- 调试困难：消息流转异步，链路追踪复杂

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Pub/Sub** | 通过 channel 完全解耦发布者和订阅者，支持主题过滤 |
| **Observer** | 被观察者直接持有观察者列表，同步通知 |
| **Mediator** | 集中管理对象间通信，双方都知道中介者 |
| **Event Bus** | Pub/Sub 的变体，通常基于事件类型路由 |

---

## 运行测试

```bash
cd 27-publish-and-subscribe
go test -v ./...
```
