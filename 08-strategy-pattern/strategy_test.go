package strategy

import (
	"reflect"
	"testing"
	"fmt"
)

func TestAddition_Apply(t *testing.T) {
	add := new(Addition)
	operator := CreateOpration(add)
	value := operator.Operate(1,2)
	if value != 3 {
		t.Error("add algorithm error, expected result is 3")
	}
}

func TestMultiplication_Apply(t *testing.T) {
	multi := new(Multiplication)
	operator := CreateOpration(multi)
	value := operator.Operate(2, 3)
	if value != 6 {
		t.Error("multi algorithm error, expected result is 6")
	}
}

func TestCreateOpration(t *testing.T) {
	add := new(Addition)
	oper :=  CreateOpration(add)
	result  := reflect.TypeOf(oper.operator).Elem().Name()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(add).Elem().Name())
	if reflect.TypeOf(add).Elem().Name() != result {
		t.Error("create error， expected result is Addition")
	}
	a := Addition{}
	//指针和值 reflect需要使用Elem()获取指针的封装
	fmt.Println(reflect.TypeOf(a).Name())
}
