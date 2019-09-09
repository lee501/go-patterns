package visitor

import "testing"

func TestElementContainer_Add(t *testing.T) {
	container := new(ElementContainer)
	a := &ElementA{}
	b := &ElementA{}
	container.Add(a)
	container.Add(b)
	if len(container.list) != 2 {
		t.Error("count error, expected amount is 2")
	}
}

func TestElementA_Accept(t *testing.T) {
	elementA := &ElementA{}
	visitorA := &ConcreteVisitorA{Name: "lee"}
	visitorb := &ConcreteVisitorB{Name: "anne"}
	elementA.Accept(visitorA)
	elementA.Accept(visitorb)
}
