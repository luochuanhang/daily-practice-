package main

import (
	"fmt"
	"time"
)
//超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序
//而言是很重要的。得益于通道和 select，在 Go中实现超时操作是简洁
//而优雅的。
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
		//等待持续时间结束后，然后在返回的通道上发送当前时间。
		//它等价于NewTimer(d). c。在定时器触发之前，
		//底层的Timer不会被垃圾回收器恢复。如果效率是一个问题，
		//使用NewTimer代替并调用Timer。如果不再需要计时器，则停止。
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c2 <- "nihao"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
