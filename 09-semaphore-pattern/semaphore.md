# Semaphore Pattern（信号量模式）

## 概述

信号量模式是一种**并发控制**模式，通过限制同时访问某一资源的 goroutine（或线程）数量，防止资源过载，保证系统稳定性。

> Use a semaphore to control access to a resource by limiting the number of concurrent users.

---

## 适用场景

- **限制并发数**：如最多允许 N 个 goroutine 同时访问数据库
- **流量控制**：限制接口的并发请求数量
- **资源池保护**：防止有限资源（连接、文件句柄）被耗尽
- **速率限制**：结合超时机制实现带超时的并发控制

---

## 结构

```
┌─────────────┐   Acquire()    ┌──────────────────────┐
│  Goroutine  │ ─────────────> │  Semaphore           │
│  (请求资源)  │                │  sem chan struct{}    │
│             │ <───────────── │  timeout Duration    │
└─────────────┘   Release()    └──────────────────────┘
                                        │
                               最多 N 个并发（chan 容量）
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Interface** | 信号量接口，定义 `Acquire()` 和 `Release()` 方法 |
| **Semaphore** | 具体实现，使用 `chan struct{}` 控制并发计数 |
| **sem** | 有缓冲 channel，容量即为最大并发数（tickets） |
| **timeout** | 获取/释放信号量的超时时间，防止永久阻塞 |

---

## Go 实现

### 接口定义

```go
var (
    ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
    ErrIllegalRelease = errors.New("semaphore: can't release semaphore without acquiring it first")
)

type Interface interface {
    Acquire() error
    Release() error
}
```

### 信号量结构体

```go
type Semaphore struct {
    sem     chan struct{}   // 有缓冲 channel，容量 = 最大并发数
    timeout time.Duration  // 超时时间
}
```

### Acquire 与 Release

```go
func (s *Semaphore) Acquire() error {
    select {
    case s.sem <- struct{}{}: // 向 channel 写入，成功则获得资源
        return nil
    case <-time.After(s.timeout): // 超时则返回错误
        return ErrNoTickets
    }
}

func (s *Semaphore) Release() error {
    select {
    case <-s.sem: // 从 channel 读取，释放一个占位
        return nil
    case <-time.After(s.timeout):
        return ErrIllegalRelease
    }
}

func New(tickets int, timeout time.Duration) Interface {
    return &Semaphore{
        sem:     make(chan struct{}, tickets),
        timeout: timeout,
    }
}
```

### 使用示例

```go
// 创建最多允许 3 个并发的信号量，超时 100ms
sem := New(3, 100*time.Millisecond)

var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        if err := sem.Acquire(); err != nil {
            fmt.Printf("goroutine %d: 获取信号量失败: %v\n", id, err)
            return
        }
        defer sem.Release()

        // 临界区：最多 3 个 goroutine 同时执行
        fmt.Printf("goroutine %d: 正在执行\n", id)
        time.Sleep(50 * time.Millisecond)
    }(i)
}
wg.Wait()
```

---

## channel 实现原理

```
sem = make(chan struct{}, 3)  // 容量为 3 的缓冲 channel

Acquire(): sem <- struct{}{} → 容量用完时阻塞（或超时）
Release(): <-sem             → 释放一个槽位，唤醒等待中的 Acquire
```

| 状态 | channel 占用 | 可并发数 |
|------|-------------|---------|
| 空闲 | 0/3 | 3 |
| 半满 | 2/3 | 1 |
| 满载 | 3/3 | 0（等待或超时） |

---

## 优缺点

### 优点

- **并发安全**：基于 channel，无需额外加锁
- **超时控制**：通过 `select + time.After` 避免永久阻塞
- **简洁实现**：利用 Go channel 语义，代码清晰
- **可扩展**：可轻松调整并发上限（修改 tickets）

### 缺点

- 超时设置不当可能导致误拒（超时太短）或资源耗尽（超时太长）
- 不支持优先级队列（等待的 goroutine 无法按优先级排序）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Semaphore** | 限制同时访问资源的并发数量，支持超时 |
| **Object Pool** | 复用有限数量的对象，获取后独占使用 |
| **Mutex** | 同一时刻只允许一个访问者（信号量 tickets=1 的特例） |
| **Rate Limiter** | 控制单位时间内的请求频率（令牌桶/漏桶） |

---

## 运行测试

```bash
cd 09-semaphore-pattern
go test -v ./...
```
