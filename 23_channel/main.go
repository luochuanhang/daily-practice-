package main

import "fmt"

func main() {
	//用 make(chan val-type) 创建一个新的通道。
	//通道类型就是他们需要传递值的类型。
	ch := make(chan string)
	go func() {
		//使用 channel <- 语法 发送 一个新的值到通道中
		ch <- "nihao"
	}()
	//使用 <-channel 语法从通道中 接收 一个值
	messages := <-ch
	//我们运行程序时，通过通道，消息 "nihao"
	//成功的从一个 Go 协程传到另一个中。
	fmt.Println(messages)
	//默认发送和接收操作是阻塞的，
	//直到发送方和接收方都准备完毕
}
