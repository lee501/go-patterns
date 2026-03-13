# Template Method Pattern（模板方法模式）

## 概述

模板方法模式是一种**行为型**设计模式，在父类中定义算法的**骨架（模板）**，将某些步骤的具体实现延迟到子类中，使得子类可以在不改变算法结构的情况下，重新定义算法的某些步骤。

> Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

---

## 适用场景

- 多个类有**相同的算法流程**，但某些步骤的具体实现不同
- 希望控制子类扩展的方式，只允许重写特定步骤，而非整个流程
- 提取公共逻辑到父类，减少代码重复
- 框架中定义处理流程，由具体实现提供细节（Hook 机制）

---

## 结构

```
┌──────────────────────────────────────┐
│              Person（父类）           │
│   name string                        │
│   Concrete Shape（子类接口引用）       │
│                                      │
│   Exit() {                           │  ← 模板方法（定义骨架）
│       p.BeforeAction()  ← 延迟到子类  │
│       fmt.Println(name + "exit")     │
│   }                                  │
└──────────────────────────────────────┘
               ▲ 匿名组合
      ┌────────┴────────┐
      │                 │
 ┌─────────┐       ┌─────────┐
 │   Boy   │       │  Girl   │
 │BeforeAction│    │BeforeAction│
 └─────────┘       └─────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Shape（接口）** | 定义可被重写的步骤：`BeforeAction()`、`SetName()`、`Exit()` |
| **Person（父类/模板）** | 持有子类接口引用 `Concrete Shape`，在 `Exit()` 中定义算法骨架 |
| **Boy / Girl（具体子类）** | 匿名组合 `Person`，重写 `BeforeAction()` 提供具体实现 |

---

## Go 实现

### 接口与父类

```go
type Shape interface {
    SetName(name string)
    BeforeAction()
    Exit()
}

// 父类 Person，持有子类的接口引用
type Person struct {
    name     string
    Concrete Shape // 子类引用，用于调用子类实现的方法
}

func (p *Person) SetName(name string) { p.name = name }

func (p *Person) BeforeAction() {
    p.Concrete.BeforeAction() // 委托给子类执行
}

// Exit 是模板方法：定义算法骨架
func (p *Person) Exit() {
    p.BeforeAction()                    // 步骤 1：可变，由子类决定
    fmt.Println(p.name + " exit")      // 步骤 2：固定
}
```

### 具体子类

```go
type Boy struct {
    Person // 匿名组合，继承父类方法
}

// 重写 BeforeAction
func (b *Boy) BeforeAction() {
    fmt.Println(b.name + ": boy before action")
}

type Girl struct {
    Person
}

func (g *Girl) BeforeAction() {
    fmt.Println(g.name + ": girl before action")
}
```

### 使用示例

```go
// 创建 Boy，将自身作为 Concrete 传给 Person
boy := &Boy{}
boy.Concrete = boy  // 关键：将子类引用注入父类
boy.SetName("Tom")
boy.Exit()
// 输出:
// Tom: boy before action
// Tom exit

// 创建 Girl
girl := &Girl{}
girl.Concrete = girl
girl.SetName("Alice")
girl.Exit()
// 输出:
// Alice: girl before action
// Alice exit
```

---

## Go 中模板方法的实现要点

Go 没有传统 OOP 的继承，实现模板方法有两个关键点：

1. **匿名组合**：子类通过 `Person` 匿名组合继承父类的方法（如 `Exit()`）
2. **接口引用注入**：父类持有 `Concrete Shape` 接口，在调用可变步骤（`BeforeAction()`）时，通过接口引用调用子类的重写方法，实现运行时多态

```
p.BeforeAction()
  → p.Concrete.BeforeAction()  // Concrete 指向子类（Boy 或 Girl）
    → boy.BeforeAction()       // 调用子类的具体实现
```

---

## 优缺点

### 优点

- **复用**：公共算法流程在父类中定义一次，所有子类复用
- **扩展点明确**：子类只能重写特定步骤（Hook），不能修改整体流程
- **开闭原则**：新增子类无需修改父类的模板方法

### 缺点

- 每个差异化的实现都需要一个子类，可能导致类的数量增多
- 子类必须了解父类的模板方法流程，父子类耦合度较高
- Go 没有原生继承，需要通过匿名组合+接口注入模拟，代码略显复杂

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Template Method** | 在父类中固化算法骨架，子类重写特定步骤 |
| **Strategy** | 封装整个算法，运行时自由替换，无固定骨架 |
| **Factory Method** | 模板方法的一种特例（创建对象的步骤延迟到子类） |
| **Hook** | 模板方法中的可选扩展点，子类可以选择性重写 |

---

## 运行测试

```bash
cd 20-template-method-pattern
go test -v ./...
```
