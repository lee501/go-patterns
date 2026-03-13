# Visitor Pattern（访问者模式）

## 概述

访问者模式是一种**行为型**设计模式，将**作用于对象的操作**从对象结构中分离出来，使得可以在不修改对象类的前提下，为对象结构中的元素定义新的操作。

> Represent an operation to be performed on elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.

---

## 适用场景

- 需要对一个对象结构中的元素执行**多种不同且不相关的操作**，且不想让这些操作污染元素类
- 对象结构较稳定，但经常需要在此结构上定义**新的操作**
- 需要对不同类型的元素进行**集中处理**（统一聚合错误、统一收集结果）
- 类似 `kubectl` 等框架中对资源的批量访问和处理

---

## 结构

```
┌──────────────┐
│   Visitor    │  ← 访问者接口
│  Visit(fn)   │
└──────┬───────┘
       │ 实现
  ┌────┴───────────────────────┐
  │                            │
┌──────────────┐    ┌────────────────────────┐
│StreamVisitor │    │    EagerVisitorList     │
│ Source string│    │  []Visitor             │
│ Visit(fn)    │    │  Visit(fn) → 遍历聚合   │
└──────────────┘    └────────────────────────┘

┌──────────────┐
│  VisitorFunc │  ← 操作的函数类型
│ func(*Info,  │
│  error)error │
└──────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Info** | 被访问的对象，包含 Namespace、Name 等信息 |
| **Visitor（接口）** | 定义 `Visit(VisitorFunc) error`，接受一个操作函数 |
| **VisitorFunc** | 操作函数类型 `func(*Info, error) error`，即具体"操作" |
| **StreamVisitor** | 从流（文件、HTTP 等）中读取 Info 并执行访问 |
| **EagerVisitorList** | 聚合多个 Visitor，统一执行并收集所有错误 |
| **URLVisitor** | 从 URL 读取资源的访问者 |

---

## Go 实现

### 核心类型定义

```go
type Info struct {
    Namespace string
    Name      string
}

type Visitor interface {
    Visit(VisitorFunc) error
}

type VisitorFunc func(*Info, error) error
```

### 流式访问者

```go
type StreamVisitor struct {
    io.Reader
    Source string
}

func (s *StreamVisitor) Visit(fn VisitorFunc) error {
    info := &Info{
        Namespace: s.Source,
        Name:      s.Source,
    }
    return fn(info, nil)
}
```

### 聚合访问者（统一执行多个 Visitor）

```go
type EagerVisitorList []Visitor

func (l EagerVisitorList) Visit(fn VisitorFunc) error {
    errs := []error(nil)
    for i := range l {
        if err := l[i].Visit(func(info *Info, err error) error {
            if err != nil {
                errs = append(errs, err)
                return nil
            }
            if err := fn(info, nil); err != nil {
                errs = append(errs, err)
            }
            return nil
        }); err != nil {
            errs = append(errs, err)
        }
    }
    return nil // 返回聚合后的错误（可用 utilerrors.NewAggregate）
}
```

### 使用示例

```go
// 创建多个访问者
v1 := &StreamVisitor{Source: "file://config.yaml"}
v2 := &StreamVisitor{Source: "http://api/resource"}

// 聚合为一个统一访问者
visitors := EagerVisitorList{v1, v2}

// 定义操作函数
printInfo := func(info *Info, err error) error {
    if err != nil {
        return err
    }
    fmt.Printf("Namespace: %s, Name: %s\n", info.Namespace, info.Name)
    return nil
}

// 统一执行操作
visitors.Visit(printInfo)
// 输出:
// Namespace: file://config.yaml, Name: file://config.yaml
// Namespace: http://api/resource, Name: http://api/resource
```

---

## 设计亮点：表面与内部分离

```
表面：某个对象执行了一个方法
    visitors.Visit(fn)

内部：对象内部调用了多个 Visitor，最后统一聚合结果
    → v1.Visit(fn) → StreamVisitor 读取文件执行 fn
    → v2.Visit(fn) → StreamVisitor 读取 HTTP 执行 fn
    → 错误统一收集到 []error
```

---

## 优缺点

### 优点

- **开闭原则**：新增操作（`VisitorFunc`）无需修改元素类
- **单一职责**：不同的操作封装在不同的 VisitorFunc 中
- **统一处理**：`EagerVisitorList` 可以聚合多个来源的访问，并统一收集错误

### 缺点

- 若被访问的元素类型频繁变化，需要修改所有 Visitor 实现
- 访问者持有元素的内部信息，可能破坏封装性

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Visitor** | 将操作从对象结构中分离，在不修改对象的情况下添加新操作 |
| **Iterator** | 遍历集合元素，不关心元素上执行什么操作 |
| **Strategy** | 封装可互换的算法，关注算法本身 |
| **Composite** | 构建树形结构，Visitor 常与 Composite 配合使用 |

---

## 运行测试

```bash
cd 15-visitor-pattern
go test -v ./...
```
