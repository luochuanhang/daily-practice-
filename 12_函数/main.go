package main

import "fmt"

//函数是go的核心
func a (a int,b int)int{
return a+b
}
//当多个连续的参数为同样的类型时，
//可以仅声明最后一个参数类型，
//忽略之前相同的类型
func b(a,b,c int)int{
return a+b+c
}
//可以通过函数名来调用参数
func main() {
	b:=a(1,2)
	fmt.Println(b)
}