package main

import "fmt"

func main() {
	//创建一个刚好可以存放5个int元素的数组a
	//元素的类型和长度都是数组类型的一部分，
	//数组默认值是零值，或者array[index]得到值
	var nums [5]int
	fmt.Println(nums)
	//我们可以使用array[index]=value语法来设置数组指定
	//位置的值，或者使用array[index]得到值
	nums[4] = 5
	a := nums[0]
	fmt.Println(a)
	fmt.Println(nums[4])
	//内置函数len返回数组的长度(存放的数据个数)
	fmt.Println(len(nums))
	//声明并初始化一个数组
	num := [4]int{1, 2, 3, 4}
	fmt.Println(num)
	//这是一个二维数组
	n := [2][2]int{{1, 2}, {2, 2}}
	fmt.Println(n)
}
