package main

import (
	"fmt"
	"math"
)
//const可以出现在任何var语句可以出现的地方
const s string = "constant"

func main() {
	fmt.Println(s)
	//const用于声明一个常量
	const a = 10
	//常数表达式可以执行任意精度的运算
	const d = 3e20 / 50000
	fmt.Println(d)
	//数值型常量没有确定的类型，直到被给定某个类型
	fmt.Println(int64(d))
	//一个数值可以根据上下文的需要，自动确定类型
	fmt.Println(math.Sin(d))

}
