package pool

import "fmt"

/*
	根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本
	*设计思想
		1.对象结构体
		2.类型为结构体指针的channel
		3.New方法, 创建新的对象放到channel中
*/
type Object struct {
	Name string
}

type Pool chan *Object

func NewPool(count int) *Pool {
	pool := make(Pool, count)
	//循环创建对象，放入pool中
	for i := 0; i < count; i++ {
		pool <- new(Object)
	}
	return &pool
}

// Acquire 从池中获取一个对象
func (p *Pool) Acquire() *Object {
	return <-(*p)
}

// Release 将对象归还到池中
func (p *Pool) Release(obj *Object) {
	(*p) <- obj
}

func (obj *Object) Do() {
	fmt.Println(&obj)
}