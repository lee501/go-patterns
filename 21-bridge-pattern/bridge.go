package bridge

import "fmt"

// 抽象接口
type SoftWare interface {
	Run()
}

// 具体类型CPU和Storage
type Cpu struct{}

func (c *Cpu) Run() {
	fmt.Println("this is cpu run")
}

type Storage struct{}

func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

// 桥接层结构体
type Phone struct {
	software SoftWare
}

// 赋值具体software
func (s *Phone) SetSoftWare(soft SoftWare) {
	s.software = soft
}

// Apple结构体
type Apple struct {
	phone Phone
}

func (p *Apple) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *Apple) Print() {
	p.phone.software.Run()
}

// HuaWei结构体
type HuaWei struct {
	phone Phone
}

func (p *HuaWei) SetShape(soft SoftWare) {
	p.phone.SetSoftWare(soft)
}

func (p *HuaWei) Print() {
	p.phone.software.Run()
}
