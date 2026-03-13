# Interpreter Pattern（解释器模式）

## 概述

解释器模式是一种**行为型**设计模式，为某种语言定义其语法表示，并提供一个解释器来处理该语法。通常用**抽象语法树（AST）**表示句子，树中每个节点对应一种语法规则，递归调用 `Interpret()` 完成求值。

> Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.

---

## 适用场景

- 需要解析并执行**简单语言或表达式**（数学运算、布尔逻辑、正则、SQL 子集）
- 处理**配置文件 / DSL**，如解析自定义规则表达式
- 语法规则相对简单，且不会频繁变化
- 需要将表达式以**对象结构**表示，便于扩展新的操作（如求值、打印、优化）

---

## 结构

```
┌──────────────────────────────────────────────┐
│       Expression（接口）                      │
│       Interpret() int                        │
└──────────────────────────────────────────────┘
              ▲
   ┌──────────┼──────────┐
   │          │          │
┌──────────┐  ┌─────────────┐  ┌─────────────┐
│NumExpres-│  │AddExpression│  │SubExpression│
│sion      │  │left, right  │  │left, right  │
│val int   │  │Interpret()  │  │Interpret()  │
│Interpret │  │= left+right │  │= left-right │
└──────────┘  └─────────────┘  └─────────────┘

┌──────────────────────────┐
│         Parser           │  ← 解析器：将字符串构建为 AST
│   exp   []string         │
│   index int              │
│   prev  Expression       │
│   Parse(str string)      │
│   Result() Expression    │
└──────────────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Expression（接口）** | 所有表达式节点的公共接口，声明 `Interpret() int` |
| **NumExpression（终结符）** | 叶子节点，直接返回整数值，不包含子表达式 |
| **AddExpression（非终结符）** | 内部节点，持有左右子表达式，递归求和 |
| **SubExpression（非终结符）** | 内部节点，持有左右子表达式，递归求差 |
| **Parser（解析器）** | 将中缀表达式字符串解析为 AST，提供 `Result()` 入口 |

---

## Go 实现

### 表达式接口

```go
type Expression interface {
    Interpret() int
}
```

### 终结符表达式（叶子节点）

```go
type NumExpression struct {
    val int
}

func (n *NumExpression) Interpret() int {
    return n.val
}
```

### 非终结符表达式（内部节点）

```go
type AddExpression struct {
    left, right Expression
}

func (n *AddExpression) Interpret() int {
    return n.left.Interpret() + n.right.Interpret()
}

type SubExpression struct {
    left, right Expression
}

func (n *SubExpression) Interpret() int {
    return n.left.Interpret() - n.right.Interpret()
}
```

### 解析器（构建 AST）

```go
type Parser struct {
    exp   []string
    index int
    prev  Expression
}

func (p *Parser) Parse(str string) {
    p.exp = strings.Split(str, " ")
    for {
        if p.index >= len(p.exp) {
            return
        }
        switch p.exp[p.index] {
        case "+":
            p.prev = p.newAddExpression()
        case "-":
            p.prev = p.newSubExpression()
        default:
            p.prev = p.newNumExpression()
        }
    }
}

func (p *Parser) newNumExpression() Expression {
    v, _ := strconv.Atoi(p.exp[p.index])
    p.index++
    return &NumExpression{val: v}
}

func (p *Parser) newAddExpression() Expression {
    p.index++
    return &AddExpression{left: p.prev, right: p.newNumExpression()}
}

func (p *Parser) newSubExpression() Expression {
    p.index++
    return &SubExpression{left: p.prev, right: p.newNumExpression()}
}

func (p *Parser) Result() Expression {
    return p.prev
}
```

### 使用示例

```go
p := &Parser{}
p.Parse("1 + 3 + 3 + 3 - 3")
result := p.Result().Interpret()
fmt.Println(result) // 7
```

---

## AST 构建过程图解

以表达式 `"1 + 3 - 3"` 为例：

```
解析 "1"  → NumExpression(1)
解析 "+"  → AddExpression( left=Num(1), right=Num(3) )   → 值: 4
解析 "-"  → SubExpression( left=Add(1+3), right=Num(3) ) → 值: 1

AST 结构:
        SubExpression
        /           \
  AddExpression    Num(3)
  /         \
Num(1)     Num(3)

求值: Sub.Interpret()
    = Add.Interpret() - 3
    = (1 + 3) - 3
    = 1
```

---

## 扩展：添加乘法

只需新增一个 `MulExpression`，无需修改已有代码：

```go
type MulExpression struct {
    left, right Expression
}

func (n *MulExpression) Interpret() int {
    return n.left.Interpret() * n.right.Interpret()
}
```

在 `Parser.Parse()` 的 `switch` 中增加 `case "*"` 即可支持乘法。

---

## 优缺点

### 优点

- **易于扩展语法**：新增语法规则只需添加新的 Expression 实现类
- **语法即结构**：AST 的对象结构与语法规则一一对应，直观清晰
- **递归求值**：利用递归自然地处理嵌套表达式

### 缺点

- **语法复杂时不适合**：每条规则一个类，规则过多时类爆炸
- **性能**：大量递归调用在表达式很深时可能有性能开销
- **不适合复杂语言**：复杂语言建议使用专业解析器生成工具（如 ANTLR、yacc）

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Interpreter** | 定义语言语法并实现解释器，AST 递归求值 |
| **Visitor** | 在不改变对象结构的情况下，对 AST 节点执行新操作（如代码生成） |
| **Composite** | 组合模式常用于构建 AST 的树形结构 |
| **Strategy** | 封装可互换算法，不涉及语法解析 |

---

## 运行测试

```bash
cd 26-interpreter-pattern
go test -v ./...
```
