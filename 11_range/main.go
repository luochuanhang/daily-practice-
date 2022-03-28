package main

import "fmt"

func main() {
	//range用来迭代各种各样的数据结构
	num := []int{1, 2, 3}
	sum := 0
	//使用range遍历一个切片或者数组会返回两个值
	//下标和对应的值，不需要可以使用空白标识符省略
	for _, v := range num {
		sum += v
	}
	fmt.Println(sum)

	//rang也可以对map进行迭代
	m := map[string]int{"chuanchuan": 12, "mingyue": 21}
	for k, v := range m {
		fmt.Printf("k:%s:v:%d", k, v)
	}

	//range也可以只遍历map的键
	for key := range m {
		fmt.Println(key)
	}
	//range在字符串中迭代，第一个返回是字符串起始字节位置
	//第二个是字符本身
	for i, s := range "go你好h" {
		fmt.Printf("i=%d,s=%c", i, s)
	}
}
