package main

import "fmt"
//可变参函数，在调用时可以传递任意数量的参数
func a(nums ...int) {
	fmt.Println(nums)
	for i, v := range nums {
		fmt.Printf("i:%d,v:%d", i, v)
	}

}

func main() {
	//调用方式和一般函数一样
	a(1, 2)
	a(4, 5, 6)
	a(7, 5, 4, 5, 6, 5, 4)
	//如果有一个含有多个值得slice，想把它作为参数
	//使用需要func(slice...)
	nums:=[]int{1,2,3,8}
	a(nums...)
}
