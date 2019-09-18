package interpreter

import (
	"fmt"
	"testing"
)

func TestEqual_Interpret(t *testing.T) {
	left := Context{"need"}
	right := Context{"must"}
	expression := CreateExpression("equal", left, right)
	if expression != nil && !expression.Interpret() {
		 fmt.Println(expression.Interpret())
	}
}

func TestContain_Interpret(t *testing.T) {
	left := Context{"need"}
	right := Context{"n"}
	expression := CreateExpression("contain", left, right)
	if expression != nil && expression.Interpret() {
		fmt.Println(expression.Interpret())
	}
}