# Mediator Pattern（中介者模式）

## 概述

中介者模式是一种**行为型**设计模式，通过引入一个**中介者对象**来封装多个对象之间的交互，使各对象不需要直接相互依赖，从而降低耦合度。

> Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.

---

## 适用场景

- 多个对象之间**交互复杂**，相互依赖关系形成网状结构
- 希望将对象间的通信逻辑**集中管理**，便于修改和扩展
- 系统中某个对象与其他对象的引用过多，难以复用
- 典型场景：聊天室（用户通过聊天室转发消息）、飞机塔台（飞机通过控制塔通信）、部门间协作（通过中介者转发消息）

---

## 结构

```
┌─────────────────────────────────────┐
│              Mediator               │  ← 中介者（集中管理通信）
│   Market  market                    │
│   Technical technical               │
│   ForwardMessage(dept, msg)         │
└──────────────┬──────────────────────┘
               │ 转发
    ┌──────────┴──────────┐
    ▼                     ▼
┌──────────┐       ┌──────────────┐
│  Market  │       │  Technical   │  ← 各部门只依赖中介者，不直接通信
│ mediator │       │  mediator    │
│SendMess()│       │ SendMess()   │
│GetMess() │       │ GetMess()    │
└──────────┘       └──────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Mediator** | 中介者结构体，持有各部门引用，实现 `ForwardMessage()` 消息转发 |
| **IDepartment（接口）** | 各部门的公共接口，声明 `SendMess()` 和 `GetMess()` |
| **Technical** | 技术部门，通过 `mediator` 发送消息给其他部门 |
| **Market** | 市场部门，通过 `mediator` 发送消息给其他部门 |

---

## Go 实现

### 部门接口

```go
type IDepartment interface {
    SendMess(message string)
    GetMess(message string)
}
```

### 中介者

```go
type Mediator struct {
    Market
    Technical
}

func (m *Mediator) ForwardMessage(department IDepartment, message string) {
    switch department.(type) {
    case *Technical:
        // Technical 发送消息 → 转发给 Market
        m.Market.GetMess(message)
    case *Market:
        // Market 发送消息 → 转发给 Technical
        m.Technical.GetMess(message)
    default:
        fmt.Println("部门不在中介者中")
    }
}
```

### 具体部门

```go
type Technical struct {
    mediator *Mediator
}

func (t *Technical) SendMess(message string) {
    t.mediator.ForwardMessage(t, message) // 通过中介者发送
}

func (t *Technical) GetMess(message string) {
    fmt.Printf("技术部收到消息: %s\n", message)
}

type Market struct {
    mediator *Mediator
}

func (m *Market) SendMess(message string) {
    m.mediator.ForwardMessage(m, message)
}

func (m *Market) GetMess(message string) {
    fmt.Printf("市场部收到消息: %s\n", message)
}
```

### 使用示例

```go
// 创建中介者
mediator := &Mediator{}

// 设置各部门的中介者引用
mediator.Technical.mediator = mediator
mediator.Market.mediator = mediator

// 技术部发消息给市场部（通过中介者）
mediator.Technical.SendMess("新功能上线了")
// 输出: 市场部收到消息: 新功能上线了

// 市场部发消息给技术部（通过中介者）
mediator.Market.SendMess("客户反馈了一个 bug")
// 输出: 技术部收到消息: 客户反馈了一个 bug
```

---

## 网状通信 vs 星形通信

**不使用中介者（网状，M×N 依赖）：**
```
Technical ←→ Market
Technical ←→ HR
Market    ←→ HR
... 每增加一个部门，依赖数量指数增长
```

**使用中介者（星形，M+N 依赖）：**
```
Technical → Mediator → Market
Market    → Mediator → Technical
HR        → Mediator → ...
所有通信集中在 Mediator 中管理
```

---

## 优缺点

### 优点

- **降低耦合**：各对象（部门）只依赖中介者，不直接依赖彼此
- **集中管理**：通信逻辑集中在 Mediator，便于维护和修改
- **易扩展**：新增部门只需修改 Mediator，无需修改其他部门

### 缺点

- 中介者可能变成**"上帝对象"**，包含过多业务逻辑
- 系统中对象越多，中介者的逻辑越复杂，可能难以维护

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Mediator** | 集中管理多对象间的复杂通信，减少直接依赖 |
| **Observer** | 一对多通知，被观察者不知道观察者的具体实现 |
| **Facade** | 为复杂子系统提供简化接口，子系统单向依赖外观 |
| **Pub/Sub** | 通过消息队列解耦，发布者和订阅者完全不知道对方 |

---

## 运行测试

```bash
cd 25-mediator-pattern
go test -v ./...
```
