# Chain of Responsibility Pattern（职责链模式）

## 概述

职责链模式是一种**行为型**设计模式，将多个处理对象连成一条**链**，请求沿链传递，直到某个对象能够处理它为止。每个处理对象只负责处理自己能处理的请求，无法处理则传递给链上的下一个对象。

> Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.

---

## 适用场景

- 有多个对象可以处理同一请求，且具体由哪个对象处理在运行时确定
- 请求的处理者不固定，需要**动态组合**处理链
- 需要向多个对象发出请求，且不明确指定接收者
- 典型场景：日志级别过滤（DEBUG→INFO→ERROR）、审批流程（员工→主管→总监）、HTTP 中间件链

---

## 结构

```
Client ──► ObjectA ──► ObjectB ──► ObjectC ──► nil（无法处理）
              │            │           │
           Level=1      Level=2     Level=3
           处理 Level=1  处理 Level=2  处理 Level=3
           的事件        的事件        的事件
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Interface** | 处理者接口，声明 `SetNext(next Interface)` 和 `HandleEvent(event Event)` |
| **ObjectA / ObjectB** | 具体处理者，匿名组合 `Interface`（持有下一个处理者的引用） |
| **Event** | 请求对象，包含 `Level` 和 `Name` |
| **SetNext()** | 设置链上的下一个处理者，构建处理链 |

---

## Go 实现

### 接口与请求

```go
type Interface interface {
    SetNext(next Interface)
    HandleEvent(event Event)
}

type Event struct {
    Level int
    Name  string
}
```

### 具体处理者

```go
type ObjectA struct {
    Interface    // 匿名组合，持有下一个处理者的引用
    Level int
    Name  string
}

func (ob *ObjectA) SetNext(next Interface) {
    ob.Interface = next
}

func (ob *ObjectA) HandleEvent(event Event) {
    if ob.Level == event.Level {
        // 自己能处理，处理并终止
        fmt.Printf("%s 处理这个事件 %s\n", ob.Name, event.Name)
    } else {
        // 自己不能处理，传递给下一个
        if ob.Interface != nil {
            ob.Interface.HandleEvent(event)
        } else {
            fmt.Println("无法处理")
        }
    }
}

// ObjectB 实现同上（略）
```

### 使用示例

```go
// 构建处理链：A(Level=1) → B(Level=2) → C(Level=3)
objA := &ObjectA{Level: 1, Name: "处理者A"}
objB := &ObjectB{Level: 2, Name: "处理者B"}
objC := &ObjectA{Level: 3, Name: "处理者C"} // 复用 ObjectA

objA.SetNext(objB)
objB.SetNext(objC)

// 发送 Level=2 的事件
objA.HandleEvent(Event{Level: 2, Name: "订单审批"})
// 输出: 处理者B 处理这个事件 订单审批

// 发送 Level=3 的事件
objA.HandleEvent(Event{Level: 3, Name: "大额审批"})
// 输出: 处理者C 处理这个事件 大额审批

// 发送无人能处理的事件
objA.HandleEvent(Event{Level: 99, Name: "未知事件"})
// 输出: 无法处理
```

---

## 职责链 vs 状态模式

| 维度 | 职责链模式 | 状态模式 |
|------|-----------|----------|
| **请求处理者** | 多个处理者可能处理同一请求 | Context 根据状态选择行为 |
| **处理结果** | 找到能处理的就停止（或全部传递） | 每个状态都必然响应 |
| **逻辑类型** | 类似 `switch-case`，按 Level 分发 | 类似 `if-elseif-else`，按状态分支 |
| **对象关系** | 链式结构，对象不知道彼此具体实现 | 状态彼此了解，可自行触发状态转移 |

---

## 优缺点

### 优点

- **解耦**：请求发送者不需要知道哪个对象处理请求
- **灵活**：可以动态添加、移除或重排链中的处理者
- **单一职责**：每个处理者只关注自己能处理的请求类型

### 缺点

- 请求不保证被处理（可能到达链尾仍无人处理）
- 调试困难：请求沿链传递，难以追踪处理过程
- 若链过长，性能会有影响

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Chain of Responsibility** | 请求沿链传递，直到被某个处理者处理 |
| **State** | Context 根据内部状态自动切换行为 |
| **Command** | 将请求封装为对象，支持撤销/排队 |
| **Decorator** | 所有装饰器都会处理请求（不会跳过） |

---

## 运行测试

```bash
cd 22-chain-of-responsibility-pattern
go test -v ./...
```
