# Command Pattern（命令模式）

## 概述

命令模式是一种**行为型**设计模式，将**请求封装为对象**，使发送方（Invoker）和接收方（Receiver）完全解耦。命令对象持有接收者引用，请求的执行、排队、撤销和重做都可以通过操作命令对象来实现。

> Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

---

## 适用场景

- 需要将**操作请求与执行者解耦**，发送方不需要知道谁来执行
- 需要支持**撤销（Undo）/ 重做（Redo）**操作
- 需要将操作**排队批量执行**（任务队列、批处理）
- 需要记录操作日志，支持**事务回滚**
- 需要实现**宏命令**（将多个命令组合成一个）

---

## 结构

```
┌─────────────┐   AddCommand()  ┌──────────────────────────────┐
│   Client    │ ──────────────> │         Invoker              │
└─────────────┘                 │   cmds []Command             │
                                │   AddCommand(c Command)      │
                                │   ExecuteCommand()           │
                                └──────────────┬───────────────┘
                                               │ Call()
                                   ┌───────────┴───────────┐
                                   │      Command (接口)    │
                                   │      Call()            │
                                   └───────────┬────────────┘
                                               │ 实现
                               ┌───────────────┴───────────────┐
                               │                               │
                    ┌──────────────────┐           ┌──────────────────┐
                    │ ConcreteCommandA │           │ ConcreteCommandB │
                    │ Receiver         │           │ Receiver         │
                    │ Call() →         │           │ Call() →         │
                    │ Receiver.Execute │           │ Receiver.Execute │
                    └────────┬─────────┘           └────────┬─────────┘
                             │                              │
                    ┌────────┴─────────┐          ┌─────────┴────────┐
                    │   ReceiverA      │          │   ReceiverB      │
                    │   Execute()      │          │   Execute()      │
                    └──────────────────┘          └──────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Receiver（接口）** | 实际执行业务逻辑的对象，声明 `Execute()` 方法 |
| **ReceiverA / ReceiverB** | 具体接收者，实现各自的执行逻辑 |
| **Command（接口）** | 命令接口，声明 `Call()` 方法 |
| **ConcreteCommandA/B** | 具体命令，持有 Receiver 引用，`Call()` 中委托给 `Receiver.Execute()` |
| **Invoker（调用者）** | 维护命令列表 `[]Command`，触发命令执行，不关心命令的具体实现 |

---

## Go 实现

### 接收者（Receiver）

```go
type Receiver interface {
    Execute()
}

type ReceiverA struct{}
func (a *ReceiverA) Execute() { fmt.Println("接收者A处理请求") }

type ReceiverB struct{}
func (b *ReceiverB) Execute() { fmt.Println("接收者B处理请求") }
```

### 命令接口与具体命令

```go
type Command interface {
    Call()
}

type ConcreteCommandA struct {
    Receiver // 匿名组合，持有接收者引用
}
func (ca *ConcreteCommandA) Call() { ca.Receiver.Execute() }

type ConcreteCommandB struct {
    Receiver
}
func (cb *ConcreteCommandB) Call() { cb.Receiver.Execute() }
```

### 调用者（Invoker）

```go
type Invoker struct {
    cmds []Command
}

func (in *Invoker) AddCommand(c Command) {
    in.cmds = append(in.cmds, c)
}

func (in *Invoker) ExecuteCommand() {
    for _, cmd := range in.cmds {
        cmd.Call()
    }
}
```

### 工厂函数

```go
type TYPE string

const (
    Acommand TYPE = "a"
    Bcommand TYPE = "b"
)

func CreateCommand(kind TYPE, receiver Receiver) Command {
    switch kind {
    case Acommand:
        return &ConcreteCommandA{receiver}
    case Bcommand:
        return &ConcreteCommandB{receiver}
    default:
        return nil
    }
}
```

### 使用示例

```go
// 创建接收者
receiverA := &ReceiverA{}
receiverB := &ReceiverB{}

// 创建命令（绑定接收者）
cmdA := CreateCommand(Acommand, receiverA)
cmdB := CreateCommand(Bcommand, receiverB)

// 调用者统一管理和执行命令
invoker := new(Invoker)
invoker.AddCommand(cmdA)
invoker.AddCommand(cmdB)
invoker.ExecuteCommand()
// 输出:
// 接收者A处理请求
// 接收者B处理请求
```

---

## 扩展：支持撤销（Undo）

在命令接口中增加 `Undo()` 方法，配合历史栈实现多步撤销：

```go
type Command interface {
    Call()
    Undo()
}

type History struct {
    done []Command
}

func (h *History) Execute(cmd Command) {
    cmd.Call()
    h.done = append(h.done, cmd)
}

func (h *History) Undo() {
    if len(h.done) == 0 {
        return
    }
    last := h.done[len(h.done)-1]
    last.Undo()
    h.done = h.done[:len(h.done)-1]
}
```

---

## 优缺点

### 优点

- **解耦**：Invoker 与 Receiver 完全解耦，互不依赖
- **扩展性**：新增命令只需添加新的 ConcreteCommand，无需修改 Invoker
- **组合命令**：可将多个命令组合成宏命令（`MacroCommand`）
- **支持撤销/重做**：命令对象天然适合保存操作历史

### 缺点

- 每个操作都需要一个 ConcreteCommand 类，类数量增多
- 简单场景下，命令模式会引入不必要的复杂度

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Command** | 将请求封装为对象，支持排队、撤销、重做 |
| **Strategy** | 封装可互换算法，关注算法本身而非请求 |
| **Chain of Responsibility** | 请求沿链传递，直到被处理；Command 是明确绑定接收者 |
| **Memento** | 保存对象状态快照；Command 的撤销也可基于 Memento 实现 |

---

## 运行测试

```bash
cd 23-command-pattern
go test -v ./...
```
