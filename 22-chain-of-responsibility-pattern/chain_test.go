package chain

import "testing"

func Test_HandleEvent_In_Chain(t *testing.T) {
	oba := &ObjectA{Level: 1, Name: "A"}
	obb := &ObjectB{Level: 2, Name: "B"}
	oba.SetNext(obb)

	event := Event{Name: "check2", Level: 2}
	oba.HandleEvent(event)

	event = Event{Name: "checkoutrange", Level: 3}
	oba.HandleEvent(event)
}
