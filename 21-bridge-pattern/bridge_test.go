package bridge

import "testing"

func TestCpu_Run(t *testing.T) {
	cpu := &Cpu{}
	phone := new(Phone)
	phone.SetShape(cpu)
	phone.Print()
}

func TestStorage_Run(t *testing.T) {
	storage := &Storage{}
	phone := new(Phone)
	phone.SetShape(storage)
	phone.Print()
}
