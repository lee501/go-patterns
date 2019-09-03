package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	event :=  NewEvent()

	investerlee := NewInvestorObserver("lee")
	investeranne := NewInvestorObserver("anne")

	share := NewShareNotifier(20.00)
	//注册观察者
	share.Register(investerlee)
	share.Register(investeranne)

	share.Notify(event)

	share.Remove(investerlee)
	fmt.Println(len(share.oblist))
	if len(share.oblist) != 1 {
		t.Error("observer pattern error")
	}
}
