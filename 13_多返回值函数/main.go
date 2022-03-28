package main

import "fmt"
//返回两个参数
func val() (int, int) {
	return 3, 7
}
func main() {
	//我们可以通过多赋值操作来使用不同的返回值
	a, b := val()
	fmt.Println(a,b)
	//如果只需要一部分，可以使用空白标识符
	_,c:=val()
	fmt.Println(c)
}
