package memento

import (
	"fmt"
	"testing"
)

func TestCaretaker_RecoverOriginator(t *testing.T) {
	originator := new(Originator)
	originator.SetState("init state")
	//管理者创建备忘录
	caretaker := &Caretaker{}
	memento := caretaker.CreateMemento(*originator)
	fmt.Println(memento.GetState())
	if memento.GetState() != originator.state {
		t.Error("create memento error")
	}
	//更改originator状态
	originator.SetState("change state")
	fmt.Println(originator.GetState())
	if memento.GetState() == originator.state {
		t.Error("change state error")
	}
	//恢复originator状态
	*originator = caretaker.RecoverOriginator(memento)
	fmt.Println(originator.GetState())
}
