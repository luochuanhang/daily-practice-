package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("hello")
	//用 go f(s) 在一个 Go 协程中调用这个函数。
	//这个新的 Go 协程将会并行的执行这个函数调用。
	go f("nihao")
	//也可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("dajiahao")
	time.Sleep(time.Second)
}
