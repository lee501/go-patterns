package proxy

import "fmt"

// 1.use proxy and object they must implement same methods
type IObject interface {
	ObjDo(action string)
}

// 2.object represents real objects which proxy will delegate data
type Object struct {
	action string
}

// 3.ObjDo implement IObject interface and handle all logic
func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can %s", action)
}

// ProxyObject represent proxy object with intercepts actions
type ProxyObject struct {
	object *Object
}

// 拦截作用
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "run" {
		p.object.ObjDo(action)
	}
}
