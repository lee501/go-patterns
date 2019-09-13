package facade

import "testing"

func TestNewFacade(t *testing.T) {
	music := Music{"love"}
	video := Video{1}
	count := Count{12, 30, 5}
	facade := NewFacade(music, count, video)
	facade.PrintServerInfo()
}
