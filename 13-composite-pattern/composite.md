# Composite Pattern（组合模式）

## 概述

组合模式是一种**结构型**设计模式，将对象组合成**树形结构**来表示"部分-整体"的层次关系，使客户端对单个对象和组合对象的使用具有**一致性**。

> Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

---

## 适用场景

- 需要表示对象的**树形层次结构**（目录/文件、菜单/菜单项、组织架构）
- 希望客户端**统一处理**叶子节点和容器节点（不需要区分单个元素和组合对象）
- 需要对整个层次结构进行**递归操作**（如计算总价、打印树形结构）

---

## 结构

```
              ┌────────────────────┐
              │  MenuComponent     │  ← 统一接口
              │  Price() float32   │
              │  Print()           │
              └─────────┬──────────┘
                        │
          ┌─────────────┴─────────────┐
          ▼                           ▼
  ┌──────────────┐           ┌──────────────────┐
  │   MenuItem   │           │      Menu        │
  │  (叶子节点)   │           │   (容器节点)      │
  │  price       │           │  child[]         │
  │  Print()     │           │  Add/Remove/Find │
  └──────────────┘           │  Price() → 求和  │
                             │  Print() → 递归  │
                             └──────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **MenuComponent（接口）** | 统一的叶子节点和容器节点接口，声明 `Price()` 和 `Print()` |
| **MenuItem（叶子节点）** | 具体菜单项，不包含子节点，有具体价格 |
| **Menu（容器节点）** | 包含 `[]MenuComponent` 子节点列表，价格为子节点之和 |
| **MenuDesc（共享属性）** | 抽离 `name` 和 `description` 公共字段，通过匿名组合复用 |
| **MenuGroup** | 管理子节点的增删查操作 |

---

## Go 实现

### 共享属性（抽离公共字段）

```go
type MenuDesc struct {
    name        string
    description string
}

func (desc *MenuDesc) Name() string        { return desc.name }
func (desc *MenuDesc) Description() string { return desc.description }
```

### 统一接口

```go
type MenuComponent interface {
    Price() float32
    Print()
}
```

### 叶子节点：MenuItem

```go
type MenuItem struct {
    MenuDesc        // 匿名组合共享属性
    price float32
}

func NewMenuItem(name, description string, price float32) *MenuItem {
    return &MenuItem{
        MenuDesc: MenuDesc{name: name, description: description},
        price:    price,
    }
}

func (item *MenuItem) Price() float32 { return item.price }

func (item *MenuItem) Print() {
    fmt.Printf("    %s, %.2f\n", item.name, item.price)
    fmt.Printf("    -- %s\n", item.description)
}
```

### 容器节点：Menu

```go
type MenuGroup struct {
    child []MenuComponent
}

func (group *MenuGroup) Add(component MenuComponent)    { group.child = append(group.child, component) }
func (group *MenuGroup) Remove(id int)                  { group.child = append(group.child[:id], group.child[id+1:]...) }
func (group *MenuGroup) Find(id int) MenuComponent      { return group.child[id] }

type Menu struct {
    MenuDesc
    MenuGroup
}

func NewMenu(name, description string) *Menu {
    return &Menu{MenuDesc: MenuDesc{name: name, description: description}}
}

// 递归求和
func (m *Menu) Price() (price float32) {
    for _, v := range m.child {
        price += v.Price()
    }
    return
}

// 递归打印
func (m *Menu) Print() {
    fmt.Printf("%s, %s, ¥%.2f\n", m.name, m.description, m.Price())
    fmt.Println("------------------------")
    for _, v := range m.child {
        v.Print()
    }
}
```

### 使用示例

```go
// 构建菜单树
mainMenu := NewMenu("今日菜单", "特色菜")

fastFood := NewMenu("快餐", "经济实惠")
fastFood.Add(NewMenuItem("汉堡", "牛肉汉堡", 15.0))
fastFood.Add(NewMenuItem("薯条", "黄金薯条", 8.0))

dessert := NewMenu("甜品", "甜蜜选择")
dessert.Add(NewMenuItem("冰淇淋", "香草冰淇淋", 12.0))

mainMenu.Add(fastFood)
mainMenu.Add(dessert)

// 统一调用 Print()，自动递归打印整棵树
mainMenu.Print()
// 输出:
// 今日菜单, 特色菜, ¥35.00
// ------------------------
// 快餐, 经济实惠, ¥23.00
// ------------------------
//     汉堡, 15.00
//     牛肉汉堡
//     薯条, 8.00
//     黄金薯条
// 甜品, 甜蜜选择, ¥12.00
// ...
```

---

## 优缺点

### 优点

- **统一处理**：客户端无需区分叶子节点和容器节点
- **递归操作**：天然支持对整棵树的递归遍历（求和、打印、查找）
- **开闭原则**：新增叶子或容器类型无需修改客户端代码

### 缺点

- 设计过于泛化，限制了类型安全（叶子节点和容器节点使用同一接口）
- 叶子节点被迫实现容器接口中不适用的方法（如 Add/Remove）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Composite** | 树形结构，统一处理单个对象和组合对象 |
| **Decorator** | 为单个对象动态添加职责，不构建树形结构 |
| **Iterator** | 遍历集合元素，不关心层次结构 |
| **Visitor** | 在不改变结构的前提下，对结构中的元素执行新操作 |

---

## 运行测试

```bash
cd 13-composite-pattern
go test -v ./...
```
