package bridge

import "testing"

func TestCpu_Run(t *testing.T) {
	cpu := &Cpu{}
	apple := new(Apple)
	apple.SetShape(cpu)
	apple.Print()
}

func TestStorage_Run(t *testing.T) {
	storage := &Storage{}
	hw := new(HuaWei)
	hw.SetShape(storage)
	hw.Print()
}
