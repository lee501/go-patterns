package interpreter

import (
	"fmt"
	"testing"
)

func TestEqual_Interpret(t *testing.T) {
	left := Context{"need"}
	right := Context{"must"}
	expression := CreateExpression(Equ, left, right)
	if expression != nil && !expression.Interpret() {
		 fmt.Println(expression.Interpret())
	}
}

func TestContain_Interpret(t *testing.T) {
	left := Context{"need"}
	right := Context{"n"}
	expression := CreateExpression(Cont, left, right)
	if expression != nil && expression.Interpret() {
		fmt.Println(expression.Interpret())
	}
}