package adapter

import "testing"

func TestPlay(t *testing.T) {
	music := &MusicPlayer{Src: "music.mp3"}
	Play(music)

	//使用Play来播放GamePlayer
	game := &GamePlayer{Src: "game.mp4"}
	adapter := &GamePlayerAdapter{Game: *game}
	Play(adapter)
}
