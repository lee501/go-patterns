# Singleton Pattern（单例模式）

## 概述

单例模式是一种**创建型**设计模式，确保一个类型**只有一个实例**，并提供一个全局访问点。

> Ensure a class has only one instance, and provide a global point of access to it.

---

## 适用场景

- 全局配置对象（只需初始化一次）
- 日志记录器（全局共享一个实例）
- 数据库连接池（避免重复创建）
- 缓存对象（全局唯一共享）
- 需要协调整个系统行为的管理器

---

## 结构

```
┌─────────────┐         ┌────────────────────┐
│   Client    │ ──New()─>│  sync.Once.Do()    │
└─────────────┘         │  if instance==nil  │
                        │    create instance │
                        └────────────┬───────┘
                                     │
                                     ▼
                             ┌──────────────┐
                             │   instance   │  ← 全局唯一实例
                             │ (singleton)  │
                             └──────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **singleton** | 被单例化的类型（此处为 `map[string]string`） |
| **instance** | 包级全局变量，保存唯一实例 |
| **once** | `sync.Once`，保证初始化代码只执行一次，线程安全 |
| **New()** | 全局访问点，返回唯一实例 |

---

## Go 实现

### 类型与全局变量

```go
type singleton map[string]string

var (
    once     sync.Once
    instance singleton
)
```

### 全局访问函数

```go
func New() singleton {
    once.Do(func() {
        instance = make(singleton)
    })
    return instance
}
```

> `sync.Once` 保证 `once.Do` 内的函数**只执行一次**，即使多个 goroutine 同时调用 `New()`，也只会创建一个实例。

### 使用示例

```go
// 多处调用 New()，返回的是同一个实例
s1 := New()
s2 := New()

s1["key"] = "value"
fmt.Println(s2["key"]) // 输出: value

fmt.Println(s1 == nil) // false
// s1 和 s2 指向同一个底层 map
```

---

## 并发安全性

| 方式 | 线程安全 | 说明 |
|------|----------|------|
| 全局变量直接赋值 | ❌ | 存在竞态条件 |
| `sync.Mutex` 双重检查 | ✅ | 代码较复杂 |
| **`sync.Once`** | ✅ | Go 惯用方式，简洁高效 |
| `init()` 函数 | ✅ | 程序启动时初始化，无法延迟加载 |

Go 中推荐使用 `sync.Once` 实现懒加载单例，它内部通过原子操作保证并发安全，性能优于加锁方案。

---

## 优缺点

### 优点

- **全局访问**：提供统一的访问入口
- **节省资源**：避免重复创建开销大的对象
- **延迟初始化**：`sync.Once` 实现懒加载，首次使用时才创建

### 缺点

- **全局状态**：引入全局状态，增加代码耦合度，不利于单元测试
- **违反单一职责**：单例类既要管理自己的业务逻辑，又要管理自己的生命周期
- **并发写入问题**：实例本身（如 `map`）的并发读写仍需自行加锁

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Singleton** | 全局只有一个实例，限制实例数量 |
| **Object Pool** | 管理一组可复用对象，限制总量但允许多个 |
| **Factory Method** | 每次调用创建新实例 |
| **Monostate** | 允许多个实例，但所有实例共享相同状态 |

---

## 运行测试

```bash
cd 04-singleton-pattern
go test -v ./...
```
