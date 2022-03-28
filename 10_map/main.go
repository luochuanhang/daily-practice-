package main

import (
	"fmt"
)

func main() {
	//要创建一个空map，需要使用内建函数make
	m := make(map[string]int)
	//可以使用name[key]=val语法来设置键值对
	m["k1"] = 5
	m["k2"] = 13
	//打印一个map，会输出所有的键值对
	fmt.Println("map", m)
	//可以通过name[key]来获取一个值
	v1 := m["k1"]
	fmt.Println("v1", v1)
	//内建函数len可以返回map的键值对数量
	fmt.Println("len", len(m))
	//内建函数delete可以从一个map中移除键值对
	delete(m, "k2")
	fmt.Println(m)
	//当从一个map中取值时，还可以选择是否接收的第二个返回值
	//这个值表明了map中是否存在这个键，
	s, ok := m["k1"]
	fmt.Printf("值是%d,是否存在%t", s, ok)

	//可以直接声明和初始化一个map
	n := map[string]int{"foo": 1, "hee": 2}
	fmt.Println(n)
}
