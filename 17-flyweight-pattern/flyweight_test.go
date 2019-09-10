package flyweight

import (
	"fmt"
	"testing"
)

func TestShapeFactory_GetCircle(t *testing.T) {
	shapeF := new(ShapeFactory)
	shape := shapeF.GetCircle("red")
	if _, ok := shapeF.circleMap["red"]; !ok {
		t.Error("map为空， 期待为1")
	}
	circle := shape.(*Circle)
	fmt.Println(circle.color)
	if circle.color != "red" {
		t.Error("expected color is red")
	}
}
