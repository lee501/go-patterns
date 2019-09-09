package decorator

import "log"

/*
	装饰器demo
	*设计思想
		将函数作为参数，并在闭包中调用此函数
*/
type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(i int) int {
		log.Println("staring inner func")
		result := fn(i)
		log.Println("completr inner func")
		return result
	}
}
