package main

import (
	"fmt"
	"math"
)
//这是一个几何接口
type geometry interface {
	area() float64
	perim() float64
}
//将rect实现这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
//要在go中实现一个接口，只需要实现接口中的所有方法
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
//如果一个变量是接口类型，那么我们可以调用这个被
//命名的接口中的方法
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}
	//结构体类型cirle和rect都实现了geometry接口
	//所以我们可以使用他们的实例作为measure的参数
	measure(r)
	measure(c)
}
