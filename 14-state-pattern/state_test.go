package state

import "testing"

func TestNewContext(t *testing.T) {
	context := NewContext(15)
	if context.HealthValue != 15 {
		t.Error("health value err, expected is 15")
	}
	if _, ok := context.State.(*NormalState); !ok {
		t.Error("state  err, expected is normal state")
	}
}

func TestContext_SetHealth(t *testing.T) {
	context := NewContext(15)
	//设置新的context health
	context.SetHealth(9)
	if context.HealthValue != 9 {
		t.Error("health value err, expected is 9")
	}
	if _, ok := context.State.(*RestrictedState); !ok {
		t.Error("state  err, expected is normal state")
	}
}
