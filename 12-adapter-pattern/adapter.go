package adapter

import "fmt"

// 音乐播放器
type Player interface {
	PlayMusic()
}

type MusicPlayer struct {
	Src string
}

func (music *MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + music.Src)
}

// 对外接口
func Play(player Player) {
	player.PlayMusic()
}

// 在音乐播放基础上实现游戏播放
// 定义游戏结构体
type GamePlayer struct {
	Src string
}

// game的方法
func (game *GamePlayer) PlaySound() {
	fmt.Println("play sound: " + game.Src)
}

// 这里要实现调用play方法的时候，实现GamePlayer的播放
// 通过组合的方式，声明一个Game的适配器
type GamePlayerAdapter struct {
	Game GamePlayer
}

// 继承Player interface, 调用GamePlayer的方法
func (game *GamePlayerAdapter) PlayMusic() {
	game.Game.PlaySound()
}
