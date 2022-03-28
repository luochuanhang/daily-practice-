package main

import (
	"fmt"
	"time"
)
//关闭 一个通道意味着不能再向这个通道发送值了
func main() {
	//我们将使用一个 jobs 通道来传递 main()中Go
	//协程任务执行的结束信息到一个工作 Go 协程中
	//当我们没有多余的任务给这个工作 Go 协程时，
	//我们将 close 这个 jobs 通道。
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
		//使用j, more := <- jobs循环的从jobs接收数据	
		//如果jobs已经关闭了，并且通道中所有的值都已经
		//接收完毕，那么ok的值将是false
			j, ok := <-jobs
			if ok {
				fmt.Println("收到", j)
			} else {
				fmt.Println("全部收到")
				//当我们完成所有的任务时，将使用这个特性通过 
				//done通道去进行通知。
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("发送", i)
		time.Sleep(time.Second)
	}
	close(jobs)
	fmt.Println("发送完毕")
	<-done
}
