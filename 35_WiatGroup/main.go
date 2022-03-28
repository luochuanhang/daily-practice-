package main

import (
	"fmt"
	"sync"
	"time"
)

//想要等待多个协程完成可以使用wait group
func worker(id int) {
	fmt.Printf("这是第%d个", id)
	time.Sleep(time.Second)
	fmt.Printf("处理完%d", id)
}

func main() {
	//WaitGroup用于等待这里启动的所有协程完成
	//如果WaitGroup显示传递到函数中，应使用*指针
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		//启动几个协程，就递增计数器
		wg.Add(1)
		//避免在每个协程闭包中重复利用i值
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	//直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成。
	wg.Wait()
}
