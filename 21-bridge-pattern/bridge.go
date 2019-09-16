package bridge

import "fmt"

/*
	分离抽象部分和实现部分
	设计思想：
		1. 一个桥接接口Interface以及实现该接口的具体struct
		2. 一个属性为桥接接口的struct Shape（抽象部分， 1，3为具体部分）
		3. 与Shape组合模式的具体struct
*/
//创建桥接接口
type SoftWare interface {
	Run()
}

//创建实现run方法的结构体
type Cpu struct {

}
func (c *Cpu) Run() {
	fmt.Println("this is cpu run")
}

type Storage struct {

}
func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

//创建Shape struct（抽象部分）
type Shape struct {
	software SoftWare
}
//赋值实现桥接接口的具体结构体
func (s *Shape) SetSoftWare(soft SoftWare) {
	s.software = soft
}

//组合模式创建具体的struct
type Phone struct {
	shape Shape
}

func (p *Phone) SetShape(soft SoftWare) {
	p.shape.SetSoftWare(soft)
}

func (p *Phone) Print() {
	p.shape.software.Run()
}
