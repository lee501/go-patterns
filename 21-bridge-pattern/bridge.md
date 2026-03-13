# Bridge Pattern（桥接模式）

## 概述

桥接模式是一种**结构型**设计模式，将**抽象部分**与其**实现部分**分离，使二者都可以独立地变化。通过组合而非继承，避免因类层次结构爆炸而导致的代码膨胀。

> Decouple an abstraction from its implementation so that the two can vary independently.

---

## 适用场景

- 希望**抽象和实现**都能独立扩展，互不影响
- 需要在多个维度上扩展，避免多重继承导致的类爆炸
- 运行时需要**动态切换实现**
- 典型场景：跨平台 UI 组件（不同 OS + 不同控件）、不同品牌手机 + 不同硬件组件

---

## 结构

```
              ┌───────────────────┐
              │   SoftWare (接口) │  ← 实现层（Implementation）
              │   Run()           │
              └─────────┬─────────┘
                        │ 实现
              ┌─────────┴─────────┐
              │                   │
         ┌─────────┐        ┌───────────┐
         │   Cpu   │        │  Storage  │
         └─────────┘        └───────────┘

┌─────────────────────────────────┐
│           Phone（桥接层）        │  ← 抽象层（Abstraction）
│   software SoftWare            │  ← 持有实现层接口引用（桥）
│   SetSoftWare(SoftWare)        │
└──────────────┬──────────────────┘
               │ 组合
    ┌──────────┴──────────┐
    │                     │
┌─────────┐         ┌──────────┐
│  Apple  │         │ HuaWei   │  ← 抽象的扩展（Refined Abstraction）
│ Print() │         │ Print()  │
└─────────┘         └──────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **SoftWare（实现接口）** | 定义底层实现的接口，如 `Run()` |
| **Cpu / Storage（具体实现）** | 实现 `SoftWare` 接口的具体类 |
| **Phone（桥接层/抽象）** | 持有 `SoftWare` 接口引用，将调用委托给具体实现 |
| **Apple / HuaWei（精化抽象）** | 组合 `Phone`，使用桥接层调用底层实现 |

---

## Go 实现

### 实现层接口

```go
type SoftWare interface {
    Run()
}

type Cpu struct{}
func (c *Cpu) Run() { fmt.Println("this is cpu run") }

type Storage struct{}
func (s *Storage) Run() { fmt.Println("this is storage run") }
```

### 桥接层

```go
type Phone struct {
    software SoftWare // 桥：持有实现层接口引用
}

func (p *Phone) SetSoftWare(soft SoftWare) {
    p.software = soft
}
```

### 精化抽象（具体品牌）

```go
type Apple struct {
    phone Phone
}

func (p *Apple) SetShape(soft SoftWare) { p.phone.SetSoftWare(soft) }
func (p *Apple) Print()                 { p.phone.software.Run() }

type HuaWei struct {
    phone Phone
}

func (p *HuaWei) SetShape(soft SoftWare) { p.phone.SetSoftWare(soft) }
func (p *HuaWei) Print()                 { p.phone.software.Run() }
```

### 使用示例

```go
// Apple 手机 + CPU
apple := &Apple{}
apple.SetShape(&Cpu{})
apple.Print() // this is cpu run

// Apple 手机 + Storage
apple.SetShape(&Storage{})
apple.Print() // this is storage run

// HuaWei 手机 + CPU（品牌与硬件独立变化）
huawei := &HuaWei{}
huawei.SetShape(&Cpu{})
huawei.Print() // this is cpu run
```

---

## 桥接模式解决的问题

若不使用桥接，品牌 × 硬件会产生 M×N 个子类：

```
Apple+Cpu, Apple+Storage, HuaWei+Cpu, HuaWei+Storage ...
```

使用桥接后，只需 M（品牌）+ N（硬件）个类，**独立扩展，互不影响**：

```
品牌维度：Apple, HuaWei, Xiaomi ...    （新增只需 +1 个品牌类）
硬件维度：Cpu, Storage, GPU ...        （新增只需 +1 个硬件类）
组合方式：运行时动态设置（SetShape）
```

---

## 优缺点

### 优点

- **避免类爆炸**：抽象和实现独立扩展，不需要 M×N 个子类
- **开闭原则**：新增品牌或硬件无需修改已有代码
- **运行时切换实现**：可以在运行时动态改变硬件（`SetSoftWare()`）

### 缺点

- 引入额外的桥接层，增加了系统的抽象性和理解难度
- 对于简单场景，桥接模式可能是过度设计

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Bridge** | 分离抽象与实现，两者独立扩展，基于组合 |
| **Adapter** | 转换已有接口，解决不兼容问题，通常是事后补救 |
| **Strategy** | 封装可互换算法，关注行为切换，不涉及层次结构 |
| **Decorator** | 动态叠加职责，不改变接口，强调功能增强 |

---

## 运行测试

```bash
cd 21-bridge-pattern
go test -v ./...
```
