# Generator Pattern（生成器模式）

## 概述

生成器模式是一种**并发**设计模式，利用 Go 的 goroutine 和 channel 实现**惰性序列生成**，调用方按需从 channel 中读取值，而不需要一次性将所有值加载到内存中。

> Use a goroutine and a channel to lazily produce a sequence of values, similar to the `yield` keyword in Python or JavaScript generators.

---

## 适用场景

- 需要生成**无限序列**或**大数据集**，不适合一次性放入内存
- 需要**惰性求值**（lazy evaluation），按需生成
- 数据生产和消费速率不一致，需要**异步解耦**
- 实现**流式处理**管道（pipeline）

---

## 结构

```
┌─────────────────────────────────┐
│   Count(start, end int)         │
│                                 │
│   ch := make(chan int)          │
│   go func() {                   │
│       for i := start; i<=end;   │
│           ch <- i               │
│       close(ch)                 │
│   }()                           │
│   return <-chan int  (只读channel)│
└────────────────┬────────────────┘
                 │
                 ▼ 调用方 range 读取
           ┌──────────┐
           │ for v := │
           │ range ch │
           └──────────┘
```

### 核心要素

| 要素 | 说明 |
|------|------|
| **返回 `<-chan T`** | 只读 channel，外部只能读取，不能写入，防止误用 |
| **内部 goroutine** | 并发生成数据并写入 channel，与消费方解耦 |
| **`close(ch)`** | 生成结束后关闭 channel，触发 `range` 循环退出 |
| **惰性生成** | 消费者不读取时，goroutine 会阻塞在 `ch <- i`，不浪费计算 |

---

## Go 实现

```go
func Count(start, end int) <-chan int {
    ch := make(chan int)

    go func(ch chan int) {
        for i := start; i <= end; i++ {
            ch <- i       // 生产者：逐个写入值
        }
        close(ch)         // 通知消费者：数据已全部生成
    }(ch)

    return ch             // 返回只读 channel
}
```

### 使用示例

```go
// 生成 1~5 的整数序列
gen := Count(1, 5)

// 方式一：range 遍历（推荐）
for v := range gen {
    fmt.Println(v) // 1, 2, 3, 4, 5
}

// 方式二：手动读取
gen2 := Count(1, 3)
fmt.Println(<-gen2) // 1
fmt.Println(<-gen2) // 2
fmt.Println(<-gen2) // 3
```

---

## 构建 Pipeline（管道）

生成器可以与其他 channel 处理函数组合，构建数据处理管道：

```go
// 生成器
func Count(start, end int) <-chan int { ... }

// 过滤器：只保留偶数
func FilterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            if v%2 == 0 {
                out <- v
            }
        }
        close(out)
    }()
    return out
}

// 翻倍处理器
func Double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            out <- v * 2
        }
        close(out)
    }()
    return out
}

// 组合 pipeline
for v := range Double(FilterEven(Count(1, 10))) {
    fmt.Println(v) // 4, 8, 12, 16, 20
}
```

---

## 优缺点

### 优点

- **内存高效**：逐个生成值，无需一次性分配大量内存
- **并发友好**：生产者与消费者天然并发，互不阻塞
- **可组合**：多个生成器/处理器可以链式组合为 pipeline
- **简洁**：比手动管理迭代器状态更简洁

### 缺点

- goroutine 泄漏风险：若消费者提前退出而不读完 channel，生产者 goroutine 会永久阻塞
  - 解决方案：使用 `context.Context` 配合 `Done()` channel 进行优雅退出
- 无法随机访问，只能顺序消费

---

## 防止 goroutine 泄漏（配合 context）

```go
func Count(ctx context.Context, start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := start; i <= end; i++ {
            select {
            case ch <- i:
            case <-ctx.Done(): // 消费者取消时，生产者退出
                return
            }
        }
    }()
    return ch
}
```

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Generator** | 惰性生成序列，生产者驱动，基于 channel |
| **Iterator** | 遍历已存在的集合，消费者驱动，调用 `Next()` |
| **Object Pool** | 复用一组固定对象，而非生成新对象 |
| **Pipeline** | 多个生成器/处理器串联，是 Generator 的扩展应用 |

---

## 运行测试

```bash
cd 10-generator-pattern
go test -v ./...
```
