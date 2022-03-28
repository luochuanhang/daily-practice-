package main

import (
	"fmt"
	"time"
)

//打点器，以固定的时间间隔重复执行
func main() {
	//打点器和定时器相似，使用一个通道来发送数据
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	//打点器和定时器一样可以被停止，停止将不能接收到值
	time.Sleep(1600 * time.Millisecond)

	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

}
