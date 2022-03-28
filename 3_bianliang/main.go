package main

import "fmt"

func main() {
	//声明一个变量
	var a = "linlin"
	fmt.Println(a)
	//声明多个变量
	var c,d int=1,2
	fmt.Println(c,d)
	//go自动推断已经有初始值得变量
	var b=true
	fmt.Println(b)
	//声明后没有赋值的变量，变量会初始化为0值
	var e float64
	fmt.Println(e)
	//:=语法是声明并初始化变量的简写
	f:="short"
	fmt.Println(f)
}
