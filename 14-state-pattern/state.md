# State Pattern（状态模式）

## 概述

状态模式是一种**行为型**设计模式，允许对象在内部状态发生改变时，**自动切换其行为**，使对象看起来好像修改了它的类。状态模式将每个状态对应的行为封装到独立的状态类中，消除大量的条件判断语句。

> Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.

---

## 适用场景

- 对象的行为依赖于其内部状态，且需要在运行时根据状态改变行为
- 存在大量基于状态的 `if-else` 或 `switch-case` 代码，难以维护
- 不同状态下的行为差异较大，希望将每种状态封装成独立的类
- 状态转移规则固定，且每个状态"知道"下一个状态是什么

---

## 结构

```
┌─────────────────────────┐
│       Context           │
│  HealthValue  int       │
│  State ActionState      │──────> ActionState (接口)
│                         │        View()  Comment()  Create()
│  SetHealth(value int)   │               ▲
│  changestate()          │       ┌───────┼───────┐
└─────────────────────────┘       │       │       │
                              NormalState RestrictedState ClosedState
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Context（上下文）** | 拥有多种状态的对象，持有 `ActionState` 接口引用，将行为委托给当前状态 |
| **ActionState（接口）** | 封装特定状态行为，声明 `View()`/`Comment()`/`Create()` |
| **NormalState** | 健康值 > 10，正常状态，所有操作均可用 |
| **RestrictedState** | 0 < 健康值 < 10，受限状态，部分操作受限 |
| **ClosedState** | 健康值 < 0，封禁状态，所有操作均不可用 |

---

## Go 实现

### 状态接口

```go
type ActionState interface {
    View()
    Comment()
    Create()
}
```

### 上下文（Context）

```go
type Context struct {
    State       ActionState
    HealthValue int
}

func (a *Context) View()    { a.State.View() }
func (a *Context) Comment() { a.State.Comment() }
func (a *Context) Create()  { a.State.Create() }

func (a *Context) SetHealth(value int) {
    a.HealthValue = value
    a.changestate() // 状态转移逻辑集中在 Context 内部
}

func (a *Context) changestate() {
    if a.HealthValue < 0 {
        a.State = &ClosedState{}
    } else if a.HealthValue > 10 {
        a.State = &NormalState{}
    } else {
        a.State = &RestrictedState{}
    }
}

func NewContext(health int) *Context {
    con := &Context{HealthValue: health}
    con.changestate()
    return con
}
```

### 具体状态类

```go
type NormalState struct{}
func (n *NormalState) View()    { fmt.Println("view normal") }
func (n *NormalState) Comment() { fmt.Println("comment normal") }
func (n *NormalState) Create()  { fmt.Println("create normal") }

type RestrictedState struct{}
func (r *RestrictedState) View()    { fmt.Println("view Restricted") }
func (r *RestrictedState) Comment() { fmt.Println("comment Restricted") }
func (r *RestrictedState) Create()  { fmt.Println("create Restricted") }

type ClosedState struct{}
func (c *ClosedState) View()    { fmt.Println("view closed") }
func (c *ClosedState) Comment() { fmt.Println("comment closed") }
func (c *ClosedState) Create()  { fmt.Println("create closed") }
```

### 使用示例

```go
// 健康值 15，正常状态
ctx := NewContext(15)
ctx.View()    // view normal
ctx.Comment() // comment normal

// 降低健康值至受限状态
ctx.SetHealth(5)
ctx.View()    // view Restricted
ctx.Create()  // create Restricted

// 封禁
ctx.SetHealth(-1)
ctx.View()    // view closed
ctx.Create()  // create closed
```

---

## 状态转移图

```
健康值 > 10          0 < 健康值 < 10        健康值 < 0
┌─────────────┐      ┌──────────────────┐   ┌─────────────┐
│ NormalState │ ◄──► │ RestrictedState  │◄──►│ ClosedState │
└─────────────┘      └──────────────────┘   └─────────────┘
```

---

## 状态模式 vs 策略模式

| 维度 | 状态模式 | 策略模式 |
|------|----------|----------|
| **切换方式** | 由 Context 内部自动切换 | 由客户端主动设置 |
| **状态感知** | 各状态知道其他状态（可自行转移） | 各策略相互独立，不知道彼此 |
| **关注点** | 对象内部状态的变化驱动行为变化 | 封装可互换的算法 |
| **客户端** | 通常不感知具体状态 | 通常主动选择策略 |

---

## 优缺点

### 优点

- **消除条件语句**：将 `if-else`/`switch-case` 转化为多态，代码更清晰
- **单一职责**：每个状态类只负责该状态下的行为
- **开闭原则**：新增状态只需新增状态类，无需修改 Context

### 缺点

- 状态类数量随状态增多而增加
- 状态间的转移逻辑分散，可能使整体流程不直观

---

## 运行测试

```bash
cd 14-state-pattern
go test -v ./...
```
