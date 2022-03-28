package main

import "fmt"

func main() {
	//与数组不同，slice的类型由它所包含的元素类型决定
	//（与元素个数无关），要创建一个长度不为0的空slice
	//需要使用内建函数make，这里我们创建一个长度为5的
	//int类型slice
	//声明一个切片
	s := make([]int, 5)
	fmt.Println("s:", s)
	//切片可以和数组一样设置和得到值
	s[0] = 25
	s[1] = 36
	fmt.Println(s)
	fmt.Println(s[1])
	//len返回slice的长度
	fmt.Println(len(s))
	//向切片添加一个元素
	//该函数会返回一个包含了一个或多个新值得slice
	s = append(s, 1)
	//向切片添加两个元素
	s = append(s, 11, 22)
	fmt.Println("add", s)
	//初始化一个切片，长度为s的长度
	c := make([]int, len(s))
	//将s的值拷贝的c
	copy(c, s)
	fmt.Println("cpy", s)
	//将s的第2个元素到5元素以前赋值给l
	l := s[2:5]
	fmt.Println("sli", l)
	//将s0到5以前的元素赋值给l
	l = s[:5]
	fmt.Println("sl2", l)
	//将s和2以后的值赋值给l
	l = s[2:]
	fmt.Println("sl3", l)
	//声明并初始化一个slice变量
	t := []int{1, 2, 5}
	fmt.Println("dcl", t)
	//创建一个二维切片
	twod := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twod[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println("2d", twod)
}
