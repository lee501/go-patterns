package abstractfactory

import (
	"fmt"
)

type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

// 具体的产品
type ConcreteProduct struct {
	Name string
}

func (conproduct *ConcreteProduct) Describe() {
	fmt.Println(conproduct.Name)
}

// 具体工厂
type ConCreteFactory struct{}

func (confactory *ConCreteFactory) CreateProduct() Product {
	return &ConcreteProduct{Name: "KG"}
}
