# Proxy Pattern（代理模式）

## 概述

代理模式是一种**结构型**设计模式，为目标对象提供一个**代理对象**，由代理对象控制对目标对象的访问。代理对象与真实对象实现相同的接口，调用方无感知地与代理交互，代理在转发请求前后可插入额外逻辑（权限校验、缓存、日志等）。

> Provide a surrogate or placeholder for another object to control access to it.

---

## 适用场景

- **访问控制**：根据条件决定是否允许调用真实对象（权限拦截）
- **延迟初始化（虚拟代理）**：真实对象创建代价高，延迟到第一次使用时才创建
- **缓存代理**：缓存真实对象的计算结果，避免重复计算
- **日志/监控代理**：在调用前后记录日志、采集指标
- **远程代理**：为远程对象提供本地代理，屏蔽网络通信细节

---

## 结构

```
┌─────────────┐        ┌──────────────────┐
│   Client    │───────>│   IObject (接口)  │
└─────────────┘        ├──────────────────┤
                       │  ObjDo(action)   │
                       └──────────────────┘
                                ▲
                   ┌────────────┴────────────┐
                   │                         │
           ┌──────────────┐        ┌──────────────────┐
           │    Object    │        │   ProxyObject    │
           │（真实对象）   │        │  object *Object  │
           │  ObjDo()     │        │  ObjDo()         │
           └──────────────┘        │  （拦截 + 转发）   │
                                   └──────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **IObject（接口）** | 定义真实对象和代理对象共同实现的接口 `ObjDo(action string)` |
| **Object（真实对象）** | 实际执行业务逻辑的对象 |
| **ProxyObject（代理对象）** | 持有 `*Object` 引用，实现相同接口，在调用前后插入拦截逻辑 |

---

## Go 实现

### 公共接口

```go
type IObject interface {
    ObjDo(action string)
}
```

### 真实对象

```go
type Object struct {
    action string
}

func (obj *Object) ObjDo(action string) {
    fmt.Printf("I can %s", action)
}
```

### 代理对象

```go
type ProxyObject struct {
    object *Object
}

func (p *ProxyObject) ObjDo(action string) {
    // 延迟初始化：真实对象未创建时才创建
    if p.object == nil {
        p.object = new(Object)
    }
    // 访问控制：只允许 "run" 操作
    if action == "run" {
        p.object.ObjDo(action)
    }
}
```

### 使用示例

```go
// 通过代理访问，客户端无需感知代理的存在
var obj IObject = &ProxyObject{}

obj.ObjDo("run")   // 允许 → 输出: I can run
obj.ObjDo("fly")   // 被拦截 → 无输出

// 也可以预先注入真实对象
realObj := &Object{action: "run"}
proxy := &ProxyObject{object: realObj}
proxy.ObjDo("run") // 输出: I can run
```

---

## 常见代理类型

| 类型 | 说明 | 本示例体现 |
|------|------|-----------|
| **保护代理** | 控制对真实对象的访问权限 | `action == "run"` 才允许调用 |
| **虚拟代理** | 延迟初始化昂贵对象 | `p.object == nil` 时才创建 |
| **缓存代理** | 缓存结果，减少重复计算 | 可在 `ObjDo` 中添加缓存逻辑 |
| **远程代理** | 封装网络通信，本地调用远程对象 | gRPC Stub 即为远程代理 |

---

## 扩展：日志代理

```go
type LogProxyObject struct {
    object *Object
}

func (p *LogProxyObject) ObjDo(action string) {
    log.Printf("[before] action: %s", action)
    p.object.ObjDo(action)
    log.Printf("[after]  action: %s done", action)
}
```

---

## 优缺点

### 优点

- **开闭原则**：无需修改真实对象即可添加控制逻辑
- **单一职责**：将访问控制、缓存、日志等横切关注点从业务逻辑中分离
- **透明性**：调用方通过接口访问，无需区分真实对象和代理对象

### 缺点

- 引入额外的代理层，增加请求处理时延
- 若代理类过多，会导致类的数量膨胀

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Proxy** | 控制对原对象的访问，接口保持一致 |
| **Decorator** | 动态为对象添加职责，可叠加多层，关注功能增强 |
| **Adapter** | 转换接口，解决不兼容问题 |
| **Facade** | 为复杂子系统提供简化的统一入口 |

---

## 运行测试

```bash
cd 06-proxy-pattern
go test -v ./...
```
