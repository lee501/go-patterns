package decorator

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
	return func(i int) int {
		log.Println("staring inner func")
		result := fn(i)
		log.Println("completr inner func")
		return result
	}
}
