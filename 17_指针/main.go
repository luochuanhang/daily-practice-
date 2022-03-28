package main

import "fmt"

//numval有一个int型参数，所以使用值传递
//将从调用它的那个函数中得到一个实参的拷贝
func numval(ival int) {
	ival = 0
}

//numptr有一个*int参数，这意味它使用了int指针
//函数类*ptr会解引用这个指针，从它的内存地址得到
//这个地址当前对应的值，对解引用的指针赋值，会改变这个
//指针引用的真实地址的值
func numptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial", i)
	numval(i)
	fmt.Println("numval", i)
	//用&i语法可以取得i的内存地址，即指向i的指针
	numptr(&i)
	//指针也可以被打印
	fmt.Println("numptr", i)

	//zeroval 在 main 函数中不能改变 i 的值，
	//但是 zeroptr 可以，因为它有这个变量的
	//内存地址的引用。
	fmt.Println(&i)
}
