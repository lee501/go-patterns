# Facade Pattern（外观模式）

## 概述

外观模式是一种**结构型**设计模式，为复杂子系统提供一个**简化的统一入口**，使客户端无需了解子系统内部的复杂性，只通过外观对象与子系统交互。

> Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

---

## 适用场景

- 子系统较复杂，希望为客户端提供**简洁的访问接口**
- 需要**分层设计**，通过外观层隔离客户端与内部实现
- 希望减少客户端与多个子系统之间的**直接依赖**
- 微服务架构中：聚合多个服务的调用，对外提供统一 API（BFF 层）

---

## 结构

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ 只调用 Facade
       ▼
┌─────────────────────────────┐
│          Facade             │  ← 外观层，统一入口
│   music Music               │
│   count Count               │
│   video Video               │
│   PrintServerInfo()         │
└──────┬──────────────────────┘
       │ 内部调用各子系统
  ┌────┴────────────┐
  │                 │
┌────────┐  ┌───────────┐  ┌─────────┐
│ Music  │  │   Video   │  │  Count  │
│GetMusic│  │GetVideoId │  │GetComment│
└────────┘  └───────────┘  └─────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Facade** | 外观类，持有各子系统的引用，提供统一访问接口 `PrintServerInfo()` |
| **Music** | 音乐子系统，提供 `GetMusic()` |
| **Video** | 视频子系统，提供 `GetVideoId()` |
| **Count** | 计数子系统，提供 `GetComment()` 等统计方法 |
| **Client** | 只依赖 `Facade`，不直接调用子系统 |

---

## Go 实现

### 子系统

```go
type Music struct {
    Name string
}
func (m *Music) GetMusic() string { return m.Name }

type Video struct {
    Id int64
}
func (v *Video) GetVideoId() int64 { return v.Id }

type Count struct {
    Comment int64
    Praise  int64
    Collect int64
}
func (c *Count) GetComment() int64 { return c.Comment }
```

### 外观层

```go
type Facade struct {
    music Music
    count Count
    video Video
}

// 统一入口：聚合调用各子系统
func (f *Facade) PrintServerInfo() {
    music := f.music.GetMusic()
    videoId := f.video.GetVideoId()
    comment := f.count.GetComment()
    fmt.Printf("Music: %s, VideoId: %d, Comment: %d\n", music, videoId, comment)
}

func NewFacade(music Music, count Count, video Video) *Facade {
    return &Facade{music: music, video: video, count: count}
}
```

### 使用示例

```go
// 初始化各子系统
music := Music{Name: "Shape of You"}
video := Video{Id: 12345}
count := Count{Comment: 1000, Praise: 5000, Collect: 200}

// 通过外观统一访问
facade := NewFacade(music, count, video)
facade.PrintServerInfo()
// 输出: Music: Shape of You, VideoId: 12345, Comment: 1000

// 客户端无需分别调用 music.GetMusic()、video.GetVideoId()、count.GetComment()
```

---

## 优缺点

### 优点

- **简化接口**：客户端无需了解子系统内部复杂性，学习成本低
- **解耦**：客户端与子系统之间通过外观层隔离，子系统内部变化不影响客户端
- **分层架构**：清晰地划分系统层次（展示层 → 外观层 → 子系统层）

### 缺点

- 外观类不符合开闭原则：若子系统新增功能，需要修改外观类
- 若使用不当，外观类可能演变成包含大量代码的"上帝类"

---

## 外观模式的实际应用

| 场景 | 描述 |
|------|------|
| **微服务 BFF** | Backend For Frontend，聚合多个微服务接口，对前端提供统一 API |
| **SDK 封装** | 将复杂的底层 API 封装为简洁的高层接口 |
| **数据库访问层** | Repository 层封装底层 ORM 的复杂查询 |
| **第三方库封装** | 对复杂的第三方库提供简化的包装接口 |

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Facade** | 为复杂子系统提供简化的统一接口 |
| **Adapter** | 转换接口，让不兼容的接口协同工作 |
| **Decorator** | 动态为对象添加职责，接口不变 |
| **Mediator** | 集中管理对象间的通信，减少直接依赖 |

---

## 运行测试

```bash
cd 19-facade-pattern
go test -v ./...
```
