# Factory Method Pattern（工厂方法模式）

## 概述

工厂方法模式是一种**创建型**设计模式，定义一个用于创建对象的接口（工厂函数），让调用方通过**类型参数**来决定实例化哪个具体类，而无需直接依赖具体类的构造细节。

> Define an interface for creating an object, but let the caller decide which class to instantiate. The factory method pattern lets a function defer instantiation to subclasses or type switches.

---

## 适用场景

- 调用方**不需要知道**所创建对象的具体类型
- 对象的创建逻辑需要**集中管理**，便于后续扩展
- 需要根据**运行时参数**决定返回哪种实现
- 希望通过**统一接口**操作不同类型的对象

---

## 结构

```
┌─────────────────────────┐
│   GeneratePayment()     │  ← 工厂函数（Factory Method）
│   (Kind, balance) →     │
│   (Payment, error)      │
└────────────┬────────────┘
             │ 按 Kind 类型分发
             ▼
     ┌───────────────┐
     │  Payment 接口  │
     ├───────────────┤
     │  Pay(float32) │
     └───────────────┘
            ▲
     ┌──────┴──────┐
     │             │
┌─────────┐  ┌──────────┐
│ CashPay │  │CreditPay │
└─────────┘  └──────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Product（产品接口）** | 定义所有产品的公共行为，如 `Payment` |
| **ConcreteProduct（具体产品）** | 实现产品接口，如 `CashPay`、`CreditPay` |
| **Factory Method（工厂函数）** | 接收类型参数，返回对应的产品实例，如 `GeneratePayment()` |
| **Kind（类型常量）** | 使用 `iota` 定义产品类型，标识需要创建的具体产品 |

---

## Go 实现

### 类型常量（Kind）

```go
type Kind int

const (
    Cash   Kind = 1 << iota // 1
    Credit                  // 2
)
```

> 使用位移 `iota` 定义类型常量，便于扩展。

### 产品接口（Product Interface）

```go
type Payment interface {
    Pay(money float32) error
}
```

### 具体产品（Concrete Products）

```go
type CashPay struct {
    Balance float32
}

func (cash *CashPay) Pay(money float32) error {
    if cash.Balance < 0 || cash.Balance < money {
        return errors.New("balance not enough")
    }
    cash.Balance -= money
    return nil
}

type CreditPay struct {
    Balance float32
}

func (credit *CreditPay) Pay(money float32) error {
    if credit.Balance < 0 || credit.Balance < money {
        return errors.New("balance not enough")
    }
    credit.Balance -= money
    return nil
}
```

### 工厂函数（Factory Method）

```go
func GeneratePayment(k Kind, balance float32) (Payment, error) {
    switch k {
    case Cash:
        cash := new(CashPay)
        cash.Balance = balance
        return cash, nil
    case Credit:
        return &CreditPay{balance}, nil
    default:
        return nil, errors.New("Payment do not support this")
    }
}
```

### 使用示例

```go
// 创建现金支付对象
payment, err := GeneratePayment(Cash, 100.0)
if err != nil {
    log.Fatal(err)
}
err = payment.Pay(20.0) // 扣款 20，余额变为 80

// 创建信用支付对象（只需改变 Kind）
payment, err = GeneratePayment(Credit, 500.0)
err = payment.Pay(100.0)

// 不支持的类型返回 error
payment, err = GeneratePayment(3, 100.0)
// err: "Payment do not support this"
// payment: nil
```

---

## 扩展：添加新的支付方式

只需新增 `ConcreteProduct` 并在工厂函数中添加一个 `case`，**无需修改调用方代码**：

```go
// 1. 定义新的类型常量
const (
    Cash   Kind = 1 << iota
    Credit
    Alipay // 新增
)

// 2. 实现新的具体产品
type AlipayPay struct {
    Balance float32
}

func (a *AlipayPay) Pay(money float32) error {
    if a.Balance < money {
        return errors.New("alipay balance not enough")
    }
    a.Balance -= money
    return nil
}

// 3. 在工厂函数中添加 case
func GeneratePayment(k Kind, balance float32) (Payment, error) {
    switch k {
    case Cash:
        return &CashPay{balance}, nil
    case Credit:
        return &CreditPay{balance}, nil
    case Alipay:
        return &AlipayPay{balance}, nil // 新增
    default:
        return nil, errors.New("Payment do not support this")
    }
}
```

---

## 优缺点

### 优点

- **解耦**：调用方只依赖接口，不依赖具体实现
- **开闭原则**：新增产品类型只需扩展，无需修改已有代码
- **集中管理**：对象创建逻辑统一在工厂函数中，便于维护
- **类型安全**：通过接口约束，确保所有产品行为一致

### 缺点

- 每新增一种产品类型，都需要修改工厂函数中的 `switch` 分支
- 若产品种类过多，工厂函数会逐渐膨胀（可考虑升级为**抽象工厂**）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Factory Method** | 通过类型参数在运行时决定创建哪种对象，单一产品族 |
| **Abstract Factory** | 创建一组相关的对象（产品族），强调产品间的兼容性 |
| **Builder** | 分步骤构建一个复杂对象，关注构建过程 |
| **Prototype** | 通过克隆现有实例创建新对象 |

---

## 运行测试

```bash
cd 02-factory-method-patterns
go test -v ./...
```
