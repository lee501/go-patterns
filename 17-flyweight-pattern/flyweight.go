package flyweight

// 创建shape接口
type Shape interface {
	SetRadius(radius int)
	SetColor(color string)
}

// 创建circle,实现shape方法
type Circle struct {
	color  string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

// 创建ShapeFactory
type ShapeFactory struct {
	circleMap map[string]Shape
}

// GetCircle 对象不存在则创建
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
