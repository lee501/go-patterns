package factory

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	k Kind = 1
	m Kind = 2
	n Kind = 3
	balance float32 = 100.00
)

func TestGeneratePayment(t *testing.T) {
	payment, _ := GeneratePayment(k, balance)
	fmt.Println(reflect.TypeOf(payment).Elem().String())
	if reflect.TypeOf(payment).Elem().String() != "factory.CashPay" {
		t.Error("factory method generate error")
	}

	payment, _ = GeneratePayment(n, balance)
	if payment != nil {
		t.Error("factory method params has error")
	}
}

func TestCashPay_Pay(t *testing.T) {
	payment, _ := GeneratePayment(1, balance)
	payment.Pay(20)
	//cash := reflect.New(reflect.TypeOf(payment).Elem()).Interface().(*CashPay)relect新的对象
	cash := payment.(*CashPay)
	fmt.Println(reflect.TypeOf(cash))
	if cash.Balance != float32(80) {
		t.Error("结算错误")
	}
}
