# Strategy Pattern（策略模式）

## 概述

策略模式是一种**行为型**设计模式，定义一系列算法，将每个算法封装起来，并使它们可以相互替换。策略模式让算法的变化独立于使用算法的客户端。

> Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

---

## 适用场景

- 需要在运行时**动态切换算法**或行为
- 有多种相似行为（算法），只是具体实现不同
- 避免使用大量 `if-else` 或 `switch-case` 判断算法类型
- 希望将算法实现从调用方中分离，降低耦合

---

## 结构

```
┌────────────────┐        ┌──────────────────┐
│   Operation    │───────>│  Operator (接口)  │
│                │        ├──────────────────┤
│ Operate(l,r)  │        │ Apply(int,int)int │
└────────────────┘        └──────────────────┘
                                   ▲
                        ┌──────────┴──────────┐
                        │                     │
                ┌───────────────┐   ┌──────────────────┐
                │   Addition    │   │  Multiplication  │
                └───────────────┘   └──────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Operator（Strategy 接口）** | 定义算法的公共接口 `Apply(int, int) int` |
| **Addition / Multiplication** | 具体策略，实现 `Operator` 接口，封装各自算法 |
| **Operation（Context）** | 持有 `Operator` 接口引用，通过 `Operate()` 委托给具体策略执行 |

---

## Go 实现

### 策略接口

```go
type Operator interface {
    Apply(int, int) int
}
```

### 上下文（包装器）

```go
type Operation struct {
    operator Operator
}

func (op *Operation) Operate(left, right int) int {
    return op.operator.Apply(left, right)
}
```

### 具体策略

```go
type Addition struct{}

func (add *Addition) Apply(left, right int) int {
    return left + right
}

type Multiplication struct{}

func (mu *Multiplication) Apply(left, right int) int {
    return left * right
}
```

### 工厂函数

```go
func CreateOpration(operator Operator) Operation {
    return Operation{operator}
}
```

### 使用示例

```go
// 使用加法策略
add := CreateOpration(&Addition{})
result := add.Operate(3, 4) // 7

// 使用乘法策略
mul := CreateOpration(&Multiplication{})
result = mul.Operate(3, 4) // 12

// 运行时切换策略
op := Operation{}
op.operator = &Addition{}
fmt.Println(op.Operate(10, 5)) // 15
op.operator = &Multiplication{}
fmt.Println(op.Operate(10, 5)) // 50
```

---

## 扩展：添加新策略

只需新增实现 `Operator` 接口的 struct，**无需修改 Operation**：

```go
type Subtraction struct{}

func (s *Subtraction) Apply(left, right int) int {
    return left - right
}

sub := CreateOpration(&Subtraction{})
fmt.Println(sub.Operate(10, 3)) // 7
```

---

## 优缺点

### 优点

- **开闭原则**：新增策略无需修改上下文代码
- **消除条件语句**：用多态替代 `if-else`/`switch-case`
- **可测试性**：每个策略可独立测试
- **运行时切换**：可动态替换算法

### 缺点

- 策略类数量会随算法增多而增加
- 客户端必须了解不同策略之间的区别，才能选择合适的策略

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Strategy** | 封装可互换的算法，运行时切换 |
| **State** | 根据对象内部状态自动切换行为，客户端无需感知 |
| **Template Method** | 在父类中定义算法骨架，子类重写特定步骤 |
| **Command** | 将请求封装为对象，支持撤销/重做 |

---

## 运行测试

```bash
cd 08-strategy-pattern
go test -v ./...
```
