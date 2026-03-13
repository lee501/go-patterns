# Flyweight Pattern（享元模式）

## 概述

享元模式是一种**结构型**设计模式，通过**共享**细粒度对象来减少内存占用。当系统中需要大量相似对象时，享元模式将对象的状态分为**内部状态**（可共享）和**外部状态**（不可共享），只共享内部状态相同的对象实例。

> Use sharing to efficiently support a large number of fine-grained objects.

---

## 适用场景

- 系统中存在**大量相似对象**，占用大量内存
- 对象的大部分状态可以**外部化**（由调用方传入）
- 去掉对象的外部状态后，可以用**较少的共享对象**取代大量对象
- 典型场景：游戏中的粒子系统（大量相同颜色的粒子）、文字处理中的字符对象、地图上的树木

---

## 结构

```
┌─────────────────────────────────────┐
│           ShapeFactory              │
│   circleMap map[string]Shape        │  ← 享元工厂（缓存池）
│   GetCircle(color string) Shape     │
└───────────────────┬─────────────────┘
                    │ 按 color 缓存
                    ▼
            ┌──────────────┐
            │    Shape     │  ← 享元接口
            ├──────────────┤
            │ SetRadius()  │
            │ SetColor()   │
            └──────┬───────┘
                   │ 实现
            ┌──────┴───────┐
            │    Circle    │  ← 具体享元对象
            │ color string │  ← 内部状态（可共享）
            │ radius int   │  ← 外部状态（每次设置）
            └──────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Shape（享元接口）** | 定义享元对象的公共接口 |
| **Circle（具体享元）** | 实现 Shape 接口，`color` 为内部状态（共享），`radius` 为外部状态（可变） |
| **ShapeFactory（享元工厂）** | 维护 `map[string]Shape` 缓存，相同 color 的 Circle 只创建一次 |

---

## Go 实现

### 享元接口

```go
type Shape interface {
    SetRadius(radius int)
    SetColor(color string)
}
```

### 具体享元对象

```go
type Circle struct {
    color  string // 内部状态：共享，由工厂管理
    radius int    // 外部状态：每次使用时设置
}

func (c *Circle) SetRadius(radius int) { c.radius = radius }
func (c *Circle) SetColor(color string) { c.color = color }
```

### 享元工厂

```go
type ShapeFactory struct {
    circleMap map[string]Shape
}

func (sh *ShapeFactory) GetCircle(color string) Shape {
    if sh.circleMap == nil {
        sh.circleMap = make(map[string]Shape)
    }
    if shape, ok := sh.circleMap[color]; ok {
        return shape // 命中缓存，返回已有对象
    }
    // 未命中，创建新对象并缓存
    circle := new(Circle)
    circle.SetColor(color)
    sh.circleMap[color] = circle
    return circle
}
```

### 使用示例

```go
factory := &ShapeFactory{}

// 相同颜色的 Circle 只创建一次
red1 := factory.GetCircle("red")
red1.SetRadius(10)

red2 := factory.GetCircle("red")
red2.SetRadius(20)

// red1 和 red2 是同一个对象
fmt.Println(red1 == red2)           // true（同一指针）
fmt.Println(red1.(*Circle).radius)  // 20（radius 被覆盖）

// 不同颜色创建不同对象
blue := factory.GetCircle("blue")
fmt.Println(red1 == blue)           // false
```

---

## 内部状态 vs 外部状态

| 状态类型 | 说明 | 示例 |
|----------|------|------|
| **内部状态（Intrinsic）** | 存储在享元对象内部，**可共享**，不随环境变化 | `color`（颜色相同则共享） |
| **外部状态（Extrinsic）** | 由客户端在使用时传入，**不可共享**，随使用场景变化 | `radius`（每次使用时设置） |

> **注意**：上述实现中 `radius` 存储在共享对象中，修改会影响所有持有该引用的调用方。生产环境中外部状态通常**不存储在享元对象内部**，而是由调用方自行维护。

---

## 优缺点

### 优点

- **节省内存**：大量相似对象被共享，显著降低内存占用
- **提升性能**：减少对象创建和垃圾回收的开销

### 缺点

- 引入享元工厂，代码复杂度增加
- 外部状态需要调用方自行管理，容易出错
- 共享对象不能随意修改内部状态（线程安全问题）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Flyweight** | 共享细粒度对象，减少内存占用 |
| **Singleton** | 全局只有一个实例 |
| **Object Pool** | 复用一组有限对象，使用完后归还 |
| **Prototype** | 通过克隆已有对象创建新对象 |

---

## 运行测试

```bash
cd 17-flyweight-pattern
go test -v ./...
```
