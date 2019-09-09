package abstractfactory

import (
	"fmt"
)

/*
	设计思想
		1. 抽象工厂接口
		2. 抽象产品接口
		3. 具体的工厂和产品struct
		4. 使用具体的工厂来创建产品，并返回接口类型值
*/
type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

//具体的产品
type ConcreteProduct struct {
	Name string
}

func (conproduct *ConcreteProduct) Describe() {
	fmt.Println(conproduct.Name)
}

//具体工厂
type ConCreteFactory struct {}

func (confactory *ConCreteFactory) CreateProduct() Product {
	return &ConcreteProduct{Name: "KG"}
}
