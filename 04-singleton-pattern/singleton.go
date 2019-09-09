package singleton

import "sync"

/*
	单类模式严格一个类只有一个实例，并提供一个全局的访问接口
	*设计思想
		1.声明一个全局变量
		2.多线程考虑线程安全，引入sync.Once
*/
type singleton map[string]string

var (
	 once 		sync.Once
	 instance 	singleton
)
func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
