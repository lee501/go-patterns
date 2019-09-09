package composite

import "fmt"

/*
	设计思想：
		struct不依赖interface
		1. 包含角色：
			1). 共同的接口MenuComponent，为root和leaf结构体共有的方法
			2). root结构体(包含leaf列表)和leaf结构体
			3). 将结构体中共同部分的数据抽离，使用匿名组合的方式实现2中的两类结构体
*/
//Menu示例如何使用组合设计模式：Menu和MenuItem

//抽离共同属性部分
type MenuDesc struct {
	name		string
	description string
}
func (desc *MenuDesc) Name() string {
	return desc.name
}
func (desc *MenuDesc) Description() string {
	return desc.description
}

//MenuItem组合，继承了MenuDesc的方法
type MenuItem struct {
	MenuDesc
	price 	float32
}
func NewMenuItem(name, description string, price float32) *MenuItem {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name: name,
			description: description,
		},
		price: price,
	}
}
//实现MenuItem Price方法和Print()
func (item *MenuItem) Price() float32 {
	return item.price
}
func (item *MenuItem) Print() {
	fmt.Printf("	%s, %0.2f\n", item.name, item.price)
	fmt.Printf("	-- %s\n", item.description)
}

//接下来实现Menu struct, Menu包含共同部分MenuDesc以及列表, 将列表部分分离出来
/*
	MenuGroup, 这里引入接口MenuComponent, 因为child类型是不确定的
	此外，由于接口为Menu和MenuItem的共同接口，所以包含Price和Print方法
*/
type MenuComponent interface {
	Price() float32
	Print()
}
type MenuGroup struct {
	child []MenuComponent
}
//MenuGroup需要实现Add、Remove、Find方法
func (group *MenuGroup) Add(component MenuComponent) {
	group.child = append(group.child, component)
}
func (group *MenuGroup) Remove(id int) {
	group.child = append(group.child[:id], group.child[id+1:]...)
}
func (group *MenuGroup) Find(id int) MenuComponent {
	return group.child[id]
}

//Menu结构体
type Menu struct {
	MenuDesc
	MenuGroup
}

//简单工厂
func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name: name,
			description: description,
		},
	}
}
//实现Price和Print方法
func (m *Menu) Price() (price float32) {
	for _, v := range m.child {
		price += v.Price()
	}
	return price
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ¥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.child {
		v.Print()
	}
	fmt.Println("结束")
}
