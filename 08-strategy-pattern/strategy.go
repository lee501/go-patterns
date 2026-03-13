package strategy

type Operator interface {
	Apply(int, int) int
}

// 包装器
type Operation struct {
	operator Operator
}

func (op *Operation) Operate(left, right int) int {
	return op.operator.Apply(left, right)
}

// Addition struct inherit Operator
type Addition struct{}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

// Multiplication struct
type Multiplication struct{}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}

func CreateOpration(operator Operator) Operation {
	return Operation{operator}
}
