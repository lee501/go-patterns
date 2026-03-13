# Object Pool Pattern（对象池模式）

## 概述

对象池模式是一种**创建型**设计模式，预先创建一批对象并放入"池"中，需要时从池中获取，用完后归还，从而避免频繁创建和销毁对象带来的性能开销。

> Reuse a set of initialized objects rather than allocating and deallocating them on demand.

---

## 适用场景

- 对象**创建代价高**（如数据库连接、网络连接、线程）
- 同一时刻只需要少量对象，但对象创建和销毁非常频繁
- 对象**可以被复用**，使用前后状态可以重置
- 需要**限制资源使用上限**（如最大连接数）

---

## 结构

```
┌─────────────┐      Acquire()      ┌───────────────┐
│   Client    │ ──────────────────> │  Pool (chan)   │
│             │ <────────────────── │               │
│             │      Release()      │  *Object × N  │
└─────────────┘                     └───────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Pool** | 基于 `chan *Object` 实现的对象池，利用 channel 的阻塞特性做并发安全的存取 |
| **Object** | 被池化管理的对象 |
| **Acquire()** | 从池中取出一个对象（channel 读取） |
| **Release()** | 将对象归还到池中（channel 写入） |

---

## Go 实现

### 对象与池类型定义

```go
type Object struct {
    Name string
}

// Pool 本质是一个带缓冲的 channel
type Pool chan *Object
```

> 利用 `chan` 的缓冲特性，天然支持并发安全，无需额外加锁。

### 创建对象池

```go
func NewPool(count int) *Pool {
    pool := make(Pool, count)
    for i := 0; i < count; i++ {
        pool <- new(Object)  // 预创建对象，放入池中
    }
    return &pool
}
```

### 获取与归还对象

```go
// Acquire 从池中获取一个对象（若池为空则阻塞）
func (p *Pool) Acquire() *Object {
    return <-(*p)
}

// Release 将对象归还到池中
func (p *Pool) Release(obj *Object) {
    (*p) <- obj
}
```

### 使用示例

```go
pool := NewPool(3) // 创建容量为 3 的对象池

// 获取对象并使用
obj := pool.Acquire()
obj.Name = "worker-1"
obj.Do()

// 用完归还
pool.Release(obj)

// 并发场景
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        obj := pool.Acquire()      // 最多 3 个对象同时被使用，其余 goroutine 阻塞等待
        defer pool.Release(obj)
        obj.Name = fmt.Sprintf("worker-%d", id)
        obj.Do()
    }(i)
}
wg.Wait()
```

---

## channel 实现的优势

| 特性 | 说明 |
|------|------|
| **并发安全** | `chan` 内置互斥，无需 `sync.Mutex` |
| **阻塞等待** | 池为空时 `Acquire()` 自动阻塞，池满时 `Release()` 自动阻塞 |
| **容量限制** | `make(Pool, count)` 直接限制池的最大对象数 |
| **零值友好** | 利用 Go 原生 channel 语义，代码极简 |

---

## 优缺点

### 优点

- **性能**：复用对象，减少 GC 压力和内存分配开销
- **并发安全**：基于 channel 无锁实现
- **资源上限**：天然限制同时使用对象的数量

### 缺点

- 若 `Acquire()` 后未调用 `Release()`，会导致对象泄漏，池逐渐耗尽
- 对象需要在归还前**重置状态**，否则可能携带上一次使用的脏数据
- 池容量固定，运行时无法动态扩缩

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Object Pool** | 复用一组预创建的对象，控制资源总量 |
| **Singleton** | 全局只有一个实例 |
| **Flyweight** | 共享细粒度对象，减少内存占用，对象通常是只读的 |
| **Factory Method** | 每次调用都创建新对象 |

---

## 运行测试

```bash
cd 03-object-pool-pattern
go test -v ./...
```
