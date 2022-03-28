package main

import (
	"fmt"
	"sync"
)

//我们可以使用一个互斥量 来在 Go 协程间
//安全的访问数据。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

//Container 中定义了 counters 的 map ，
//由于我们希望从多个 goroutine 同时更新它，
//因此我们添加了一个 互斥锁Mutex 来同步访问。
//请注意不能复制互斥锁，如果需要传递这个
//struct，应使用指针完成。
func (c *Container) inc(name string) {
	//在访问 counters 之前锁定互斥锁；
	//使用 [defer]（defer） 在函数结束时解锁。
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	//这个函数在循环中递增对 name 的计数
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	//同时运行多个 goroutines; 请注意，它们都访问相同的
	// Container，其中两个访问相同的计数器。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait()
	fmt.Println(c.counters)
}
