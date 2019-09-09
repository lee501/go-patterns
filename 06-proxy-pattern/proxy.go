package proxy

import "fmt"

/*
	proxy pattern provide object control access to another object, intercepting all calls
	//设计思想
		1. 代理inteface
		2. 真实对象Object struct
		3. 代理对象ProxyObject struct，属性为Object, 拦截所有的action
		4. 方法ObjDo处理所有的逻辑
*/

//1.use proxy and object they must implement same methods
type IObject interface {
	ObjDo(action string)
}
//2.object represents real objects which proxy will delegate data
type Object struct {
	action string
}
//3.ObjDo implement IObject interface and handle all logic
func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can %s", action)
}
//ProxyObject represent proxy object with intercepts actions
type ProxyObject struct {
	object *Object
}
//拦截作用
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "run" {
		p.object.ObjDo(action)
	}
}
