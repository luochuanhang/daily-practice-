package main

import "os"

func main() {

	//panic 意味着有些出乎意料的错误发生。 通常我们用它来
	//表示程序正常运行中不应该出现的错误， 或者我们不准备优雅处理的错误。
	//panic("a problem")

	//panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。
	_, err := os.Create("/tmp/file1")
	if err != nil {
		panic(err)
	}
	//与某些使用 exception 处理错误的语言不同， 在 Go 中，通常会尽可能的使用返回值来标示错误。
}
