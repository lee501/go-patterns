# Observer Pattern（观察者模式）

## 概述

观察者模式是一种**行为型**设计模式，定义对象之间的**一对多依赖关系**，当被观察对象（Subject）的状态发生变化时，所有依赖它的观察者都会自动收到通知并更新。

> Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

---

## 适用场景

- 一个对象的状态变化需要**通知多个其他对象**
- 对象之间需要**松耦合**，不希望观察者和被观察者直接依赖
- 消息广播、事件系统、股票行情推送等场景
- 需要动态添加或移除观察者

---

## 结构

```
┌────────────────────────┐
│   Notifier (接口)       │
├────────────────────────┤
│ Register(Observer)     │
│ Remove(Observer)       │
│ Notify(Event)          │
└────────────┬───────────┘
             │ 实现
     ┌───────┴────────┐
     │ ShareNotifier  │
     │ oblist[]       │──────┐ Notify
     └────────────────┘      ▼
                     ┌──────────────┐
                     │  Observer    │  ← 接口
                     │  Receive()   │
                     └──────┬───────┘
                            │ 实现
                   ┌────────┴────────┐
                   │InvestorObserver │
                   │ Name string     │
                   └─────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Event** | 事件数据结构，携带通知内容 |
| **Observer（接口）** | 观察者接口，定义 `Receive(Event)` 方法 |
| **Notifier（接口）** | 被观察者接口，定义 `Register`/`Remove`/`Notify` 方法 |
| **InvestorObserver** | 具体观察者，收到通知后执行响应逻辑 |
| **ShareNotifier** | 具体被观察者，维护观察者列表并广播事件 |

---

## Go 实现

### 事件与接口定义

```go
type Event struct {
    Info string
}

type Observer interface {
    Receive(event Event)
}

type Notifier interface {
    Register(observer Observer)
    Remove(observer Observer)
    Notify(event Event)
}
```

### 具体观察者

```go
type InvestorObserver struct {
    Name string
}

func (investor *InvestorObserver) Receive(event Event) {
    fmt.Printf("%s 收到事件通知 %s\n", investor.Name, event.Info)
}
```

### 具体被观察者

```go
type ShareNotifier struct {
    Price  float64
    oblist []Observer
}

func (share *ShareNotifier) Register(observer Observer) {
    share.oblist = append(share.oblist, observer)
}

func (share *ShareNotifier) Remove(observer Observer) {
    for i, ob := range share.oblist {
        if ob == observer {
            share.oblist = append(share.oblist[:i], share.oblist[i+1:]...)
        }
    }
}

func (share *ShareNotifier) Notify(event Event) {
    for _, ob := range share.oblist {
        ob.Receive(event)
    }
}
```

### 使用示例

```go
// 创建被观察者（股票）
share := NewShareNotifier(100.0)

// 创建观察者（投资人）
investor1 := NewInvestorObserver("张三")
investor2 := NewInvestorObserver("李四")

// 注册观察者
share.Register(investor1)
share.Register(investor2)

// 触发事件，所有观察者收到通知
share.Notify(NewEvent())
// 输出:
// 张三 收到事件通知 价格变动通知
// 李四 收到事件通知 价格变动通知

// 移除观察者
share.Remove(investor1)
share.Notify(NewEvent())
// 输出:
// 李四 收到事件通知 价格变动通知
```

---

## 优缺点

### 优点

- **松耦合**：Subject 和 Observer 只通过接口依赖，互不了解具体实现
- **动态扩展**：可在运行时动态添加/移除观察者
- **开闭原则**：新增观察者无需修改 Subject 代码

### 缺点

- 观察者较多时，`Notify` 的性能开销线性增长
- 若观察者之间存在依赖，通知顺序可能引发意外问题
- 可能产生循环通知问题（A 通知 B，B 又通知 A）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Observer** | 一对多通知，被观察者主动推送事件 |
| **Pub/Sub** | 通过中间消息队列解耦，发布者和订阅者完全不知道对方存在 |
| **Mediator** | 通过中介者集中管理多个对象之间的通信 |
| **Event Bus** | Observer 的变体，通过事件总线广播消息 |

---

## 运行测试

```bash
cd 07-observer-pattern
go test -v ./...
```
