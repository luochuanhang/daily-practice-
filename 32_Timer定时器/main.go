package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器表示在未来某一时刻的独立运行事件
	//设置一个等待的时间，然后提供一个用于通知
	//的通道
	timer := time.NewTimer(2 * time.Second)
	//<-timer.c会一直阻塞，直到定时器的通道c
	//明确的发送了定时器失效的值
	<-timer.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	//time stop，可以在定时器触发之前取消
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
	//第一个定时器将在程序开始后大约 2s 触发，
	// 但是第二个定时器还未触发就停止了。
}
