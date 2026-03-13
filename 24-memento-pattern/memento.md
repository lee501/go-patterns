# Memento Pattern（备忘录模式）

## 概述

备忘录模式是一种**行为型**设计模式，在不破坏封装性的前提下，捕获并保存一个对象的内部状态，以便在需要时将对象恢复到之前的某个状态（撤销/回滚）。

> Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.

---

## 适用场景

- 需要实现**撤销（Undo）/重做（Redo）**功能
- 需要保存对象的**历史快照**，支持回滚
- 事务操作：执行失败时回滚到执行前的状态
- 游戏存档、文本编辑器撤销、数据库事务回滚

---

## 结构

```
┌─────────────────────┐  CreateMemento()  ┌──────────────────┐
│     Originator      │ ───────────────>  │     Memento      │
│   state string      │                   │  Originator      │
│   GetState()        │ <──────────────── │  GetState()      │
│   SetState()        │  RecoverOriginator│  SetState()      │
└─────────────────────┘                   └──────────────────┘
          ▲                                        ▲
          │ 恢复                                   │ 保存
          │                                        │
          └─────────── Caretaker ─────────────────┘
                      CreateMemento()
                      RecoverOriginator()
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Originator（发起人）** | 拥有内部状态的对象，提供 `GetState()`/`SetState()` |
| **Memento（备忘录）** | 存储发起人的状态快照，通过匿名组合 `Originator` 实现 |
| **Caretaker（管理者）** | 负责保存和恢复备忘录，不直接操作备忘录内容 |

---

## Go 实现

### 发起人（Originator）

```go
type Originator struct {
    state string
}

func (o *Originator) GetState() string        { return o.state }
func (o *Originator) SetState(state string)   { o.state = state }
```

### 备忘录（Memento）

```go
// 通过匿名组合 Originator，保存其完整状态
type Memento struct {
    Originator
}

func (m *Memento) GetState() string             { return m.Originator.state }
func (m *Memento) SetState(originator Originator) { m.Originator = originator }
```

### 管理者（Caretaker）

```go
type Caretaker struct{}

// 保存：从发起人获取状态，存入备忘录
func (c *Caretaker) CreateMemento(originator Originator) Memento {
    return Memento{originator}
}

// 恢复：从备忘录取出状态，还原给发起人
func (c *Caretaker) RecoverOriginator(memento Memento) Originator {
    return memento.Originator
}
```

### 使用示例

```go
caretaker := &Caretaker{}
originator := Originator{}

// 设置初始状态
originator.SetState("state-A")
fmt.Println("当前状态:", originator.GetState()) // state-A

// 保存快照
snapshot := caretaker.CreateMemento(originator)

// 修改状态
originator.SetState("state-B")
fmt.Println("修改后:", originator.GetState()) // state-B

// 回滚到保存的状态
originator = caretaker.RecoverOriginator(snapshot)
fmt.Println("回滚后:", originator.GetState()) // state-A
```

---

## 多状态历史（栈式备忘录）

实际开发中，支持多步撤销通常使用**栈**保存历史快照：

```go
type History struct {
    snapshots []Memento
}

func (h *History) Save(m Memento) {
    h.snapshots = append(h.snapshots, m)
}

func (h *History) Undo() (Memento, bool) {
    if len(h.snapshots) == 0 {
        return Memento{}, false
    }
    last := h.snapshots[len(h.snapshots)-1]
    h.snapshots = h.snapshots[:len(h.snapshots)-1]
    return last, true
}
```

---

## 优缺点

### 优点

- **封装性**：备忘录对象对外隐藏发起人的内部状态，不破坏封装
- **易实现撤销**：通过保存多个快照，轻松实现多步撤销/重做
- **职责分离**：状态保存逻辑（Caretaker）与业务逻辑（Originator）分离

### 缺点

- 频繁创建备忘录会消耗大量内存（尤其是状态体积大时）
- Caretaker 不应操作备忘录内容，但在 Go 中由于结构体可访问，需遵守约定

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Memento** | 保存对象状态快照，支持回滚/撤销 |
| **Prototype** | 克隆对象创建新实例，关注对象创建而非状态历史 |
| **Command** | 封装操作，支持撤销（通过记录逆操作而非状态快照） |
| **State** | 管理对象当前状态的切换，不关注历史状态 |

---

## 运行测试

```bash
cd 24-memento-pattern
go test -v ./...
```
