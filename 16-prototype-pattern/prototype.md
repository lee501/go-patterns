# Prototype Pattern（原型模式）

## 概述

原型模式是一种**创建型**设计模式，通过**复制（克隆）已有实例**来创建新对象，而不是通过构造函数重新创建。

> Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.

---

## 适用场景

- 对象创建**代价较高**（复杂初始化、数据库查询等），通过克隆已有对象更高效
- 需要创建与某个对象**相似但略有不同**的新对象
- 系统中对象的具体类型在运行时才能确定，需要动态创建
- 希望避免构造函数的复杂参数，通过**复制+修改**的方式创建新对象

---

## 结构

```
┌────────────────────────────┐
│         Example            │
│   Description  string      │
├────────────────────────────┤
│   Clone() *Example         │  ← 克隆方法
│   New(des string) *Example │  ← 工厂方法
└────────────────────────────┘
         │
         │ Clone()
         ▼
┌────────────────────────────┐
│    新的 *Example 实例       │
│   （独立的内存副本）         │
└────────────────────────────┘
```

### 核心要素

| 要素 | 说明 |
|------|------|
| **Clone()** | 核心方法，通过值拷贝创建一个独立的新实例 |
| **res := *e** | 对结构体解引用，得到值副本，与原对象独立 |
| **return &res** | 返回副本的指针，新对象与原对象内存独立 |

---

## Go 实现

```go
type Example struct {
    Description string
}

// Clone 通过值拷贝创建独立副本
func (e *Example) Clone() *Example {
    res := *e     // 解引用，获得值副本
    return &res   // 返回副本的指针
}

func New(des string) *Example {
    return &Example{Description: des}
}
```

### 使用示例

```go
// 创建原型
original := New("原始对象")

// 克隆出新对象
clone := original.Clone()

// 修改克隆对象，不影响原对象
clone.Description = "克隆对象"

fmt.Println(original.Description) // 原始对象
fmt.Println(clone.Description)    // 克隆对象

// 验证两者内存独立
fmt.Println(original == clone)    // false（不同指针）
fmt.Println(*original == *clone)  // false（内容也不同了）
```

---

## 深拷贝 vs 浅拷贝

Go 中 `res := *e` 执行的是**浅拷贝**，对于包含指针、slice、map 的结构体，需要手动实现深拷贝：

```go
type ComplexExample struct {
    Description string
    Tags        []string          // slice：浅拷贝后共享底层数组
    Meta        map[string]string // map：浅拷贝后共享同一 map
}

// 浅拷贝（不安全，修改 slice/map 会影响原对象）
func (e *ComplexExample) ShallowClone() *ComplexExample {
    res := *e
    return &res
}

// 深拷贝（安全）
func (e *ComplexExample) DeepClone() *ComplexExample {
    newTags := make([]string, len(e.Tags))
    copy(newTags, e.Tags)

    newMeta := make(map[string]string, len(e.Meta))
    for k, v := range e.Meta {
        newMeta[k] = v
    }

    return &ComplexExample{
        Description: e.Description,
        Tags:        newTags,
        Meta:        newMeta,
    }
}
```

---

## 优缺点

### 优点

- **高效**：直接复制内存，比重新执行初始化逻辑更快
- **灵活**：可以在克隆后按需修改部分字段，快速创建相似对象
- **解耦**：客户端无需知道被克隆对象的具体类型

### 缺点

- 包含循环引用的对象**深度克隆**较为复杂
- 若对象包含私有字段，克隆可能无法访问这些字段
- 浅拷贝容易引发"共享状态"的 bug，需要特别注意

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Prototype** | 通过克隆已有对象创建新对象，避免重复初始化 |
| **Factory Method** | 通过工厂函数创建全新对象 |
| **Builder** | 分步骤构建复杂对象 |
| **Memento** | 保存对象状态快照用于恢复，而非创建新对象 |

---

## 运行测试

```bash
cd 16-prototype-pattern
go test -v ./...
```
