package strategy

/*
	设计思想：
		1. 一个Operator接口
		2. 属性为Operator接口的struct
		3. 实现结构体的自身的方法
		4. 封装算法到实现接口Operator具体的struct中
*/

type Operator interface {
	Apply(int, int) int
}

//包装器
type Operation struct {
	operator Operator
}

func (op *Operation) Operate(left, right int) int {
	return op.operator.Apply(left, right)
}

//Addition struct inherit Operator
type Addition struct {}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

//Multiplication struct
type Multiplication struct {}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}

func CreateOpration(operator Operator) Operation {
	return Operation{operator}
}
