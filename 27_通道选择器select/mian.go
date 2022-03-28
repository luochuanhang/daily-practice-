package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		//各个通道将在若干时间后接收一个值
		//这个用来模拟go协程中阻塞的rpc操作
		time.Sleep(time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		//使用select关键字来同时等待这两个值
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

}
