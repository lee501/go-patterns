package composite

import (
	"fmt"
	"testing"
)

func TestNewMenuItem(t *testing.T) {
	menuitem := NewMenuItem("奥尔良套餐", "赠送薯条", 12.5)
	//可以直接调用匿名组合的方法
	fmt.Println(menuitem.Name())
}

func TestMenuGroup_Add(t *testing.T) {
	menu := NewMenu("超值套餐", "奥尔良和辣宝")
	menuitem := NewMenuItem("奥尔良套餐", "赠送薯条", 12.5)
	menuitem1 := NewMenuItem("辣宝套餐", "赠送薯条", 12.5)
	menu.Add(menuitem)
	menu.Add(menuitem1)
	if len(menu.child) != 2 {
		t.Error("数量错误，期望为2")
	}
}

func TestMenuItem_Price(t *testing.T) {
	menu := NewMenu("超值套餐", "奥尔良和辣宝")
	menuitem := NewMenuItem("奥尔良套餐", "赠送薯条", 12.5)
	menuitem1 := NewMenuItem("辣宝套餐", "赠送薯条", 12.5)
	menu.Add(menuitem)
	menu.Add(menuitem1)
	if menu.Price() != 25 {
		t.Error("价格错误，期望为25")
	}
}

func TestMenu_Print(t *testing.T) {
	menu := NewMenu("超值套餐", "奥尔良和辣宝")
	menuitem := NewMenuItem("奥尔良套餐", "赠送薯条", 12.5)
	menuitem1 := NewMenuItem("辣宝套餐", "赠送薯条", 12.5)
	menu.Add(menuitem)
	menu.Add(menuitem1)
	menu.Print()
}
