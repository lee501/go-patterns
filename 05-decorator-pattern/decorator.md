# Decorator Pattern（装饰器模式）

## 概述

装饰器模式是一种**结构型**设计模式，通过对象组合的方式**动态地**为对象添加额外职责，相比继承更加灵活。

> Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.

---

## 适用场景

- 需要在**不修改原有类**的情况下扩展功能
- 需要**动态地、可撤销地**为对象添加职责
- 功能扩展组合方式多样，继承会产生大量子类
- 希望将核心职责与装饰职责分离（如日志、缓存、权限校验）

---

## 结构

```
┌──────────────┐
│  Component   │  ← 公共接口（Describe / GetCount）
└──────┬───────┘
       │ 实现            ┌──────────────────────┐
  ┌────┴────┐            │   AppleDecorator     │
  │  Fruit  │            ├──────────────────────┤
  │(基础组件) │            │ Component (匿名组合) │
  └─────────┘            │ Type  string         │
                         │ Num   int            │
       ▲                 └──────────────────────┘
       │ 包装（组合）              │
       └──────────────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Component（接口）** | 定义被装饰对象和装饰器的公共接口 |
| **ConcreteComponent（基础组件）** | 实现 Component 接口的原始对象，如 `Fruit` |
| **Decorator（装饰器）** | 通过**匿名组合**持有 Component 接口引用，并实现同一接口 |
| **ConcreteDecorator** | 具体装饰器，重写接口方法并在调用原方法的基础上扩展，如 `AppleDecorator` |

---

## Go 实现

### 方式一：结构体组合装饰

#### 公共接口

```go
type Component interface {
    Describe() string
    GetCount() int
}
```

#### 基础组件

```go
type Fruit struct {
    Count       int
    Description string
}

func (f *Fruit) Describe() string { return f.Description }
func (f *Fruit) GetCount() int    { return f.Count }
```

#### 装饰器（匿名组合 Component）

```go
type AppleDecorator struct {
    Component       // 匿名组合，持有被装饰对象的引用
    Type string
    Num  int
}

func (apple *AppleDecorator) Describe() string {
    // 在原有描述基础上追加信息
    return fmt.Sprintf("%s, %s", apple.Component.Describe(), apple.Type)
}

func (apple *AppleDecorator) GetCount() int {
    // 在原有数量基础上叠加
    return apple.Component.GetCount() + apple.Num
}

func CreateAppleDecorator(c Component, t string, n int) Component {
    return &AppleDecorator{c, t, n}
}
```

#### 使用示例

```go
// 原始水果
fruit := &Fruit{Count: 5, Description: "Fruit basket"}

// 用苹果装饰器包装
apple := CreateAppleDecorator(fruit, "Apple", 3)
fmt.Println(apple.Describe())  // "Fruit basket, Apple"
fmt.Println(apple.GetCount())  // 8

// 可以叠加多层装饰
orange := CreateAppleDecorator(apple, "Orange", 2)
fmt.Println(orange.Describe()) // "Fruit basket, Apple, Orange"
fmt.Println(orange.GetCount()) // 10
```

---

### 方式二：函数装饰（高阶函数）

Go 中还可以用**高阶函数**实现轻量级装饰器，适合对函数本身进行增强：

```go
type Object func(int) int

func LogDecorate(fn Object) Object {
    return func(i int) int {
        log.Println("starting inner func")
        result := fn(i)
        log.Println("complete inner func")
        return result
    }
}
```

#### 使用示例

```go
double := func(i int) int { return i * 2 }

// 用日志装饰器包装函数
loggedDouble := LogDecorate(double)
result := loggedDouble(5)
// 输出:
// starting inner func
// complete inner func
// result = 10
```

---

## 优缺点

### 优点

- **开闭原则**：不修改原有代码即可扩展功能
- **灵活组合**：可以叠加多个装饰器，运行时动态组合
- **单一职责**：将核心逻辑和附加功能解耦

### 缺点

- 多层装饰后，调试时调用链较深，不易追踪
- 装饰器顺序影响结果，使用时需注意叠加顺序

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Decorator** | 动态为对象添加职责，支持叠加，不改变接口 |
| **Proxy** | 控制对原对象的访问（如权限、缓存），通常只有一层 |
| **Composite** | 将对象组合成树形结构，关注部分-整体层次 |
| **Adapter** | 转换接口，使不兼容的接口能协同工作 |

---

## 运行测试

```bash
cd 05-decorator-pattern
go test -v ./...
```
