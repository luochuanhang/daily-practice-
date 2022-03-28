package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var num uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			//如果我们使用非原子的 ops++ 来增加计数器，
			// 由于多个协程会互相干扰，运行时值会改变，
			//可能会导致我们得到一个不同的数字。
			for c := 0; c < 1000; c++ {
				//num++
				atomic.AddUint64(&num, 1)
			}
			wg.Done()
		}()
	}
	//可以在运行时通过-race来检测出现问题的位置
	wg.Wait()
	fmt.Println(num)
}
