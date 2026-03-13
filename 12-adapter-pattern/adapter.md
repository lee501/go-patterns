# Adapter Pattern（适配器模式）

## 概述

适配器模式是一种**结构型**设计模式，将一个类的接口转换成客户端期望的另一个接口，使原本不兼容的接口能够协同工作。

> Convert the interface of a class into another interface that clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

---

## 适用场景

- 需要使用**已有类**，但其接口与目标接口不兼容
- 想复用第三方库，但其接口与系统规范不一致
- 系统需要统一调用多种不同来源的对象（统一接口）
- 在不修改原有类的前提下，为其提供新的接口形式

---

## 结构

```
┌─────────────┐        ┌────────────────────┐
│   Client    │──Play()│   Player (目标接口)  │
│             │───────>│   PlayMusic()       │
└─────────────┘        └──────────┬──────────┘
                                  │ 实现
                       ┌──────────┴───────────┐
                       │                      │
               ┌───────────────┐   ┌──────────────────────┐
               │  MusicPlayer  │   │  GamePlayerAdapter   │
               │（直接实现）    │   │  Game GamePlayer     │
               └───────────────┘   │  PlayMusic() →       │
                                   │  Game.PlaySound()    │
                                   └──────────────────────┘
                                              │ 组合
                                   ┌──────────┴──────────┐
                                   │    GamePlayer       │
                                   │    PlaySound()      │
                                   └─────────────────────┘
```

### 核心角色

| 角色 | 说明 |
|------|------|
| **Target（目标接口）** | 客户端期望的接口，如 `Player` |
| **Adaptee（被适配者）** | 已有的、接口不兼容的类，如 `GamePlayer` |
| **Adapter（适配器）** | 通过组合 `Adaptee`，实现 `Target` 接口，如 `GamePlayerAdapter` |
| **Client** | 只依赖 `Target` 接口，不感知适配器的存在 |

---

## Go 实现

### 目标接口（Target）

```go
type Player interface {
    PlayMusic()
}

// 统一调用入口，只依赖 Player 接口
func Play(player Player) {
    player.PlayMusic()
}
```

### 直接实现目标接口的类

```go
type MusicPlayer struct {
    Src string
}

func (music *MusicPlayer) PlayMusic() {
    fmt.Println("play music: " + music.Src)
}
```

### 被适配者（接口不兼容）

```go
type GamePlayer struct {
    Src string
}

// GamePlayer 的方法名与 Player 接口不一致
func (game *GamePlayer) PlaySound() {
    fmt.Println("play sound: " + game.Src)
}
```

### 适配器（Adapter）

```go
// 通过组合 GamePlayer，实现 Player 接口
type GamePlayerAdapter struct {
    Game GamePlayer
}

func (adapter *GamePlayerAdapter) PlayMusic() {
    adapter.Game.PlaySound() // 将 PlayMusic() 调用转发给 PlaySound()
}
```

### 使用示例

```go
// 客户端只调 Play()，不关心内部实现
music := &MusicPlayer{Src: "song.mp3"}
Play(music)
// 输出: play music: song.mp3

// 通过适配器，让 GamePlayer 也能被 Play() 使用
game := &GamePlayerAdapter{
    Game: GamePlayer{Src: "game.wav"},
}
Play(game)
// 输出: play sound: game.wav
```

---

## 两种适配器实现方式

### 对象适配器（Object Adapter）— 本示例使用

通过**组合**持有被适配者实例，灵活性更高：

```go
type GamePlayerAdapter struct {
    Game GamePlayer  // 组合被适配者
}
```

### 类适配器（Class Adapter）— Go 中通过匿名嵌入实现

通过**匿名嵌入**继承被适配者方法，再重写目标接口方法：

```go
type GamePlayerAdapter2 struct {
    GamePlayer  // 匿名嵌入
}

func (a *GamePlayerAdapter2) PlayMusic() {
    a.PlaySound() // 调用嵌入的方法
}
```

---

## 优缺点

### 优点

- **复用已有代码**：无需修改被适配者，即可复用其功能
- **开闭原则**：新增适配器无需修改已有代码
- **解耦**：客户端与被适配者完全解耦，只依赖目标接口

### 缺点

- 引入额外的适配器层，增加代码复杂度
- 过多的适配器会使系统结构混乱，不易理解

---

## 与其他模式的区别

| 模式 | 关注点 |
|------|--------|
| **Adapter** | 转换接口，使不兼容的类可以协同工作 |
| **Decorator** | 在不改变接口的前提下，动态扩展对象功能 |
| **Proxy** | 控制对原对象的访问，接口保持一致 |
| **Facade** | 为复杂子系统提供简化的统一入口 |

---

## 运行测试

```bash
cd 12-adapter-pattern
go test -v ./...
```
