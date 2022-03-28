package main

import "fmt"

func main() {
	//这是基本例子
	if 5/2 == 0 {
		fmt.Println("这是奇数")
	} else {
		fmt.Println("这是双数")
	}
	//可以不用else
	if 8%4 == 0 {
		fmt.Println("这个数能被4整除")
	}
	//在条件语句前可以有一个声明语句，在这
	//里声明的变量可以在这个语句所有的条件
	//分支中使用

	if num := 9; num < 0 {
		fmt.Println("这是负数")
	} else if num < 10 {
		fmt.Println("小于10的整数")
	} else {
		fmt.Println("很大的数")
	}
}
