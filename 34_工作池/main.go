package main

import (
	"fmt"
	"time"
)

//使用协程和通道实现一个工作池
func main() {
	const workerJobs = 5
	jobs := make(chan int, workerJobs)
	result := make(chan int, workerJobs)
	//启动3个工作线程
	for i := 0; i < 3; i++ {
		go worker(i, jobs, result)
	}
	//发送任务
	for i := 1; i <= workerJobs; i++ {
		jobs <- i
	}
	//关闭任务通道
	close(jobs)
	//等待执行结束
	for a := 1; a < workerJobs; a++ {
		<-result
	}

}

//这是一个worker程序，我们会并发的运行多个worker
func worker(id int, jobs <-chan int, results chan<- int) {
	//等待工作任务
	for j := range jobs {
		//获取到一个任务
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
