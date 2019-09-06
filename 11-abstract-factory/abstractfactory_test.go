package abstractfactory

import "testing"

func TestConCreteFactory_CreateProduct(t *testing.T) {
	conFactory := &ConCreteFactory{}

	product := conFactory.CreateProduct()

	conProduct := product.(*ConcreteProduct)

	if conProduct.Name != "KG" {
		t.Error("abstract factory can not create the concreate product")
	}
}
