package main

import "fmt"

//for和range为基本的数据结构提供了迭代的功能。
//我们也可以使用这个语法来遍历从通道中取得的值。
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	//range迭代从queue中得到的每个值。因为我们在前面
	//close了这个通道，这个迭代会在接收完2个值之后结束。
	//如果我们没有close它，我们将在这个循环中继续阻塞
	//执行，等待接收第三个值
	for elem := range queue {
		//一个非空的通道也是可以关闭的，
		//但是通道中剩下的值仍然可以被接收到。
		fmt.Println(elem)
	}
}
