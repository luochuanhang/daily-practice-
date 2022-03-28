package main

import "fmt"

//go支持递归
//递归就是自己调用自己
func a(n int) int {
	if n == 0 {
		return 1
	}
	return n * a(n-1)
}

func main() {
	fmt.Println(a(5))
	var fib func(n int) int
	//闭包也可以递归，需要在定义闭包之前用类型化
	//的var显示声明闭包
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
