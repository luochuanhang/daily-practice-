package main

import (
	"fmt"
)

type person struct {
	age int
}

func (p *person) str() int {
	return p.age
}

func add(a int) int {
	i := 10
	func() {
		i += 1
	}()
	return i
}

func main() {
	fmt.Println(add(10))

}
