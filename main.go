package main

import "fmt"

type P struct {
	Name string
}
func main() {
	q := make(chan int, 3)
	q <- 1
	q <- 2
	q <- 3
	close(q)
	for v := range q {
		fmt.Println(v)
	}
	fmt.Println("jie")

	var m struct{}
	m = P{}
}
