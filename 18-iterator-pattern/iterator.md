# Iterator Pattern（迭代器模式）

## 概述

迭代器模式是一种**行为型**设计模式，提供一种方法**顺序访问集合中的各个元素**，而无需暴露集合的内部表示（数组、链表、树等）。

> Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

---

## 适用场景

- 需要一种**统一方式**遍历不同类型的集合
- 希望**隐藏集合的内部结构**，只暴露遍历接口
- 需要支持**多种遍历方式**（正序、逆序、过滤等）
- 集合类型较多，不想在业务代码中写各种类型的遍历逻辑

---

## 结构

```
┌─────────────────────────┐
│       Container         │  ← 集合容器
│   list []Visitor        │
│   Add(Visitor)          │
│   Remove(index int)     │
└────────────┬────────────┘
             │ 组合
┌────────────▼────────────┐
│        Iterator         │  ← 迭代器
│   index int             │
│   Container             │
│   Next()  Visitor       │
│   HasNext() bool        │
└─────────────────────────┘
                    │ 迭代的元素类型
             ┌──────┴──────┐
             │   Visitor   │  ← 元素接口
             │   Visit()   │
             └──────┬──────┘
                    │ 实现
          ┌─────────┴─────────┐
          │                   │
    ┌───────────┐       ┌────────────┐
    │  Teacher  │       │  Analysis  │
    └───────────┘       └────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Visitor（元素接口）** | 集合中存储的元素类型，声明 `Visit()` 方法 |
| **Container（容器）** | 存储元素的集合，提供 `Add`/`Remove` 操作 |
| **Iterator（迭代器）** | 持有 `Container` 引用和当前索引，提供 `Next()`/`HasNext()` |
| **Teacher / Analysis** | 具体元素，实现 `Visitor` 接口 |

---

## Go 实现

### 元素接口

```go
type Visitor interface {
    Visit()
}

type Teacher struct{}
func (t *Teacher) Visit() { fmt.Println("this is teacher visitor") }

type Analysis struct{}
func (a *Analysis) Visit() { fmt.Println("this is analysis visitor") }
```

### 容器

```go
type Container struct {
    list []Visitor
}

func (c *Container) Add(visitor Visitor) {
    c.list = append(c.list, visitor)
}

func (c *Container) Remove(index int) {
    if index < 0 || index > len(c.list) {
        return
    }
    c.list = append(c.list[:index], c.list[index+1:]...)
}
```

### 迭代器

```go
type Iterator struct {
    index int
    Container
}

func (i *Iterator) HasNext() bool {
    return i.index < len(i.list)
}

func (i *Iterator) Next() Visitor {
    visitor := i.list[i.index]
    i.index++
    return visitor
}

func NewIterator() *Iterator {
    return &Iterator{index: 0, Container: Container{}}
}
```

### 使用示例

```go
iter := NewIterator()
iter.Add(&Teacher{})
iter.Add(&Analysis{})
iter.Add(&Teacher{})

// 顺序遍历
for iter.HasNext() {
    v := iter.Next()
    v.Visit()
}
// 输出:
// this is teacher visitor
// this is analysis visitor
// this is teacher visitor
```

---

## 优缺点

### 优点

- **封装集合结构**：客户端无需了解集合的内部表示
- **统一遍历接口**：不同类型的集合可以使用相同的迭代方式
- **单一职责**：遍历逻辑从集合类中分离到迭代器中

### 缺点

- 对于简单的集合（如 slice），直接 `range` 比迭代器更简洁
- 迭代器不支持随机访问，也无法在遍历时修改集合（否则 index 可能失效）

---

## Go 原生迭代

Go 内置 `range` 关键字通常已足够，迭代器模式更适用于：
- 集合结构复杂（如树、图）
- 需要多个并发迭代器各自维护独立位置
- 需要自定义遍历顺序（逆序、跳步等）

```go
// Go 原生方式
list := []Visitor{&Teacher{}, &Analysis{}}
for _, v := range list {
    v.Visit()
}
```

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Iterator** | 顺序遍历集合，隐藏内部结构 |
| **Visitor** | 在不修改元素的前提下，对元素执行新操作 |
| **Composite** | 构建树形结构，Iterator 常用于遍历 Composite 结构 |
| **Generator** | 惰性生成序列，基于 channel，是 Iterator 的并发变体 |

---

## 运行测试

```bash
cd 18-iterator-pattern
go test -v ./...
```
