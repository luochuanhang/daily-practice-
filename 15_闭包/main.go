package main

import "fmt"
//go支持匿名函数，并能用其构造闭包
//a函数返回一个在其函数体内定义的匿名函数
//返回的函数使闭包的方式隐藏变量i，返回
//函数隐藏变量i形成闭包
func a() func() int {
	//闭包就是内部的匿名函数引用外部的变量
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	next := a()
	//没调用一次返回的值就加1
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
