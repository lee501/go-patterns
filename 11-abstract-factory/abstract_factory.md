# Abstract Factory Pattern（抽象工厂模式）

## 概述

抽象工厂模式是一种**创建型**设计模式，提供一个创建**一系列相关对象**的接口，而无需指定它们的具体类。与工厂方法模式的区别在于，抽象工厂关注**产品族**（一组相互关联的产品），而工厂方法只关注**单个产品**。

> Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

---

## 适用场景

- 系统需要创建**一组相互关联的对象**（产品族），并保证它们的兼容性
- 需要在多套产品族之间切换（如换主题、换数据库驱动）
- 不希望客户端依赖具体的产品类，只依赖抽象接口
- 产品族扩展频繁，希望通过替换工厂来切换整套实现

---

## 结构

```
┌──────────────────┐       ┌──────────────────┐
│  Factory (接口)   │       │  Product (接口)   │
├──────────────────┤       ├──────────────────┤
│ CreateProduct()  │       │ Describe()       │
└──────────┬───────┘       └──────────┬───────┘
           │ 实现                      │ 实现
           ▼                          ▼
┌──────────────────┐       ┌──────────────────┐
│ ConcreteFactory  │──────>│ ConcreteProduct  │
│                  │       │  Name string     │
└──────────────────┘       └──────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Factory（抽象工厂接口）** | 声明创建产品族的方法 `CreateProduct()` |
| **Product（抽象产品接口）** | 声明产品的公共行为 `Describe()` |
| **ConcreteFactory** | 实现 `Factory` 接口，负责创建具体产品 |
| **ConcreteProduct** | 实现 `Product` 接口，是工厂生产的最终对象 |

---

## Go 实现

### 抽象接口

```go
type Factory interface {
    CreateProduct() Product
}

type Product interface {
    Describe()
}
```

### 具体产品

```go
type ConcreteProduct struct {
    Name string
}

func (p *ConcreteProduct) Describe() {
    fmt.Println(p.Name)
}
```

### 具体工厂

```go
type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct() Product {
    return &ConcreteProduct{Name: "KG"}
}
```

### 使用示例

```go
var factory Factory = &ConcreteFactory{}
product := factory.CreateProduct()
product.Describe() // 输出: KG
```

---

## 扩展：多产品族

抽象工厂的核心价值在于支持整套产品族的切换，下面演示扩展多个工厂和产品族：

```go
// 工厂 B（另一个产品族）
type FactoryB struct{}

func (f *FactoryB) CreateProduct() Product {
    return &ConcreteProduct{Name: "FactoryB-Product"}
}

// 客户端代码只依赖 Factory 接口，切换工厂即切换整套产品
factories := []Factory{&ConcreteFactory{}, &FactoryB{}}
for _, f := range factories {
    f.CreateProduct().Describe()
}
// 输出:
// KG
// FactoryB-Product
```

---

## 抽象工厂 vs 工厂方法

| 特性 | 工厂方法 | 抽象工厂 |
|------|----------|----------|
| 创建对象数量 | 单个产品 | 多个相关产品（产品族） |
| 扩展方式 | 新增工厂子类 | 新增工厂 + 对应产品族 |
| 关注点 | 对象创建的抽象 | 产品族的一致性 |
| 代码复杂度 | 较简单 | 较复杂 |

---

## 优缺点

### 优点

- **产品族一致性**：同一工厂创建的产品保证互相兼容
- **开闭原则**：新增产品族只需新增工厂类，不修改已有代码
- **解耦**：客户端只依赖抽象接口，完全屏蔽具体实现

### 缺点

- **扩展产品类型困难**：如果要为现有产品族新增一种产品，需要修改所有工厂接口及其实现
- **类数量增多**：每个产品族都需要一套对应的工厂和产品类

---

## 运行测试

```bash
cd 11-abstract-factory
go test -v ./...
```
