package main

import "fmt"

type rect struct {
	width, height int
}
//area有一个接收者 rect
func (r *rect) area() int {
	return r.width * r.height
}
//可以为值类型或者指针类型的接收者定义方法
func (r rect) perim() int {
	return r.width + 2*r.height
}

func main() {
	r := rect{width: 5, height: 9}
	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())
	//go自动处理方法调用时的值和指针之间的转化
	//可以使用指针来调用方法来避免在方法调用时产生一个拷贝
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())

}
