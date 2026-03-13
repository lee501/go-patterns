# Builder Pattern（建造者模式）

## 概述

建造者模式是一种**创建型**设计模式，将一个复杂对象的**构建过程**与其**表示**分离，使得同样的构建过程可以创建不同的表示。

> Separate the construction of a complex object from its representation so that the same construction process can create different representations.

---

## 适用场景

- 需要创建的对象**结构复杂**，包含多个组成部分
- 对象的构建步骤固定，但每个步骤的**具体实现可以不同**
- 希望通过**链式调用**来构造对象，提升代码可读性
- 不同配置的同类对象需要用同一套构建流程生成

---

## 结构

```
┌─────────────┐        ┌──────────────────┐
│  Director   │───────>│   Builder (接口)  │
│             │        ├──────────────────┤
│ Construct() │        │ SetWheels()      │
│ SetBuilder()│        │ SetSeats()       │
└─────────────┘        │ SetStructure()   │
                       │ GetVehicle()     │
                       └──────────────────┘
                              ▲
                              │ 实现
                   ┌──────────┴──────────┐
                   │        Car          │
                   │        Bike         │
                   │        Truck        │
                   └─────────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Builder（建造者接口）** | 声明所有构建步骤的抽象方法，方法返回 `Builder` 本身以支持链式调用 |
| **ConcreteBuilder（具体建造者）** | 实现 `Builder` 接口，为每个步骤提供具体实现，并持有最终产品 |
| **Director（指挥者）** | 持有 `Builder` 引用，定义构建步骤的**调用顺序**，不关心具体实现 |
| **Product（产品）** | 被构建的复杂对象 |

---

## Go 实现

### 产品（Product）

```go
type Vehicle struct {
    Wheels    Wheel
    Seats     int
    Structure string
}
```

### 建造者接口（Builder Interface）

```go
type Builder interface {
    SetWheels()    Builder
    SetSeats()     Builder
    SetStructure() Builder
    GetVehicle()   Vehicle
}
```

> 每个 `Set` 方法返回 `Builder` 本身，支持**链式调用**。

### 具体建造者（Concrete Builder）

```go
type Car struct {
    vehicle Vehicle
}

func (car *Car) SetWheels() Builder {
    car.vehicle.Wheels = 4
    return car
}

func (car *Car) SetSeats() Builder {
    car.vehicle.Seats = 4
    return car
}

func (car *Car) SetStructure() Builder {
    car.vehicle.Structure = "Car"
    return car
}

func (car *Car) GetVehicle() Vehicle {
    return car.vehicle
}
```

### 指挥者（Director）

```go
type Director struct {
    builder Builder
}

func (director *Director) SetBuilder(builder Builder) {
    director.builder = builder
}

// Construct 定义构建顺序（链式调用）
func (director *Director) Construct() {
    director.builder.SetWheels().SetSeats().SetStructure()
}
```

### 使用示例

```go
director := Director{}

// 构建一辆 Car
car := &Car{}
director.SetBuilder(car)
director.Construct()
vehicle := director.builder.GetVehicle()
fmt.Println(vehicle) // {4 4 Car}

// 切换为 Bike（只需替换 Builder）
bike := &Bike{}
director.SetBuilder(bike)
director.Construct()
vehicle = director.builder.GetVehicle()
fmt.Println(vehicle) // {2 1 Bike}
```

---

## 扩展：添加新的具体建造者

只需新增一个实现 `Builder` 接口的 struct，**无需修改 Director 或其他代码**：

```go
type Bike struct {
    vehicle Vehicle
}

func (b *Bike) SetWheels() Builder {
    b.vehicle.Wheels = 2
    return b
}

func (b *Bike) SetSeats() Builder {
    b.vehicle.Seats = 1
    return b
}

func (b *Bike) SetStructure() Builder {
    b.vehicle.Structure = "Bike"
    return b
}

func (b *Bike) GetVehicle() Vehicle {
    return b.vehicle
}
```

---

## 优缺点

### 优点

- **单一职责**：将复杂构建逻辑从业务代码中分离
- **开闭原则**：新增产品类型只需添加新的 `ConcreteBuilder`，无需修改现有代码
- **链式调用**：`Builder` 方法返回自身，构建步骤流畅易读
- **复用构建流程**：`Director` 可以复用同一构建过程生产不同产品

### 缺点

- 代码量增加：每种产品都需要对应的 `ConcreteBuilder`
- 若产品差异过大，`Builder` 接口会变得臃肿

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Builder** | 分步构建一个复杂对象，构建过程可复用 |
| **Abstract Factory** | 创建一系列相关对象（产品族），强调产品之间的兼容性 |
| **Factory Method** | 由子类决定实例化哪个类，侧重于单个对象的创建 |
| **Prototype** | 通过克隆现有对象来创建新对象 |

---

## 运行测试

```bash
cd 01-builder-patterns
go test -v ./...
```
