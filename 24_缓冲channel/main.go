package main

import "fmt"

func main() {
	// make 了一个通道，最多允许缓存 2 个值
	ch := make(chan int, 2)
	//因为这个通道是有缓冲区的，即使没有一个对应的并发接收方，
	//我们仍然可以发送这些值。
	ch <- 1
	ch <- 2
	//然后我们可以像前面一样接收这两个值。
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
