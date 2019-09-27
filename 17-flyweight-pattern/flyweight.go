package flyweight

/*
	享元模式核心是创建一个map属性的结构体
	设计思想:
		1. 创建Shape接口
		2. 创建实现接口Shape的实体struct Circle
		3. 创建ShapeFactory, 属性为Circle的map
*/
//创建shape接口
type Shape interface {
	SetRadius(radius int)
	SetColor(color string)
}
//创建circle,实现shape方法
type Circle struct {
	color string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

//创建ShapeFactory
type ShapeFactory struct {
	circleMap map[string]Shape
}

//GetCircle 对象不存在则创建
func (sh *ShapeFactory) GetCircle(color string) Shape {
	if sh.circleMap == nil {
		sh.circleMap = make(map[string]Shape)
	}
	if shape, ok := sh.circleMap[color]; ok {
		return shape
	}
	circle := new(Circle)
	circle.SetColor(color)
	sh.circleMap[color] = circle
	return circle
}
