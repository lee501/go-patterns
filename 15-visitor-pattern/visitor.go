package visitor

import "fmt"

/*	根据外部环境更改这个类所能执行的行为, 传入不同的visitor，输出不同的行为
	设计思路：
		1. Visitor interface
		2. ConcreteVisitor struct
		3. Element interface(accept Visitor)
		4. ConcreteElement struct
		5. ElementContainer（包含Element list）非必须
*/
//visitor 接口
type Visitor interface {
	Visit()
}
//具体Visitor对象
type ConcreteVisitorA struct {
	Name string
}

func (conV *ConcreteVisitorA) Visit() {
	fmt.Println("this is visitor A")
}

//Visitor B
type ConcreteVisitorB struct {
	Name string
}
func (conV *ConcreteVisitorB) Visit() {
	fmt.Println("this is visitor B")
}

//创建元素接口
type Element interface {
	Accept(visitor Visitor)
}

//元素对象
type ElementA struct {}

func (e *ElementA) Accept(visitor Visitor)  {
	visitor.Visit()
}

//Element容器
type ElementContainer struct {
	list []Element
}

//实现容器的元素的添加和移除
func (container *ElementContainer) Add(element Element)  {
	if container == nil || element == nil {
		return
	}
	container.list = append(container.list, element)
}

func (container *ElementContainer) Delete(element Element) {
	for i, val := range container.list {
		if val == element {
			container.list = append(container.list[:i], container.list[i+1:]...)
			break
		}
	}
}