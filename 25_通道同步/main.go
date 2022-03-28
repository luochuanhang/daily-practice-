package main

import (
	"fmt"
	"time"
)
//可以使用通道来同步 Go 协程间的执行状态。
//这里是一个使用阻塞的接受方式来等待一个
// Go 协程的运行结束
func work(done chan bool) {
	fmt.Println("开始工作")
	time.Sleep(time.Second)
	fmt.Println("结束工作")
	//发送一个值通知完工
	done <- true
}
func main() {
	ch := make(chan bool)
	//运行一个工作协程
	go work(ch)
	//程序在接收通道中worker发出的通知前一直阻塞
	<-ch
	//如果这行代码从程序中移除，
	//程序甚至会在 worker还没开始运行时就结束了。
}
